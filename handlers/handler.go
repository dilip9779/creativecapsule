package handlers

import (
	"creativecapsule/business/scratchcard"
	"creativecapsule/business/users"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func API(e *echo.Echo, logger zerolog.Logger, pdb *sqlx.DB) {

	user := users.New(pdb, logger)
	scratchCard := scratchcard.New(pdb, logger, user)
	userHandlers := userHandlers{
		user: user,
	}

	e.GET("/v1/all-users", userHandlers.getAllUser)
	e.GET("/v1/active-users", userHandlers.getActiveUser)
	e.POST("/v1/user", userHandlers.create)
	e.GET("/v1/user/:id", userHandlers.get)
	e.PUT("/v1/user/:id", userHandlers.update)
	e.DELETE("/v1/user/:id", userHandlers.delete)
	e.PATCH("/v1/user/in-actived", userHandlers.inActived)
	scratchCardHandlers := scratchCardHandlers{
		scratchcard: scratchCard,
	}

	e.POST("/v1/scrach-card", scratchCardHandlers.create)
	e.GET("/v1/scrach-card/active", scratchCardHandlers.getActiveCard)
	e.POST("/v1/transaction", scratchCardHandlers.transaction)
	e.POST("/v1/transaction/all", scratchCardHandlers.gettransaction)
}
