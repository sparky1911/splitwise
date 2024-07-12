package models

import (
	"log"
	"splitwise/db"
	"time"
)

type Group struct {
	ID        int64
	Name      string `binding:"required"`
	CreatedAt time.Time
}

func (g *Group) Create() error {
	query := `INSERT INTO groups(name,created_at) VALUES($1,$2) `
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(g.Name, g.CreatedAt)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	g.ID = id
	return nil

}
