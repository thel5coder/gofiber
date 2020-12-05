package file

import (
	"io"
	"os"

	"github.com/labstack/echo"
)

func UploadFile(ctx echo.Context, destinationFolder, fileInputName string) (destinationFile, fileName string, err error) {
	// Source
	file, err := ctx.FormFile(fileInputName)
	if err != nil {
		return destinationFolder, file.Filename, err
	}

	src, err := file.Open()
	if err != nil {
		return destinationFolder, file.Filename, err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(destinationFolder + "" + file.Filename)
	if err != nil {
		return destinationFolder, file.Filename, err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return destinationFolder, file.Filename, err
	}

	return destinationFolder, file.Filename, err
}
