package database

import (
	"database/sql"
)
import "github.com/google/uuid"

type Category struct {
	DB          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{DB: db}
}

func (c *Category) Create(name, description string) (*Category, error) {
	id := uuid.New().String()
	_, err := c.DB.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description)
	if err != nil {
		return nil, err
	}
	return &Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.DB.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	var Categories []Category
	for rows.Next() {
		var Category Category
		if err := rows.Scan(&Category.ID, &Category.Name, &Category.Description); err != nil {
			return nil, err
		}
		Categories = append(Categories, Category)
	}
	return Categories, nil
}

func (c *Category) FindByCourseID(id_course string) (*Category, error) {
	var id string
	var name string
	var description string
	row, err := c.DB.Query("SELECT ca.id,ca.name,ca.description FROM categories ca join course co on ca.id = co.category_id where co.id = $1;", id_course)
	if err != nil {

		return nil, err
	}
	if row.Next() {
		if err := row.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
	}
	return &Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}
