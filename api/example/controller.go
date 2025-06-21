package reportgenerator

import (
	"net/http"
	"my-golang-starter/report-generator/api/example/model"

	"github.com/gin-gonic/gin"
)

func ExecutivePrompt(context *gin.Context) {
	var body model.ExecutiveReportBody

	if err := context.BindJSON(&body); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	
	ExecutivePromptService(&body)

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Hello world"})
}

func Welcome(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Welcome to report generator AI"})
}