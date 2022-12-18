package handlers

import (
	"creativecapsule/business/scratchcard"
	"creativecapsule/handlers/web"

	"github.com/labstack/echo/v4"
)

type scratchCardHandlers struct {
	scratchcard scratchcard.ScratchCard
}

func (s scratchCardHandlers) create(c echo.Context) error {
	card := scratchcard.CreateCard{}
	err := c.Bind(&card)
	if err != nil {
		return web.BadRequest(err.Error())
	}

	info, err := s.scratchcard.Create(c.Request().Context(), card)
	if err != nil {
		return web.Error(err)
	}

	return web.OK(c, info)
}

func (s scratchCardHandlers) getActiveCard(c echo.Context) error {
	info, err := s.scratchcard.GetActiveCards(c.Request().Context(), scratchcard.Info{
		IsScratched: true,
		IsActive:    true,
	})
	if err != nil {
		return web.Error(err)
	}

	return web.OK(c, info)
}

func (s scratchCardHandlers) transaction(c echo.Context) error {
	card := scratchcard.Filter{}
	err := c.Bind(&card)
	if err != nil {
		return web.BadRequest(err.Error())
	}

	err = s.scratchcard.TransactionUpdate(c.Request().Context(), card)
	if err != nil {
		return web.Error(err)
	}

	return web.OK(c, map[string]string{"Messgae": "Successfully Transaction"})
}

func (s scratchCardHandlers) gettransaction(c echo.Context) error {
	filter := scratchcard.Filter{}
	err := c.Bind(&filter)
	if err != nil {
		return web.BadRequest(err.Error())
	}

	info, err := s.scratchcard.GetTransaction(c.Request().Context(), filter)
	if err != nil {
		return web.Error(err)
	}

	return web.OK(c, info)
}
