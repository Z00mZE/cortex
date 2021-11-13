package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

const (
	defaultMaxPoolSize        = 10
	defaultConnectionAttempts = 10
	defaultConnectionTimeout  = time.Second
)

// Postgres -.
type Postgres struct {
	maxPoolSize               int32
	connectionAttempts        int
	connectionAttemptsTimeout time.Duration
	Pool                      *pgxpool.Pool
}

//NewPostgres -.
func NewPostgres(url string, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:               defaultMaxPoolSize,
		connectionAttempts:        defaultConnectionAttempts,
		connectionAttemptsTimeout: defaultConnectionTimeout,
	}

	//	Some custom Postgre connection settings
	for _, opt := range opts {
		opt(pg)
	}

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - pgxpool.ParseConfig: %w", err)
	}

	poolConfig.MaxConns = pg.maxPoolSize

	for ; pg.connectionAttempts > 0; pg.connectionAttempts-- {
		pg.Pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", pg.connectionAttempts)

		time.Sleep(pg.connectionAttemptsTimeout)
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - connectionAttempts == 0: %w", err)
	}

	return pg, nil
}

//	Close db connection
func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
