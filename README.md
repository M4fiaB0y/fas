# Subdomain Enumeration Script

This script performs subdomain enumeration using various tools such as `findomain`, `amass`, `subfinder`, and `rusolver`. It helps security professionals and penetration testers discover subdomains associated with a target domain.

## Features

- Parallel execution of subdomain enumeration jobs to speed up the process.
- Support for custom resolvers for improved DNS resolution.
- Option to include external subdomains in the enumeration.
- Easy setup of the required environment with a single command.
- Overwrite protection for the output file.
- Detailed progress tracking and reporting.

## Installation

Before using the script, you need to set up the required environment, which includes installing the necessary dependencies. Run the following command to set up the environment:

```shell
go run subdomain_scanner.go -s
