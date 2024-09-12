package handler

import (
	"fmt"
	"myproject/application/services"
	"myproject/dto"
	"myproject/transform"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lucsky/cuid"
)

type StaffHandler struct {
	StaffService *services.Service
}

func NewStaffHandler(staffService *services.Service) *StaffHandler {
	return &StaffHandler{StaffService: staffService}
}

func (h *StaffHandler) GetStaff(c echo.Context) error {
	name := c.QueryParam("name")
	role := c.QueryParam("role")

	count := 0
	if name != "" {
		count++
	}
	if role != "" {
		count++
	}
	if count > 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Only one query parameter is allowed: title, body, or date",
		})
	}

	staffs, err := h.StaffService.GetStaff(name, role)
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	staffDtos := transform.EntityToDTO_Staffs(staffs)
	return c.JSON(http.StatusOK, staffDtos)
}

func (h *StaffHandler) GetByStaff_ID(c echo.Context) error {
	id := c.Param("id")

	staff, err := h.StaffService.GetByStaff_ID(id)
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

	staffDto := transform.EntityToDTO_Staff(staff)
	return c.JSON(http.StatusOK, staffDto)
}

func (h *StaffHandler) CreateStaff(c echo.Context) error {
	var staffDto dto.Staff
	if err := c.Bind(&staffDto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	staffDto.ID = cuid.New()
	staffDto.CreatedAt = time.Now()
	staffDto.UpdatedAt = time.Now()

	staffEntity := transform.DtoToEntity_Staff(&staffDto) //dtoからentityへ
	if err := h.StaffService.CreateStaff(&staffEntity); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, staffDto)
}

func (h *StaffHandler) UpdateStaff(c echo.Context) error {
	var bodyData map[string]interface{}
	if err := c.Bind(&bodyData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	id := c.Param("id")
	name, _ := bodyData["name"].(string)
	role, _ := bodyData["role"].(string)

	// スタッフを更新
	staffUpdated, err := h.StaffService.UpdateStaff(id, name, role)
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

	// 更新されたスタッフをDTOに変換して返す
	staffUpdatedDto := transform.EntityToDTO_Staff(staffUpdated)
	return c.JSON(http.StatusOK, staffUpdatedDto)
}

func (h *StaffHandler) DeleteStaff(c echo.Context) error {
	id := c.Param("id")

	if err := h.StaffService.DeleteStaff(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Staff deleted successfully",
	})
}
