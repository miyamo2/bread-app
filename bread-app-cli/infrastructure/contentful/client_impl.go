package contentful

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Contentfulクライアント 実装
type contentfulClientImpl struct {
	// Contentful トークン
	token string
	// httpクライアント
	client *http.Client
}

func (r *contentfulClientImpl) GetEntry(id string) (*ContentfulEntry, error) {
	// TODO: プロセスIDの生成、ログ出力
	log.Print("[INFO]ContentfulClient#GetEntry START")
	defer log.Print("[INFO]ContentfulClient#GetEntry END")
	url := fmt.Sprintf(
		"https://cdn.contentful.com/spaces/2vskphwbz4oc/entries/%v?access_token=%v",
		id,
		r.token,
	)
	log.Printf("[DEBUG]ContentfulClient#GetEntry URL: %v", url)
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := r.client.Do(req)
	if err != nil {
		log.Printf("[WARN]ContentfulClient#GetEntry ERROR %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[WARN]ContentfulClient#GetEntry ERROR %v", err)
		return nil, err
	}

	var entry ContentfulEntry
	json.Unmarshal(body, &entry)
	return &entry, nil
}

// ContentfulClientコンストラクタ
func NewClient(token string, client *http.Client) ContentfulClient {
	return &contentfulClientImpl{token: token, client: client}
}
