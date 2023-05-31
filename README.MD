# IP and Server Lookup

A command-line application for retrieving IP addresses and server names from the internet.

- [Installation](#installation)
- [Prerequisites](#prerequisites)
- [Usage](#usage)
  - [Command Line Options](#command-line-options)
  - [Commands](#commands)
    - [Retrieve IP addresses](#retrieve-ip-addresses)
    - [Retrieve server names](#retrieve-server-names)
- [License](#license)

## Prerequisites

- Go (version 1.13 or higher) installed on your machine

## Installation

1. Clone the repository:

```shell
  git clone https://github.com/nicktheodoro/ip-and-server-lookup.git
```

2. Open a terminal and navigate to the project directory:

```shell
  cd ip-and-server-lookup
```

3. Build the application:

```shell
  go build
```

## Usage

### Command Line Options

```shell
  --host: Specify the target host (default: github.com)
```

### Commands

#### Retrieve IP addresses

To retrieve IP addresses associated with a host, use the following command:

```shell
  ./ip-and-server-lookup.exe ip --host example.com
```

#### Retrieve server names

To retrieve server names associated with a host, use the following command:

```shell
  ./ip-and-server-lookup.exe servers --host example.com
```

## License

This project is licensed under the <strong><a href="https://mit-license.org/">MIT License</a></strong>.
