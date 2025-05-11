package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"video-subscriber/internal/playlist"
	"video-subscriber/internal/sync/youtube"
	"video-subscriber/internal/ytdl"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:54321/video_subscriber?sslmode=disable")
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		panic(err)
	}

	conn, err := db.Conn(ctx)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		panic(err)
	}

	bin := ytdl.NewBinary("youtube-dl")
	repo := playlist.NewRepository(conn)
	syncer := youtube.NewSyncer(bin, repo, "https://www.youtube.com/playlist?list=PLHL0FtQN1IkhPkfuY5vuSIitWauC96W4w")
	fmt.Println(syncer.Sync(ctx))
}
