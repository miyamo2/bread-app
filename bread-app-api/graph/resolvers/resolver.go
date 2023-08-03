package resolvers

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"github.com/miyamo2theppl/bread-app/bread-app-api/graph/services"
	"github.com/miyamo2theppl/bread-app/bread-app-api/infrastructure/dao"
	"github.com/miyamo2theppl/bread-app/bread-app-common/core"
)

type Resolver struct{ BSvc services.BreadService }

func NewResolver() *Resolver {
	bdao := dao.NewBreadDao(core.GetContext(), core.GetClient())
	bsvc := services.NewBreadService(bdao)
	return &Resolver{BSvc: bsvc}
}
