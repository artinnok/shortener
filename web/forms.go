package web

type ShortifyForm struct {
	LongLink string `json:"link"`
}

type LongifyForm struct {
	ShortLink string `json:"link"`
}
