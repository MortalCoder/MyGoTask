package service

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetWordById ищем слово по id
// localhost:8000/api/word/:id
func (s *Service) GetWordById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.wordsRepo
	word, err := repo.RGetWordById(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, Response{Object: word})
}

// CreateWords добавляем в базу новые слова в базу
// localhost:8000/api/words
func (s *Service) CreateWords(c echo.Context) error {
	var wordSlice []Word
	err := c.Bind(&wordSlice)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.wordsRepo
	for _, word := range wordSlice {
		err = repo.CreateNewWords(word.Title, word.Translation)
	}
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}

// обновляем слово по id
// PUT localhost:8000/api/word/:id
func (s *Service) UpdateWord(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	var word Word

	if err := c.Bind(&word); err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.wordsRepo
	if err := repo.UpdateWord(id, word.Title, word.Translation); err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, Response{Object: "Word updated successfully"})
}

// удаляем слово по id
// DELETE localhost:8000/api/word/:id
func (s *Service) DeleteWord(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.wordsRepo
	if err := repo.DeleteWord(id); err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, Response{Object: "Word deleted successfully"})
}

// создание репорта
func (s *Service) CreateReport(c echo.Context) error {
	var req Report
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := s.wordsRepo.CreateReport(req.Title, req.Description); err != nil {
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

	report, err := s.wordsRepo.GetReportById(id)
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

	if err := s.wordsRepo.UpdateReport(id, req.Title, req.Description); err != nil {
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

	if err := s.wordsRepo.DeleteReport(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete report"})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "report deleted"})
}
