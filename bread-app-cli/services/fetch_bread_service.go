package services

type FetchBreadService interface {
	// パンのデータを取得し、登録・更新を行う.
	Fetch(ids []string) error
}
