package model

import (
	"errors"
	//"github.com/asaskevich/govalidator"
	"text/template"
	"net/url"
	"github.com/lucashenriqueteixeira/list_go/lib/db"
	
)

/*
* @param Dados map
* Recebe os dados digitado pelo usuario Filtra e salva no banco
*/
func Cadastra(Dados url.Values) (error) {
	
	Quant := len(Dados)
	if Quant <= 2 {
		DadosFiltrados := make(map[string]interface{},Quant)
		DadosFiltrados["nome"] = template.HTMLEscapeString(Dados["nome"][0])
		DadosFiltrados["telefone"] = template.HTMLEscapeString(Dados["telefone"][0])

		db.CadastroSave(DadosFiltrados)

		return nil
	}

	//Caso chegue mais de um parametro que pode ser alterado pela url facilmente.
	return errors.New("Ouve uma alteração proibida nos parametros do formulario")
}