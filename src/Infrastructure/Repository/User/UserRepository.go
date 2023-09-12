package user_repository

import (
	"context"
	"database/sql"
	"fmt"
	user "go-football/src/Domain/User/Model"
	infrastructure "go-football/src/Infrastructure"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (rcv *UserRepository) FindAll() ([]*user.User, error) {
	rows, err := rcv.db.Query(fmt.Sprintf("SELECT * FROM `%s`", infrastructure.UserTable))

	if err != nil {
		fmt.Printf("FindAll repository %+v\n", err)
		return nil, err
	}

	defer rows.Close()

	return rcv.rowsToModel(rows)
}

func (rcv *UserRepository) IsUserExist(userId int32) (bool, error) {
	var exist bool
	err := rcv.db.QueryRow(fmt.Sprintf("SELECT exists (SELECT id FROM `%s` where `remoteId` = ?)", infrastructure.UserTable), userId).Scan(&exist)
	if err != nil {
		fmt.Printf("FindAll repository %+v\n", err)
		return false, err
	}

	fmt.Println(exist)
	return exist, nil
}

func (rcv *UserRepository) GetUser(remoteId int32) (*user.User, error) {
	user := new(user.User)

	row := rcv.db.QueryRow(fmt.Sprintf("SELECT id, name, remoteId FROM `%s` where `remoteId` = ?", infrastructure.UserTable), remoteId)

	if err := row.Scan(
		&user.Id,
		&user.Name,
		&user.RemoteId,
	); err != nil {
		return nil, err
	}

	return user, nil
}

func (rcv *UserRepository) Add(item *user.User) (*user.User, error) {
	query := fmt.Sprintf("INSERT INTO `%s` (`name`, `remoteId`) VALUES (?,?)", infrastructure.UserTable)
	insertResult, err := rcv.db.ExecContext(context.Background(), query, item.Name, item.RemoteId)
	if err != nil {
		log.Fatalf("impossible insert: %s", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}
	item.Id = int32(id)

	return item, nil
}

func (rcv *UserRepository) rowsToModel(rows *sql.Rows) ([]*user.User, error) {
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
