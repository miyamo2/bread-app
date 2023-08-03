package contentful

// Contentfulクライアント Interface
type ContentfulClient interface {
	// IDで一意にEntryの取得を行う.
	GetEntry(id string) (*ContentfulEntry, error)
}
