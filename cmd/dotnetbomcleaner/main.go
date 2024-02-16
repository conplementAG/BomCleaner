package main

import (
	"dotnetbomcleaner/internal/bom"
	"dotnetbomcleaner/internal/deps"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	if len(os.Args) < 2 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		printHelp()
		os.Exit(1)
	}

	fmt.Println("Start...")
	workingDir, _ := os.Getwd()
	fmt.Println("Current working directory: ", workingDir)
	fmt.Println("Given file path: ", os.Args[1])

	bomFilePath, depsFilepath, customValues := readInput()

	runtimeDependencies, errRun := deps.GetRuntimeDependencies(depsFilepath)
	checkError(errRun)
	runtimeDependencies = append(runtimeDependencies, *customValues...)

	errBom := bom.CleanupBom(bomFilePath, runtimeDependencies)
	checkError(errBom)

	fmt.Println("Finished successfully!")
	os.Exit(0)
}

func readInput() (string, string, *[]string) {
	bomFilePath := filepath.Clean(os.Args[1])
	depsFilepath := filepath.Clean(os.Args[2])

	_, err2 := os.Stat(bomFilePath)
	if err2 != nil {
		fmt.Println("Could not find file at path: %w", bomFilePath)
		os.Exit(1)
	}

	_, err := os.Stat(depsFilepath)
	if err != nil {
		fmt.Println("Could not find file at path: %w", depsFilepath)
		os.Exit(1)
	}

	customValues := make([]string, 0, 50)

	if len(os.Args) == 4 {
		customValues = append(customValues, strings.Split(os.Args[3], ",")...)
	}

	return bomFilePath, depsFilepath, &customValues
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Finished with error: %s\n", err)
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Println("This tool will remove <component> entries from a already created CycloneDX-Bom.xml file.")
	fmt.Println("The basic idea is to remove dependencies that are actually part of the dotnet runtime " +
		"environment.")
	fmt.Println("For example a System.Buffers/4.5.1 dependency will be remove because it will use " +
		"the System.Buffers.dll from the installed runtime.")
	fmt.Println("A custom list of components to remove can also be used.")
	fmt.Println("")
	fmt.Println("Parameter: ")
	fmt.Println("The file path to the bom.xml is mandatory.")
	fmt.Println("The file path to the {project name}.deps.json file is mandatory. It is needed to identify the runtime dependencies.")
	fmt.Println("To remove a custom list of components add a comma separated string of 'Name/Version' values as last argument")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Example: ")
	fmt.Println("	./dotnetbomcleaner {xml-filepath} {deps.json-filepath} {custom component string} ")
	fmt.Println("")
	fmt.Println("	./dotnetbomcleaner bom.xml myProject.deps.json XSerializer/0.4.2,Ardalis.Result/3.7.100.14")
	fmt.Println("----------------------------------------------------------------------------------")
}
