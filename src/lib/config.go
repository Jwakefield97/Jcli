package lib

import (
	"bytes"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"text/template"

	"github.com/BurntSushi/toml"
)

/*
	Structure of config file (toml)

	Username = "myUserName"
	Apikey = "23423lk4j234lk234l234234"
	JenkinsUrl = "https://ci.jenkins.com/"
*/

// DefaultConfigName - the default name of the config file
const DefaultConfigName = ".jcli"

// JcliConfig - the config that holds credentials for jenkins
type JcliConfig struct {
	Username   string
	Apikey     string
	JenkinsUrl string
	Path       string
}

// NewJcliConfig - constructor for JcliConfig
func NewJcliConfig() JcliConfig {
	config := JcliConfig{}
	config.Username = ""
	config.Apikey = ""
	config.JenkinsUrl = ""
	config.Path = filepath.Join(getHomePath(), DefaultConfigName) //create the default path of config (~/.jcli)
	return config
}

// IsValid - a function to determine if the config is valid
func (conf *JcliConfig) IsValid() bool {
	if IsBlank(conf.Username) || IsBlank(conf.Apikey) || IsBlank(conf.JenkinsUrl) { //if either of the fields are "" then the conf is invalid
		return false
	}
	return true
}

// getHomePath - a function to get the home path of the current user
func getHomePath() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

// GetConfigFromFile - a function to get the jenkins config from a file
func GetConfigFromFile(path string) JcliConfig {
	conf := NewJcliConfig()

	if IsBlank(path) { //if the path parameter was not passed in look for it in the home dir
		path = conf.Path
	}

	if _, err := toml.DecodeFile(path, &conf); err != nil {
		fmt.Printf("Either the config file at %s does not exist or the toml fields Username, and/or Apikey do not exist.", path)
		panic(err)
	}

	return conf
}

// GenerateConfigFile - a function to generate a jcli config from user input
func GenerateConfigFile(config *JcliConfig, path string) {
	//if a blank path is provided, the use default location ~/.jcli
	if IsBlank(path) {
		path = config.Path
	}

	//create template to fill out with user data
	template, err := template.New("config").Parse(ConfigFileTemplate)
	if err != nil {
		fmt.Printf("Error creating a template with config provided %+v", config)
		panic(err)
	}

	//fill out config template with user supplied info
	var parsedTemplate bytes.Buffer
	err = template.Execute(&parsedTemplate, config)
	if err != nil {
		fmt.Printf("Error creating a template with config provided %+v", config)
		panic(err)
	}

	//save template to config file
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error: unable to create file at %s", path)
		panic(err)
	}
	defer file.Close() //close file when function is done executing

	if _, err := file.Write(parsedTemplate.Bytes()); err != nil {
		fmt.Printf("Error: unable to save config to file at %s with config %+v", path, config)
		panic(err)
	}

}
