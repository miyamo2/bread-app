package dao

import (
	"github.com/miyamo2theppl/bread-app/bread-app-cli/domain/entity"
)

// Bread Data Access Object Interface
type BreadDao interface {
	// 未登録の場合登録し、登録されている場合更新する.
	Upsert([]entity.Bread) error
}
