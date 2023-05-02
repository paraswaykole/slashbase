#!/bin/bash
set -eu

exists() {
  command -v "$1" 1>/dev/null 2>&1
}

check_dependencies() {
	if ! exists curl; then
		echo "curl is not installed."
		exit 1
	fi

	if ! exists docker; then
		echo "docker is not installed."
		exit 1
	fi

	if ! exists docker compose; then
		echo "docker compose is not installed."
		exit 1
	fi
}

download() {
	curl --fail --silent --location --output "$2" "$1"
}

generate_variables() {
    auth_secret=$(tr -dc A-Za-z0-9 </dev/urandom | head -c 64 ; echo '')
    crypted_data_secret=$(openssl rand -hex 32)
    app_db_host=slashbase-db
    app_db_port=5432
    app_db_user=$postgres_root_user
    app_db_pass=$postgres_root_password
    app_db_name=$postgres_root_user
}

get_containers() {
	echo "fetching docker-compose.yml from slashbase repo"
	download https://raw.githubusercontent.com/slashbaseide/slashbase/main/deploy/docker-compose.yml docker-compose.yml
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

generate_app_config_file() {
    echo "fetching env file from slashbase repo"
	download https://raw.githubusercontent.com/slashbaseide/slashbase/main/deploy/server.env.sample server.env.sample
    
    rm -f replace.sed
    touch replace.sed

    variables=( "auth_secret"
                $auth_secret
                "crypted_data_secret"
                $crypted_data_secret
                "slashbase_root_email"
                $slashbase_root_email
                "slashbase_root_password"
                $slashbase_root_password
                "app_db_host"
                $app_db_host
                "app_db_port"
                $app_db_port
                "app_db_user"
                $app_db_user
                "app_db_pass"
                $app_db_pass
                "app_db_name"
                $app_db_name
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

    sed -f replace.sed server.env.sample > app.env
    rm -f replace.sed server.env.sample
}

wait_for_containers_start() {
    local timeout=$1

    # The while loop is important because for-loops don't work for dynamic values
    while [[ $timeout -gt 0 ]]; do
        status_code="$(curl -s -o /dev/null -w "%{http_code}" http://localhost:3000/api/v1/health || true)"
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

check_dependencies
read_postgres_username
read_postgres_password
read_rootuser_email
read_rootuser_password
generate_variables
generate_app_config_file
generate_docker_env_file
get_containers

echo ""
echo "Pulling the latest container images"
sudo docker compose pull

echo ""
echo "Starting the Slashbase containers"
sudo docker compose up --detach --remove-orphans || true

wait_for_containers_start 30
echo ""

if [[ $status_code -ne 200 ]]; then
    echo "+++++++++++ ERROR ++++++++++++++++++++++"
    exit 1
fi


if [[ $status_code -eq 200 ]]; then
    echo "++++++++++++++++++ SUCCESS ++++++++++++++++++++++"
    echo "Slashbase app is running on port 3000"
    echo "+++++++++++++++++++++++++++++++++++++++++++++++++"
fi