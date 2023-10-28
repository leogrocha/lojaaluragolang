package models

import "github.com/lojaaluragolang/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {

	db := db.ConectaComBancoDeDados()
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

		prod.Id = id
		prod.Nome = nome
		prod.Descricao = descricao
		prod.Quantidade = quantidade
		prod.Preco = preco

		produtos = append(produtos, prod)

	}

	defer db.Close()
	return produtos
}

func CriarNovoProduto(produto Produto) {
	db := db.ConectaComBancoDeDados()

	nome := produto.Nome
	descricao := produto.Descricao
	preco := produto.Preco
	quantidade := produto.Quantidade

	insereNovoProdutoNoBancoDeDados, err := db.Prepare("insert into tb_produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	_, err = insereNovoProdutoNoBancoDeDados.Exec(nome, descricao, preco, quantidade)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func DeletarProduto(id string) {
	db := db.ConectaComBancoDeDados()

	deletarProdutoNoBancoDeDados, err := db.Prepare("DELETE FROM tb_produtos tp WHERE produto_id = $1")
	if err != nil {
		panic(err.Error())
	}

	_, err = deletarProdutoNoBancoDeDados.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
