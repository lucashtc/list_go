package util

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	DataBase
	Admin
}

type DataBase struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	User     string `yaml:"user"`
	name     string `yaml:"name"`
}

type Admin struct {
	Port string `yaml:"port"`
}

//carrega arquivo de configuração
//Retorna fatal caso nao encontre o arquivo config.yml
func ReadConfig() {
	
	FileConfig, err := ioutil.ReadFile("./config.yml")
	VerifyErr(err)

	config := Config{}
	yaml.Unmarshal(FileConfig,&config)

	os.Setenv("HOST",config.Host)
	os.Setenv("PORT",config.Port)
	os.Setenv("PASSWORD",config.Password)
	os.Setenv("NAME",config.Password)
	os.Setenv("USER",config.User)

	os.Setenv("PORT_HTTP",config.Admin.Port)
	
}