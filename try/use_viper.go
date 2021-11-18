package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	getConfig()
	{
	}
}

func getConfig() {
	viper.AutomaticEnv()
	viper.Set("sun.name", "sct")
	fmt.Println(viper.Get("sun.name"))
	//viper.SetConfigName("config") // name of config file (without extension)
	//viper.SetConfigType("env") // REQUIRED if the config file does not have the extension in the name
	////viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
	//viper.AddConfigPath("D:/workSpace/hello/configs/.env") // call multiple times to add many search paths
	////viper.AddConfigPath(".")                                   // optionally look for config in the working directory
	//err := viper.ReadInConfig() // Find and read the config file
	//
	//if err != nil { // Handle errors reading the config file
	//	panic(fmt.Errorf("Fatal error config file: %w \n", err))
	//}
}
