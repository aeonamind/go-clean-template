package folder

import (
	"github.com/google/wire"
	"github.com/koujiman/gobox/domain"
	"github.com/koujiman/gobox/infrastructure/repository"
	usecase "github.com/koujiman/gobox/usecase/folder"
)

var FolderProviderSet = wire.NewSet(
	ProvideFolderController,
	usecase.ProvideCreateFolderUseCase,
	usecase.ProvideGetAllFoldersUseCase,
	repository.ProvideFolderRepository,

	wire.Bind(new(domain.FolderRepository), new(*repository.FolderRepository)),
	wire.Bind(new(domain.CreateFolderUseCase), new(*usecase.CreateFolderUseCase)),
	wire.Bind(new(domain.GetAllFoldersUseCase), new(*usecase.GetAllFoldersUseCase)),
)
