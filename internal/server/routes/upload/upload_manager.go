package upload

import (
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.canergulay/blogbackend/internal/core"
	"github.canergulay/blogbackend/internal/server/endpoints"
	"github.com/labstack/echo/v4"
)

type UploadManager struct {
}

func NewUploadManager() *UploadManager {
	return &UploadManager{}
}

func (h UploadManager) Handler(c echo.Context) error {
	file, err := c.FormFile("file")

	if err != nil {
		c.String(http.StatusInternalServerError, "an unexpected error occured...")
		return err
	}

	fileName, saveErr := saveFile(file)
	if saveErr != nil {
		return saveErr
	}
	fileNameWithExtension := core.ASSET_EXTENSION + fileName
	return c.HTML(http.StatusOK, fileNameWithExtension)
}

func (h UploadManager) GetUploadEndpoint() endpoints.Endpoint {
	return endpoints.NewEndpoint("/upload", h.Handler, "POST")
}

func saveFile(file *multipart.FileHeader) (string, error) {
	fileNameUniqifed := fmt.Sprintf("%d_%d_%s", rand.Intn(1000000), time.Now().Unix(), file.Filename)
	fileNameWithAssetPath := fmt.Sprintf("%s%s", core.ASSET_PATH, fileNameUniqifed)
	src, err := file.Open()

	defer func() {
		cerr := src.Close()

		if err == nil {
			err = cerr
		}
	}()
	dst, err := os.Create(fileNameWithAssetPath)

	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return fileNameUniqifed, nil
}
