package cache_service

import (
	"blog/pkg/e"
	"strconv"
	"strings"
)

type Tag struct {
	Id    int
	Name  string
	State int

	PageNum  int
	PageSize int
}

func (t *Tag) GetTagsKey() string {
	keys := []string{
		e.CAHCE_TAG,
		"LIST",
	}
	if t.Id >= 0 {
		keys = append(keys, strconv.Itoa(t.Id))
	}
	if t.Name != "" {
		keys = append(keys, t.Name)
	}
	if t.State >= 0 {
		keys = append(keys, strconv.Itoa(t.State))
	}
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, "_")
}
