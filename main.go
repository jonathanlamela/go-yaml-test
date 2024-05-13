package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type MailConfig struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
}

type Config struct {
	Mail MailConfig `yaml:"mail"`
}

func main() {

	var config Config
	var mailConfig MailConfig

	fileContent, err := os.ReadFile("config.yaml")

	if err != nil {
		fmt.Println("Impossibile leggere il file")
	}

	err = yaml.Unmarshal(fileContent, &config)

	if err != nil {
		fmt.Println("Impossibile eseguire il parsing del file yaml")
	}

	mailConfig = config.Mail

	fmt.Println(mailConfig.Username)
	fmt.Println(mailConfig.Password)
	fmt.Println(mailConfig.Host)
	fmt.Println(mailConfig.From)

}
