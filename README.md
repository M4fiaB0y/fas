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

 **Environment Setup:** Before using the script, ensure you have the required dependencies installed on your system. You can set up the environment with a single command:

```
go run subdomain_scanner.go -s
```
This command will install the necessary tools (findomain, amass, subfinder, rusolver) if they are not already present on your system. Make sure you have Go (Golang) installed to run this script.
## Usage
![image](https://github.com/M4fiaB0y/fas/assets/95071636/3ac6199f-07a4-4f44-af9c-a51bb7255fad)

To utilize this script for subdomain enumeration, follow these steps:

### Perform Subdomain Enumeration
Execute the script with the following command, customizing the command-line arguments as needed:
```
go run subdomain_scanner.go -p <number of jobs> -f <input file> -t <number of threads> -r <custom resolvers file> -o <output file> -e
```
#### Replace the command-line arguments as follows:
`-p`: Number of parallel jobs to run.
`-f`: Path to a file containing a list of hosts to scan.
`-t`: Number of threads to use for the resolver.
`-r`: Path to a custom resolvers file (optional).
`-o`: Output file name.
`-e`: Use external subdomains (optional).
### Additional Options
For more details on available command-line options, you can use the `-h` flag:
```
go run subdomain_scanner.go -h
```
### Example
Here's an example command for subdomain enumeration:
```
go run subdomain_scanner.go -p 10 -f targets.txt -t 4 -o subdomains.txt -e
```
This command performs subdomain enumeration with 10 parallel jobs, utilizes a custom resolvers file (if provided), and saves the results in the subdomains.txt file. External subdomains will be included in the enumeration.
## Contact 

`https://www.facebook.com/mafiab0yy`
`https://twitter.com/mafiab0yy`
 

