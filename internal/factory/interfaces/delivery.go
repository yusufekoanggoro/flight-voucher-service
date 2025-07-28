package interfaces

import "github.com/gin-gonic/gin"

type RestHandler interface {
	RegisterRoutes(router gin.IRoutes)
}
