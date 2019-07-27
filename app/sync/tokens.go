package sync

import (
	"encoding/json"
	"net/http"

	
)

func GetAllTokens(db *gorm.DB, resp http.ResponseWriter, req *http.Request) {
	tokens := []model.Tokens{}
	db.Find(&tokens)
	responseJSON(resp, http.StatusOK, tokens)
}

func CreateTokens(db *gorm.DB, resp http.ResponseWriter, req *http.Request) {
	tokens := model.Tokens{}

	parse := json.NewDecoder(req.Body)
	if err := parse.Decode(&tokens); err != nil {
		responseError(resp, http.StatusBadRequest, err.Error())
		return
	}
	defer req.Body.Close()

	if err := db.Save(&tokens).Error; err != nil {
		responseError(resp, http.StatusInternalServerError, err.Error())
		return
	}
	responseJSON(resp, http.StatusCreated, tokens)
}

func GetToken(db *gorm.DB, resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	hash := vars["hash"]
	tokens := getTokenOr404(db, hash, resp, req)
	if tokens == nil {
		return
	}
	responseJSON(resp, http.StatusOK, tokens)
}

func DeleteToken(db *gorm.DB, resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	hash := vars["hash"]
	hashes := getTokenOr404(db, hash, resp, req)
	if hashes == nil {
		return
	}
	if err := db.Delete(&hash).Error; err != nil {
		responseError(resp, http.StatusInternalServerError, err.Error())
		return
	}
	responseJSON(resp, http.StatusNoContent, nil)
}

func getTokenOr404(db *gorm.DB, hash string, resp http.ResponseWriter, req *http.Request) *model.Tokens {
	token := model.Tokens{}
	if err := db.First(&token, model.Tokens{Hash: hash}).Error; err != nil {
		responseError(resp, http.StatusNotFound, err.Error())
		return nil
	}
	return &token
}
