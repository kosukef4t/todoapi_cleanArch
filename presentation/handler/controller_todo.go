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

type ToDoHandler struct {
	ToDoService *services.Service
}

func NewToDoHandler(todoService *services.Service) *ToDoHandler {
	return &ToDoHandler{ToDoService: todoService}
}

func (h *ToDoHandler) GetByID(c echo.Context) error {
	id := c.Param("id")

	todo, err := h.ToDoService.GetByID(id)
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

	todoDto := transform.Entity_To_DTO(todo)
	return c.JSON(http.StatusOK, todoDto)
}

func (h *ToDoHandler) GetToDo(c echo.Context) error {
	title := c.QueryParam("title")
	body := c.QueryParam("body")
	startDate := c.QueryParam("startdate")
	endDate := c.QueryParam("enddate")

	count := 0
	if title != "" {
		count++
	}
	if body != "" {
		count++
	}
	if startDate != "" || endDate != "" {
		count++
	}

	if count > 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Only one query parameter is allowed: title, body, or date",
		})
	}

	if (startDate != "" && endDate == "") || (startDate == "" && endDate != "") {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid date range",
		})
	}

	todos, err := h.ToDoService.GetToDo(title, body, startDate, endDate)
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

	todoDto := transform.EntityToDTO(todos)
	return c.JSON(http.StatusOK, todoDto)
}

func (h *ToDoHandler) CreateToDo(c echo.Context) error {
	var todoDto dto.ToDo
	if err := c.Bind(&todoDto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	todoDto.ID = cuid.New()
	todoDto.CreatedAt = time.Now()
	todoDto.UpdatedAt = time.Now()

	if todoDto.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Title is required",
		})
	}

	todoEntity := transform.DtoToEntity(&todoDto)
	if err := h.ToDoService.CreateToDo(&todoEntity); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, todoDto)
}

func (h *ToDoHandler) UpdateToDo(c echo.Context) error {
	var bodyData map[string]interface{}
	if err := c.Bind(&bodyData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	id := c.Param("id")
	title, _ := bodyData["title"].(string)
	body, _ := bodyData["body"].(string)
	duedate, _ := bodyData["duedate"].(string)
	completedAt, _ := bodyData["completedAt"].(string)

	todoUpdated, err := h.ToDoService.UpdateTodo(id, title, body, duedate, completedAt)
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

	todoUpdatedDto := transform.Entity_To_DTO(todoUpdated)
	return c.JSON(http.StatusOK, todoUpdatedDto)
}

type Response struct {
	Message string `json:"message"`
}

func (h *ToDoHandler) CompletedToDo(c echo.Context) error {
	id := c.Param("id")

	todo, err := h.ToDoService.CompletedToDo(id)
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

	todoDto := transform.Entity_To_DTO(todo)
	time_now := time.Now()

	var message string
	if todoDto.CompletedAt.Valid && todoDto.CompletedAt.Time.Before(time_now) {
		message = "タスクは既に完了しました！"
	} else {
		message = "タスクはまだ終わってません！"
	}

	response := Response{Message: message}
	return c.JSON(http.StatusOK, response)
}

func (h *ToDoHandler) DuplicateToDo(c echo.Context) error {
	id := c.Param("id")

	todo, err := h.ToDoService.DuplicateToDo(id)
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

	todoDto := transform.Entity_To_DTO(todo)
	return c.JSON(http.StatusOK, todoDto)
}

func (h *ToDoHandler) DeleteToDo(c echo.Context) error {
	id := c.Param("id")

	if err := h.ToDoService.DeleteToDoByID(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Staff deleted successfully",
	})
}
