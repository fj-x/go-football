package repository

import (
	"database/sql"
	"fmt"
	team "go-football/Domain/Team"
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
	rows, err := rcv.db.Query(fmt.Sprintf("SELECT * FROM `%s` ORDER BY `id` LIMIT ?", TableName))

	if err != nil {
		fmt.Printf("FindAll repository %+v\n", err)
		return nil, err
	}

	defer rows.Close()

	return rcv.rowsToModel(rows)
}

func (rcv *teamRepository) Add(item *team.Team) (int32, error) {
	stmt, stmtErr := rcv.db.Prepare(fmt.Sprintf("INSERT INTO `%s` SET `name`=?, `remoteId`=?", TableName))

	if stmtErr != nil {
		return item.Id, stmtErr
	}

	defer stmt.Close()

	_, execErr := stmt.Exec(item.Name, item.Id)

	if execErr != nil {
		return item.Id, execErr
	}

	return item.Id, nil
}

func (rcv *teamRepository) rowsToModel(rows *sql.Rows) ([]*team.Team, error) {
	items := make([]*team.Team, 0)

	for rows.Next() {
		item := new(team.Team)

		if err := rows.Scan(
			&item.Id,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}
