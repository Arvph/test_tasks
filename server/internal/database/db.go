package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// DB представляет структуру для подключения к базе данных.
type DB struct {
	Pool       *pgxpool.Pool
	DbPort     string
	DbHost     string
	DbName     string
	DbUser     string
	DbPassword string
}

// PoolConnects создает пул подключений к экземпляру базы данных.
func PoolConnects(log *logrus.Logger, db *DB) error {
	log.Printf("Trying to connect to DB")
	ctx := context.Background()
	connDNS := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db.DbHost, db.DbPort, db.DbUser, db.DbPassword, db.DbName)

	var err error
	db.Pool, err = pgxpool.New(ctx, connDNS)
	if err != nil {
		log.Printf("Unable to connect to DB: %v", err)
		return err
	}
	if err = db.Pool.Ping(ctx); err != nil {
		log.Printf("Unable to ping to DB: %v", err)
		return err
	}

	log.Printf("Connect to DB is done")
	return nil
}

// ClosePool закрывает пул соединений к экземпляру базы данных.
func (db *DB) ClosePool(log *logrus.Logger) {
	if db.Pool != nil {
		db.Pool.Close()
	}
	log.Printf("Connection to DB is closed")
}
