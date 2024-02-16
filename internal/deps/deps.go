package deps

import (
	"dotnetbomcleaner/internal/types"
	"encoding/json"
	"fmt"
	"os"
)

func GetRuntimeDependencies(filepath string) ([]string, error) {

	content, errFile := os.ReadFile(filepath)
	if errFile != nil {
		return nil, fmt.Errorf("GetRuntimeDependencies: could not read file: %w", errFile)
	}

	var p types.DotnetDeps
	errJson := json.Unmarshal(content, &p)

	if errJson != nil {
		return nil, fmt.Errorf("GetRuntimeDependencies: could not parse content to dotnet deps struct: %w", errJson)
	}

	dependencies := listOfRuntimeDependencies(&p)

	return dependencies, nil

}

func listOfRuntimeDependencies(p *types.DotnetDeps) []string {
	runtimeDependencies := make([]string, 0, 20)

	for _, item := range p.Targets {
		for key, subItem := range item {
			if subItem.Runtime == nil {
				runtimeDependencies = append(runtimeDependencies, key)
			}
		}
	}
	return runtimeDependencies

}
