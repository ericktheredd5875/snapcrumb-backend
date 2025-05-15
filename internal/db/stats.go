package db

import (
	"database/sql"
	"time"
)

func LogVisit(db *sql.DB, shortcode string, ip string, userAgent string) error {
	query := `
		INSERT INTO url_visits (shortcode, visited_at, ip_address, user_agent)
		VALUES ($1, $2, $3, $4)
	`

	_, err := db.Exec(query, shortcode, time.Now(), ip, userAgent)
	return err
}

func GetStats(db *sql.DB, shortcode string) (count int, lastVisit sql.NullTime, err error) {
	query := `
		SELECT COUNT(*), MAX(visited_at) FROM url_visits 
		WHERE shortcode = $1;
	`

	err = db.QueryRow(query, shortcode).Scan(&count, &lastVisit)
	return
}
