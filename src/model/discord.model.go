package model

type DiscordReq struct {
	Host       string
	Url        string
	Query      map[string]string
	Body       map[string]interface{}
	Token      *string
	SignDate   *float64
	ExpireDate *float64
}
