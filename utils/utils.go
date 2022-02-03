package utils

import (
	"io/ioutil"
	"os"
	"text/template"
)

func GenerateRenderedFile(
	outputFile string,
	templateContents string,
	templateInput interface{},
	generatedData string,
) {
	temp := template.New("template")
	funcs := make(map[string]interface{})
	funcs["getData"] = func() string {
		return generatedData
	}
	temp, err := temp.Funcs(funcs).Parse(templateContents)
	if err != nil {
		return
	}
	temp, err = temp.New(outputFile).Parse(templateContents)
	if err != nil {
		return
	}
	f, err := os.Create(outputFile)
	if err != nil {
		return
	}
	err = temp.Execute(f, templateInput)
	f.Close()
	if err != nil {
		return
	}
}

func CreateFileWithContents(outputFile string, fileContents string) {
	err := ioutil.WriteFile(outputFile, []byte(fileContents), 0755)
	if err != nil {
		panic(err)
	}
}
