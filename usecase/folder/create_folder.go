package usecase

import "github.com/koujiman/go-clean-template/domain"

type CreateFolderUseCase struct {
	repo domain.FolderRepository
}

func ProvideCreateFolderUseCase(repo domain.FolderRepository) *CreateFolderUseCase {
	return &CreateFolderUseCase{repo: repo}
}

func (u *CreateFolderUseCase) Execute(data *domain.Folder) (*domain.Folder, error) {
	return u.repo.CreateFolder(data)
}
