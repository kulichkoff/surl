package entities

import "surl/pkg/random"

type Link struct {
	original  string
	shortened string
}

func NewLink(original string) *Link {
	slug := random.GetRandomString(8)

	return &Link{
		original:  original,
		shortened: slug,
	}
}

func (l *Link) GetOrinial() string {
	return l.original
}

func (l *Link) GetShotened() string {
	return l.shortened
}
