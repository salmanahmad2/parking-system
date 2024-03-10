package api

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	r.POST("/user", AddUserHandler)
	r.PUT("/user/:id", UpdateUserHandler)
	r.GET("/user/:id", GetUserHandler)
	r.DELETE("/user/:id", DeleteUserHandler)
	r.GET("/user", GetAllUsersHandler)

	r.POST("/lot", AddLotHandler)
	r.PUT("/lot/:id", UpdateLotHandler)
	r.GET("/lot/:id", GetLotHandler)
	r.DELETE("/lot/:id", DeleteLotHandler)
	r.GET("/lot", GetAllLotsHandler)

	r.POST("/slot", AddSlotHandler)
	r.PATCH("/slot/:id", UpdateSlotStatusHandler)
	r.GET("/slot/:id", GetSlotHandler)
	r.DELETE("/slot/:id", DeleteSlotHandler)
	r.GET("/slot", GetAllSlotsHandler)

	r.POST("/book", AddBookSlotHandler)
	r.PATCH("/book/:id", UpdateBookSlotHandler)
	r.GET("/book/:id", GetBookSlotHandler)
	r.DELETE("/book/:id", DeleteBookSlotHandler)
	r.GET("/book", GetAllBookSlotsHandler)
	r.GET("/book/stats", GetBookedSlotsStatsHandler)

}
