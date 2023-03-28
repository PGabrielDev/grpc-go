package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Course struct {
	DB          *sql.DB
	ID          string
	Name        string
	Description string
	Category_ID string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{
		DB: db,
	}
}

func (c *Course) Create(name, description, category_id string) (*Course, error) {
	id := uuid.New().String()
	_, err := c.DB.Exec("INSERT INTO course (id, name, description, category_id) VALUES ($1, $2, $3, $4)", id, name, description, category_id)
	if err != nil {
		return nil, err
	}
	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.DB.Query("SELECT * FROM course")
	if err != nil {
		return nil, err
	}
	var courses []Course
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.Category_ID); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	rows, err := c.DB.Query("SELECT co.id, co.name, co.description, co.category_id FROM course co join categories ca on co.category_id = ca.id where co.category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	var courses []Course
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.Category_ID); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}
