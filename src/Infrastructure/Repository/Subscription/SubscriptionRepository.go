package subscription_repository

import (
	"context"
	"database/sql"
	"fmt"
	subscription "go-football/src/Domain/Subscription"
	footballdataapi "go-football/src/Infrastructure/Service/footballDataApi"
	"log"
)

var (
	TableName = "subscription"
)

type subscriptionRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *subscriptionRepository {
	return &subscriptionRepository{db: db}
}

func (rcv *subscriptionRepository) FindAll() ([]*subscription.Subscription, error) {
	rows, err := rcv.db.Query(fmt.Sprintf("SELECT * FROM `%s`", TableName))
	if err != nil {
		fmt.Printf("FindByUser repository %+v\n", err)
		return nil, err
	}
	defer rows.Close()

	return rcv.rowsToModel(rows)
}

func (rcv *subscriptionRepository) FindUnqueSubscribedTeams() ([]int32, error) {
	rows, err := rcv.db.Query(fmt.Sprintf("SELECT DISTINCT t.`remoteId` FROM `%s` s JOIN `team` t ON s.teamId = t.id ", TableName))
	if err != nil {
		fmt.Printf("FindByUser repository %+v\n", err)
		return nil, err
	}
	defer rows.Close()

	var values []int32
	for rows.Next() {
		var teamId int32

		// read the row on the table, and assign them to the variable declared above
		err := rows.Scan(&teamId)
		if err != nil {
			return nil, err
		}

		// appending the row data to the slice
		values = append(values, teamId)
	}

	return values, nil
}

func (rcv *subscriptionRepository) FindMatchSubscribers(match footballdataapi.Match) ([]int32, error) {
	rows, err := rcv.db.Query(fmt.Sprintf("SELECT DISTINCT `userId` FROM `%s` WHERE `teamId` IN (SELECT `id` FROM `team` WHERE `remoteId` IN (?, ?))", TableName), match.HomeTeam.Id, match.AwayTeam.Id)
	if err != nil {
		fmt.Printf("FindByUser repository %+v\n", err)
		return nil, err
	}
	defer rows.Close()

	var values []int32
	for rows.Next() {
		var teamId int32

		// read the row on the table, and assign them to the variable declared above
		err := rows.Scan(&teamId)
		if err != nil {
			return nil, err
		}

		// appending the row data to the slice
		values = append(values, teamId)
	}

	return values, nil
}

func (rcv *subscriptionRepository) FindByUser(userId int32) ([]*subscription.Subscription, error) {
	rows, err := rcv.db.Query(fmt.Sprintf("SELECT * FROM `%s` WHERE `userId`= ?", TableName), userId)
	if err != nil {
		fmt.Printf("FindByUser repository %+v\n", err)
		return nil, err
	}
	defer rows.Close()

	return rcv.rowsToModel(rows)
}

func (rcv *subscriptionRepository) Add(item *subscription.Subscription) (*subscription.Subscription, error) {
	query := fmt.Sprintf("INSERT INTO `%s` (`userId`, `teamId`) VALUES (?, ?)", TableName)
	insertResult, err := rcv.db.ExecContext(context.Background(), query, item.UserId, item.TeamId)
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

func (rcv *subscriptionRepository) Delete(item *subscription.Subscription) error {
	stmt, err := rcv.db.Prepare(fmt.Sprintf("DELETE FROM `%s` WHERE `userId` = ? AND `teamId = ?", TableName))

	if err != nil {
		return err
	}

	if _, err := stmt.Exec(item.UserId, item.TeamId); err != nil {
		return err
	}

	defer stmt.Close()

	return nil
}

func (rcv *subscriptionRepository) rowsToModel(rows *sql.Rows) ([]*subscription.Subscription, error) {
	items := make([]*subscription.Subscription, 0)

	for rows.Next() {
		item := new(subscription.Subscription)

		if err := rows.Scan(
			&item.Id,
			&item.UserId,
			&item.TeamId,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}
