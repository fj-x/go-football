package subscription_repository

import (
	"context"
	"database/sql"
	"fmt"
	subscription "go-football/src/Domain/Subscription/Model"
	notification_repository "go-football/src/Infrastructure/Repository/Notification"
	"log"
)

var (
	TableName = "subscription"
)

type SubscriptionRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (rcv *SubscriptionRepository) FindAll() ([]*subscription.Subscription, error) {
	rows, err := rcv.db.Query(fmt.Sprintf("SELECT * FROM `%s`", TableName))
	if err != nil {
		fmt.Printf("FindByUser repository %+v\n", err)
		return nil, err
	}
	defer rows.Close()

	return rcv.rowsToModel(rows)
}

func (rcv *SubscriptionRepository) FindUnqueSubscribedTeams() ([]int32, error) {
	rows, err := rcv.db.Query(fmt.Sprintf("SELECT DISTINCT t.`remoteId` FROM `%s` s JOIN `%s` t ON s.teamId = t.id ", TableName, notification_repository.TableName))
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

func (rcv *SubscriptionRepository) FindMatchSubscribers(homeTeam, awayTeam int32) ([]int32, error) {
	rows, err := rcv.db.Query(fmt.Sprintf("SELECT DISTINCT `userId` FROM `%s` WHERE `teamId` IN (SELECT `id` FROM `%s` WHERE `remoteId` IN (?, ?))", TableName, notification_repository.TableName), homeTeam, awayTeam)
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

func (rcv *SubscriptionRepository) FindByUser(userId int32) ([]*subscription.Subscription, error) {
	rows, err := rcv.db.Query(fmt.Sprintf("SELECT * FROM `%s` WHERE `userId`= ?", TableName), userId)
	if err != nil {
		fmt.Printf("FindByUser repository %+v\n", err)
		return nil, err
	}
	defer rows.Close()

	return rcv.rowsToModel(rows)
}

func (rcv *SubscriptionRepository) Add(item *subscription.Subscription) (*subscription.Subscription, error) {
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

func (rcv *SubscriptionRepository) Delete(item *subscription.Subscription) error {
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

func (rcv *SubscriptionRepository) rowsToModel(rows *sql.Rows) ([]*subscription.Subscription, error) {
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
