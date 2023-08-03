package dao

import "github.com/miyamo2theppl/bread-app/bread-app-api/domain/entity"

// Bread Data Access Object Interface
type BreadDao interface {
	// 全検索.
	FindAll() ([]entity.Bread, []error)
	// IDで一意に検索する.
	FindByID(id string) (*entity.Bread, error)
}
