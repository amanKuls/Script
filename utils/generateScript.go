package utils

import _ "embed" //stores file contents to string

//go:embed templates/copy-commands.txt
var templateCopyCommands string

//go:embed templates/shell-script.txt
var templateShellScript string

func GenerateSQLAndShellFiles(outputDirectory string) {
	CreateFileWithContents(outputDirectory+"/copy_commands.sql", []string{templateCopyCommands})
	CreateFileWithContents(outputDirectory+"/copy_commands.sh", []string{templateShellScript})

}
