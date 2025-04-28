package services

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strconv"

	"github.com/AliMumtaz001/GoTask/api/repositories"
	model "github.com/AliMumtaz001/GoTask/models"
	"github.com/AliMumtaz001/GoTask/utils"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type ResultsService struct {
	Repo *repositories.ResultsRepository
}

func NewResultsService(repo *repositories.ResultsRepository) *ResultsService {
	return &ResultsService{Repo: repo}
}

func (s *ResultsService) GetPaginatedResults(userID string, page, pageSize int) (string, int, error) {
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return "", 0, errors.New("invalid user_id")
	}

	offset := (page - 1) * pageSize

	totalRecords, err := s.Repo.GetTotalRecords(userIDInt)
	if err != nil {
		return "", 0, err
	}

	results, err := s.Repo.GetResults(userIDInt, pageSize, offset)
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

func (s *ResultsService) SaveResult(result model.Multiples, userID int) error {
	if userID <= 0 {
		return errors.New("invalid user ID")
	}

	return s.Repo.SaveResult(result, userID)
}

func (s *ResultsService) ProcessFile(ctx *gin.Context, file *multipart.FileHeader, email string) (model.Multiples, error) {
	// create upload dir
	if err := os.MkdirAll("upload", os.ModePerm); err != nil {
		return model.Multiples{}, errors.New("unable to create upload directory")
	}

	// save file
	path := filepath.Join("upload", file.Filename)
	if err := ctx.SaveUploadedFile(file, path); err != nil {
		return model.Multiples{}, errors.New("unable to save file")
	}

	// read content
	content, err := os.ReadFile(path)
	if err != nil {
		return model.Multiples{}, errors.New("unable to read the uploaded file")
	}

	// get user_id from email
	userID, err := s.Repo.GetUserIDByEmail(email)
	if err != nil {
		return model.Multiples{}, errors.New("failed to retrieve user ID")
	}

	// process file 
	data := string(content)
	result := utils.CombineFunc(data)

	// save result
	if err := s.Repo.SaveResult(result, userID); err != nil {
		return model.Multiples{}, errors.New("unable to save result to database")
	}

	return result, nil
}