package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=user dbname=alura_loja password=1234 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// Encapsulando os arquivos da pasta de templates
var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	// db := conectaComBancoDeDados()
	// defer db.Close()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// produtos := []Produto{
	// 	{Nome: "Camiseta", Descricao: "Azul", Preco: 39, Quantidade: 5},
	// 	{"Tênis", "Adidas", 390, 4},
	// 	{"Notebook", "Dell", 3900, 40},
	// }

	db := conectaComBancoDeDados()
	selectDeTodosOsProdutos, err := db.Query("select * from tb_produtos")
	if err != nil {
		panic(err.Error())
	}

	prod := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &quantidade, &preco)
		if err != nil {
			panic(err.Error())
		}

		prod.Nome = nome
		prod.Descricao = descricao
		prod.Quantidade = quantidade
		prod.Preco = preco

		produtos = append(produtos, prod)

	}

	templates.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
