package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lucashenriqueteixeira/list_go/lib/controller"
	"github.com/lucashenriqueteixeira/list_go/lib/server"
	"net/http"

)

func Routers() {
	Router := httprouter.New()
	Router.GET("/teste",controller.Teste)
	Router.GET("/cadastro",controller.HandleCadastro)

	//Permite acesso a pasta /public/
	Router.ServeFiles("/public/*filepath",http.Dir("public/"))

	server.InitServer(":80",Router)

}