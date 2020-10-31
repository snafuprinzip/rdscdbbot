package config

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

var (
	Token     string
	CmdPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `yaml:"Token"`
	CmdPrefix string `yaml:"CommandPrefix"`
}

func ReadConfig() error {
	fmt.Println("reading config file...")

	file, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	CmdPrefix = config.CmdPrefix

	return nil
}
