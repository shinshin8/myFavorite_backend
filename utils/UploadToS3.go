package utils

import (
	"io"
	"net/http"
	"os"
)

// UploadToS3 uploads and saves images in AWS S3.
func UploadToS3(w http.ResponseWriter, r *http.Request) {
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		if part.FileName() == "" {
			continue
		}

		uploadFile, err := os.Create("/Users/shintarokanno/" + part.FileName())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			uploadFile.Close()
			return
		}

		_, err = io.Copy(uploadFile, part)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			uploadFile.Close()
			return
		}
	}
	w.WriteHeader(200)
}
