package controllers

import (
	"github.com/marcosouzatech/items-api/api/src/database"
	"github.com/marcosouzatech/items-api/api/src/models"
	"github.com/marcosouzatech/items-api/api/src/repositories"
	"github.com/marcosouzatech/items-api/api/src/responses"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateItem inserts a user into the database
func CreateItem(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var item models.Item
	if err = json.Unmarshal(corpoRequest, &item); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = item.Prepare("registration"); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeItems(db)
	item.ID, err = repositorio.Create(item)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, item)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		// Registrar o evento GET no stdout
		w.Write([]byte("Health Check OK!"))
		log.Println("INFO: Requisição processada com sucesso no endpoint /")
	}
}
func SearchItems(w http.ResponseWriter, r *http.Request) {
	ProductOuName := strings.ToLower(r.URL.Query().Get("item"))
	log.Println("INFO: Query the list of /Items")

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		log.Println("WARNING: Failed to connect to database de dados")
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeItems(db)
	items, err := repositorio.Search(ProductOuName)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		log.Println("WARNING: Failed to create new item")
		return
	}
	responses.JSON(w, http.StatusOK, items)

}

// Search um item salvo into the database
func SearchItem(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	itemID, err := strconv.ParseUint(parametros["itemId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeItems(db)
	item, err := repositorio.SearchPorID(itemID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	itemID, err := strconv.ParseUint(parametros["itemId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var item models.Item
	if err = json.Unmarshal(corpoRequisicao, &item); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = item.Prepare("edit"); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeItems(db)
	if err = repositorio.Update(itemID, item); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	itemID, err := strconv.ParseUint(parametros["itemId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeItems(db)
	if err = repositorio.Delete(itemID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
