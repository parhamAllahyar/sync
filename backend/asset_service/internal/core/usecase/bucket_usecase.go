package usecase

import (
	// "asset/internal/core/dto"
	// "asset/internal/core/entity"
	"asset/internal/core/port/driven/dao"
	"asset/internal/core/port/driven/storage"
	// "asset/pkg/authenticator"
	// filePkg "asset/pkg/file"
	// "fmt"
	// "github.com/google/uuid"
)

type BucketUsecase interface {
	// List(dto.ListFileRequest) (dto.ListFileResponse, error)

	// Save(dto.SaveFileRequest) (dto.Response, error)
	// Remove(dto.RemoveFileRequest) (dto.Response, error)

	// DownloadByID(dto.DownloadFileRequest) (dto.DownloadFileResponse, error)
	// DownloadByChatID(dto.DownloadFileRequest) (dto.DownloadFileResponse, error)
}

type bucketUsecase struct {
	// Logging    logging.Logging
	// Monitoring monitoring.Monitoring
	DAO     dao.FileDAO
	Storage storage.Storage
}

func NewBucketUsecase(DAO dao.FileDAO, Storage storage.Storage) BucketUsecase {
	return fileUsecase{
		// Logging:    Logging,
		// Monitoring: Monitoring,
		DAO:     DAO,
		Storage: Storage,
	}
}

// bucket List

//Bucket cors Crud

//Add lable on bucket
