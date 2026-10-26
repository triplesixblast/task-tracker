package app

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"

	"task-tracker/internal/dto"
	"task-tracker/internal/storage"
)

type taskHandler struct {
	store storage.Storage
}

func BuildApp() *echo.Echo {
	e := echo.New()

	h := &taskHandler{store: storage.NewMemoryStorage()}

	e.POST("/tasks", h.create)
	e.GET("/tasks/:id", h.get)
	e.PUT("/tasks/:id", h.update)
	e.DELETE("/tasks/:id", h.delete)

	return e
}

func (h *taskHandler) create(c *echo.Context) error {
	var task dto.Task
	if err := c.Bind(&task); err != nil {
		return err
	}

	created, err := h.store.CreateTask(c.Request().Context(), task)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, created)
}

func (h *taskHandler) get(c *echo.Context) error {
	id, err := parseID(c)
	if err != nil {
		return err
	}

	task, err := h.store.GetTask(c.Request().Context(), id)
	if err != nil {
		return mapError(err)
	}

	return c.JSON(http.StatusOK, task)
}

func (h *taskHandler) update(c *echo.Context) error {
	id, err := parseID(c)
	if err != nil {
		return err
	}

	var task dto.Task
	if err := c.Bind(&task); err != nil {
		return err
	}

	if err := h.store.UpdateTask(c.Request().Context(), id, task); err != nil {
		return mapError(err)
	}

	updated, err := h.store.GetTask(c.Request().Context(), id)
	if err != nil {
		return mapError(err)
	}

	return c.JSON(http.StatusOK, updated)
}

func (h *taskHandler) delete(c *echo.Context) error {
	id, err := parseID(c)
	if err != nil {
		return err
	}

	if err := h.store.DeleteTask(c.Request().Context(), id); err != nil {
		return mapError(err)
	}

	return c.NoContent(http.StatusNoContent)
}

func parseID(c *echo.Context) (int64, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return 0, echo.ErrBadRequest.Wrap(err)
	}

	return id, nil
}

func mapError(err error) error {
	if errors.Is(err, storage.ErrTaskNotFound) {
		return echo.ErrNotFound.Wrap(err)
	}

	return err
}
