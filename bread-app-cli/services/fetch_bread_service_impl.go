package services

import (
	"log"
	"sync"

	idao "github.com/miyamo2theppl/bread-app/bread-app-cli/domain/dao"
	"github.com/miyamo2theppl/bread-app/bread-app-cli/domain/entity"
	"github.com/miyamo2theppl/bread-app/bread-app-cli/infrastructure/contentful"
)

type fetchBreadServiceImpl struct {
	// Contentfulクライアント
	cclient contentful.ContentfulClient
	// BreadテーブルDao
	bdao idao.BreadDao
}

// Contentfulクライアントでパンの取得を行う.
func (r *fetchBreadServiceImpl) request(id string, wg *sync.WaitGroup, reschan chan<- contentful.ContentfulEntry, errchan chan<- error) {
	// TODO: プロセスIDの生成、ログ出力
	log.Print("[INFO]FetchBreadService#request START")
	defer log.Print("[INFO]FetchBreadService#request END")
	defer wg.Done()
	resp, err := r.cclient.GetEntry(id)

	errchan <- err
	reschan <- *resp
}

func (r *fetchBreadServiceImpl) Fetch(ids []string) error {
	// TODO: プロセスIDの生成、ログ出力
	log.Print("[INFO]FetchBreadService#Fetch START")
	defer log.Printf("[INFO]FetchBreadService#Fetch END")

	tlen := len(ids)
	data := make([]entity.Bread, 0)

	wg := &sync.WaitGroup{}
	wg.Add(tlen)

	// レスポンス用チャネル
	respch := make(chan contentful.ContentfulEntry, tlen)
	// エラー用チャネル
	errch := make(chan error, tlen)

	// 並列リクエスト
	for _, v := range ids {
		go r.request(v, wg, respch, errch)
	}
	wg.Wait()
	close(respch)
	close(errch)

	// Firestore登録用データ生成
	for resp := range respch {
		sys := resp.Sys
		id := sys.ID
		if sys.Type == "Error" {
			log.Printf("[WARN]ContentfulClient#request FAILED: Cause=%v", id)
			continue
		}
		createdAt := sys.CreatedAt
		name := resp.Fields.Name
		data = append(data, entity.NewBread(id, name, createdAt.String()))
	}

	// エラーチェック
	for err := range errch {
		if err != nil {
			log.Printf("[WARN]ContentfulClient#request FAILED: Cause=%v", err.Error())
		}
	}

	err := r.bdao.Upsert(data)

	return err
}

// FetchBreadService コンストラクタ
func NewFetchBreadService(cclient contentful.ContentfulClient, bdao idao.BreadDao) FetchBreadService {
	return &fetchBreadServiceImpl{cclient: cclient, bdao: bdao}
}
