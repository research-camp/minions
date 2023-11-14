package minion

import (
	"fmt"
	"github.com/amirhnajafiz/minions/pkg/enum"
	"log"
	"net/http"
	"os"

	"github.com/amirhnajafiz/minions/internal/storage"

	"github.com/gofiber/fiber/v2"
)

const LocalDir = "./tmp/local"

type Handler struct {
	Router string
	MinIO  *storage.Storage
}

func (h Handler) Notify(code string) {
	url := fmt.Sprintf("%s?signal=%s", h.Router, code)
	request, _ := http.NewRequest(fiber.MethodGet, url, nil)
	client := http.Client{}

	if _, err := client.Do(request); err != nil {
		log.Println(fmt.Errorf("failed to send signal: %w", err))
	}
}

func (h Handler) Download(ctx *fiber.Ctx) error {
	name := ctx.Query("file", "")
	if len(name) == 0 {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	code := enum.HitSignal
	path := fmt.Sprintf("%s/%s", LocalDir, name)

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if er := h.MinIO.Get(name, path); er != nil {
				return ctx.SendStatus(fiber.StatusInternalServerError)
			}

			code = enum.MissSignal
		} else {
			log.Println(fmt.Errorf("failed to check file: %w", err))

			return ctx.SendStatus(fiber.StatusNotFound)
		}
	}

	go h.Notify(code)

	return ctx.SendFile(path)
}

func (h Handler) Upload(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return fmt.Errorf("failed to get multipart form: %w", err)
	}

	for _, file := range form.File["file"] {
		path := fmt.Sprintf("%s/%s", LocalDir, file.Filename)
		if er := ctx.SaveFile(file, path); er != nil {
			return fmt.Errorf("failed to save file on local: %w", er)
		}

		if er := h.MinIO.Put(file.Filename, path); er != nil {
			return fmt.Errorf("failed to send on MinIo: %w", er)
		}
	}

	return ctx.SendStatus(fiber.StatusOK)
}
