package generator

import (
	"fmt"
	"os"
	"text/template"
)

func DockerConfigFiles(goVersion, projectName string) {
	files := []string{
		"dockerfile",
		"docker-compose.yml",
		"docker-compose.override.yml.example",
	}

	data := map[string]string{
		"GoVersion": goVersion,
		"ProjectName": projectName,
	}

	for _, file := range files {
		generateFile(file, data)
	}
}

func generateFile(fileName string, data map[string]string) {
	fmt.Printf("Generating %s\n", fileName)
	tmpl, err := template.ParseFiles(
		fmt.Sprintf("internal/generator/templates/%s.tmpl", fileName),
	)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(fmt.Sprintf("%s/%s", data["ProjectName"], fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = tmpl.Execute(file, data);
	if err != nil {
		panic(err)
	}
}
