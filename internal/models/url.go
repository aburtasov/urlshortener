package models

type Url struct {
	LongLink  string
	ShortLink string
}

func (u *Url) GetLongLink() string {
	return u.LongLink
}

func (u *Url) GetShortLink() string {
	return u.ShortLink
}
