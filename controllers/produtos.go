package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/lojaaluragolang/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

const codigoRedirect = 301

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	templates.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat64, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade", err)
		}

		novoProduto := models.Produto{
			Nome:       nome,
			Descricao:  descricao,
			Preco:      precoConvertidoParaFloat64,
			Quantidade: quantidadeConvertidaParaInt,
		}

		models.CriarNovoProduto(novoProduto)
	}

	http.Redirect(w, r, "/", codigoRedirect)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletarProduto(idDoProduto)
	http.Redirect(w, r, "/", codigoRedirect)
}
