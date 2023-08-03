package services

import (
	"log"

	"github.com/miyamo2theppl/bread-app/bread-app-api/domain/dao"
	"github.com/miyamo2theppl/bread-app/bread-app-api/graph/model"
)

// BreadService実装
type breadServiceImpl struct {
	// BreadDao
	bdao dao.BreadDao
}

func (r *breadServiceImpl) GetAllBreads() ([]*model.Bread, []error) {
	// TODO: プロセスIDの生成、ログ出力
	log.Print("[INFO]BreadDao#GetAllBreads START")
	defer log.Printf("[INFO]BreadDao#GetAllBreads END")

	res := make([]*model.Bread, 0)
	ary, errs := r.bdao.FindAll()
	for _, v := range ary {
		res = append(res, &model.Bread{
			ID:        v.ID,
			Name:      v.Name,
			CreatedAt: v.CreatedAt,
		})
	}

	return res, errs
}

func (r *breadServiceImpl) GetBreadById(id string) (*model.Bread, error) {
	// TODO: プロセスIDの生成、ログ出力
	log.Print("[INFO]BreadDao#GetBreadById START")
	defer log.Printf("[INFO]BreadDao#GetBreadById END")

	entity, err := r.bdao.FindByID(id)
	var res *model.Bread = nil
	if err != nil {
		return nil, err
	}
	if entity != nil {
		res = &model.Bread{
			ID:        entity.ID,
			Name:      entity.Name,
			CreatedAt: entity.CreatedAt,
		}
	}
	return res, nil
}

// BreadServiceコンストラクタ
func NewBreadService(bdao dao.BreadDao) BreadService {
	return &breadServiceImpl{bdao: bdao}
}
