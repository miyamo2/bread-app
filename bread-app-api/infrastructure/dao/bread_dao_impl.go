package dao

import (
	"log"
	"strings"

	idao "github.com/miyamo2theppl/bread-app/bread-app-api/domain/dao"
	"github.com/miyamo2theppl/bread-app/bread-app-api/domain/entity"
	"github.com/miyamo2theppl/bread-app/bread-app-common/core"
	"github.com/miyamo2theppl/bread-app/bread-app-common/errors"
	"google.golang.org/api/iterator"
)

const (
	// コレクション名
	collectionName = "bread"
)

// BreadDao実装
type BreadDaoImpl struct {
	ctx    core.Context
	client core.Client
}

// See: github.com/miyamo2theppl/bread-app/bread-app-api/domain/dao/BreadDao.go#FindAll
func (r *BreadDaoImpl) FindAll() ([]entity.Bread, []error) {
	// TODO: プロセスIDの生成、ログ出力
	log.Print("[INFO]BreadDao#Upsert START")
	defer log.Printf("[INFO]BreadDao#Upsert END")

	res := make([]entity.Bread, 0)
	errs := make([]error, 0)
	iter := r.client.Collection(collectionName).Documents(r.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			errs = append(errs, err)
		}
		var ent entity.Bread
		doc.DataTo(&ent)

		res = append(res, ent)
	}

	return res, errs
}

// See: github.com/miyamo2theppl/bread-app/bread-app-api/domain/dao/BreadDao.go#FindByID
func (r *BreadDaoImpl) FindByID(id string) (*entity.Bread, error) {
	// TODO: プロセスIDの生成、ログ出力
	log.Print("[INFO]BreadDao#FindByID START")
	defer log.Printf("[INFO]BreadDao#FindByID END")

	dsnap, err := r.client.Collection(collectionName).Doc(id).Get(r.ctx)
	if err != nil {
		if strings.Contains(err.Error(), "code = NotFound") {
			return nil, errors.NewDataNotFoundError(collectionName, id)
		}
		return nil, err
	}
	var res entity.Bread
	dsnap.DataTo(&res)
	res.ID = id

	return &res, nil
}

func NewBreadDao(ctx core.Context, client core.Client) idao.BreadDao {
	return &BreadDaoImpl{ctx: ctx, client: client}
}
