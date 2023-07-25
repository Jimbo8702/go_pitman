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

func (b * StructBuilder) GenerateStruct(schema map[string]interface{}) error {
	structDefinition := b.GenerateStructDefinition(schema, b.StructName)

	err := b.WriteGoFile(structDefinition)
	if err != nil {
		return err
	}

	return nil
}

func (b *StructBuilder) GenerateStructDefinition(data map[string]interface{}, structName string) string {
	var structDefBuilder strings.Builder

	structDefBuilder.WriteString(fmt.Sprintf("type %s struct {\n", structName))

	for key, value := range data {
		fieldName := strings.Title(key)
		fieldType := b.GenerateFieldType(value)
		structDefBuilder.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"`\n", fieldName, fieldType, key))
	}

	structDefBuilder.WriteString("}\n\n")
	return structDefBuilder.String()
}

func (b *StructBuilder) GenerateFieldType(value interface{}) string {
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
		return err
	}

	filePath := filepath.Join(b.OutputFolder, b.OutFileName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "package main\n\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(file, "%s", structDef)
	if err != nil {
		return err
	}

	return nil
}
