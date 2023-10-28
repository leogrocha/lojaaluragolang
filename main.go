package main

import (
	"net/http"

	"github.com/lojaaluragolang/routes"
)

// Encapsulando os arquivos da pasta de templates

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
