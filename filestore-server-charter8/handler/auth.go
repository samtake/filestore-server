package handler

import (
	"filestore-server/common"
	"filestore-server/util"
	"net/http"
)

// HTTPInterceptor : http请求拦截器
func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			username := r.Form.Get("username")
			token := r.Form.Get("token")

			//验证登录token是否有效
			if len(username) < 3 || !IsTokenValid(token) {
				// token校验失败则跳转到直接返回失败提示
				resp := util.NewRespMsg(
					int(common.StatusInvalidToken),
					"token无效",
					nil,
				)
				w.Write(resp.JSONBytes())
				return
			}
			h(w, r)
		})
}
