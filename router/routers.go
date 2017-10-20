package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lucashenriqueteixeira/list_go/controller"
	"github.com/lucashenriqueteixeira/list_go/server"
)

func Routers() {
	Router := httprouter.New()
	Router.GET("/teste",controller.Teste)

	server.InitServer(":80",Router)

}