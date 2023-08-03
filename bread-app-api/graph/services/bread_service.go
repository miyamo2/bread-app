package services

import (
	"github.com/miyamo2theppl/bread-app/bread-app-api/graph/model"
)

// BreadService Interface
type BreadService interface {
	// 全てのBreadを取得する
	GetAllBreads() ([]*model.Bread, []error)
	// IDから一意にBreadを取得する
	GetBreadById(id string) (*model.Bread, error)
}
