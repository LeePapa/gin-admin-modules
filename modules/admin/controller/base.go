package admin_controller

import (
	"github.com/gin-gonic/gin"
)

func OutputJson(ct *gin.Context, code int, msg string, data interface{}) {
	ct.JSON(200, gin.H{"code": code, "msg": msg, "data": data})
}
