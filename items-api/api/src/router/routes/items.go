package routes

import (
	"github.com/marcosouzatech/items-api/api/src/controllers"
	"net/http"
)

var routesItems = []Route{
	{
		URI:                "/",
		Method:             http.MethodGet,
		Function:             controllers.HealthCheck,
		RequiresAuthentication: false,
	},
	{
		URI:                "/items",
		Method:             http.MethodPost,
		Function:             controllers.CreateItem,
		RequiresAuthentication: false,
	},
	{
		URI:                "/items",
		Method:             http.MethodGet,
		Function:             controllers.SearchItems,
		RequiresAuthentication: false,
	},
	{
		URI:                "/items/{itemId}",
		Method:             http.MethodGet,
		Function:             controllers.SearchItem,
		RequiresAuthentication: false,
	},
	{
		URI:                "/items/{itemId}",
		Method:             http.MethodPut,
		Function:             controllers.UpdateItem,
		RequiresAuthentication: false,
	},
	{
		URI:                "/items/{itemId}",
		Method:             http.MethodDelete,
		Function:             controllers.DeleteItem,
		RequiresAuthentication: false,
	},
}
