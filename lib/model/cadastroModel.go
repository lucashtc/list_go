package model

import (
	"errors"
	//"github.com/asaskevich/govalidator"
	"text/template"
	"net/url"
	"github.com/lucashenriqueteixeira/list_go/lib/db"
	
)

type Contato struct {
	Nome string
	Telefone []string
}

/*
* @param Dados map
* Recebe os dados digitado pelo usuario Filtra e salva no banco
*/
func Cadastra(Dados url.Values) (error) {
	
	Quant := len(Dados)
	if Quant <= 2 {

		c := Contato {
			Nome : Dados.Get("nome"),
			Telefone : Dados["telefone"],
		}

	
		//Gambiarra???? pode ser
		Result := FiltraDados(&c)
		Nome := Result.Nome
		Telefone := Result.Telefone
		db.CadastroSave(Nome,Telefone)
		//Fim da gambiarra?
		
		return nil
	}

	//Caso chegue mais de um parametro que pode ser alterado pela url facilmente.
	return errors.New("Ouve uma alteração proibida nos parametros do formulario")
}

func FiltraDados(c *Contato) (Contato) {

	C := *c
	C.Nome = template.HTMLEscapeString(C.Nome)
	
	//Pecorre os valores do slice e Escape cada um
	var DadosTelefone []string
	for Index, Value := range C.Telefone {
		DadosTelefone[Index] = template.HTMLEscapeString(Value)
	}

	C.Telefone = DadosTelefone
	

	return C
}