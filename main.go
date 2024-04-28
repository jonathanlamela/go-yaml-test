package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func main() {

	var config map[string]interface{}

	fileContent, err := os.ReadFile("config.yaml")

	if err != nil {
		fmt.Println("Impossibile leggere il file")
	}

	err = yaml.Unmarshal(fileContent, &config)

	if err != nil {
		fmt.Println("Impossibile eseguire il parsing del file yaml")
	}

	fmt.Println(config)

}
