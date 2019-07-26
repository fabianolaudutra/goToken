package sync

import (
	"encoding/json"
	"net/http"

	"github.com/fabianolaudutra/goToken/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllTokens(db *gorm.DB, response http.ResponseWriter, request *http.Request) {
	tokens := []model.Tokens{}
	db.Find(&tokens)
	responseJSON(response, http.StatusOK, tokens)
}

func CreateTokens(db *gorm.DB, response http.ResponseWriter, request *http.Request) {
	tokens := model.Tokens{}

	parse := json.NewDecoder(request.Body)
	if err := parse.Decode(&tokens); err != nil {
		responseError(response, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()

	if err := db.Save(&tokens).Error; err != nil {
		responseError(request, http.StatusInternalServerError, err.Error())
		return
	}
	responseJSON(response, http.StatusCreated, tokens)
}

func GetToken(db *gorm.DB, response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	hash := vars["hash"]
	hash := getTokenOr404(db, hash, response, request)
	if tokens == nil {
		return
	}
	responseJSON(response, http.StatusOK, token)
}

func DeleteToken(db *gorm.DB, response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	hash := vars["hash"]
	hash := getTokenOr404(db, hash, response, request)
	if hash == nil {
		return
	}
	if err := db.Delete(&hash).Error; err != nil {
		respondError(response, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(response, http.StatusNoContent, nil)
}

func getTokenOr404(db *gorm.DB, hash string, response http.ResponseWriter, request *http.Request) *model.Tokens {
	token := model.Tokens{}
	if err := db.First(&token, model.Tokens{Hash: hash}).Error; err != nil {
		respondError(response, http.StatusNotFound, err.Error())
		return nil
	}
	return &token
}
