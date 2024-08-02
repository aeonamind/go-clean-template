package domain

import (
	"time"
)

type (
	Folder struct {
		ID        *string   `json:"id"`
		Name      string    `json:"name"`
		ParentID  int       `json:"parent_id"`
		OwnerID   int       `json:"owner_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	FolderRepository interface {
		CreateFolder(data *Folder) (*Folder, error)
		GetAllFolders(params interface{}) ([]*Folder, error)
	}

	CreateFolderUseCase interface {
		Execute(data *Folder) (*Folder, error)
	}

	GetAllFoldersUseCase interface {
		Execute(params interface{}) ([]*Folder, error)
	}
)
