package router

import (
	"EftServer/app/service/response"
	"EftServer/app/service/router"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
}

func (c *Controller) List(r *ghttp.Request) {
	var req router.Req
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	list, err := req.List()
	if err != nil {
		response.JsonExit(r, 2, err.Error())

	}
	//if err := req.Register(); err != nil {
	//	response.JsonExit(r, 2, err.Error())
	//}
	response.JsonExit(r, 0, "", list)

}
