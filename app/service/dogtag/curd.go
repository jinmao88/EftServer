package dogtag

import (
	"EftServer/app/dao"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

type Req struct {
	Page     int `p:"page"`
	PageSize int `p:"page_size"`
	Query
}
type Query struct {
	Id         int        `p:"id"`
	PlayerName string     `p:"player_name"`
	KillerName string     `p:"killer_name"`
	Kill_at    gtime.Time `p:"kill_at"`
	Weapon     string     `p:"weapon"`
	Camp       string     `p:"camp"`
	Level      int        `p:"level"`
	Map        string     `p:"map"`
	Uuid       int        `p:"uuid"`
}

func (r Req) List() (g.Map, error) {
	m := g.DB().Model("dogtag")
	if r.PlayerName != "" {
		m = m.Where("player_name", r.PlayerName)
	}
	if r.KillerName != "" {
		m = m.Where("killer_name", r.KillerName)
	}
	count, err := m.Count()
	if err != nil {
		return nil, err
	}
	all, err := m.Page(r.Page, r.PageSize).All()
	if err != nil {
		return nil, err
	}
	return g.Map{
		"items":    all,
		"total":    count,
		"page":     r.Page,
		"pageSize": r.PageSize,
	}, nil
}

func (r Req) Add() error {
	if _, err := dao.Dogtag.Data(r).Save(); err != nil {
		return err
	}
	return nil
}

func (r Req) Edit() error {
	panic("implement me")
}

func (r Req) Del() error {
	panic("implement me")
}
