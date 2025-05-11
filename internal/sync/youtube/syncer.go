package youtube

import (
	"context"
)

type Binary interface {
	ListVideos(context.Context, string) ([]string, error)
}

type Repository interface {
	ListVideos(context.Context, string) (map[string]struct{}, error)
}

type Syncer struct {
	bin  Binary
	url  string
	repo Repository
}

func NewSyncer(bin Binary, repo Repository, url string) *Syncer {
	return &Syncer{
		bin:  bin,
		url:  url,
		repo: repo,
	}
}

func (s *Syncer) Sync(ctx context.Context) error {
	urls, err := s.bin.ListVideos(ctx, s.url)
	if err != nil {
		return err
	}

	existing, err := s.repo.ListVideos(ctx, s.url)
	if err != nil {
		return err
	}

	var result []string
	for _, url := range urls {
		if _, ok := existing[url]; !ok {
			result = append(result, url)
		}
	}
	// TODO: Use the new found videos to do something
	return nil
}
