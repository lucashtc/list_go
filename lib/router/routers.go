package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lucashenriqueteixeira/list_go/lib/controller"
	"github.com/lucashenriqueteixeira/list_go/lib/server"
)

func Routers() {
	Router := httprouter.New()
	Router.GET("/teste",controller.Teste)
	Router.GET("/cadastro",controller.ViewCadastro)


	//Preciso tirar isso daqui.
	server.InitServer(":80",Router)

}