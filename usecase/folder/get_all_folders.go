package usecase

import "github.com/koujiman/gobox/domain"

type GetAllFoldersUseCase struct {
	repo domain.FolderRepository
}

func ProvideGetAllFoldersUseCase(repo domain.FolderRepository) *GetAllFoldersUseCase {
	return &GetAllFoldersUseCase{repo: repo}
}

func (u *GetAllFoldersUseCase) Execute(params interface{}) ([]*domain.Folder, error) {
	return u.repo.GetAllFolders(params)
}
