package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	ConfigureApi(*gin.Engine)
	ConfigureViews(*gin.Engine)
}
