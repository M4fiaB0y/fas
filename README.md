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
This command will install the required tools (findomain, amass, subfinder, rusolver) if they are not already installed on your system. Make sure to have Go (Golang) installed on your system to run this script.
```
Usage
To perform subdomain enumeration, you can use the following command:

go run subdomain_scanner.go -p <number of jobs> -f <input file> -t <number of threads> -r <custom resolvers file> -o <output file> -e
Replace the command-line arguments as follows:

-p: Number of parallel jobs to run.
-f: Path to a file containing a list of hosts to scan.
-t: Number of threads to use for the resolver.
-r: Path to a custom resolvers file (optional).
-o: Output file name.
-e: Use external subdomains (optional).
For more details on command-line options, you can use the -h flag:

shell
Copy code
go run subdomain_scanner.go -h
Example
Here's an example command for subdomain enumeration:

shell
Copy code
go run subdomain_scanner.go -p 10 -f targets.txt -t 4 -o subdomains.txt -e
This command will perform subdomain enumeration using 10 parallel jobs, a custom resolvers file (if provided), and save the results in the subdomains.txt file. External subdomains will be included in the enumeration.
