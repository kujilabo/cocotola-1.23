package service

import "context"

type RepositoryFactory interface {
	NewAudioRepository(ctx context.Context) AudioRepository
}
