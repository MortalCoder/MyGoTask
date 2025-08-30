package service

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// создание репорта
func (s *Service) CreateReport(c echo.Context) error {
	var req Report
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := s.reportsRepo.CreateReport(req.Title, req.Description); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create report"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"status": "report created"})
}

// получение репорта
func (s *Service) GetReportById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	report, err := s.reportsRepo.GetReportById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "report not found"})
	}

	return c.JSON(http.StatusOK, report)
}

// обновление репорта
func (s *Service) UpdateReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	var req Report
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := s.reportsRepo.UpdateReport(id, req.Title, req.Description); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update report"})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "report updated"})
}

// удаление репорта
func (s *Service) DeleteReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	if err := s.reportsRepo.DeleteReport(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete report"})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "report deleted"})
}
