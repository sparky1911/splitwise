package models

import (
	"database/sql"
	"log"
	"splitwise/db"
	"time"
)

type Group struct {
	ID        int64
	Name      string `binding:"required"`
	CreatedAt time.Time
}

type Membership struct {
	ID       int64
	UserID   int64 `binding:"required"`
	GroupID  int64 `binding:"required"`
	JoinedAt time.Time
}

// Create method for Group
func (g *Group) Create() error {
	query := `INSERT INTO groups(name, created_at) VALUES($1, $2)`
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

// Create method for Membership
func (m *Membership) Create() error {
	query := `INSERT INTO memberships(user_id, group_id, joined_at) VALUES($1, $2, $3)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(m.UserID, m.GroupID, m.JoinedAt)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	m.ID = id
	return nil
}

// FetchMemberships retrieves all memberships
func FetchMemberships() ([]Membership, error) {
	query := `SELECT id, user_id, group_id, joined_at FROM memberships`
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var memberships []Membership
	for rows.Next() {
		var m Membership
		if err := rows.Scan(&m.ID, &m.UserID, &m.GroupID, &m.JoinedAt); err != nil {
			return nil, err
		}
		memberships = append(memberships, m)
	}
	return memberships, nil
}

// Update method for Membership
func (m *Membership) Update() error {
	query := `UPDATE memberships SET user_id=$1, group_id=$2, joined_at=$3 WHERE id=$4`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(m.UserID, m.GroupID, m.JoinedAt, m.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete method for Membership
func (m *Membership) Delete() error {
	query := `DELETE FROM memberships WHERE id=$1`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(m.ID)
	if err != nil {
		return err
	}

	return nil
}

// FetchMembershipByID retrieves a single membership by ID
func FetchMembershipByID(id int64) (*Membership, error) {
	query := `SELECT id, user_id, group_id, joined_at FROM memberships WHERE id=$1`
	row := db.DB.QueryRow(query, id)

	var m Membership
	if err := row.Scan(&m.ID, &m.UserID, &m.GroupID, &m.JoinedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No result found
		}
		return nil, err
	}

	return &m, nil
}

// FetchMembershipsByGroupID retrieves memberships by group ID
func FetchMembershipsByGroupID(groupID int64) ([]Membership, error) {
	query := `SELECT id, user_id, group_id, joined_at FROM memberships WHERE group_id=$1`
	rows, err := db.DB.Query(query, groupID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var memberships []Membership
	for rows.Next() {
		var m Membership
		if err := rows.Scan(&m.ID, &m.UserID, &m.GroupID, &m.JoinedAt); err != nil {
			return nil, err
		}
		memberships = append(memberships, m)
	}
	return memberships, nil
}
