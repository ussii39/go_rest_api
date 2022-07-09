package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetUsers(ctx context.Context, db *sql.DB) ([]*User, error) {
	fmt.Println(ctx)
	conn, err := db.Conn(ctx)
	if err != nil {
		return nil, err
	}

	var Users []*User
	rows, err := conn.QueryContext(
		ctx,
		`
      SELECT 
        id,                               
        uuid, 
        name,
		email,
        created_at, 
        updated_at 
       from users`)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var User User
		var createdDatetime string
		var updateDatetime string
		err := rows.Scan(
			&(User.ID),
			&(User.UUID),
			&(User.Name),
			&(User.Email),
			&createdDatetime,
			&updateDatetime)

		if err != nil {
			return nil, err
		}

		User.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdDatetime) // 日時はこの日付じゃないといけない
		if err != nil {
			return nil, err
		}
		User.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updateDatetime) // 日時はこの日付じゃないといけない
		if err != nil {
			return nil, err
		}
		Users = append(Users, &User)
	}

	return Users, nil
}

func GetUser(ctx context.Context, db *sql.DB, UserUUID string) (*User, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	var User User
	var createdDatetime string
	var updateDatetime string

	err = conn.QueryRowContext(
		ctx,
		`
      SELECT 
        id,                               
        uuid, 
        name,
		email,
        created_at, 
        updated_at 
       from users
       where uuid = ?`,
		UserUUID).Scan(
		&(User.ID),
		&(User.UUID),
		&(User.Name),
		&(User.Email),
		&createdDatetime,
		&updateDatetime,
	)
	if err != nil {
		return nil, err
	}

	User.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdDatetime) // 日時はこの日付じゃないといけない
	if err != nil {
		return nil, err
	}
	User.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updateDatetime) // 日時はこの日付じゃないといけない
	if err != nil {
		return nil, err
	}

	return &User, nil
}

func CheckUserExist(ctx context.Context, db *sql.DB, UserUUID string) (bool, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return false, err
	}
	var fetchRecordCount int
	err = conn.QueryRowContext(
		ctx,
		`
      SELECT
       count(*)
       from users
       where uuid = ?`,
		UserUUID).Scan(
		&fetchRecordCount,
	)
	if err != nil {
		return false, err
	}

	if fetchRecordCount > 0 {
		return true, nil
	}
	return false, nil
}

func CreateUser(ctx context.Context, db *sql.DB, User *User) (int64, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return 0, err
	}
	result, err := conn.ExecContext(
		ctx,
		`INSERT INTO users (
                                   uuid, 
                                   name, 
                                   email, 
                                   created_at,
                                   updated_at
                                   ) VALUES (?, ?, ?, ?, ?) `,
		uuid.New(),
		User.Name,
		User.Email,
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

func GetUserByID(ctx context.Context, db *sql.DB, UserID int64) (*User, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	var User User
	var createdDatetime string
	var updateDatetime string

	err = conn.QueryRowContext(
		ctx,
		`
      SELECT 
        id,                               
        uuid, 
        name,
  	    email,
        created_at, 
        updated_at 
       from users
       where id = ?`,
		UserID).Scan(
		&(User.ID),
		&(User.UUID),
		&(User.Name),
		&(User.Email),
		&createdDatetime,
		&updateDatetime,
	)
	if err != nil {
		return nil, err
	}

	User.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdDatetime) // 日時はこの日付じゃないといけない
	if err != nil {
		return nil, err
	}
	User.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updateDatetime) // 日時はこの日付じゃないといけない
	if err != nil {
		return nil, err
	}

	return &User, nil
}

func UpdateUser(ctx context.Context, db *sql.DB, User *User, UserUUID string) error {
	conn, err := db.Conn(ctx)
	if err != nil {
		return err
	}
	_, err = conn.ExecContext(
		ctx,
		`UPDATE users 
                   SET name = ?, 
                       email = ?, 
                       updated_at = ?
                 WHERE uuid = ?`,
		User.Name,
		User.Email,
		time.Now(),
		UserUUID,
	)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(ctx context.Context, db *sql.DB, UserUUID string) error {
	conn, err := db.Conn(ctx)
	if err != nil {
		return err
	}
	_, err = conn.ExecContext(
		ctx,
		`DELETE FROM users 
                 WHERE uuid = ?`,
		UserUUID,
	)
	if err != nil {
		return err
	}
	return nil
}
