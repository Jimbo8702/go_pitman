package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type StructBuilder struct {
	StructName   	string
	OutFileName 	string
	OutputFolder	string
}

func NewStructBuilder(structName, outfile, outfolder string) *StructBuilder {
	return &StructBuilder{
		StructName: structName,
		OutFileName: outfile,
		OutputFolder: outfolder,
	}
}

func (b *StructBuilder) GenerateStructDefinition(data map[string]interface{}) string {
	var structDefBuilder strings.Builder

	structDefBuilder.WriteString(fmt.Sprintf("type %s struct {\n", b.StructName))

	for key, value := range data {
		fieldName := strings.Title(key)
		fieldType := b.generateFieldType(value)
		structDefBuilder.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"`\n", fieldName, fieldType, key))
	}

	structDefBuilder.WriteString("}\n\n")
	return structDefBuilder.String()
}

func (b *StructBuilder) generateFieldType(value interface{}) string {
	switch value.(type) {
	case float64:
		return "float64"
	case string:
		return "string"
	case bool:
		return "bool"
	case map[string]interface{}:
		return "map[string]interface{}"
	case []interface{}:
		return "[]interface{}"
	default:
		return "interface{}"
	}
}

func (b *StructBuilder) WriteGoFile(structDef string) error {
	err := os.MkdirAll(b.OutputFolder, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error checking directory: %s", err)
	}

	filePath := filepath.Join(b.OutputFolder, b.OutFileName)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %s", err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "package main\n\n")
	if err != nil {
		return fmt.Errorf("error writing package to go file: %s", err)
	}

	_, err = fmt.Fprintf(file, "%s", structDef)
	if err != nil {
		return fmt.Errorf("error writing package to go file: %s", err)
	}

	return nil
}
