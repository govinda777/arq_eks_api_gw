package repositories

import (
	"github.com/marcosouzatech/items-api/api/src/models"
	"database/sql"
	"fmt"
)

type Items struct {
	db *sql.DB
}

// NovoRepositorioDeItems creates a repository of items
func NovoRepositorioDeItems(db *sql.DB) *Items {
	return &Items{db}
}

// Create inserts a item into the database
func (repositorio Items) Create(item models.Item) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"insert into items (Product, Name, Category, token) values(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(item.Product, item.Name, item.Category, item.Token)
	if err != nil {
		return 0, err
	}
	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(ultimoIDInserido), nil

}

// Search brings all items or brings based on the filter Product ou Name
func (repositorio Items) Search(ProductOuName string) ([]models.Item, error) {
	ProductOuName = fmt.Sprintf("%%%s%%", ProductOuName)

	linhas, err := repositorio.db.Query(
		"select id, Product, Name, Category, createdAt from items where Product LIKE ? or Name LIKE ?",
		ProductOuName, ProductOuName,
	)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	var items []models.Item

	for linhas.Next() {
		var item models.Item

		if err = linhas.Scan(
			&item.ID,
			&item.Product,
			&item.Name,
			&item.Category,
			&item.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// SearchPorID brings an item from the database
func (repositorio Items) SearchPorID(ID uint64) (models.Item, error) {
	linhas, err := repositorio.db.Query(
		"select id, Product, Name, Category, createdAt from items where id = ? ",
		ID,
	)
	if err != nil {
		return models.Item{}, err
	}
	defer linhas.Close()

	var item models.Item

	if linhas.Next() {
		if err = linhas.Scan(
			&item.ID,
			&item.Product,
			&item.Name,
			&item.Category,
			&item.CreatedAt,
		); err != nil {
			return models.Item{}, err
		}
	}
	return item, nil
}

// Update altera as informações de um item into the database
func (repositorio Items) Update(ID uint64, item models.Item) error {
	statement, err := repositorio.db.Prepare(
		"update items set Product = ?, Name = ?, Category = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(item.Product, item.Name, item.Category, ID); err != nil {
		return err
	}
	return nil
}

// Delete exclui as informações de um item into the database
func (repositorio Items) Delete(ID uint64) error {
	statement, err := repositorio.db.Prepare("delete from items where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}
