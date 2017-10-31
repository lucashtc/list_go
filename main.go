package main

import (
	"github.com/lucashenriqueteixeira/list_go/lib/router"
	"github.com/lucashenriqueteixeira/list_go/lib/util"
)

func main() {

	//Carrega configurações e cria variaveis de ambiente
	Dados := util.ReadConfig()
	util.Env(Dados)


	//Nao coloque nada depois dessa funcão
	//Maldita nao deixa nada ser executado depois dela apenas concorrencia
	router.Routers()
}
