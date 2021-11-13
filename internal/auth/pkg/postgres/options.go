package postgres

import "time"

type Option func(*Postgres)

// MaxPoolSize set max pool connections size
func MaxPoolSize(size int32) Option {
	return func(c *Postgres) {
		c.maxPoolSize = size
	}
}

//	ConnectionAttempts set connection attempts
func ConnectionAttempts(attempts int) Option {
	return func(c *Postgres) {
		c.connectionAttempts = attempts
	}
}

//	ConnectionAttemptsTimeout set connection attempts timeout
func ConnectionAttemptsTimeout(timeout time.Duration) Option {
	return func(c *Postgres) {
		c.connectionAttemptsTimeout = timeout
	}
}
