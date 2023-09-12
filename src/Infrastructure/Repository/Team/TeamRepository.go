package team_repository

import (
	"context"
	"database/sql"
	"fmt"
	team "go-football/src/Domain/Team/Model"
	infrastructure "go-football/src/Infrastructure"
	"log"
)

type TeamRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

func (rcv *TeamRepository) FindAll() ([]*team.Team, error) {
	rows, err := rcv.db.Query(fmt.Sprintf("SELECT * FROM `%s`", infrastructure.TeamTable))

	if err != nil {
		fmt.Printf("FindAll repository %+v\n", err)
		return nil, err
	}

	defer rows.Close()

	return rcv.rowsToModel(rows)
}

func (rcv *TeamRepository) FindUsersTeams(userId int32) ([]*team.Team, error) {
	query := "SELECT t.id, t.name, t.remoteId FROM `%s` t JOIN `%s` s ON t.id = s.teamId AND s.userId = ?"
	rows, err := rcv.db.Query(fmt.Sprintf(query, infrastructure.TeamTable, infrastructure.SubscriptionTable), userId)

	if err != nil {
		fmt.Printf("FindAll repository %+v\n", err)
		return nil, err
	}

	defer rows.Close()

	return rcv.rowsToModel(rows)
}

func (rcv *TeamRepository) Add(item *team.Team) (int64, error) {
	query := fmt.Sprintf("INSERT INTO `%s` (`name`, `remoteId`) VALUES (?, ?)", infrastructure.TeamTable)
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

func (rcv *TeamRepository) rowsToModel(rows *sql.Rows) ([]*team.Team, error) {
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
