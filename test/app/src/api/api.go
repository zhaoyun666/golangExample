package api

import (
	"github.com/adolphlxm/atc"
)

type ApiHandler struct {
	atc.Handler
}
func (this *ApiHandler) Get() {
	this.JSON(map[string]interface{}{"atc":"The ATC restful API runs successfully"})
}


type Api2Handler struct {
	atc.Handler
}
func (this *Api2Handler) Get() {
	userid := this.Ctx.ParamInt64("userid")
	this.JSON(map[string]interface{}{
		"atc":"The ATC restful API2 runs successfully",
		"userid":userid,
	})
}


type Api2TestHandler struct {
	atc.Handler
}
func (this *Api2TestHandler) Get() {
	this.JSON(map[string]interface{}{
		"atc":"The ATC restful API3 runs successfully",
	})
}

