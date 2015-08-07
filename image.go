package main

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func NewImage(c *gin.Context) {
	err := newImage(c)
	if err != nil {
		c.JSON(err.Status, err)
	}
}

func newImage(c *gin.Context) *Error {
	contentType := getContentType(c)

	if contentType == "multipart/form-data" {
		return newImageFromFormData(c)
	}

	return newImageFromOctetStream(c)
}

func newImageFromFormData(c *gin.Context) *Error {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return NewError(http.StatusBadRequest, "Bad Request", &err)
	}
	filename := createFileName(filepath.Ext(header.Filename))

	out, err := os.Create(filepath.Join("./files", filename))
	if err != nil {
		return NewError(http.StatusInternalServerError, "Can not write file.", &err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return NewError(http.StatusInternalServerError, "Can not write file.", &err)
	}

	url := getBaseUrl() + "/files/" + filename

	c.JSON(http.StatusCreated, url)

	return nil
}

func newImageFromOctetStream(c *gin.Context) *Error {
	contentType := getContentType(c)

	buff := make([]byte, 4096)

	size, err := c.Request.Body.Read(buff)
	if err != nil {
		return NewError(http.StatusInternalServerError, "Can not read request body.", &err)
	}

	ext := getExt(contentType)
	if ext == "" {
		contentType = http.DetectContentType(buff)
		ext = getExt(contentType)
	}

	filename := createFileName(ext)
	out, err := os.Create(filepath.Join("./files", filename))
	if err != nil {
		return NewError(http.StatusInternalServerError, "Can not write file.", &err)
	}
	defer out.Close()

	_, err = out.Write(buff[0:size])
	if err != nil {
		return NewError(http.StatusInternalServerError, "Can not write file.", &err)
	}

	_, err = io.Copy(out, c.Request.Body)
	if err != nil {
		return NewError(http.StatusInternalServerError, "Can not write file.", &err)
	}

	url := getBaseUrl() + "/files/" + filename

	c.JSON(http.StatusCreated, url)

	return nil
}

func getExt(contentType string) string {
	ext := ""
	switch contentType {
	case "image/png":
		ext = ".png"
	case "image/jpeg":
		ext = ".jpg"
	case "image/gif":
		ext = ".gif"
	}

	return ext
}

func createFileName(ext string) string {
	return uuid.New() + ext
}

func getContentType(c *gin.Context) string {
	return strings.Split(c.Request.Header.Get("Content-Type"), ";")[0]
}
