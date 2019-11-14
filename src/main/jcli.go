package main

import (
	jcli "Jcli/src/lib"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var conf jcli.JcliConfig
	params := jcli.GetParameters()

	//either generate the config or load it from a file
	if params.Generate {
		clConf, path := getConfigFromCommandLine()
		conf = clConf
		jcli.GenerateConfigFile(&conf, path)
	} else {
		conf = jcli.GetConfigFromFile(params.Config) //pass in .Config. if user didn't supply this param then ~/.jcli is assumed
		//validate config
		if conf.IsValid() {
			//validate passed in params
			if params.HasMinParams() {
				jcli.StartJenkinsCommands(&conf, &params) //where all the jenkins commands are executed from
			} else {
				fmt.Printf("Error: you must at least pass in the -job or -jobs parameters to identify the jenkins job.\nParams Defaults:\n%+v\n", params)
				os.Exit(1)
			}
		} else {
			fmt.Printf("Error: the config file is invalid. run \"./jcli -generate\" to create a valid one.\nCurrent Config:\n %+v\n", conf)
			os.Exit(1)
		}
	}
}

// getConfigFromCommandLine - get config input from user
func getConfigFromCommandLine() (jcli.JcliConfig, string) {
	conf := jcli.NewJcliConfig()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("------------------------Jenkins Config-----------------------------")

	conf.Username = jcli.ReadLine("Username: ", reader)
	conf.Apikey = jcli.ReadLine("Apikey: ", reader)
	conf.JenkinsUrl = jcli.ReadLine("JenkinsUrl: ", reader)
	filePath := jcli.ReadLine("Config File Path (leave blank if you want ~/.jcli): ", reader)

	fmt.Println("--------------------------End Config-------------------------------")

	return conf, filePath //return user built config and file path to store the config at
}
