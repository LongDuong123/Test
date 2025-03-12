package api

import (
	"github.com/gin-gonic/gin"
)

func Routes(ctr UserController) *gin.Engine {
	router := gin.Default()

	empGroup := router.Group("/user")
	{
		empGroup.POST("/join", ctr.UserJoin)
		empGroup.POST("/checkin", ctr.CheckIn)
		empGroup.POST("/checkout", ctr.CheckOut)
		empGroup.POST("/attendance", ctr.UserAttendance)
		empGroup.POST("/salary", ctr.UserSalary)
	}

	rootGroup := router.Group("/root")
	rootGroup.Use(RootAuthMiddleware)
	{
		rootGroup.POST("/updateuser", ctr.RootUpdateUser)
		rootGroup.POST("/deleteuser", ctr.RootDeleteUser)
		rootGroup.GET("/refcode", ctr.RootCreateRefCode)
		rootGroup.GET("/overview", ctr.RootCompanyOverview)
	}

	return router
}
