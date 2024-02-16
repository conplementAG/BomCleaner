package bom

import (
	"bytes"
	"fmt"
	cdx "github.com/CycloneDX/cyclonedx-go"
	"os"
	"slices"
)

func CleanupBom(filePath string, rd []string) error {

	content, errFile := os.ReadFile(filePath)
	if errFile != nil {
		return fmt.Errorf("CleanupBom: failed to read file: %w", errFile)
	}

	bom := new(cdx.BOM)
	decoder := cdx.NewBOMDecoder(bytes.NewReader(content), cdx.BOMFileFormatXML)
	errBom := decoder.Decode(bom)

	if errBom != nil {
		return fmt.Errorf("CleanupBom: failed to decode bom xml file: %w", errBom)
	}

	validDependencies := make([]cdx.Component, 0, 20)

	for _, item := range *bom.Components {

		if slices.Contains(rd, item.Name+"/"+item.Version) {
			continue
		}
		validDependencies = append(validDependencies, item)
	}

	bom.Components = &validDependencies

	f, errWrite := os.OpenFile("./cleanbom.xml", os.O_WRONLY|os.O_CREATE, 0600)
	if errWrite != nil {
		return fmt.Errorf("CleanupBom: failed to open output file: %w", errWrite)
	}

	errEncode := cdx.NewBOMEncoder(f, cdx.BOMFileFormatXML).SetPretty(true).Encode(bom)
	if errEncode != nil {
		return fmt.Errorf("CleanupBom: failed to write encoded output: %w", errEncode)
	}

	return f.Close()
}
