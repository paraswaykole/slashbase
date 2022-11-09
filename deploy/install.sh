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

	if ! exists docker-compose; then
		echo "docker-compose is not installed."
		exit 1
	fi
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
        answer="$default"
    else
        echo
    fi

    [[ yY =~ $answer ]]
}

download() {
	curl --fail --silent --location --output "$2" "$1"
}

generate_variables() {
    if confirm y "Would you like to share anonymous usage data and receive better support?"; then
        telemetry_id=$(curl -s 'https://api64.ipify.org')
    else
        telemetry_id="disabled"
    fi
    auth_secret=$(tr -dc A-Za-z0-9 </dev/urandom | head -c 64 ; echo '')
    crypted_data_secret=$(openssl rand -hex 32)
}

get_containers() {
	echo "fetching docker-compose.yml from slashbase repo"
	download https://raw.githubusercontent.com/slashbaseide/slashbase/main/deploy/docker-compose.yaml docker-compose.yaml
}

generate_app_config_file() {
    echo "fetching env file from slashbase repo"
	download https://raw.githubusercontent.com/slashbaseide/slashbase/main/deploy/production.env.sample production.env.sample
    
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
                "telemetry_id"
                $telemetry_id
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

    sed -f replace.sed production.env.sample > app.env
    rm -f replace.sed production.env.sample
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

# Generate slashbase app.env file
read_rootuser_email
read_rootuser_password
generate_variables
generate_app_config_file
get_containers

echo ""
echo "Pulling the latest container images"
sudo docker-compose pull

echo ""
echo "Starting the Slashbase containers"
sudo docker-compose up --detach --remove-orphans || true

wait_for_containers_start 60
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

echo ""
if confirm y "Would you like to share your email to receive updates from Slashbase?"; then
    read -rp 'Email: ' email
    curl -XPOST -H "Content-type: application/json" -d '{
        "api_key": "phc_XSWvMvnTUEH9pLJDVmYfaKaKH8QZtK5fJO8NIiFoNwv",
        "event": "New Installation",
        "properties": {
        "distinct_id": "'$SLASHBASE_INSTALLATION_ID'",
        "email": "'$email'",
        "type": "docker"
        }
    }' 'https://app.posthog.com/capture/' > /dev/null 2>&1
fi
