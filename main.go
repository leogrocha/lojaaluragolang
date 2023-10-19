package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=1234 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// Encapsulando os arquivos da pasta de templates
var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	db := conectaComBancoDeDados()
	defer db.Close()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul", Preco: 39, Quantidade: 5},
		{"Tênis", "Adidas", 390, 4},
		{"Notebook", "Dell", 3900, 40},
	}

	templates.ExecuteTemplate(w, "Index", produtos)
}
