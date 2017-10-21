package controller

import(
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

func Teste(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w,"Teste de codigo")
}