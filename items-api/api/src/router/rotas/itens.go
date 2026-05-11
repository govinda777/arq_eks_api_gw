package rotas

import (
	"github.com/marcosouzatech/items-api/api/src/controllers"
	"net/http"
)

var rotasItens = []Rota{
	{
		URI:                "/",
		Metodo:             http.MethodGet,
		Funcao:             controllers.HealthCheck,
		RequerAutenticacao: false,
	},
	{
		URI:                "/items",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarItem,
		RequerAutenticacao: false,
	},
	{
		URI:                "/items",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarItens,
		RequerAutenticacao: false,
	},
	{
		URI:                "/items/{itemId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarItem,
		RequerAutenticacao: false,
	},
	{
		URI:                "/items/{itemId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarItem,
		RequerAutenticacao: false,
	},
	{
		URI:                "/items/{itemId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarItem,
		RequerAutenticacao: false,
	},
}
