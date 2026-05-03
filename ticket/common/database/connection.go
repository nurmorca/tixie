package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConnectionPool(context context.Context, config Config) *pgxpool.Pool {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s pool_max_conns=%s pool_max_conn_idle_time=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.DBname,
		config.MaxConnection,
		config.MaxConnectionIdleTime)

	connConfig, parseConfigError := pgxpool.ParseConfig(connString) // checks if connString has problems
	if parseConfigError != nil {
		panic(parseConfigError)
	}

	var pool *pgxpool.Pool
	var err error

	// Retry loop to handle the "Fast Shutdown" window
	for i := 0; i < 15; i++ {
		pool, err = pgxpool.NewWithConfig(context, connConfig)
		if err == nil {
			// Ping actually attempts a connection
			err = pool.Ping(context)
			if err == nil {
				return pool // Success!
			}
		}

		if pool != nil {
			pool.Close()
		}
		fmt.Printf("Database starting up... retrying in 1s (attempt %d)\n", i+1)
		time.Sleep(1 * time.Second)
	}

	panic(fmt.Sprintf("Could not connect to database after retries: %v", err))
	return nil
}
