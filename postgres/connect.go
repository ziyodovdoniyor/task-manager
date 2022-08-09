package postgres

import (
	"alfa1/config"
	"database/sql"
	"fmt"
)

func Connection(cn config.Config) (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cn.PostgresHost, cn.PostgresPort, cn.PostgresUser, cn.PostgresPassword, cn.PostgresDB,
		),
	)

	if err != nil {
		return &sql.DB{}, err
	}

	if err = db.Ping(); err != nil {
		return &sql.DB{}, err
	}

	return db, nil
}
