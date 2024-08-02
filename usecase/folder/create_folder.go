package usecase

import "github.com/koujiman/gobox/domain"

type CreateFolderUseCase struct {
	repo domain.FolderRepository
}

func ProvideCreateFolderUseCase(repo domain.FolderRepository) *CreateFolderUseCase {
	return &CreateFolderUseCase{repo: repo}
}

func (u *CreateFolderUseCase) Execute(data *domain.Folder) (*domain.Folder, error) {
	return u.repo.CreateFolder(data)
}
