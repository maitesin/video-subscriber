package ytdl

import (
	"context"
	"os/exec"
)

type Binary struct {
	binaryPath string
}

func NewBinary(binaryPath string) *Binary {
	return &Binary{binaryPath: binaryPath}
}

func (b *Binary) ListVideos(ctx context.Context, url string) ([]string, error) {
	data, err := exec.CommandContext(ctx, b.binaryPath, "--print-json", url).Output()
	if err != nil {
		return nil, err
	}
	results := make([]string, len(data))
	for i, d := range data {
		results[i] = string(d)
	}
	return results, nil
}
