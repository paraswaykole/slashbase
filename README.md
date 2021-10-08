<p align="center">
  <a href="https://slashbase.com" alt="Slashbase">
    <img src="https://slashbase.com/github-logo.png" alt="Slashbase" width="220">
  </a>
  <p align="center">The open-source collaborative IDE for your databases in your browser.</p>
</p>
<p align="center">
  <img src="https://img.shields.io/github/license/slashbase/slashbase-server" alt="License">
  <img src="https://img.shields.io/github/commit-activity/m/slashbase/slashbase-server" alt="Commits per month">
  <img src="https://img.shields.io/github/go-mod/go-version/slashbase/slashbase-server.svg" alt="Go verison">
  <img src="https://img.shields.io/github/release/slashbase/slashbase-server.svg" alt="Release version">
</p>

___

<img src="https://slashbase.com/screenshot.png" alt="Slashbase" width="100%">

## About

Slashbase is an open-source collaborative IDE for your databases in your browser. Connect to your database, browse data, run a bunch of SQL commands or share SQL queries with your team, right from your browser!

It's written in Golang and Nextjs React Framework (SPA) and runs as a single Linux binary with PostgreSQL. Documentation is currently WIP.

## Demo

Checkout demo at: [https://demo.slashbase.com](https://demo.slashbase.com)

## Installation

The install script is tested on DigitalOcean Ubuntu 20.04 droplet. Requires postgres to run. Install it on your cloud instance or use a managed postgres database. To install Postgres on instance follow [this tutorial](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-ubuntu-20-04). 

Then run the following commands:

1. `mkdir slashbase && cd slashbase`
2. `wget https://raw.githubusercontent.com/slashbase/slashbase-server/main/deploy/install.sh && chmod +x install.sh`
3. `./install.sh`
4. Enter required configs & follow the instructions till the setup completes.
5. Visit your url on browser and enter the root user credentials you entered while installation to login.

## License

See the [LICENSE file](LICENSE.txt) for license rights and limitations.
