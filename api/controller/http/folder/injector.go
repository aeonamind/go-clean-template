//go:build wireinject

package folder

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InjectFolder(*gorm.DB) *FolderController {
	panic(wire.Build(FolderProviderSet))
}
