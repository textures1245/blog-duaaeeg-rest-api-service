package entities

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"

	"log"

	"github.com/gin-gonic/gin"
	errorEntity "github.com/textures1245/BlogDuaaeeg-backend/pkg/error/entity"
)

type FileUploaderReq struct {
	FileName string `json:"file_name" form:"file_nae" validate:"required" binding:"required"`
	FileData string `json:"file_data" form:"file_data" validate:"base64" binding:"required"`
	FileType string `json:"file_type" form:"file_type" validate:"required" binding:"required"`
}

type File struct {
	Id        int    `json:"id" db:"id"`
	FileName  string `json:"file_name" db:"file_name"`
	FileData  string `json:"file_data" db:"file_data"`
	FileType  string `json:"file_type" db:"file_type"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

func (f *File) Base64toPng(c *gin.Context) (*string, *string, error) {

	if len(f.FileData) == 0 || f.FileType != "PNG" {
		return nil, nil, errors.New("Invalid file data or file type, expected PNG file type but got " + f.FileType)
	}

	hasher := sha256.New()
	hasher.Write([]byte(f.FileData))
	hash := hex.EncodeToString(hasher.Sum(nil))

	path := "public/image/"
	pngFilename := path + hash + ".png"

	if _, err := os.Stat(pngFilename); os.IsNotExist(err) {
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(f.FileData))
		m, _, err := image.Decode(reader)
		if err != nil {
			return nil, nil, err
		}
		// bounds := m.Bounds()
		// fmt.Println(bounds, formatString)

		osFile, errOnOpenFIle := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
		if errOnOpenFIle != nil {
			return nil, nil, err
		}
		err = png.Encode(osFile, m)
		if err != nil {
			return nil, nil, err
		}
		buffer := new(bytes.Buffer)
		errWhileEncoding := png.Encode(buffer, m) // img is your image.Image
		if errWhileEncoding != nil {
			return nil, nil, err
		}
		base64url := fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(buffer.Bytes()))
		filePathData := fmt.Sprintf("%s/%s", c.FullPath(), pngFilename)
		log.Println("Create new PNG file name: ", pngFilename, "as the output")

		return &base64url, &filePathData, nil
	}

	data, err := os.ReadFile(pngFilename)
	if err != nil {
		return nil, nil, err
	}

	base64url := "data:image/png;base64," + base64.StdEncoding.EncodeToString(data)
	filePathData := fmt.Sprintf("%s/%s", c.FullPath(), pngFilename)
	log.Println("Reusing exist PNG file name: ", pngFilename, "as the output")

	return &base64url, &filePathData, nil

}

func (f *File) Base64toJpg(c *gin.Context) (*string, *string, error) {

	if len(f.FileData) == 0 || f.FileType != "JPG" {
		return nil, nil, errors.New("Invalid file data or file type, expected PNG file type but got " + f.FileType)
	}

	hasher := sha256.New()
	hasher.Write([]byte(f.FileData))
	hash := hex.EncodeToString(hasher.Sum(nil))

	jpgFilename := "public/image/" + hash + ".jpg"

	if _, err := os.Stat(jpgFilename); os.IsNotExist(err) {
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(f.FileData))
		m, formatString, err := image.Decode(reader)
		if err != nil {
			return nil, nil, err
		}
		bounds := m.Bounds()
		fmt.Println("base64toJpg", bounds, formatString)

		osFile, err := os.OpenFile(jpgFilename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return nil, nil, err
		}

		err = jpeg.Encode(osFile, m, &jpeg.Options{Quality: 75})
		if err != nil {
			return nil, nil, err
		}

		buffer := new(bytes.Buffer)
		errWhileEncoding := jpeg.Encode(buffer, m, nil) // img is your image.Image
		if errWhileEncoding != nil {
			log.Fatal(errWhileEncoding)
		}
		base64url := fmt.Sprintf("data:image/jpeg;base64,%s", base64.StdEncoding.EncodeToString(buffer.Bytes()))
		filePathData := fmt.Sprintf("%s/%s", c.FullPath(), jpgFilename)
		log.Println("Create new JPG file name: ", jpgFilename, "as the output")

		return &base64url, &filePathData, nil
	}

	data, err := os.ReadFile(jpgFilename)

	if err != nil {
		return nil, nil, err
	}

	base64url := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(data)
	filePathData := fmt.Sprintf("%s/%s", c.FullPath(), jpgFilename)

	log.Println("Reusing exist JPG file name: ", jpgFilename, "as the output")

	return &base64url, &filePathData, nil
}

func (f *File) Base64toFile(c *gin.Context, includeDomain bool) (*string, *string, error) {
	if len(f.FileData) == 0 || f.FileType != "PDF" || f.FileType != "MARKDOWN_FILE" {
		return nil, nil, errors.New("Invalid file data or file type, expected PDF file type but got " + f.FileType)
	}

	// encode blob to string
	hasher := sha256.New()
	hasher.Write([]byte(f.FileData))
	hash := hex.EncodeToString(hasher.Sum(nil))

	var fileName string
	switch f.FileType {
	case "PDF":
		fileName = "public/file/" + hash + ".pdf"
	case "MARKDOWN_FILE":
		fileName = "public/file/" + hash + ".md"
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {

		data, err := base64.StdEncoding.DecodeString(f.FileData)
		if err != nil {
			return nil, nil, err
		}

		err = os.WriteFile(fileName, data, 0644)
		if err != nil {
			return nil, nil, err
		}

		srcFile := fmt.Sprintf("data:file/%s;base64,%s", strings.ToLower(f.FileType), base64.StdEncoding.EncodeToString(data))

		log.Println("Reusing exist ", f.FileType, " file name: ", fileName, "as the output")

		filePathData := fileName
		if includeDomain {
			filePathData = fmt.Sprintf("%s/%s", c.FullPath(), fileName)
		}
		return &srcFile, &filePathData, nil
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, nil, err
	}

	srcFile := fmt.Sprintf("data:file/%s;base64,%s", strings.ToLower(f.FileType), base64.StdEncoding.EncodeToString(data))
	filePathData := fileName
	if includeDomain {
		filePathData = fmt.Sprintf("%s/%s", c.FullPath(), fileName)
	}
	log.Println("Reusing exist ", f.FileType, " file name: ", fileName, "as the output")

	return &srcFile, &filePathData, nil

}

func (file *File) EncodeBase64toFile(c *gin.Context, domainIncludeOnFile bool) (*string, *string, int, *errorEntity.CError) {
	var (
		base64urlRes string
		fPathDatRes  string
	)
	switch file.FileType {
	case "PNG":
		base64url, fPathDat, err := file.Base64toPng(c)
		if err != nil {
			return nil, nil, http.StatusInternalServerError, &errorEntity.CError{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}
		base64urlRes = *base64url
		fPathDatRes = *fPathDat
	case "JPG":
		base64url, fPathDat, err := file.Base64toJpg(c)
		if err != nil {
			return nil, nil, http.StatusInternalServerError, &errorEntity.CError{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}
		base64urlRes = *base64url
		fPathDatRes = *fPathDat
	case "PDF":
		base64url, fPathDat, err := file.Base64toFile(c, domainIncludeOnFile)
		if err != nil {
			return nil, nil, http.StatusInternalServerError, &errorEntity.CError{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}
		base64urlRes = *base64url
		fPathDatRes = *fPathDat
	default:
		return nil, nil, http.StatusUnsupportedMediaType, &errorEntity.CError{
			StatusCode: http.StatusUnsupportedMediaType,
			Err:        errors.New("Only except for PNG, JPG, PDF AND MD for now"),
		}
	}

	return &base64urlRes, &fPathDatRes, http.StatusOK, nil
}
