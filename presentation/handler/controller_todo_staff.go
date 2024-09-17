package handler

import (
	"myproject/application/services"
	"myproject/dto"
	"myproject/transform"
	"net/http"
	"strings"
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

func (h *ToDo_StaffHandler) GetAssign(c echo.Context) error {
	todo_id := c.QueryParam("todo_id")
	staff_id := c.QueryParam("staff_id")

	count := 0
	if todo_id != "" {
		count++
	}
	if staff_id != "" {
		count++
	}
	if count > 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Only one query parameter is allowed: title, body, or date",
		})
	}

	todo_staff, err := h.ToDo_StaffService.GetAssign(todo_id, staff_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	todo_staffDto := transform.EntityToDto_ToDoStaffs(todo_staff)
	return c.JSON(http.StatusOK, todo_staffDto)
}

func (h *ToDo_StaffHandler) GetAssignByID(c echo.Context) error {
	id := c.Param("id")

	todo_staff, err := h.ToDo_StaffService.GetAssignByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	staffDto := transform.EntityToDto_ToDoStaff(todo_staff)
	return c.JSON(http.StatusOK, staffDto)
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
func (h *ToDo_StaffHandler) UpdateAssign(c echo.Context) error {
	var bodyData map[string]interface{}
	if err := c.Bind(&bodyData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	id := c.Param("id")
	todo_id, _ := bodyData["todo_id"].(string)
	staff_id, _ := bodyData["staff_id"].(string)

	AssignUpdated, err := h.ToDo_StaffService.UpdateAssign(id, todo_id, staff_id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	AssignUpdatedDto := transform.EntityToDto_ToDoStaff(AssignUpdated)
	return c.JSON(http.StatusOK, AssignUpdatedDto)
}

func (h *ToDo_StaffHandler) DeleteAssign(c echo.Context) error {
	id := c.Param("id")

	if err := h.ToDo_StaffService.DeleteAssign(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Assign ToDo deleted successfully",
	})
}
