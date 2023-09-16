package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

)

func usage() {
	fmt.Println("Usage: go run main.go -p <number of jobs> -f <input file> -t <number of threads> -r <custom resolvers file> -o <output file> -e -s -g -l <resolver retries> -h")
	fmt.Println("  -p = number of jobs to run in parallel")
	fmt.Println("  -f = file containing a list of hosts to scan")
	fmt.Println("  -t = number of threads to use for the resolver")
	fmt.Println("  -r = path to custom resolvers file")
	fmt.Println("  -o = output file name")
	fmt.Println("  -e = use external subdomains")
	fmt.Println("  -s = setup the environment")
	fmt.Println("  -g = overwrite the output file if it exists")
	fmt.Println("  -l = number of retries for the resolver")
	fmt.Println("  -h = display this help message")
	os.Exit(1)
}

func setupEnv() {
	fmt.Println("Setting up the environment...")

	cmd := exec.Command("sudo", "apt-get", "update")
	cmd.Run()

	cmd = exec.Command("sudo", "apt-get", "install", "-y", "curl", "wget", "git", "tar", "unzip")
	cmd.Run()

	cmd = exec.Command("sh", "-c", `cd $(mktemp -d) && curl -LO https://github.com/findomain/findomain/releases/latest/download/findomain-linux.zip && unzip findomain-linux.zip && chmod +x findomain && sudo mv findomain /usr/local/bin/findomain`)
	cmd.Run()

	cmd = exec.Command("sh", "-c", `cd $(mktemp -d) && curl -LO https://github.com/owasp-amass/amass/releases/latest/download/amass_Linux_amd64.zip && unzip amass_Linux_amd64.zip && chmod +x amass_Linux_amd64/amass && sudo mv amass_Linux_amd64/amass /usr/local/bin/amass`)
	cmd.Run()

	cmd = exec.Command("sh", "-c", `cd $(mktemp -d) && curl -LO https://github.com/projectdiscovery/subfinder/releases/download/v2.5.8/subfinder_2.5.8_linux_amd64.zip && unzip subfinder_2.5.8_linux_amd64.zip && chmod +x subfinder && sudo mv subfinder /usr/local/bin/subfinder`)
	cmd.Run()

	cmd = exec.Command("sh", "-c", `cd $(mktemp -d) && curl -LO https://github.com/Edu4rdSHL/rusolver/releases/latest/download/rusolver-linux && chmod +x rusolver-linux && sudo mv rusolver-linux /usr/local/bin/rusolver`)
	cmd.Run()
}

func runFindomainAndRusolver(target, output string, useExternalSubdomains bool, rusolverThreads int, customResolversFile string, resolverRetries int) {
	excludedSubdomains := "www."
	var rusolverCommand []string

	if customResolversFile != "" {
		rusolverCommand = []string{"rusolver", "-t", fmt.Sprintf("%d", rusolverThreads), "--retries", fmt.Sprintf("%d", resolverRetries), "--no-verify", "-r", customResolversFile}
	} else {
		rusolverCommand = []string{"rusolver", "-t", fmt.Sprintf("%d", rusolverThreads), "--retries", fmt.Sprintf("%d", resolverRetries), "--no-verify"}
	}

	findomainCommand := []string{"findomain", "-qt", target, "--exclude", excludedSubdomains}

	if useExternalSubdomains {
		findomainCommand = append(findomainCommand, "--external-subdomains")
	}

	findomainCmd := exec.Command(findomainCommand[0], findomainCommand[1:]...)
	rusolverCmd := exec.Command(rusolverCommand[0], rusolverCommand[1:]...)

	rusolverCmd.Stdin, _ = findomainCmd.StdoutPipe()
	rusolverCmd.Stdout, _ = os.Create(output)

	rusolverCmd.Start()
	findomainCmd.Run()
	rusolverCmd.Wait()
}

func runJobs(maxParallelJobs int, inputFile, outputFile string, withExternalSubdomains bool, rusolverThreads int, customResolversFile string, rusolverRetries int) {
	var currentPool []*exec.Cmd
	totalLines := 0
	currentLine := 0

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		totalLines++
	}

	file, err = os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		target := scanner.Text()

		target = strings.TrimSpace(target)

		target = strings.TrimPrefix(target, "http://")
		target = strings.TrimPrefix(target, "https://")

		if len(target) == 0 {
			continue
		}

		for len(currentPool) >= maxParallelJobs {
			var newPool []*exec.Cmd
			for _, cmd := range currentPool {
				if cmd.ProcessState == nil || !cmd.ProcessState.Exited() {
					newPool = append(newPool, cmd)
				}
			}
			currentPool = newPool
		}

		fmt.Printf("Scanning %s\n", target)

		go func(target string) {
			runFindomainAndRusolver(target, outputFile, withExternalSubdomains, rusolverThreads, customResolversFile, rusolverRetries)
			currentPool = append(currentPool, nil)
		}(target)

		currentLine++
		percentage := float64(currentLine) * 100.0 / float64(totalLines)
		fmt.Printf("Progress: %.2f%%\r", percentage)
	}

	for _, cmd := range currentPool {
		if cmd != nil {
			cmd.Wait()
		}
	}
	fmt.Println("\nSubdomain enumeration completed.")
}

func main() {
	var maxParallelJobs int
	var inputFile string
	var outputFile string
	var withExternalSubdomains bool
	var rusolverThreads int
	var customResolversFile string
	var rusolverRetries int
	var setupEnvFlag bool
	var overwriteOutput bool

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-p":
			maxParallelJobs = atoi(args[i+1])
			i++
		case "-f":
			inputFile = args[i+1]
			i++
		case "-t":
			rusolverThreads = atoi(args[i+1])
			i++
		case "-r":
			customResolversFile = args[i+1]
			i++
		case "-o":
			outputFile = args[i+1]
			i++
		case "-e":
			withExternalSubdomains = true
		case "-s":
			setupEnvFlag = true
		case "-g":
			overwriteOutput = true
		case "-l":
			rusolverRetries = atoi(args[i+1])
			i++
		case "-h":
			usage()
		default:
			usage()
		}
	}

	if maxParallelJobs == 0 && inputFile == "" && rusolverThreads == 0 && customResolversFile == "" && outputFile == "" && !withExternalSubdomains && !setupEnvFlag && !overwriteOutput && rusolverRetries == 0 {
		usage()
	}

	if setupEnvFlag {
		setupEnv()
		return
	}

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Printf("Input file does not exist: %s\n", inputFile)
		os.Exit(1)
	}

	if _, err := os.Stat(customResolversFile); customResolversFile != "" && os.IsNotExist(err) {
		fmt.Printf("Custom resolvers file does not exist: %s\n", customResolversFile)
		os.Exit(1)
	}

	if _, err := os.Stat(outputFile); err == nil {
		fmt.Printf("Output file already exists: %s. Moving it to %s.old\n", outputFile, outputFile)
		if _, err := os.Stat(outputFile + ".old"); err == nil {
			if overwriteOutput {
				fmt.Printf("Old output file already exists but overwriting is enabled. Overwriting it.\n")
			} else {
				fmt.Printf("Old output file already exists: %s. Please move/remove it and try again or use the -g option to overwrite it.\n", outputFile+".old")
				os.Exit(1)
			}
		}
		err := os.Rename(outputFile, outputFile+".old")
		if err != nil {
			fmt.Printf("Error renaming the existing output file: %s\n", err)
			os.Exit(1)
		}
	}

	if _, err := os.Stat(outputFile + "_total"); err == nil {
		os.Rename(outputFile+"_total", outputFile+"_total.old")
	}

	runJobs(maxParallelJobs, inputFile, outputFile, withExternalSubdomains, rusolverThreads, customResolversFile, rusolverRetries)

	fmt.Println("\n----------------------------------------")
	fmt.Printf("Subdomain enumeration completed.\n")
	fmt.Printf("Results saved to: %s\n", outputFile)
	fmt.Println("----------------------------------------")
}

func atoi(s string) int {
	n, err := fmt.Sscanf(s, "%d", new(int))
	if err != nil || n != 1 {
		return 0
	}
	return *new(int)
}

