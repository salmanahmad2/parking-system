package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"parking/lib/utils"
	"parking/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func AddUserHandler(ctx *gin.Context) {
	var body models.AddUserReq
	var resp models.UserResp
	var user models.User

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.Val.Validate(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = fmt.Sprintf("%s is invalid", fieldError.Field())
		}
		ctx.JSON(http.StatusBadRequest, errorMessages)
		return
	}

	copier.Copy(&user, body)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	user.Password = string(hashedPassword)

	userDb, err := c.DB.AddUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			ctx.JSON(http.StatusConflict, "email already exists")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, userDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "user added successfully", Data: resp})
}

func UpdateUserHandler(ctx *gin.Context) {
	var body models.UpdateUserReq
	var resp models.UserResp
	var user models.User

	id := ctx.Param("id")
	if id == "" || !utils.IsValidUUID(id) {
		ctx.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.Val.Validate(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = fmt.Sprintf("%s is invalid", fieldError.Field())
		}
		ctx.JSON(http.StatusBadRequest, errorMessages)
		return
	}

	copier.Copy(&user, body)

	_, err := c.DB.GetUserByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "invalid id")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userDb, err := c.DB.UpdateUser(id, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, userDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "user updated successfully", Data: resp})
}

func GetUserHandler(ctx *gin.Context) {
	var resp models.UserResp

	id := ctx.Param("id")
	if id == "" || !utils.IsValidUUID(id) {
		ctx.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	userDb, err := c.DB.GetUserByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "invalid id")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, userDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "user data fetched successfully", Data: resp})
}

func DeleteUserHandler(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" || !utils.IsValidUUID(id) {
		ctx.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	_, err := c.DB.GetUserByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "invalid id")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = c.DB.DeleteUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Message: "user deleted successfully", Data: nil})
}

func GetAllUsersHandler(ctx *gin.Context) {
	var resp []models.UserResp

	userDb, err := c.DB.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, userDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "all users data fetched successfully", Data: resp})
}
