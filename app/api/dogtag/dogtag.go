package dogtag

import (
	"EftServer/app/service/response"
	"EftServer/app/service/user"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
}

func Image(r *ghttp.Request) {
	post, err := ghttp.NewClient().SetContentType("application/json").Post("http://47.116.71.73:8866/predict/ocr_system", g.Map{"images": []string{r.GetString("image")}})
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	res := gjson.New(post.ReadAllString())
	if err != nil {
		response.JsonExit(r, 2, err.Error())

	}
	response.JsonExit(r, 0, "", res.GetArray("results")[0])

}

func (c *Controller) Info(r *ghttp.Request) {
	var req *user.Req
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	//if err != nil {
	//	response.JsonExit(r, 2, err.Error())
	//}
	response.JsonExit(r, 0, "", g.Map{"username": "vben", "roles": []string{"admin"}})
}
