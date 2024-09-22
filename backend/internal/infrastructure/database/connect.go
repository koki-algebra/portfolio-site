package database

import (
	"context"
	"database/sql"
	"fmt"
	"runtime"
	"time"

	"github.com/jmoiron/sqlx"

	"backend/internal/config"

	_ "github.com/lib/pq"
)

func Open(ctx context.Context) (*sql.DB, error) {
	// connect to database
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		config.Env.DBUser,
		config.Env.DBPassword,
		config.Env.DBHost, config.Env.DBPort,
		config.Env.DBDatabase,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(time.Minute * 15)
	db.SetConnMaxLifetime(time.Hour * 12)

	maxOpenConns := 4 * runtime.GOMAXPROCS(0)
	db.SetMaxIdleConns(maxOpenConns)
	db.SetMaxOpenConns(maxOpenConns)

	// verify connection
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func NewSqlxDB(db *sql.DB) *sqlx.DB {
	return sqlx.NewDb(db, "postgres")
}
