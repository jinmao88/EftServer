package weapon

import (
	"EftServer/app/dao"
	"github.com/gogf/gf/frame/g"
)

type Req struct {
	Page     int `p:"page"`
	PageSize int `p:"page_size"`
	Query
}
type Query struct {
	Id   int `p:"id"`
	Type int `p:"type"`
}

func (r Req) List() (g.Map, error) {
	m := g.DB().Model("weapon")
	if r.Type != 0 {
		m = m.Where("type", r.Type)
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
