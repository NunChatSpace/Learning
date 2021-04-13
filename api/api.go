package api

import (
	"main/db"

	"github.com/gin-gonic/gin"
)

// Setup - Grouping request and setup db
func Setup(router *gin.Engine) {
	db.SetupDB()

	SetupAuthenAPI(router)
	SetupProductAPI(router)
	SetupTransactionAPI(router)
}
