package main

import (
	"github.com/spf13/viper"
	"github.com/spf13/pflag"
	"fmt"
)

var conf = viper.New()

func initConfig() {
	configPath := pflag.String("config", "", "Please input the config path")
	pflag.Parse()
	conf.SetConfigType("toml")
	if(*configPath == "") {
		conf.SetConfigName("config")
		conf.AddConfigPath(".")
	} else {
		conf.SetConfigFile(*configPath)
	}

	err := conf.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error conig file: %s \n", err))
	}

	file := conf.ConfigFileUsed()
	if(file != "") {
		fmt.Println("Use config file: ", file)
	}

}


