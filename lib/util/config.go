package util

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	DataBase `yaml:"database"`
	Admin `yaml:"admin"`
}

type DataBase struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	User     string `yaml:"user"`
	Name     string `yaml:"name"`
}

type Admin struct {
	PortAdmin string `yaml:"port"`
}

//Carrega as configurações
//Retorna fatal caso nao encontre o arquivo padrao config.yml
func ReadConfig() Config {

	FileConfig, err := ioutil.ReadFile("./config.yml")
	VerifyErr(err)

	config := Config{}

	yaml.Unmarshal(FileConfig, &config)

	return config
}

//Create variaveis de ambiente
func Env(config Config) {

	os.Setenv("HOST", config.Host)
	os.Setenv("PORT", config.Port)
	os.Setenv("PASSWORD", config.Password)
	os.Setenv("USER", config.User)
	os.Setenv("NAME", config.Name)

	os.Setenv("PORT_ADMIN", config.Admin.PortAdmin)

}