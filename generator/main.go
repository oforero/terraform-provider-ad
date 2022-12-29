package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	//go:embed templates/*.tmpl
	templates embed.FS

	input      *string
	output     *string
	outputFile *os.File
	generate   *string
)

const (
	unset string = "<UNSET>"
)

func init() {
	input = flag.String("i", unset, "The input file name")
	output = flag.String("o", unset, "The output file name")
	generate = flag.String("generate", unset, "The type of code to generate datasources|converters")
}

type ParamsFieldsStruct struct {
	Name    string `json:"name"`
	ApiName string `json:"api_name"`
	Type    string `json:"type"`
}

type ParamsStruct struct {
	Name   string               `json:"name"`
	Fields []ParamsFieldsStruct `json:"fields"`
}

type TerraformArguments struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
	Computed    bool   `json:"computed,omitempty"`
	Optional    bool   `json:"optional,omitempty"`
	Required    bool   `json:"required,omitempty"`
}

type Terraform struct {
	Name      string               `json:"name"`
	Arguments []TerraformArguments `json:"arguments,omitempty"`
}

type ApiArguments struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
}

type ApiFunction struct {
	Name            string         `json:"name"`
	ResultIsPointer bool           `json:"isPointer"`
	Arguments       []ApiArguments `json:"arguments"`
	ApiParams       ParamsStruct   `json:"params"`
}

type Operation struct {
	ApiFunction             ApiFunction `json:"api_function"`
	ResultField             string      `json:"result_field"`
	ResultId                string      `json:"result_id"`
	Name                    string
	ApiToTerraformFunction  string `json:"api_to_terraform"`
	ResultWrapperFunction   string `json:"result_wrapper"`
	SchemaFunction          string `json:"schema_function"`
	SchemaFunctionArguments string `json:"schema_function_arguments"`
	ElementName             string
}

type Datasource struct {
	Name                    string               `json:"name"`
	Terraform               Terraform            `json:"terraform"`
	ElementName             string               `json:"element_name"`
	SchemaFunction          string               `json:"schema_function"`
	SchemaFunctionArguments string               `json:"schema_function_arguments"`
	CRUD                    map[string]Operation `json:"crud"`
}

type InputData struct {
	ApiAlias    string       `json:"api_alias"`
	ApiPackage  string       `json:"api_package"`
	Package     string       `json:"package"`
	DataSources []Datasource `json:"data_sources"`
	Resources   []Datasource `json:"resources"`
}

// Main function
func main() {
	flag.Parse()

	if *input == unset || *generate == unset {
		flag.CommandLine.Usage()
		os.Exit(0)
	}

	inputFile, err := ioutil.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	data := InputData{}
	err = json.Unmarshal([]byte(inputFile), &data)
	if err != nil {
		panic(err)
	}

	if *output == unset {
		outputFile = os.Stdout
	} else {
		file, err := os.OpenFile(*output, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		outputFile = file
	}

	tmplName := fmt.Sprintf("%s.go.tmpl", *generate)
	tmplFile := fmt.Sprintf("templates/%s", tmplName)
	caser := cases.Title(language.Und, cases.NoLower)
	tmpl, err := template.New(tmplName).Funcs(template.FuncMap{
		"isPointer": func(p bool) string {
			if p {
				return "*"
			} else {
				return ""
			}
		},
		"contains": func(k string, m map[string]interface{}) bool {
			_, ok := m[k]
			return ok
		},
		"default": func(d string, v string) string {
			if v != "" {
				return v
			} else {
				return d
			}
		},
		"ToLower": strings.ToLower,
		"toTitle": caser.String,
		"enrichOperation": func(op Operation, n string, sF string, sFA string, en string) Operation {
			op.Name = n
			op.SchemaFunction = sF
			op.SchemaFunctionArguments = sFA
			op.ElementName = en
			return op
		},
		"apiToTerraformType": func(t string) string {
			switch t {
			case "int":
				return "schema.TypeInt"
			case "string":
				return "schema.TypeString"
			case "boolean":
				return "schema.TypeBool"
			default:
				panic(fmt.Sprintf("Unsupported type: %s", t))
			}
		},
	}).ParseFS(templates, tmplFile)
	if err != nil {
		panic(err)
	}
	fmt.Println(tmpl.Name())

	if err := tmpl.Execute(outputFile, data); err != nil {
		panic(err)
	}
}
