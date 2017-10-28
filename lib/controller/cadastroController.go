package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lucashenriqueteixeira/list_go/lib/util"
	"github.com/lucashenriqueteixeira/list_go/lib/model"
	"net/http"
	"text/template"
)


func HandleCadastro(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	tpl, err := template.ParseFiles("view/cadastro.html", "view/base.html")
	util.VerifyErr(err)

	err = tpl.ExecuteTemplate(w, "base", nil)
	util.VerifyErr(err)

	r.ParseForm()
	
	err = model.Cadastra(r.Form)
	util.VerifyErr(err)

	tpl.Execute(w, r.Form)

	
}