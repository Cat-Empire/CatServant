package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ConfigStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

var ConfigInfo ConfigStruct

func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &ConfigInfo)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
