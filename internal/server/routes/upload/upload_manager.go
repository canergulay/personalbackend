package upload

import (
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"

	"github.canergulay/blogbackend/internal/server/endpoints"
	"github.com/labstack/echo/v4"
)

type UploadManager struct {
}

func NewUploadManager() *UploadManager {
	return &UploadManager{}
}

func (h UploadManager) Handler(c echo.Context) error {
	fmt.Println("geldi")
	file, err := c.FormFile("file")
	fmt.Println("geldi", file.Filename)

	if err != nil {
		c.String(http.StatusInternalServerError, "an unexpected error occured...")
		return err
	}

	saveErr := saveFile(file)
	if saveErr != nil {
		return saveErr
	}
	return c.HTML(http.StatusOK, "upload has been completed successfully !")
}

func (h UploadManager) GetUploadEndpoint() endpoints.Endpoint {
	return endpoints.NewEndpoint("/upload", h.Handler, "POST")
}

func saveFile(file *multipart.FileHeader) error {
	n1, n2, n3, n4, n5 := rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(1000000)
	nAll := fmt.Sprintf("%d-%d-%d-%d-%d", n1, n2, n3, n4, n5)
	fileNameUniqifed := fmt.Sprintf("%s-%s", nAll, file.Filename)
	fmt.Println(fileNameUniqifed)
	src, err := file.Open()
	fmt.Println(err)

	defer func() {
		cerr := src.Close()
		fmt.Println(cerr)

		if err == nil {
			err = cerr
		}
	}()
	dst, err := os.Create(fileNameUniqifed)
	fmt.Println(err)

	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	fmt.Println(err)

	return err
}
