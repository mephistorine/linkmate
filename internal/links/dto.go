package links

import "time"

type SingleLinkDto struct {
	Id         string    `json:"id"`
	Key        string    `json:"key"`
	Url        string    `json:"url"`
	UserId     int       `json:"userId"`
	CreateTime time.Time `json:"createTime"`
}
