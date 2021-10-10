#!/bin/bash
export WORK_DIR=$(pwd)

wget "https://github.com/slashbase/slashbase-web/releases/download/v1.0.0-beta/web-release.tar.gz" && wget "https://github.com/slashbase/slashbase-server/releases/download/v1.0.0-beta/server-release.tar.gz"
tar -xvf web-release.tar.gz && tar -xvf server-release.tar.gz

touch template.yaml & cat << EOF > template.yaml
name: production

server:
    port: ':3000'

database:
    host: \${db_host}
    user: \${db_user}
    database: \${db_name}
    password: \${db_pass}

root_user:
    email: \${root_email}
    password: \${root_pass}

secret:
    auth_token_secret: \${auth_secret}
    crypted_data_secret_key: \${crypted_data_secret}

constants:
    api_host: "https://\${domain}"
    app_host: "https://\${domain}"
EOF

echo "Enter the postgres host (default 127.0.0.1):"
read db_host
echo "Enter the postgres database name (default slashbase):"
read db_name
echo "Enter the postgres db user (default slashbase):"
read db_user
echo "Enter the postgres db password:"
read db_pass
echo "Enter the domain/host:"
read domain
echo "Enter the root user email (your/admin email):"
read root_email
echo "Enter the root user password:"
read root_pass

auth_secret=$(tr -dc A-Za-z0-9 </dev/urandom | head -c 64 ; echo '')
crypted_data_secret=$(openssl rand -hex 32)

# create backend configs
rm -f replace.sed
touch replace.sed
variables=( "db_host"
            $db_host
            "db_user"
            $db_user
            "db_name"
            $db_name
            "db_pass"
            $db_pass
            "auth_secret"
            $auth_secret
            "crypted_data_secret"
            $crypted_data_secret
            "domain"
            $domain
            "root_email"
            $root_email
            "root_pass"
            $root_pass
          )

flip=0
for i in "${variables[@]}"; do
  if [ $flip -eq 0 ]
  then
    printf "s/\${$i}/" >> replace.sed
    flip=1
  else
    printf "$i/\n" >> replace.sed
    flip=0
  fi
done

sed -f replace.sed template.yaml > production.yaml
rm -f replace.sed template.yaml

# create frontend configs
touch replace.sed
variables=( "#API_HOST#"
            "https:\/\/${domain}"
          )

flip=0
for i in "${variables[@]}"; do
  if [ $flip -eq 0 ]
  then
    printf "s/$i/" >> replace.sed
    flip=1
  else
    printf "$i/\n" >> replace.sed
    flip=0
  fi
done
sed -i -f replace.sed out/config.js
rm -f replace.sed

# create, install and start slashbase service
touch slashbase.service
cat << EOF > slashbase.service
[Unit]
Description=Slashbase HTTP API
After=network.target

[Service]
Type=simple
WorkingDirectory=${WORK_DIR}
ExecStart=${WORK_DIR}/backend -e production
Restart=on-failure
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
EOF
sudo mv slashbase.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl start slashbase.service
sudo systemctl enable slashbase.service
rm -f slashbase.service

# install nginx
sudo apt update
sudo apt install nginx

# create frontend html folder
sudo mkdir -p /var/www/$domain/html
sudo chown -R $USER:$USER /var/www/$domain/html
sudo chmod -R 755 /var/www/$domain
sudo mv $WORK_DIR/out/* /var/www/$domain/html/

# setup nginx configs
touch $domain
cat << EOF > $domain
server {
        listen 80;
        listen [::]:80;

        root /var/www/${domain}/html;
        index index.html index.htm index.nginx-debian.html;

        server_name ${domain};

        location / {
            try_files \$uri /index.html;
        }

        location /api {
            proxy_pass http://localhost:3000;
            proxy_http_version 1.1;
            proxy_set_header Upgrade \$http_upgrade;
            proxy_set_header Connection 'upgrade';
            proxy_set_header Host \$host;
            proxy_cache_bypass \$http_upgrade;
        }
}
EOF
sudo mv $domain /etc/nginx/sites-available/
sudo ln -s /etc/nginx/sites-available/$domain /etc/nginx/sites-enabled/
sudo systemctl restart nginx

# install certbot for SSL
sudo snap install --classic certbot
sudo ln -s /snap/bin/certbot /usr/bin/certbot
sudo certbot --nginx