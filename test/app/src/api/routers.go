package api

import (
	"github.com/adolphlxm/atc"
)

func init() {
	v1 := atc.Route.Group("V1")
	{
		// V1版本过滤器, 根据路由规则加载。
		// 支持三种过滤器：
		//      1. EFORE_ROUTE                    //匹配路由之前
		//      2. BEFORE_HANDLER                 //匹配到路由后,执行Handler之前
		//      3. AFTER                          //执行完所有逻辑后
		//v1.AddFilter(atc.BEFORE_ROUTE, ".*", V1.AfterAuth)

		// V1版本测试
		v1.AddRouter("api", &ApiHandler{})
		v1.AddRouter("api2.{userid:[0-9]*}", &Api2Handler{})
		v1.AddRouter("api2.test", &Api2TestHandler{})
	}

}
