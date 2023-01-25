package model

import (
	"context"
	"database/sql"
	"time"
)

type HouseHold struct {
	ID        int       `json:"id"`
	Cost      string    `json:"cost"`
	CostName  string    `json:"costName"`
	User_Id   int    `json:"user_id"`
	IsSolidCost bool    `json:"isSolidCost"`
	ResiteredAt time.Time `json:"resistered_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetHouseHolds(ctx context.Context, db *sql.DB) ([]*HouseHold, error) {
	
	conn, err := db.Conn(ctx)
	if err != nil {
		return nil, err
	}

	var houseHolds []*HouseHold
	rows, err := conn.QueryContext(
		ctx,
		`
      SELECT 
        id,                               
        cost, 
        costName, 
        user_id, 
        issolidCost,
		resistered_at,
        created_at, 
        updated_at 
       from households`)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var houseHold HouseHold
		var createdDatetime string
		var updateDatetime string
		var resisteredDatetime string
		err := rows.Scan(
			&(houseHold.ID),
			&(houseHold.Cost),
			&(houseHold.CostName),
			&(houseHold.User_Id),
			&(houseHold.IsSolidCost),
			&resisteredDatetime,
			&createdDatetime,
			&updateDatetime)

		if err != nil {
			return nil, err
		}

		houseHold.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdDatetime) // 日時はこの日付じゃないといけない
		if err != nil {
			return nil, err
		}
		houseHold.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updateDatetime) // 日時はこの日付じゃないといけない
		if err != nil {
			return nil, err
		}
		houseHolds = append(houseHolds, &houseHold)
	}

	return houseHolds, nil
}

func CreateHouseHold(ctx context.Context, db *sql.DB, houseHold *HouseHold) (int64, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return 0, err
	}
	result, err := conn.ExecContext(
		ctx,
		`INSERT INTO households (
			cost, 
			costName, 
			user_id, 
			issolidCost,
			resistered_at,
			created_at, 
			updated_at
			) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		houseHold.Cost,
		houseHold.CostName,
		houseHold.User_Id,
		houseHold.IsSolidCost,
		time.Now(),
        time.Now(),
		time.Now(),
	)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertID, nil
}

func GetHouseHoldByID(ctx context.Context, db *sql.DB, houseHoldID int64) (*HouseHold, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	var houseHold HouseHold
	var createdDatetime string
	var updateDatetime string
	var resisteredDatetime string

	err = conn.QueryRowContext(
		ctx,
		`
      SELECT 
	   id,                               
	   cost, 
	   costName, 
	   user_id, 
	   issolidCost,
	   resistered_at,
	   created_at, 
	   updated_at 
       from households
       where id = ?`,
	   houseHoldID).Scan(
			&(houseHold.ID),
			&(houseHold.Cost),
			&(houseHold.CostName),
			&(houseHold.User_Id),
			&(houseHold.IsSolidCost),
			&resisteredDatetime,
			&createdDatetime,
			&updateDatetime)

	if err != nil {
		return nil, err
	}

	houseHold.ResiteredAt, err = time.Parse("2006-01-02 15:04:05", createdDatetime) // 日時はこの日付じゃないといけない
	if err != nil {
		return nil, err
	}

	houseHold.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdDatetime) // 日時はこの日付じゃないといけない
	if err != nil {
		return nil, err
	}
	houseHold.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updateDatetime) // 日時はこの日付じゃないといけない
	if err != nil {
		return nil, err
	}

	return &houseHold, nil
}