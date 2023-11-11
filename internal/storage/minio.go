package storage

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Storage struct {
	client *minio.Client
	cfg    Config
}

func New(cfg Config) (*Storage, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.SSL,
	})
	if err != nil {
		return nil, err
	}

	return &Storage{
		client: client,
		cfg:    cfg,
	}, nil
}

func (s Storage) Put(name string, path string) error {
	ctx := context.Background()

	_, err := s.client.FPutObject(ctx, s.cfg.Bucket, name, path, minio.PutObjectOptions{ContentType: "application/octet-stream"})

	return err
}

func (s Storage) Get(name string) (string, error) {
	return "", nil
}
