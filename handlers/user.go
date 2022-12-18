package handlers

import (
	"creativecapsule/business/users"
	"creativecapsule/handlers/web"

	"github.com/labstack/echo/v4"
)

type userHandlers struct {
	user users.User
}

func (u userHandlers) create(c echo.Context) error {
	filter := users.Info{}
	err := c.Bind(&filter)
	if err != nil {
		return web.BadRequest(err.Error())
	}

	info, err := u.user.Create(c.Request().Context(), filter)
	if err != nil {
		return web.Error(err)
	}

	return web.OK(c, info)
}

func (u userHandlers) get(c echo.Context) error {
	filter := users.Info{}
	err := c.Bind(&filter)
	if err != nil {
		return web.BadRequest(err.Error())
	}

	info, err := u.user.Get(c.Request().Context(), filter.ID)
	if err != nil {
		return web.Error(err)
	}

	return web.OK(c, info)
}

func (u userHandlers) update(c echo.Context) error {
	filter := users.Info{}
	err := c.Bind(&filter)
	if err != nil {
		return web.BadRequest(err.Error())
	}

	err = u.user.Update(c.Request().Context(), filter)
	if err != nil {
		return web.Error(err)
	}

	return web.OK(c, map[string]string{"Messgae": "Successfully Updated"})
}

func (u userHandlers) delete(c echo.Context) error {
	filter := users.Info{}
	err := c.Bind(&filter)
	if err != nil {
		return web.BadRequest(err.Error())
	}

	err = u.user.Delete(c.Request().Context(), filter.ID)
	if err != nil {
		return web.Error(err)
	}

	return web.OK(c, map[string]string{"Messgae": "Successfully Deleted"})
}

func (u userHandlers) inActived(c echo.Context) error {
	ids := []uint64{}
	err := c.Bind(&ids)
	if err != nil {
		return web.BadRequest(err.Error())
	}

	err = u.user.InActived(c.Request().Context(), ids)
	if err != nil {
		return web.Error(err)
	}

	return web.OK(c, map[string]string{"Messgae": "Successfully InActived"})
}

func (u userHandlers) getAllUser(c echo.Context) error {
	info, err := u.user.GetAll(c.Request().Context())
	if err != nil {
		return web.Error(err)
	}

	return web.OK(c, info)
}

func (u userHandlers) getActiveUser(c echo.Context) error {
	info, err := u.user.GetActiveUser(c.Request().Context())
	if err != nil {
		return web.Error(err)
	}

	return web.OK(c, info)
}
