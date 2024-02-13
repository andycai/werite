package media

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/andycai/werite/constant"
	"github.com/andycai/werite/library/random"
	"github.com/andycai/werite/log"
	"github.com/andycai/werite/model"
)

const (
	KEY_CMS_UPLOAD_DIR        = "./data/uploads/"
	KEY_CMS_EXTERNAL_UPLOADER = ""
)

func removeFile(path, name string) error {
	if name == "" {
		return errors.New("invalid path and name")
	}

	media, err := getMedia(path, name)
	if err != nil {
		return err
	}

	if media.External {
		return nil
	}

	uploadDir := KEY_CMS_UPLOAD_DIR
	fullPath := filepath.Join(uploadDir, media.StorePath)
	if err := os.Remove(fullPath); err != nil {
		return err
	}
	return nil
}

func getLatest(path string, count int) ([]model.Media, error) {
	var medias []model.Media
	tx := db.Model(&model.Media{}).Where("path", path).Order("created_at desc").Limit(count)
	r := tx.Find(&medias)
	if r.Error != nil {
		return nil, r.Error
	}
	return medias, nil
}

func getMedia(path, name string) (*model.Media, error) {
	var obj model.Media
	if len(path) > 1 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	tx := db.Model(&model.Media{}).Where("path", path).Where("name", name)
	r := tx.First(&obj)
	if r.Error != nil {
		return nil, r.Error
	}
	return &obj, nil
}

func makeMediaPublish(siteID, path, name string, obj any, publish bool) error {
	tx := db.Model(obj).Where("site_id", siteID).Where("path", path).Where("name", name)
	vals := map[string]any{"published": publish}
	vals["published"] = publish
	return tx.Updates(vals).Error
}

func prepareStoreLocalDir() (string, error) {
	uploadDir := KEY_CMS_UPLOAD_DIR
	if uploadDir == "" {
		return "", errors.New("uploads dir not configured")
	}

	if _, err := os.Stat(uploadDir); err != nil {
		if os.IsNotExist(err) {
			log.Infof("upload dir not exist, create it: %s", uploadDir)
			if err = os.MkdirAll(uploadDir, 0755); err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}
	return uploadDir, nil
}

func storeLocal(uploadDir, storePath string, data []byte) (string, error) {
	storePath = filepath.Join(uploadDir, storePath)
	err := os.WriteFile(storePath, data, 0644)
	if err != nil {
		return "", err
	}
	return storePath, nil
}

func storeExternal(externalUploader, path, name string, data []byte) (string, error) {
	buf := new(bytes.Buffer)
	form := multipart.NewWriter(buf)
	form.WriteField("path", path)
	form.WriteField("name", name)

	fileField, _ := form.CreateFormFile("file", name)
	fileField.Write(data)
	form.Close()

	resp, err := http.Post(externalUploader, form.FormDataContentType(), buf)
	if err != nil {
		log.Infof("upload to external server failed: %s, %s", err, externalUploader)
		return "", err
	}

	defer resp.Body.Close()
	respBody := bytes.NewBuffer(nil)
	io.Copy(respBody, resp.Body)
	body := respBody.Bytes()
	if resp.StatusCode != http.StatusOK {
		log.Infof("upload to external server failed: %s, %s, %s", resp.StatusCode, externalUploader, string(body))
		return "", fmt.Errorf("upload to external server failed, code:%d %s", resp.StatusCode, string(body))
	}
	var remoteResult model.UploadResult
	json.Unmarshal(body, &remoteResult)
	return remoteResult.StorePath, nil
}

func uploadFile(path, name string, reader io.Reader) (*model.UploadResult, error) {
	var r model.UploadResult
	r.Path = path
	r.Name = name
	r.Ext = strings.ToLower(filepath.Ext(name))

	canGetDimension := false

	switch r.Ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		canGetDimension = true
		fallthrough
	case ".webp", ".svg", ".ico", ".bmp":
		r.ContentType = constant.ContentTypeImage
	case ".mp3", ".wav", ".ogg", ".aac", ".flac":
		r.ContentType = constant.ContentTypeAudio
	case ".mp4", ".webm", ".avi", ".mov", ".wmv", ".mkv":
		r.ContentType = constant.ContentTypeVideo
	default:
		r.ContentType = constant.ContentTypeFile
	}
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	r.Size = int64(len(data))

	externalUploader := KEY_CMS_EXTERNAL_UPLOADER
	if externalUploader != "" {
		storePath, err := storeExternal(externalUploader, path, name, data)
		if err != nil {
			return nil, err
		}
		r.StorePath = storePath
		r.External = true
	} else {
		storePath := fmt.Sprintf("%s%s", random.RandText(10), r.Ext)
		r.StorePath = storePath
		r.External = false
		uploadDir, err := prepareStoreLocalDir()
		if err != nil {
			return nil, err
		}
		_, err = storeLocal(uploadDir, storePath, data)
		if err != nil {
			return nil, err
		}
	}

	if canGetDimension {
		config, _, err := image.DecodeConfig(bytes.NewReader(data))
		if err == nil {
			r.Dimensions = fmt.Sprintf("%dX%d", config.Width, config.Height)
		} else {
			log.Infof("decode image config error: %s", err)
			r.Dimensions = "X"
		}
	}
	return &r, nil
}
