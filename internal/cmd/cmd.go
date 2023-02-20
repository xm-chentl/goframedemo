package cmd

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/xm-chentl/goframedemo/api"
	"github.com/xm-chentl/goframedemo/internal/consts"
	"github.com/xm-chentl/goframedemo/internal/controller"
	"github.com/xm-chentl/goframedemo/utility"
	"github.com/xm-chentl/goframedemo/utility/apicontainer"

	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/xm-chentl/gocore/iocex"
)

var handlers = []interface{}{
	controller.Hello,
}

type resp struct {
	Message string      `json:"msg"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				// group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(handlers...)
			})

			serverName, _ := g.Cfg().Get(context.TODO(), "server.name")
			endPoint, _ := g.Cfg().Get(context.TODO(), "jaeger.endPoint")

			s.BindHandler("/mobile/:module/:action", func(r *ghttp.Request) {
				ctx := context.Background()
				var err error
				resp := resp{}
				traceProvider, err := jaeger.Init(serverName.String(), endPoint.String())
				if err != nil {
					return
				}

				defer func() {
					if rErr := recover(); rErr != nil {
						err = utility.NewCustomError(500, "panic: %v", rErr)
					}
					if err != nil {
						var ok bool
						var vErr utility.CustomError
						if vErr, ok = err.(utility.CustomError); !ok {
							vErr = utility.CustomError{
								Code:    500,
								Message: err.Error(),
							}
						}
						resp.Code = vErr.Code
						resp.Message = vErr.Message
						traceProvider.Shutdown(ctx)
					}
					r.Response.WriteJson(resp)
				}()

				// 提取路由
				path := fmt.Sprintf("mobile/%s/%s", r.Get("module"), r.Get("action"))
				api, ok := apicontainer.Get(path)
				if !ok {
					err = utility.NewCustomError(504, "api (%s) not exists", path)
					return
				}

				apiInterface := reflect.New(reflect.TypeOf(api).Elem()).Interface()
				// 方法
				method := gmeta.Get(api, "method")
				if !strings.EqualFold(strings.ToLower(method.String()), strings.ToLower(r.Method)) {
					err = utility.NewCustomError(504, "api (%s) not exists", path)
					return
				}
				// trace

				traceID := r.Header.Get(consts.TraceIDHeader)
				fmt.Sprintln("trace-id: ", traceID)
				ctx, span := gtrace.NewSpan(ctx, path)
				fmt.Println("span1 >>> ", span.SpanContext().TraceID().String())
				defer span.End()
				ctx, span2 := gtrace.NewSpan(ctx, path+"2")
				fmt.Println("span2 >>> ", span2.SpanContext().TraceID().String())
				defer span2.End()

				// 传入生成的traceID, 用于组件中传递
				ctx = context.WithValue(ctx, consts.TraceIDKey, span.SpanContext().TraceID())

				// 参数(赋值、校验)
				if err = r.Parse(apiInterface); err != nil {
					err = utility.NewCustomError(501, "request parameters is failed: %s", err.Error())
					return
				}
				// 注入对象
				if err = iocex.Inject(apiInterface); err != nil {
					err = utility.NewCustomError(502, "ioc inject is failed: %s", err.Error())
					return
				}
				resp.Data, err = apiInterface.(apicontainer.APIHandler).Call(ctx)
			})
			s.Run()
			return
		},
	}
)
