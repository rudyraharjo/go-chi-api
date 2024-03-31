package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rudyraharjo/api-rud/application"
)

var DB *pgxpool.Pool

func ConnectPostgres(cfg application.Config) (*pgxpool.Pool, error) {

	// Konfigurasi koneksi ke database
	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s pool_max_conn_idle_time=%ds pool_max_conn_lifetime=%ds pool_max_conns=%d pool_min_conns=%d",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.ConnectionPool.MaxIdleConnection,
		cfg.DB.ConnectionPool.MaxLifetimeConnection,
		cfg.DB.ConnectionPool.MaxOpenConnetcion,
		cfg.DB.ConnectionPool.MaxIdleConnection)

	// Buat konfigurasi untuk koneksi pool ke database
	configDB, err := pgxpool.ParseConfig(dbConnectionString)
	if err != nil {
		return nil, err
	}

	// Buat koneksi pool ke database
	dbPool, err := pgxpool.NewWithConfig(context.Background(), configDB)
	if err != nil {
		return nil, err
	}

	DB = dbPool
	return dbPool, nil
}
