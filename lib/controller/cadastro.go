package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lucashenriqueteixeira/list_go/lib/util"
	"html/template"
	"net/http"
)

func ViewCadastro(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	tpl, err := template.ParseFiles("view/cadastro.html")

	util.VerifyErr(err)

	if r.Method == "GET" {

		tpl.Execute(w, nil)

	}

	// caso POST vai validar os dados e salvar em banco de dados

}