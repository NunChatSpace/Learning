func SetupTransactionAPI(router *gin.Engine) {
	transactionAPI := router.Group("api/v2")
	{
		transactionAPI.GET("/transaction", ListTransaction)
		transactionAPI.POST("/tramsaction", CreateTransaction)
	}
}

func ListTransaction(c *gin.Context) {
	c.String(http.StatusOK, "List Transaction")
}

func CreateTransaction(c *gin.Context) {
	c.String(http.StatusOK, "Create Transaction")
}