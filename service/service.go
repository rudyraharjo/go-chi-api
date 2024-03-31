package service

import "github.com/jackc/pgx/v5/pgxpool"

type Service struct {
	db *pgxpool.Pool
}
