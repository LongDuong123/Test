package api

import (
	"net/http"

	"example.com/models"
	"example.com/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	UserJoin(ctx *gin.Context)
	CheckIn(ctx *gin.Context)
	CheckOut(ctx *gin.Context)
	UserAttendance(ctx *gin.Context)
	UserSalary(ctx *gin.Context)
	RootUpdateUser(ctx *gin.Context)
	RootDeleteUser(ctx *gin.Context)
	RootCreateRefCode(ctx *gin.Context)
	RootCompanyOverview(ctx *gin.Context)
}

type userController struct {
	serv service.UserService
}

func NewUserController(serv service.UserService) UserController {
	c := &userController{}
	c.serv = serv
	return c
}

func (c *userController) UserJoin(ctx *gin.Context) {
	var req models.UserJoinRequest
	if err := ctx.ShouldBindJSON(&req); err != nil { // Bind the request to the struct
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	walletAddr, walletKey, err := c.serv.UserJoin(req.Name, req.Refcode) // Call service to register the user
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"wallet_addr": walletAddr,
		"wallet_key":  walletKey,
	})
}

func (c *userController) CheckIn(ctx *gin.Context) {
	var req models.UserCheckInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil { // Bind the request to the struct
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.serv.UserCheckIn(req.WalletAddr) // Call service to check in the user
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.UserCheckInResponse{
		Success: true,
	})
}

func (c *userController) CheckOut(ctx *gin.Context) {
	var req models.UserCheckOutRequest
	if err := ctx.ShouldBindJSON(&req); err != nil { // Bind the request to the struct
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.serv.UserCheckOut(req.WalletAddr) // Call service to check out the user
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.UserCheckOutResponse{
		Success: true,
	})
}

func (c *userController) UserAttendance(ctx *gin.Context) {
	var req models.UserAttendanceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil { // Bind the request to the struct
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attendance, err := c.serv.UserAttendance(req.WalletAddr) // Call service to get the user's attendance
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, attendance)
}

func (c *userController) UserSalary(ctx *gin.Context) {
	var req models.UserSalaryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil { // Bind the request to the struct
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	salary, err := c.serv.UserSalary(req.WalletAddr) // Call service to get the user's salary
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.UserSalaryResponse{
		Salary: salary,
	})
}

func (c *userController) RootUpdateUser(ctx *gin.Context) {
	var req models.RootUpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil { // Bind the request to the struct
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.serv.RootUpdateUser(req.Name, req.WalletAddr, req.Role, req.Salary) // Call service to update the user
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.RootUpdateUserResponse{
		Success: true,
	})
}

func (c *userController) RootDeleteUser(ctx *gin.Context) {
	var req models.RootDeleteUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil { // Bind the request to the struct
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.serv.RootDeleteUser(req.WalletAddr) // Call service to delete the user
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.RootDeleteUserResponse{
		Success: true,
	})
}

func (c *userController) RootCreateRefCode(ctx *gin.Context) {
	refcode, err := c.serv.RootCreateRefCode() // Call service to create a referral code
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.RootCreateRefCodeResponse{
		Refcode: refcode,
	})
}

func (c *userController) RootCompanyOverview(ctx *gin.Context) {
	overview, err := c.serv.RootCompanyOverview() // Call service to get the company overview
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, overview)
}
