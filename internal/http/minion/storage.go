package minion

type MinIO interface {
	Put(name, path string) error
	Get(name, path string) error
}
