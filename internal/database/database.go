package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Database struct {
	Pool   *pgxpool.Pool
	logger *zerolog.Logger
}

// DatabaseTimeout is the timeout in seconds for database operations
const DatabaseTimeout = 10
