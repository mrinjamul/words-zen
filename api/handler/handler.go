package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/words-zen/models"
	"github.com/mrinjamul/words-zen/utils"
)

func GetHandler(ctx *gin.Context) {
	query := ctx.Param("phrase")
	phrase := utils.PhraseResolver(query)
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"phrase": phrase,
		},
	)
}

func PostHandler(ctx *gin.Context) {
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	var query models.Query
	err = json.Unmarshal(bytes, &query)
	if err != nil {
		log.Fatal(err)
	}
	phrase := utils.PhraseResolver(query.Query)
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"phrase": phrase,
		},
	)
}
