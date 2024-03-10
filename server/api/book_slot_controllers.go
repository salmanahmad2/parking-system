package api

import (
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"parking/lib/utils"
	"parking/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

func AddBookSlotHandler(ctx *gin.Context) {
	var body models.AddBookSlotReq
	var resp models.AddBookSlotResp
	var book models.BookSlot

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

	copier.Copy(&book, body)

	_, err := c.DB.GetLotByID(body.LotID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "invalid lot id")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	slotDb, err := c.DB.GetNearestSlotByLotID(body.LotID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "no available slots")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	book.SlotID = slotDb.ID

	bookDb, err := c.DB.AddBookSlot(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	copier.Copy(&resp, bookDb)
	resp.SlotNumber = slotDb.Number

	_, err = c.DB.UpdateSlotStatus(bookDb.SlotID, BookedSlot)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "failed to update slot status")
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Message: "slot booked successfully", Data: resp})
}

func UpdateBookSlotHandler(ctx *gin.Context) {
	var resp models.GetBookSlotResp
	var update models.BookSlot

	id := ctx.Param("id")
	if id == "" || !utils.IsValidUUID(id) {
		ctx.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	getBook, err := c.DB.GetBookSlotByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "invalid id")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !getBook.IsParked {
		ctx.JSON(http.StatusUnprocessableEntity, "vehicle already unparked")
		return
	}

	lotDb, err := c.DB.GetLotBySlotID(getBook.SlotID)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	parkTime := getBook.CreatedAt
	hourlyFee := lotDb.HourlyRate
	currentTime := time.Now()

	utcTime := time.Date(
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second(),
		currentTime.Nanosecond(),
		time.UTC)

	update.EndedAt = &utcTime

	parkingFee := calculateFee(parkTime, utcTime, hourlyFee)
	update.BillAmount = &parkingFee
	update.IsParked = false

	bookDb, err := c.DB.UpdateBookSlot(id, update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	copier.Copy(&resp, bookDb)

	_, err = c.DB.UpdateSlotStatus(bookDb.SlotID, AvailableSlot)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "failed to update slot status")
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Message: "booked slot updated successfully", Data: resp})
}

func GetBookSlotHandler(ctx *gin.Context) {
	var resp models.GetBookSlotResp

	id := ctx.Param("id")
	if id == "" || !utils.IsValidUUID(id) {
		ctx.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	bookDb, err := c.DB.GetBookSlotByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "invalid id")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, bookDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "book slot data fetched successfully", Data: resp})
}

func DeleteBookSlotHandler(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" || !utils.IsValidUUID(id) {
		ctx.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	bookDb, err := c.DB.GetBookSlotByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "invalid id")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = c.DB.DeleteBookSlotByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	_, err = c.DB.UpdateSlotStatus(bookDb.SlotID, AvailableSlot)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "failed to update slot status")
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Message: "booked slot deleted successfully", Data: nil})
}

func GetBookedSlotsStatsHandler(ctx *gin.Context) {
	var resp models.GetBookStatsResp

	day := ctx.Query("day")
	if day == "" {
		ctx.JSON(http.StatusBadRequest, "invalid day")
		return
	}

	parsedDay, err := time.Parse("2006-01-02", day)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	startOfDay := time.Date(parsedDay.Year(), parsedDay.Month(), parsedDay.Day(), 0, 0, 0, 0, parsedDay.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	bookDb, err := c.DB.GetBookedSlotsStats(startOfDay, endOfDay)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, bookDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "booked slots stats data fetched successfully", Data: resp})
}

func GetAllBookSlotsHandler(ctx *gin.Context) {
	var resp []models.BookSlot

	lotId := ctx.Query("lot_id")
	if lotId == "" || !utils.IsValidUUID(lotId) {
		ctx.JSON(http.StatusBadRequest, "invalid lot id")
		return
	}

	bookSlotsDb, err := c.DB.GetAllBookSlotsByLotID(lotId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, bookSlotsDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "all parked vehicles data fetched successfully", Data: resp})
}

func calculateFee(startTime, endTime time.Time, hourlyRate float64) float64 {
	diff := endTime.Sub(startTime)
	totalHours := math.Ceil(diff.Hours())
	totalFee := totalHours * hourlyRate

	return totalFee
}
