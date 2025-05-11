package playlist

import (
	"context"
	"database/sql"
)

type Repository struct {
	conn *sql.Conn
}

func NewRepository(conn *sql.Conn) *Repository {
	return &Repository{conn: conn}
}

func (r *Repository) ListVideos(ctx context.Context, id string) (map[string]struct{}, error) {
	query := "SELECT * FROM playlist_videos WHERE playlist_id = $1"
	rows, err := r.conn.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result map[string]struct{}
	for rows.Next() {
		var video string
		if err := rows.Scan(&video); err != nil {
			return nil, err
		}
		result[video] = struct{}{}
	}

	return result, nil
}
