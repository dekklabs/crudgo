package api

import (
	"config"
	"encoding/json"
	"models"
	"net/http"
)

//FindAllProduct api respuesta, trae todos los datos
func FindAllProduct(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDatabase()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithjson(response, http.StatusOK, products)
		}
	}
}

/* Tools */
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithjson(w, code, map[string]string{"error": msg})
}

func respondWithjson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
