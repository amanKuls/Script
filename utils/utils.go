package utils

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"log"
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

func CreateFileWithContents(outputFile string, fileContents []string) {
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	datawriter := bufio.NewWriter(file)
	for _, data := range fileContents {
		_, _ = datawriter.WriteString(data)
	}
	datawriter.Flush()
	file.Close()
	// err := ioutil.WriteFile(outputFile, ConvertStringArrayToByteArray(fileContents), 0755)
	// if err != nil {
	// 	panic(err)
	// }
}

func ConvertStringArrayToByteArray(input []string) []byte {
	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(input)
	bs := buf.Bytes()
	return bs
}
