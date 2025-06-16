package utils

import (
	"fmt"
	"os/exec"
)

func ValidateXMLWithXSD(xmlPath, xsdPath string) error {
	cmd := exec.Command("xmllint", "--noout", "--schema", xsdPath, xmlPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("XML validation failed: %s\nDetails: %v", string(output), err)
	}
	return nil
}
