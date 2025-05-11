package sync

import "context"

type Syncer interface {
	Sync(ctx context.Context) error
}
