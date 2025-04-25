package services

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/AliMumtaz001/GoTask/repositories"
	"github.com/AliMumtaz001/GoTask/utils"
)

type ResultsService struct {
	Repo *repositories.ResultsRepository
}

func NewResultsService(repo *repositories.ResultsRepository) *ResultsService {
	return &ResultsService{Repo: repo}
}

func (service *ResultsService) GetPaginatedResults(userID string, page, pageSize int) (string, int, error) {
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return "", 0, errors.New("invalid user_id")
	}

	offset := (page - 1) * pageSize

	totalRecords, err := service.Repo.GetTotalRecords(userIDInt)
	if err != nil {
		return "", 0, err
	}

	results, err := service.Repo.GetResults(userIDInt, pageSize, offset)
	if err != nil {
		return "", 0, err
	}

	if len(results) == 0 {
		return "", totalRecords, nil
	}

	jsonData, err := json.Marshal(results)
	if err != nil {
		return "", 0, err
	}

	return string(jsonData), totalRecords, nil
}

func (service *ResultsService) SaveResult(result utils.Multiples, userID int) error {
	if userID <= 0 {
		return errors.New("invalid user ID")
	}

	return service.Repo.SaveResult(result, userID)
}