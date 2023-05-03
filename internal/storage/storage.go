package storage

type IStorage interface {
	GetShortUrl(url string) string
	GetFullURL(shortURL string) string
}