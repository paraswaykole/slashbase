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
2. Go to the project root directory and
    - Copy the file at `frontend/.env.local.txt` and paste as `frontend/.env.local`.
    - Copy the file at `development.env.sample` and paste as `development.env` in the root directory of the project.
    - Edit `development.env` file and update the root user email and password.
3. Open terminal at the project root directory & run `mkdir data` for local sqlite database.
4. Open terminal at the project root directory & run `go run main.go` to run go backend server & `cd frontend & yarn dev` to run frontend nextjs server. Go server is running at `http://localhost:3001` & Frontend client at `http://localhost:3000`
5. Create a new branch and make changes to the code.
6. Test and make sure the code runs
7. Push your code and send PR.

## Report a bug

You can report bugs or any issues in GitHub Issues. Put relevant `bug` & `backend` or `frontend` or both labels to the issue.

## File a feature request

You can file feature request in GitHub Issues. Put `feature request` label to the issue.
