package repository

import (
	"gorm.io/gorm"

	"github.com/koujiman/gobox/domain"
)

type FolderRepository struct {
	Db *gorm.DB
}

func ProvideFolderRepository(db *gorm.DB) *FolderRepository {
	return &FolderRepository{db}
}

func (r *FolderRepository) CreateFolder(f *domain.Folder) (*domain.Folder, error) {
	result := r.Db.Create(f)
	return f, result.Error
}

func (r *FolderRepository) GetAllFolders(params any) ([]*domain.Folder, error) {
	var folders []*domain.Folder
	r.Db.Find(folders)

	return folders, nil
}
