package links

import "time"

type SingleLinkDto struct {
	Id         int       `json:"id"`
	Key        string    `json:"key"`
	Url        string    `json:"url"`
	UserId     int       `json:"userId"`
	CreateTime time.Time `json:"createTime"`
	TagIds     []int     `json:"tagIds"`
}
