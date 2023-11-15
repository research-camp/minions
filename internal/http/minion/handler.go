package minion

const LocalDir = "./tmp/local"

type Handler struct {
	Router string
	MinIO  MinIO
}
