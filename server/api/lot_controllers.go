package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"parking/lib/utils"
	"parking/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

func AddLotHandler(ctx *gin.Context) {
	var body models.LotReq
	var resp models.LotResp
	var lot models.Lot

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

	copier.Copy(&lot, body)

	lotDb, err := c.DB.AddLot(lot)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, lotDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "lot added successfully", Data: resp})
}

func UpdateLotHandler(ctx *gin.Context) {
	var body models.LotReq
	var resp models.LotResp
	var lot models.Lot

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

	copier.Copy(&lot, body)

	_, err := c.DB.GetLotByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "invalid id")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	lotDb, err := c.DB.UpdateLot(id, lot)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, lotDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "lot updated successfully", Data: resp})
}

func GetLotHandler(ctx *gin.Context) {
	var resp models.LotResp

	id := ctx.Param("id")
	if id == "" || !utils.IsValidUUID(id) {
		ctx.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	lotDb, err := c.DB.GetLotByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "invalid id")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, lotDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "lot data fetched successfully", Data: resp})
}

func DeleteLotHandler(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" || !utils.IsValidUUID(id) {
		ctx.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	_, err := c.DB.GetLotByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "invalid id")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = c.DB.DeleteLotByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = c.DB.DeleteBookedSlotsByLotID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = c.DB.DeleteSlotsByLotID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Message: "lot deleted successfully", Data: nil})
}

func GetAllLotsHandler(ctx *gin.Context) {
	var resp []models.LotResp

	lotDb, err := c.DB.GetAllLots()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, lotDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "all lots data fetched successfully", Data: resp})
}
