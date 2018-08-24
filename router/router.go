package router

import (
    "gitee.com/johng/gf/g"
    "gitee.com/johng/gf-home/app/ctl/doc"
    "gitee.com/johng/gf/g/net/ghttp"
)

// 统一路由注册
func init() {
    // 开发文档
    g.Server("doc").BindHandler("/*path", ctldoc.Index)
    g.Server("doc").BindHandler("/hook",  ctldoc.UpdateHook)
    g.Server("doc").EnableAdmin("/admin")
    g.Server("doc").BindHookHandler("/admin/*", ghttp.HOOK_BEFORE_SERVE, func(r *ghttp.Request) {
        user := g.Config().GetString("doc.admin.user")
        pass := g.Config().GetString("doc.admin.pass")
        if !r.BasicAuth(user, pass) {
            r.Exit()
        }
    })
}