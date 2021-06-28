// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// RouterDao is the manager for logic model data accessing and custom defined data operations functions management.
type RouterDao struct {
	gmvc.M               // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      routerColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB        // DB is the raw underlying database management object.
	Table  string        // Table is the underlying table name of the DAO.
}

// RouterColumns defines and stores column names for table router.
type routerColumns struct {
	Id        string //
	Path      string //
	Name      string //
	Redirect  string //
	Title     string //
	Icon      string //
	Component string //
	Parent    string //
	OrderNo   string //
	Status    string //
	CreateAt  string //
	UpdateAt  string //
	DeleteAt  string //
}

// NewRouterDao creates and returns a new DAO object for table data access.
func NewRouterDao() *RouterDao {
	columns := routerColumns{
		Id:        "id",
		Path:      "path",
		Name:      "name",
		Redirect:  "redirect",
		Title:     "title",
		Icon:      "icon",
		Component: "component",
		Parent:    "parent",
		OrderNo:   "orderNo",
		Status:    "status",
		CreateAt:  "create_at",
		UpdateAt:  "update_at",
		DeleteAt:  "delete_at",
	}
	return &RouterDao{
		C:     columns,
		M:     g.DB("default").Model("router").Safe(),
		DB:    g.DB("default"),
		Table: "router",
	}
}