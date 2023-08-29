package user_repository

import (
	"context"
	"database/sql"
	"fmt"
	user "go-football/Domain/User"
	"log"
)

var (
	TableName = "user"
)

type userRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (rcv *userRepository) FindAll() ([]*user.User, error) {
	rows, err := rcv.db.Query(fmt.Sprintf("SELECT * FROM `%s`", TableName))

	if err != nil {
		fmt.Printf("FindAll repository %+v\n", err)
		return nil, err
	}

	defer rows.Close()

	return rcv.rowsToModel(rows)
}

func (rcv *userRepository) Add(item *user.User) (int64, error) {
	query := fmt.Sprintf("INSERT INTO `%s` (`name`) VALUES (?)", TableName)
	insertResult, err := rcv.db.ExecContext(context.Background(), query, item.Name)
	if err != nil {
		log.Fatalf("impossible insert: %s", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}

	return id, nil
}

func (rcv *userRepository) rowsToModel(rows *sql.Rows) ([]*user.User, error) {
	items := make([]*user.User, 0)

	for rows.Next() {
		item := new(user.User)

		if err := rows.Scan(
			&item.Id,
			&item.Name,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}
