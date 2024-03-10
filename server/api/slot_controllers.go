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

func AddSlotHandler(ctx *gin.Context) {
	var body models.AddSlotReq
	var resp models.SlotResp
	var slot models.Slot

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

	copier.Copy(&slot, body)

	_, err := c.DB.GetLotByID(body.LotID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "invalid lot id")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	lastSlotNumber, err := c.DB.GetLastSlotNumber(body.LotID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	slot.Number = lastSlotNumber + 1

	slotDb, err := c.DB.AddSlot(slot)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, slotDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "slot added successfully", Data: resp})
}

func UpdateSlotStatusHandler(ctx *gin.Context) {
	var body models.UpdateSlotStatusReq
	var resp models.SlotResp

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

	_, err := c.DB.GetSlotByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "invalid id")
			return
		}

		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if body.Status == UnavailableSlot {
		isBooked, err := c.DB.IsSlotBooked(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		if isBooked {
			ctx.JSON(http.StatusUnprocessableEntity, "slot is currently occupied")
			return
		}

	}

	slotDb, err := c.DB.UpdateSlotStatus(id, body.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, slotDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "slot status updated successfully", Data: resp})
}

func GetSlotHandler(ctx *gin.Context) {
	var resp models.SlotResp

	id := ctx.Param("id")
	if id == "" || !utils.IsValidUUID(id) {
		ctx.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	slotDb, err := c.DB.GetSlotByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "invalid id")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, slotDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "slot data fetched successfully", Data: resp})
}

func DeleteSlotHandler(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" || !utils.IsValidUUID(id) {
		ctx.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	_, err := c.DB.GetSlotByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "invalid id")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = c.DB.DeleteSlotByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err = c.DB.DeleteBookedSlotsBySlotID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Message: "slot deleted successfully", Data: nil})
}

func GetAllSlotsHandler(ctx *gin.Context) {
	var resp []models.SlotResp

	lotId := ctx.Query("lot_id")
	if lotId == "" || !utils.IsValidUUID(lotId) {
		ctx.JSON(http.StatusBadRequest, "invalid lot id")
		return
	}

	slotStatus := ""
	qStatus := ctx.Query("status")
	if qStatus == AvailableSlot || qStatus == BookedSlot || qStatus == UnavailableSlot {
		slotStatus = qStatus
	}
	slotDb, err := c.DB.GetAllSlotsByLotID(lotId, slotStatus)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, slotDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "all slots data fetched successfully", Data: resp})
}
