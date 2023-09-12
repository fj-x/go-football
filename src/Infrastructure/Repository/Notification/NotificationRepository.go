package notification_repository

import (
	"context"
	"database/sql"
	"fmt"
	notification "go-football/src/Domain/Notification/Model"
	infrastructure "go-football/src/Infrastructure"
	"log"
)

type NotificationRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (rcv *NotificationRepository) FindBySubscription(subscriptionId int32) ([]*notification.Notification, error) {
	rows, err := rcv.db.Query(fmt.Sprintf("SELECT * FROM `%s` WHERE `subscriptionId`= ?", infrastructure.NotificationTable), subscriptionId)
	if err != nil {
		fmt.Printf("FindByUser repository %+v\n", err)
		return nil, err
	}
	defer rows.Close()

	return rcv.rowsToModel(rows)
}

func (rcv *NotificationRepository) Add(item *notification.Notification) (*notification.Notification, error) {
	query := fmt.Sprintf("INSERT INTO `%s` (`subscriptionId`, `type`) VALUES (?, ?)", infrastructure.NotificationTable)
	insertResult, err := rcv.db.ExecContext(context.Background(), query, item.SubscriptionId, item.Type)
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

func (rcv *NotificationRepository) rowsToModel(rows *sql.Rows) ([]*notification.Notification, error) {
	items := make([]*notification.Notification, 0)

	for rows.Next() {
		item := new(notification.Notification)

		if err := rows.Scan(
			&item.Id,
			&item.SubscriptionId,
			&item.Type,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}
