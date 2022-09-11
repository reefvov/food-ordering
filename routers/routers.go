package routers

import (
	"github.com/gin-gonic/gin"
	_"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	h := CustomerHandler{}
	h.Initialize()

	r.GET("/customers", h.GetAllCustomer)
	r.GET("/customers/:id", h.GetCustomer)
	r.POST("/customers", h.SaveCustomer)
	r.PUT("/customers/:id", h.UpdateCustomer)
	r.DELETE("/customers/:id", h.DeleteCustomer)

	return r
}
