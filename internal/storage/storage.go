package storage

type IStorage interface {
	ShortUrl(url string) string
	FullURL(shortURL string) (string, error)
}