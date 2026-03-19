package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/gommon/log"
)

func GetConnectionPool(context context.Context, config Config) *pgxpool.Pool {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s pool_max_conns=%s pool_max_conn_idle_time=%s",
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

	connConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe
	conn, err := pgxpool.NewWithConfig(context, connConfig) // actually makes the connection.
	if err != nil {
		log.Error("unable to connect db: %v\n", err)
		panic(err)
	}

	return conn
}
