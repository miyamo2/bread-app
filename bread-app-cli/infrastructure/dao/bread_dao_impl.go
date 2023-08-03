package dao

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	idao "github.com/miyamo2theppl/bread-app/bread-app-cli/domain/dao"
	"github.com/miyamo2theppl/bread-app/bread-app-cli/domain/entity"
	"github.com/miyamo2theppl/bread-app/bread-app-common/core"
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

func (r *BreadDaoImpl) Upsert(entities []entity.Bread) error {
	// TODO: プロセスIDの生成、ログ出力
	log.Print("[INFO]BreadDao#Upsert START")
	defer log.Printf("[INFO]BreadDao#Upsert END")

	err := r.client.RunTransaction(r.ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		var err error
		for _, v := range entities {
			item := r.client.Collection(collectionName).Doc(v.ID)
			err = tx.Set(item, v)
			if err != nil {
				break
			}
		}
		return err
	})
	return err
}

// BreadDaoコンストラクタ
func NewBreadDao(ctx core.Context, client core.Client) idao.BreadDao {
	return &BreadDaoImpl{ctx: ctx, client: client}
}
