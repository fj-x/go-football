package subscription_repository

import (
	"context"
	"database/sql"
	"fmt"
	subscription "go-football/Domain/Subscription"
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

func (rcv *subscriptionRepository) FindByUser(userId int32) ([]*subscription.Subscription, error) {
	rows, err := rcv.db.Query(fmt.Sprintf("SELECT * FROM `%s` WHERE `userId`= ?", TableName), userId)
	if err != nil {
		fmt.Printf("FindByUser repository %+v\n", err)
		return nil, err
	}
	defer rows.Close()

	return rcv.rowsToModel(rows)
}

func (rcv *subscriptionRepository) Add(item *subscription.Subscription) (int64, error) {
	query := fmt.Sprintf("INSERT INTO `%s` (`userId`, `teamId`) VALUES (?, ?)", TableName)
	insertResult, err := rcv.db.ExecContext(context.Background(), query, item.UserId, item.TeamId)
	if err != nil {
		log.Fatalf("impossible insert: %s", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}

	return id, nil
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
