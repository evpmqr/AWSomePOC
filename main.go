package main

import (
	"fmt"

	"github.com/awsome/aws/resources"
	"github.com/awsome/file"
)

func main() {
	fileClient := file.NewClient()
	builder := resources.NewBuilder(fileClient)

	reader, err := fileClient.GetFile("testing.yml")
	if err != nil {
		fmt.Errorf("error getting file reader: %s", err)
	}

	template, err := builder.Unmarshal(reader)
	if err != nil {
		fmt.Errorf("error unmarshaling config: %s", err)
	}

	err = builder.ToTemplateFile(template)
	if err != nil {
		fmt.Errorf("error building template file: %s", err)
	}
}
