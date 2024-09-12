package handler

import (
	"myproject/application/services"
	"myproject/dto"
	"myproject/transform"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lucsky/cuid"
)

type ToDo_StaffHandler struct {
	ToDo_StaffService *services.Service
}

func NewToDo_StaffHandler(todo_staffService *services.Service) *ToDo_StaffHandler {
	return &ToDo_StaffHandler{ToDo_StaffService: todo_staffService}
}

func (h *ToDo_StaffHandler) Assigntodo(c echo.Context) error {
	var todo_staffDto dto.ToDo_Staff
	if err := c.Bind(&todo_staffDto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	todo_staffDto.ID = cuid.New()
	todo_staffDto.CreatedAt = time.Now()
	todo_staffDto.UpdatedAt = time.Now()

	todo_staffEntity := transform.DtoToEntity_ToDoStaff(&todo_staffDto) //dtoからentityへ
	if err := h.ToDo_StaffService.Assigntodo(&todo_staffEntity); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, todo_staffDto)
}
