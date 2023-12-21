package repositories

import (
	"database/sql"
	"fmt"

	"git.sula.io/solevis/poultracker/internal/database"
	"git.sula.io/solevis/poultracker/internal/models"
)

type CollectionStore struct{}

// FetchAll retrieves all collections in DB
func (store *CollectionStore) FetchAll() ([]models.Collection, error) {
	db := database.GetDB()
	var collections []models.Collection

	query := "SELECT id, DATE(laid_date) AS laid_date, number FROM collection ORDER BY DATE(laid_date)"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("FetchAllCollections: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var collection models.Collection
		if err := rows.Scan(&collection.ID, &collection.LaidDate, &collection.Number); err != nil {
			return nil, fmt.Errorf("FetchAllCollections: %v", err)
		}

		collections = append(collections, collection)
	}

	return collections, nil
}

// FetchOne retrieves the collection with the specified laid date.
func (store *CollectionStore) FetchOne(laidDate string) (models.Collection, error) {
	db := database.GetDB()

	// An collection to hold data from the returned row.
	var collection models.Collection

	query := "SELECT id, DATE(laid_date) AS laid_date, number FROM collection WHERE laid_date = $1"
	row := db.QueryRow(query, laidDate)
	if err := row.Scan(&collection.ID, &collection.LaidDate, &collection.Number); err != nil {
		if err == sql.ErrNoRows {
			return collection, fmt.Errorf("FetchOneCollection %s: no such collection", laidDate)
		}

		return collection, fmt.Errorf("FetchOneCollection %s: %v", laidDate, err)
	}

	return collection, nil
}

// Create adds new collection in DB
func (store *CollectionStore) Create(collection models.Collection) (models.Collection, error) {
	db := database.GetDB()

	query := "INSERT INTO collection (laid_date, number) VALUES (DATE($1), $2) RETURNING id"
	err := db.QueryRow(query, collection.LaidDate, collection.Number).Scan(&collection.ID)
	if err != nil {
		return collection, fmt.Errorf("SaveCollection: %v", err)
	}

	return collection, nil
}

// Delete delete an collection from DB
func (store *CollectionStore) Delete(id int) (int64, error) {
	db := database.GetDB()

	query := "DELETE FROM collection WHERE id = $1 RETURNING id"
	result, err := db.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("DeleteCollection: %v", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("DeleteCollection: %v", err)
	}

	return affectedRows, nil
}

// Update updates existing collection in DB
func (store *CollectionStore) Update(collection models.Collection) (int64, error) {
	db := database.GetDB()

	query := "UPDATE collection SET laid_date = $1, number = $2 WHERE id = $3"
	result, err := db.Exec(query, collection.LaidDate, collection.Number, collection.ID)
	if err != nil {
		return 0, fmt.Errorf("UpdateCollection: %v", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("UpdateCollection: %v", err)
	}

	return affectedRows, nil
}
