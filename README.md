#Pequena aplicação web que vai cadastrar contatos

#Ferramentas

Go/Golang
Mysql
dep


#instalação

go get "github.com/lucashenriqueteixeira/list_go"

Na pasta "github.com/lucashenriqueteixeira/list_go" encontra-se o arquivo sql

Nas pasta raiz o config.json deve ser configurado informações do banco de dados e nome da pasta raiz etc

dep ensure

cd src/github.com/lucashenriqueteixeira/list_go/lib

go build main.go