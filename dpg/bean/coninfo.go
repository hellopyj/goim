package bean

type ConInfo struct {
	Uid int `json:"uid"`
	Token string `json:"token"`
	Platform string `json:"platform"`
	Accepts []int32 `json:"accepts"`
	Room Room `json:"room"`
	CreateTime int64 `json:"create_time"`
}
type Room struct {
	Type string `json:"type"`
	ID int `json:"id"`
}
