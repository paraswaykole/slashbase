<p align="center">
  <a href="https://slashbase.com" alt="Slashbase">
    <img src="https://raw.githubusercontent.com/slashbaseide/.github/main/banner.png" alt="Slashbase" width="100%">
  </a>
</p>
<p align="center">
  <img alt="GitHub" src="https://img.shields.io/github/license/slashbaseide/slashbase">
  <img src="https://img.shields.io/github/go-mod/go-version/slashbaseide/slashbase.svg" alt="Go verison">
  <a href="https://github.com/slashbaseide/slashbase/releases">
    <img src="https://img.shields.io/github/release/slashbaseide/slashbase.svg" alt="Release version">
  </a>
  <a href="https://discord.gg/U6fXgm3FAX">
    <img src="https://img.shields.io/discord/1039799991776067615?label=discord" alt="Discord">
  </a>
</p>
<p align="center">
  <a href="https://discord.gg/U6fXgm3FAX">Join Discord Community</a>
  ·
  <a href="https://slashbase.bip.wiki">Explore the docs</a>
  ·
  <a href="https://slashbase.com/updates">Read Updates</a>
  <br/><br/>
  <a href="#installation" rel="dofollow"><strong>Install Now »</strong></a>
</p>


# About

Slashbase is a modern in-browser database IDE & CLI for your dev/data workflows. Use Slashbase to connect to any of your database, browse data and schema, write, run and save queries, create charts, right from your browser. Supports PostgreSQL and MongoDB.

It is in beta (v0.4), help us make it better by sending your feedback and reach a stable (v1.0) version.

## Features:

- **🧑‍💻 Browser & CLI based**: Use as IDE in your browser or as CLI in terminal.
- **🪄 Modern Interface**: With a modern interface, it is easy to use.
- **🪶 Lightweight**: Doesn't take much space on your system.
- **⚡️ Quick Browse**: Quickly filter, sort & browse data and schema with a low-code UI.
- **💾 Save Queries**: Write and Save queries to re-run in future.
- **📊 Create Charts**: Create charts from your query results.
- **🗂 Projects**: Organise all database connections into various projects.
- **📕 Query Cheatsheets**: Search for query commands right inside IDE, no need to search online when you forget query syntax.
- **✅ Database Support**: PostgreSQL and MongoDB.

### In browser
<img src="https://raw.githubusercontent.com/slashbaseide/.github/main/screenshot.png" alt="Slashbase IDE" width="100%">

### In terminal
<p align="center">
  <img src="https://raw.githubusercontent.com/slashbaseide/.github/main/slashbase-cli.gif" alt="Slashbase CLI" width="75%">
</p>

# Installation

The app should be installed on your desktop/laptop and should used as a part of your local development enviroment. Follow the steps below to download & start the app:

## Mac OS

### Using Homebrew
```
brew install slashbaseide/brew/slashbase
```

### Download Binary
1. Download the [latest release](https://github.com/slashbaseide/slashbase/releases) and extract the zip.
2. Run `./slashbase` on terminal to start the app.

## Linux

### Download Binary
1. Download the [latest release](https://github.com/slashbaseide/slashbase/releases) and extract the zip.
2. Run `./slashbase` on terminal to start the app.

## Windows

### Using Scoop

```
scoop bucket add kulfi-scoop https://github.com/Animesh-Ghosh/kulfi-scoop
scoop install slashbase
```

If you don't have scoop installed on windows:
```
iwr -useb get.scoop.sh | iex
```

### Download Binary
1. Download the [latest release](https://github.com/slashbaseide/slashbase/releases) and extract the zip.
2. Run `slashbase` on cmd terminal to start the app.

# Usage

## Slashbase CLI

Run slashbase:
```
slashbase
```

Check slashbase version:
```
slashbase version
```

View slashbase help:
```
slashbase help
```

## Slashbase IDE
To connect to Slashbase IDE:
1. Make sure Slashbase is running
2. Visit http://localhost:22022

# Documentation

Detailed documentation is available on [slashbase guide](https://slashbase.bip.wiki).

# Community

Join our community on [discord](https://discord.gg/U6fXgm3FAX) and [bip](https://bip.so/slashbase/feed).

# Roadmap

### Database Support
- ✅ PostgreSQL Query Engine
- ✅ MongoDB Query Engine
- ☑️ MySQL Query Engine
- ☑️ SQLite Query Engine
- ☑️ Redis Query Engine

### Features
- ✅ Query Cheatsheets
- ☑️ Add/delete Data Models (Table/collections)
- ☑️ Manage Views
- ☑️ Export/import data


# Contributing

Read our [contribution guide](CONTRIBUTING.md) for getting started on contributing to the project.

# Support

If you face any issues installing or using slashbase, send us a mail on slashbaseide@gmail.com or contact support chat on our website [slashbase.com](https://slashbase.com).

# License

See the [LICENSE file](LICENSE.txt) for license rights and limitations.
