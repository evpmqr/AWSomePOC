package resources

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"text/template"
	template2 "text/template"

	"github.com/awsome/file"
	"gopkg.in/yaml.v2"
)

// Template is struct representation of config
type Template struct {
	Name      string            `yaml:"name"`
	Namespace string            `yaml:"namespace"`
	Params    map[string]string `yaml:"params,omitempty"`
	Functions []Lambda          `yaml:"functions,omitempty"`
}

// Policy is the struct representation of the Policy used literally everywhere
type Policy struct {
	Action    []string `yaml:"action"`
	Effect    string   `yaml:"effect"`
	Resources []string `yaml:"resources"`
}

// Builder is a method accessor for accessing Tempalte methods
type Builder struct {
	fileClient file.Client
}

// NewBuilder creates a new instance of Builder
func NewBuilder(fileClient file.Client) Builder {
	return Builder{fileClient: fileClient}
}

// Unmarshal turns an io.Reader into a Template Struct
func (b *Builder) Unmarshal(templateReader io.Reader) (Template, error) {
	var template Template

	templateBytes, err := ioutil.ReadAll(templateReader)
	if err != nil {
		return template, err
	}

	err = yaml.Unmarshal(templateBytes, &template)
	return template, err
}

// ToTemplateFile writes stuff to a file
func (b *Builder) ToTemplateFile(template Template) error {
	filename := fmt.Sprintf("%s.yaml", template.Name)
	DefaultTemplateHeader := "AWSTemplateFormateVersion: 2010-09-09\nTransform: AWS::Serverless-2016-10-31\n\n"
	b.fileClient.WriteToFile(filename, DefaultTemplateHeader)

	// Write Params

	// Write Resources
	b.fileClient.WriteToFile(filename, "Resources:\n")
	// If template has functions, do function stuff
	if len(template.Functions) > 0 {
		template = b.formatFunctions(template)
		// Create function role
		for _, function := range template.Functions {
			err := b.createFunctionRole(function, filename)
			if err != nil {
				fmt.Errorf("error could not create function role: %s", err)
				return err
			}

			// Create function
			err = b.createFunction(function, filename)
			if err != nil {
				fmt.Errorf("error could not create function: %s", err)
				return err
			}
		}
	}
	return nil
}

// Sets the namespace of the function to be templateNamespace-functionName
func (b *Builder) formatFunctions(template Template) Template {
	for index, function := range template.Functions {
		template.Functions[index].Name = fmt.Sprintf("%s-%s", template.Namespace, function.Name)
	}
	return template
}

func (b *Builder) createFunctionRole(function Lambda, filename string) error {
	var nrParsed *template2.Template
	nrParsed = template2.Must(template.New("").Parse(`
  {{ .ResourceName}}Role:
    Type: AWS::IAM::Role
    Properties:
      RoleName: {{ .Name}}-role
      AssumeRolePolicyDocument:
        Version: 2017-10-17
          Statement:
          - Action: sts:AssumeRole
            Effect: Allow
            Principal:
            Service:
            - lambda.amazonaws.com
        ManagedPolicyArns:
        - !Sub someDefaultPolicy{Env}
        - someDefaultPolicy
        Policies:
        - PolicyName: some-policy
          PolicyDocument:
          Version: 2012-10-17
          {{range .Policies}}- Effect: {{ .Effect}}
            Action:
            {{range .Action}}- {{ .}}
            {{end -}}
            Resources:
            {{range .Resources}}- {{ .}}
            {{end -}}
          {{end -}}`))
	var sb strings.Builder
	err := nrParsed.Execute(&sb, function)
	if err != nil {
		return err
	}
	err = b.fileClient.WriteToFile(filename, sb.String())
	return err
}

func (b *Builder) createFunction(function Lambda, filename string) error {
	var nrParsed *template2.Template
	nrParsed = template2.Must(template.New("").Parse(`
  {{ .ResourceName}}:
    Type: AWS::Serverless::Function
    Properties:
      CodeUri: {{ .CodeURI}}
      Description: {{ .Description}}
      FunctionName: {{ .Name}}
      Handler: {{ .Handler}}
      Role: !GettAtt {{ .ResourceName}}Role.Arn
      Runtime: {{ .Runtime}}
      Timeout: {{ .Timeout}}
      {{if .Environment}}Environment:
        Variables:
          {{range $key, $value := .Environment}}{{$key}}: {{$value}}
          {{end -}}
      {{end -}}
      {{if .Event}}Events:
        {{ .Event.EventName}}
          Type: {{ .Event.Type}}
          Properties:
            {{if eq .Event.Type "Api"}}Path: {{ .Event.Properties.Path}}
            Method: {{ .Event.Properties.Method}}
            RestApiId: {{ .Event.Properties.RestAPIID}}
            {{end -}}
      {{end -}}
`))

	var sb strings.Builder
	err := nrParsed.Execute(&sb, function)
	if err != nil {
		return err
	}
	err = b.fileClient.WriteToFile(filename, sb.String())
	return err
}
