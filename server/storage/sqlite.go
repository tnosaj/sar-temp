package storage

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	db *sql.DB
}

func NewSQLiteStorage(dsn string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	s := &SQLiteStorage{db: db}
	s.init()
	return s, nil
}

func (s *SQLiteStorage) init() {
	s.db.Exec(`CREATE TABLE IF NOT EXISTS temperatures (
		client_id TEXT,
		timestamp DATETIME,
		temperature_c REAL
	)`)
}

func (s *SQLiteStorage) StoreTemperature(r TemperatureReading) error {
	_, err := s.db.Exec(`INSERT INTO temperatures (client_id, timestamp, temperature_c) VALUES (?, ?, ?)`,
		r.ClientID, r.Timestamp, r.TemperatureC)
	return err
}

func (s *SQLiteStorage) GetTodaysTemperatures(clientID string) ([]TemperatureReading, error) {
	rows, err := s.db.Query(`SELECT client_id, timestamp, temperature_c FROM temperatures
		WHERE client_id = ? AND date(timestamp) = date('now') ORDER BY timestamp`, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var readings []TemperatureReading
	for rows.Next() {
		var r TemperatureReading
		if err := rows.Scan(&r.ClientID, &r.Timestamp, &r.TemperatureC); err == nil {
			readings = append(readings, r)
		}
	}
	return readings, nil
}

func (s *SQLiteStorage) GetLastReadingTime(clientID string) (time.Time, error) {
	var ts time.Time
	err := s.db.QueryRow(`SELECT MAX(timestamp) FROM temperatures WHERE client_id = ?`, clientID).Scan(&ts)
	return ts, err
}
