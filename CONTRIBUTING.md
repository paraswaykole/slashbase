# Contributing to Slashbase

Thanks & Cheers for taking time and interest in contributing to Slashbase 🙌 🙏 👏

The following is a set of guidelines for contributing to Slashbase.

## Table Of Contents

- [Important Resources](#important-resources)
- [Setup Development Environment](#setup-development-environment)
- [Report a bug](#report-a-bug)

## Important Resources:

- [docs.slashbase.com](https://docs.slashbase.com)

## Setup Development Environment

To contribute code to the repository, you need to setup the development environment. To do that, you can follow the following steps:

1. Clone or Fork-Clone the GitHub repo and open it in your preferred IDE (VS-Code recommended).
2. You need docker installed on your desktop, follow [this link](https://docs.docker.com/desktop/) to install.
3. Go to the project root directory and
    - Copy the file at `frontend/.env.local.txt` and paste as `frontend/.env.local`
    - Copy the file at `src/config/development.yaml.txt` and paste as `src/config/development.yaml`
    - Edit `src/config/development.yaml` file and update the root user email and password.
5. Open terminal at the project root directory and run `docker-compose up`. Once it starts, server is running at `http://localhost:3001` & Client at `http://localhost:3000`
7. Create a new branch and make changes to the code.
8. Test and make sure the code runs (run `docker-compose up` or restart `docker-compose restart slashbase-server` to reflect go code changes)
9. Push your code and send PR.

## Report a bug

You can report bugs or any issues in GitHub Issues. Put relevant `bug` & `backend` or `frontend` or both labels to the issue.

## File a feature request

You can file feature request in GitHub Issues. Put `feature request` label to the issue.
