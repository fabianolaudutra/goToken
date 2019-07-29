package sync

import (
	"encoding/json"
	"net/http"
	"github.com/fabianolaudutra/goToken/app/model"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	"io/ioutil"
	"time"
	"golang.org/x/crypto/sha3"
	"encoding/hex"
	
)

type retToken struct {
	Token string `json:"token"`
	Hashe string `json:"hashe"`
	Created_at time.Time `json:"created_at"`
}

type responseToken struct {
	
	Hashe string `json:"hashe"`
	Created_at time.Time `json:"created_at"`
}


func GetAllTokens(db *gorm.DB, resp http.ResponseWriter, req *http.Request) {
	tks := []model.Tokens{}
	db.Find(&tks)
		
	responseJSON(resp, http.StatusOK, tks)
}

func CreateTokens(db *gorm.DB, resp http.ResponseWriter, req *http.Request) {
	tokens_ := responseToken{}
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(resp, err.Error(), 500)
		return
	}
	
	var ret retToken
	err = json.Unmarshal(b, &ret)
	if err != nil {
		http.Error(resp, err.Error(), 500)
		return
	}	
	
	aStringToHash := []byte( ret.Token)
	token := sha3.Sum256(aStringToHash)
	tk := model.Tokens{Hash:string(hex.EncodeToString(token[:])),Token:ret.Token}
	
	tokens_.Hashe= string(hex.EncodeToString(token[:]))
	tokens_.Created_at = time.Now()
	
	output, err := json.Marshal(tokens_)
	if err != nil {
		http.Error(resp, err.Error(), 500)
		return
	}
	tokensValidade := getTokenOr404(db, string(hex.EncodeToString(token[:])), resp, req)
	if tokensValidade != nil {
		responseError(resp, http.StatusInternalServerError,"Duplicate Hash")
		return
	}

	resp.Header().Set("content-type", "application/json")
	if err := db.Save(&tk).Error; err != nil {
		responseError(resp, http.StatusInternalServerError, err.Error())
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(output)
	
}

func GetToken(db *gorm.DB, resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	token := vars["token"]

	tokens := getTokenOr404(db, token, resp, req)
	if tokens == nil {
		return
	}
	responseJSON(resp, http.StatusOK, tokens)
}

func DeleteToken(db *gorm.DB, resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	token := vars["token"]
	hashes := getTokenOr404(db, token, resp, req)
	if hashes == nil {
		return
	}
	if err := db.Delete(&token).Error; err != nil {
		responseError(resp, http.StatusInternalServerError, err.Error())
		return
	}
	responseJSON(resp, http.StatusNoContent, nil)
}



func getTokenOr404(db *gorm.DB, hash string, resp http.ResponseWriter, req *http.Request) *model.Tokens {
	token := model.Tokens{}
	if err := db.First(&token, model.Tokens{Hash: hash}).Error; err != nil {
		
		return nil
	}
	return &token
}


