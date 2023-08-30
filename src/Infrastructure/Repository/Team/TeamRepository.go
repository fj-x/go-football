package team_repository

import (
	"context"
	"database/sql"
	"fmt"
	team "go-football/src/Domain/Team"
	"log"
)

var (
	TableName = "team"
)

type teamRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *teamRepository {
	return &teamRepository{db: db}
}

func (rcv *teamRepository) FindAll() ([]*team.Team, error) {
	rows, err := rcv.db.Query(fmt.Sprintf("SELECT * FROM `%s`", TableName))

	if err != nil {
		fmt.Printf("FindAll repository %+v\n", err)
		return nil, err
	}

	defer rows.Close()

	return rcv.rowsToModel(rows)
}

func (rcv *teamRepository) Add(item *team.Team) (int64, error) {
	query := fmt.Sprintf("INSERT INTO `%s` (`name`, `remoteId`) VALUES (?, ?)", TableName)
	insertResult, err := rcv.db.ExecContext(context.Background(), query, item.Name, item.RemoteId)
	if err != nil {
		log.Fatalf("impossible insert: %s", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}

	return id, nil
}

func (rcv *teamRepository) rowsToModel(rows *sql.Rows) ([]*team.Team, error) {
	items := make([]*team.Team, 0)

	for rows.Next() {
		item := new(team.Team)

		if err := rows.Scan(
			&item.Id,
			&item.Name,
			&item.RemoteId,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}
