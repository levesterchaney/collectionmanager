package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func Migrate(db *sql.DB) error {
	migrations := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash VARCHAR(255) NOT NULL,
			first_name VARCHAR(100),
			last_name VARCHAR(100),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		
		`CREATE TABLE IF NOT EXISTS collections (
			id SERIAL PRIMARY KEY,
			user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
			name VARCHAR(255) NOT NULL,
			description TEXT,
			category VARCHAR(100),
			is_public BOOLEAN DEFAULT false,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		
		`CREATE TABLE IF NOT EXISTS items (
			id SERIAL PRIMARY KEY,
			collection_id INTEGER REFERENCES collections(id) ON DELETE CASCADE,
			name VARCHAR(255) NOT NULL,
			description TEXT,
			category VARCHAR(100),
			condition VARCHAR(50),
			purchase_price DECIMAL(10,2),
			purchase_date DATE,
			photos TEXT[],
			metadata JSONB,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		
		`CREATE TABLE IF NOT EXISTS valuations (
			id SERIAL PRIMARY KEY,
			item_id INTEGER REFERENCES items(id) ON DELETE CASCADE,
			estimated_value DECIMAL(10,2) NOT NULL,
			valuation_date DATE NOT NULL,
			source VARCHAR(100),
			notes TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		
		`CREATE INDEX IF NOT EXISTS idx_collections_user_id ON collections(user_id)`,
		`CREATE INDEX IF NOT EXISTS idx_items_collection_id ON items(collection_id)`,
		`CREATE INDEX IF NOT EXISTS idx_valuations_item_id ON valuations(item_id)`,
	}

	for _, migration := range migrations {
		if _, err := db.Exec(migration); err != nil {
			return fmt.Errorf("failed to execute migration: %w", err)
		}
	}

	return nil
}