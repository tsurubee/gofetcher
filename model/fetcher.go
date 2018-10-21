package model

type Fetcher interface {
	Fetch(string) error
}
