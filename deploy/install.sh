#!/bin/bash
export WORK_DIR=$(pwd)

is_command_present() {
    type "$1" >/dev/null 2>&1
}

# Check whether 'wget' command exists.
has_wget() {
    has_cmd wget
}

# Check whether 'curl' command exists.
has_curl() {
    has_cmd curl
}

# Check whether the given command exists.
has_cmd() {
    command -v "$1" > /dev/null 2>&1
}

is_mac() {
    [[ $OSTYPE == darwin* ]]
}

is_arm64(){
    [[ `uname -m` == 'arm64' ]]
}

check_os() {
    if is_mac; then
        package_manager="brew"
        desired_os=1
        os="Mac"
        return
    fi

    os_name="$(cat /etc/*-release | awk -F= '$1 == "NAME" { gsub(/"/, ""); print $2; exit }')"

    case "$os_name" in
        Ubuntu*)
            desired_os=1
            os="ubuntu"
            package_manager="apt-get"
            ;;
        Amazon\ Linux*)
            desired_os=1
            os="amazon linux"
            package_manager="yum"
            ;;
        Debian*)
            desired_os=1
            os="debian"
            package_manager="apt-get"
            ;;
        Linux\ Mint*)
            desired_os=1
            os="linux mint"
            package_manager="apt-get"
            ;;
        Red\ Hat*)
            desired_os=1
            os="red hat"
            package_manager="yum"
            ;;
        CentOS*)
            desired_os=1
            os="centos"
            package_manager="yum"
            ;;
        SLES*)
            desired_os=1
            os="sles"
            package_manager="zypper"
            ;;
        openSUSE*)
            desired_os=1
            os="opensuse"
            package_manager="zypper"
            ;;
        *)
            desired_os=0
            os="Not Found: $os_name"
    esac
}

confirm() {
    local default="$1"  # Should be `y` or `n`.
    local prompt="$2"

    local options="y/N"
    if [[ $default == y || $default == Y ]]; then
        options="Y/n"
    fi

    local answer
    read -rp "$prompt [$options] " answer
    if [[ -z $answer ]]; then
        # No answer given, the user just hit the Enter key. Take the default value as the answer.
        answer="$default"
    else
        # An answer was given. This means the user didn't get to hit Enter so the cursor on the same line. Do an empty
        # echo so the cursor moves to a new line.
        echo
    fi

    [[ yY =~ $answer ]]
}


# This function checks if the relevant ports required by Slashbase are available or not
# The script should error out in case they aren't available
check_ports_occupied() {
    local port_check_output
    local ports_pattern="80|443"

    if is_mac; then
        port_check_output="$(netstat -anp tcp | awk '$6 == "LISTEN" && $4 ~ /^.*\.('"$ports_pattern"')$/')"
    elif is_command_present ss; then
        # The `ss` command seems to be a better/faster version of `netstat`, but is not available on all Linux
        # distributions by default. Other distributions have `ss` but no `netstat`. So, we try for `ss` first, then
        # fallback to `netstat`.
        port_check_output="$(ss --all --numeric --tcp | awk '$1 == "LISTEN" && $4 ~ /^.*:('"$ports_pattern"')$/')"
    elif is_command_present netstat; then
        port_check_output="$(netstat --all --numeric --tcp | awk '$6 == "LISTEN" && $4 ~ /^.*:('"$ports_pattern"')$/')"
    fi

    if [[ -n $port_check_output ]]; then
        echo "+++++++++++++++ ERROR ++++++++++++++++++"
        echo "Slashbase requires ports 80 & 443 to be open. Please shut down any other service(s) that may be running on these ports."
        echo "++++++++++++++++++++++++++++++++++++++++"
        echo ""
        exit 1
    fi
}

check_wget_curl() {
    has_wget_curl=0
    if has_curl; then
        has_wget_curl=1
    elif has_wget; then
        has_wget_curl=1
    else
        echo "+++++++++++ IMPORTANT READ ++++++++++++++++++++++"
        echo "curl or wget not found. please install and try again."
        echo "+++++++++++++++++++++++++++++++++++++++++++++++++"
        exit 1
    fi
}

install_docker() {
    echo "++++++++++++++++++++++++"
    echo "Setting up docker repos"

    if [[ $package_manager == apt-get ]]; then
        apt_cmd="sudo apt-get --yes --quiet"
        $apt_cmd update
        $apt_cmd install software-properties-common gnupg-agent
        curl -fsSL "https://download.docker.com/linux/$os/gpg" | sudo apt-key add -
        sudo add-apt-repository \
            "deb [arch=amd64] https://download.docker.com/linux/$os $(lsb_release -cs) stable"
        $apt_cmd update
        echo "Installing docker"
        $apt_cmd install docker-ce docker-ce-cli containerd.io

    elif [[ $package_manager == zypper ]]; then
        zypper_cmd="sudo zypper --quiet --no-gpg-checks --non-interactive"
        echo "Installing docker"
        if [[ $os == sles ]]; then
            os_sp="$(cat /etc/*-release | awk -F= '$1 == "VERSION_ID" { gsub(/"/, ""); print $2; exit }')"
            os_arch="$(uname -i)"
            sudo SUSEConnect -p "sle-module-containers/$os_sp/$os_arch" -r ''
        fi
        $zypper_cmd install docker docker-runc containerd
        sudo systemctl enable docker.service

    else
        yum_cmd="sudo yum --assumeyes --quiet"
        $yum_cmd install yum-utils
        os_in_repo_link="$os"
        if [[ $os == rhel ]]; then
            # For RHEL, there's no separate repo link. We can use the CentOS one though.
            os_in_repo_link=centos
        fi
        sudo yum-config-manager --add-repo "https://download.docker.com/linux/$os_in_repo_link/docker-ce.repo"
        echo "Installing docker"
        $yum_cmd install docker-ce docker-ce-cli containerd.io

    fi
}

install_docker_compose() {
    if [[ $package_manager == "apt-get" || $package_manager == "zypper" || $package_manager == "yum" ]]; then
        if [[ ! -f /usr/bin/docker-compose ]];then
            echo "++++++++++++++++++++++++"
            echo "Installing docker-compose"
            sudo curl -L "https://github.com/docker/compose/releases/download/1.27.4/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
            sudo chmod +x /usr/local/bin/docker-compose
            sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
            echo "docker-compose installed!"
            echo ""
        fi
    else
        echo "+++++++++++ IMPORTANT READ ++++++++++++++++++++++"
        echo "docker-compose not found! Please install docker-compose first and then continue with this installation."
        echo "Refer https://docs.docker.com/compose/install/ for installing docker-compose."
        echo "+++++++++++++++++++++++++++++++++++++++++++++++++"
        exit 1
    fi
}


generate_secrets() {
    if [ $os = 'Mac' ]; then
        auth_secret = $(env LC_CTYPE=C LC_ALL=C tr -dc "a-zA-Z0-9" < /dev/urandom | head -c 64; echo '')
    else
        auth_secret=$(tr -dc A-Za-z0-9 </dev/urandom | head -c 64 ; echo '')
    fi
    crypted_data_secret=$(openssl rand -hex 32)
}


generate_app_config_file() {
    touch template.yaml & cat << EOF > template.yaml
name: production

server:
    port: ':3001'

database:
    host: slashbase-db
    port: 5432
    user: \${postgres_root_user}
    database: \${postgres_root_user}
    password: \${postgres_root_password}

root_user:
    email: \${slashbase_root_email}
    password: \${slashbase_root_password}

secret:
    auth_token_secret: \${auth_secret}
    crypted_data_secret_key: \${crypted_data_secret}

constants:
    api_host: "https://\${domain_name}"
    app_host: "https://\${domain_name}"

telemetry:
  id: \${telemetry_id}
  enabled: \${enable_telemetry}
EOF

    rm -f replace.sed
    touch replace.sed

    variables=( "postgres_root_user"
                $postgres_root_user
                "postgres_root_password"
                $postgres_root_password
                "auth_secret"
                $auth_secret
                "crypted_data_secret"
                $crypted_data_secret
                "domain_name"
                $domain_name
                "slashbase_root_email"
                $slashbase_root_email
                "slashbase_root_password"
                $slashbase_root_password
                "telemetry_id"
                $telemetry_id
                "enable_telemetry"
                "$enable_telemetry"
            )

    flip=0
    for i in "${variables[@]}"; do
        if [ $flip -eq 0 ]; then
            printf "s/\${$i}/" >> replace.sed
            flip=1
        else
            printf "$i/\n" >> replace.sed
            flip=0
        fi
    done

    sed -f replace.sed template.yaml > production.yaml
    rm -f replace.sed template.yaml
}

generate_docker_env_file() {
    touch docker.env & cat << EOF > docker.env
POSTGRES_DB=\${postgres_root_user}
POSTGRES_USER=\${postgres_root_user}
POSTGRES_PASSWORD=\${postgres_root_password}
EOF

    rm -f replace.sed
    touch replace.sed

    variables=( "postgres_root_user"
                $postgres_root_user
                "postgres_root_password"
                $postgres_root_password
            )

    flip=0
    for i in "${variables[@]}"; do
        if [ $flip -eq 0 ]; then
            printf "s/\${$i}/" >> replace.sed
            flip=1
        else
            printf "$i/\n" >> replace.sed
            flip=0
        fi
    done

    sed -f replace.sed docker.env > .env
    rm -f replace.sed docker.env
}

generate_nginx_config_file() {
    if has_curl; then
        curl -o default.template https://raw.githubusercontent.com/slashbase/slashbase/main/deploy/docker/nginx/default.conf
    elif has_wget; then
        wget -O default.template https://raw.githubusercontent.com/slashbase/slashbase/main/deploy/docker/nginx/default.conf
    fi

    sed "s/domain_name/$domain_name/g" default.template > default.conf
    rm -f default.template
}


start_docker() {
    echo "Starting Docker ..."
    if [ $os = "Mac" ]; then
        open --background -a Docker && while ! docker system info > /dev/null 2>&1; do sleep 1; done
    else 
        if ! sudo systemctl is-active docker.service > /dev/null; then
            echo "Starting docker service"
            sudo systemctl start docker.service
        fi
    fi
}

wait_for_containers_start() {
    local timeout=$1

    # The while loop is important because for-loops don't work for dynamic values
    while [[ $timeout -gt 0 ]]; do
        status_code="$(curl -s -o /dev/null -w "%{http_code}" http://localhost/api/v1 || true)"
        if [[ status_code -eq 401 ]]; then
            break
        else
            echo -ne "Waiting for all containers to start. This check will timeout in $timeout seconds...\r\c"
        fi
        ((timeout--))
        sleep 1
    done

    echo ""
}

# This function prompts the user for an input for a non-empty postgres username.
read_postgres_username() {
    read -rp 'Set the postgres root user: ' postgres_root_user
    while [[ -z $postgres_root_user ]]; do
        echo ""
        echo "+++++++++++ ERROR ++++++++++++++++++++++"
        echo "The postgres username cannot be empty. Please input a valid username string."
        echo "++++++++++++++++++++++++++++++++++++++++"
        echo ""
        read -rp 'Set the postgres root user: ' postgres_root_user
    done
}

# This function prompts the user for an input for a non-empty postgres root password.
read_postgres_password() {
    read -srp 'Set the postgres password: ' postgres_root_password
    while [[ -z $postgres_root_password ]]; do
        echo ""
        echo ""
        echo "+++++++++++ ERROR ++++++++++++++++++++++"
        echo "The postgres password cannot be empty. Please input a valid password string."
        echo "++++++++++++++++++++++++++++++++++++++++"
        echo ""
        read -srp 'Set the postgres password: ' postgres_root_password
    done
    echo ""
}

# This function prompts the user for an input for a non-empty domain host.
read_domain() {
    read -rp 'Enter your custom domain: ' domain_name
    while [[ -z $domain_name ]]; do
        echo ""
        echo "+++++++++++ ERROR ++++++++++++++++++++++"
        echo "The custom domain cannot be empty. Please input a valid custom domain string."
        echo "++++++++++++++++++++++++++++++++++++++++"
        echo ""
        read -rp 'Enter your custom domain: ' domain_name
    done
    echo ""
    echo "+++++++++++ IMPORTANT PLEASE READ ++++++++++++++++++++++"
    echo "Please update your DNS records A record to point to this instance with your domain registrar"
    echo "+++++++++++++++++++++++++++++++++++++++++++++++"
    echo ""
    echo -e "This script will try to generate SSL certificates via LetsEncrypt and serve slashbase at
    "https://$domain_name". Proceed further once you have pointed your DNS to the IP of this instance."
    echo ""
    read -p 'Do you wish to proceed? (yes or no): ' proceed_ssl
    if [ $proceed_ssl == "no" ]; then
        exit 1
    fi
}

# This function prompts the user for an input for a non-empty slashbase root user email.
read_rootuser_email() {
    read -rp 'Set the slashbase root user email (admin user): ' slashbase_root_email
    while [[ -z $slashbase_root_email ]]; do
        echo ""
        echo ""
        echo "+++++++++++ ERROR ++++++++++++++++++++++"
        echo "The slashbase user email cannot be empty. Please input a valid slashbase user email string."
        echo "++++++++++++++++++++++++++++++++++++++++"
        echo ""
        read -rp 'Set the slashbase root user email (admin user): ' slashbase_root_email
    done
}

# This function prompts the user for an input for a non-empty slashbase root user password.
read_rootuser_password() {
    read -srp 'Set the slashbase root user password (admin user): ' slashbase_root_password
    while [[ -z $slashbase_root_password ]]; do
        echo ""
        echo ""
        echo "+++++++++++ ERROR ++++++++++++++++++++++"
        echo "The slashbase user password cannot be empty. Please input a valid slashbase user password string."
        echo "++++++++++++++++++++++++++++++++++++++++"
        echo ""
        read -srp 'Set the slashbase root user password (admin user): ' slashbase_root_password
    done
    echo ""
}

ask_telemetry() {
    echo ""
    echo "+++++++++++ IMPORTANT ++++++++++++++++++++++"
    echo -e "Slashbase sends anonymous data (just version and telemetry ID) to our analytics service"
    echo -e "This helps us know the number of users per version. (No user data collected)"
    echo -e ""
    if confirm y 'Would you like to enable telemetry and receive better support?'; then
        enable_telemetry="true"
    else
        enable_telemetry="false"
    fi
    echo "++++++++++++++++++++++++++++++++++++++++++++"
}

bye() {  # Prints a friendly good bye message and exits the script.
    if [ "$?" -ne 0 ]; then
        set +o errexit
        echo "Please share your email if you wish to receive support with the installation"
        read -rp 'Email: ' email

        curl -XPOST -H "Content-type: application/json" -d '{
            "api_key": "phc_XSWvMvnTUEH9pLJDVmYfaKaKH8QZtK5fJO8NIiFoNwv",
            "event": "Install Failed",
            "properties": {
            "distinct_id": "'$SLASHBASE_INSTALLATION_ID'",
            "email": "'$email'",
            "domain": "'$domain_name'",
            "type": "docker"
            }
        }' 'https://app.posthog.com/capture/' > /dev/null 2>&1


        echo ""
        echo -e "\nWe will reach out to you at the email provided shortly, Exiting for now. Bye! 👋 \n"
        exit 0
    fi
}

# Checking OS and assigning package manager
desired_os=0
os=""
echo -e "Detecting your OS"
check_os

SLASHBASE_INSTALLATION_ID=$(curl -s 'https://api64.ipify.org')

# Run bye if failure happens
trap bye EXIT

if [[ $desired_os -eq 0 ]];then
    echo ""
    echo "This script does not support your OS."
    exit 1
fi

curl -XPOST -H "Content-type: application/json" -d '{
    "api_key": "phc_XSWvMvnTUEH9pLJDVmYfaKaKH8QZtK5fJO8NIiFoNwv",
    "event": "Install Started",
    "properties": {
      "distinct_id": "'$SLASHBASE_INSTALLATION_ID'",
      "type": "docker"
    }
  }' 'https://app.posthog.com/capture/' > /dev/null 2>&1

check_ports_occupied
check_wget_curl
read_domain

# Check is Docker daemon is installed and available. If not, the install & start Docker for Linux machines. We cannot automatically install Docker Desktop on Mac OS
if ! is_command_present docker; then
    if [[ $package_manager == "apt-get" || $package_manager == "zypper" || $package_manager == "yum" ]]; then
        install_docker
    else
        echo ""
        echo "+++++++++++++++ IMPORTANT READ ++++++++++++++++++"
        echo "Docker Desktop must be installed manually on Mac OS to proceed. Docker can only be installed automatically on Ubuntu / openSUSE / SLES / Redhat / Cent OS"
        echo "https://docs.docker.com/docker-for-mac/install/"
        echo "++++++++++++++++++++++++++++++++++++++++++++++++"
        exit 1
    fi
fi

# Install docker-compose
if ! is_command_present docker-compose; then
    install_docker_compose
fi


# Starting docker service
if [[ $package_manager == "yum" || $package_manager == "zypper" || $package_manager == "apt-get" ]]; then
    start_docker
fi

# Generate server config yaml file
read_postgres_username
read_postgres_password
read_rootuser_email
read_rootuser_password
telemetry_id=$SLASHBASE_INSTALLATION_ID
ask_telemetry
generate_secrets
generate_app_config_file
generate_docker_env_file

curl -XPOST -H "Content-type: application/json" -d '{
    "api_key": "phc_XSWvMvnTUEH9pLJDVmYfaKaKH8QZtK5fJO8NIiFoNwv",
    "event": "Configs Generated",
    "properties": {
      "distinct_id": "'$SLASHBASE_INSTALLATION_ID'",
      "root_email": "'$slashbase_root_email'",
      "domain": "'$domain_name'",
      "type": "docker"
    }
  }' 'https://app.posthog.com/capture/' > /dev/null 2>&1

mkdir -p "$WORK_DIR/data/app/config"
mkdir -p "$WORK_DIR/data/nginx/conf.d"
mkdir -p "$WORK_DIR/data/ssl/conf"
mkdir -p "$WORK_DIR/data/ssl/www"
mv production.yaml "$WORK_DIR/data/app/config/"
touch default.conf & cat << EOF > default.conf
server {
    listen      80;
    listen      [::]:80;
    server_name _;
    
    # ACME-challenge
    location ^~ /.well-known/acme-challenge/ {
        root /var/www/certbot;
    }
}
EOF
mv default.conf "$WORK_DIR/data/nginx/conf.d/"

if has_curl; then
    curl -o docker-compose.yaml https://raw.githubusercontent.com/slashbase/slashbase/main/deploy/docker/docker-compose.yaml
elif has_wget; then
    wget -O docker-compose.yaml https://raw.githubusercontent.com/slashbase/slashbase/main/deploy/docker/docker-compose.yaml
fi


echo ""
echo "Pulling the latest container images"
sudo docker-compose pull

echo ""
echo "Starting the Slashbase containers"
# The docker-compose command does some nasty stuff for the `--detach` functionality. So we add a `|| true` so that the
# script doesn't exit because this command looks like it failed to do it's thing.
sudo docker-compose up --detach --remove-orphans || true

wait_for_containers_start 60
echo ""

if [[ $status_code -ne 404 ]]; then
    echo "+++++++++++ ERROR ++++++++++++++++++++++"
    exit 1
fi

echo ""
echo "Installing SSL certificates..."
sudo docker run -it --rm --name certbot \
    -v "$WORK_DIR/data/nginx/conf.d:/etc/nginx/conf.d" \
    -v "$WORK_DIR/data/ssl/conf:/etc/letsencrypt" \
    -v "$WORK_DIR/data/ssl/www:/var/www/certbot" \
    certbot/certbot certonly --webroot -w /var/www/certbot -d $domain_name

generate_nginx_config_file
mv default.conf "$WORK_DIR/data/nginx/conf.d/"
docker exec slashbase-web nginx -s reload

if [[ $status_code -eq 404 ]]; then
    echo "++++++++++++++++++ SUCCESS ++++++++++++++++++++++"
    echo "Slashbase installation is complete!"
    echo "Your app is running at https://$domain_name"
    echo "Use the root email and password you specified to login."
    echo "+++++++++++++++++++++++++++++++++++++++++++++++++"
fi

curl -XPOST -H "Content-type: application/json" -d '{
    "api_key": "phc_XSWvMvnTUEH9pLJDVmYfaKaKH8QZtK5fJO8NIiFoNwv",
    "event": "Install Success",
    "properties": {
      "distinct_id": "'$SLASHBASE_INSTALLATION_ID'",
      "root_email": "'$slashbase_root_email'",
      "domain": "'$domain_name'",
      "type": "docker"
    }
  }' 'https://app.posthog.com/capture/' > /dev/null 2>&1