package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// UploadingImage is a function that uploading images to AWS S3.
func UploadingImage(w http.ResponseWriter, r *http.Request) {
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, utils.CorsWildCard)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.ArrowMethods, utils.Methods)
	w.Header().Set(utils.Credential, utils.True)

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
	// set values in structs
	resultjson := dto.SimpleResutlJSON{
		Status:    true,
		ErrorCode: utils.SuccessCode,
	}
	// convert structs to json
	res, err := json.Marshal(resultjson)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
