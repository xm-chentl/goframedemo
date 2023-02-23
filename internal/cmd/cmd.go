package cmd

import (
	"context"
	"fmt"
	"io"
	"reflect"
	"strings"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	_ "github.com/xm-chentl/goframedemo/api"
	"github.com/xm-chentl/goframedemo/internal/controller"
	_ "github.com/xm-chentl/goframedemo/internal/logic"
	"github.com/xm-chentl/goframedemo/utility"
	"github.com/xm-chentl/goframedemo/utility/apicontainer"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
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
		Func: func(_ context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				// group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(handlers...)
			})

			s.BindHandler("/:module/:action", baseHandler(func(r *ghttp.Request) string {
				return fmt.Sprintf("%s/%s", r.Get("module"), r.Get("action"))
			}))
			s.BindHandler("/mobile/:module/:action", baseHandler(func(r *ghttp.Request) string {
				return fmt.Sprintf("mobile/%s/%s", r.Get("module"), r.Get("action"))
			}))
			s.Run()
			return
		},
	}
)

func baseHandler(route func(r *ghttp.Request) string) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		ctx := context.Background()
		var err error
		resp := resp{}
		// tracer, tracerCloser, err := InitJaeger()
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

			}
			// tracerCloser.Close()
			r.Response.WriteJson(resp)
		}()

		// 提取路由
		path := route(r)
		// path := fmt.Sprintf("mobile/%s/%s", r.Get("module"), r.Get("action"))
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

		// root span
		// spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		// var span opentracing.Span
		// if spanCtx != nil {
		// 	// 存在则创建子span
		// 	span = tracer.StartSpan(
		// 		path,
		// 		opentracing.ChildOf(spanCtx),
		// 	)
		// } else {
		// 	span = tracer.StartSpan(path)
		// 	err = tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		// 	if err != nil {
		// 		return
		// 	}
		// }
		// defer span.Finish()
		resp.Data, err = apiInterface.(apicontainer.APIHandler).Call(ctx)
	}
}

func InitJaeger() (tracer opentracing.Tracer, closer io.Closer, err error) {
	serviceName, _ := g.Cfg().Get(context.TODO(), "server.name")
	endPoint, _ := g.Cfg().Get(context.TODO(), "jaeger.endPoint")
	cfg := &config.Configuration{
		ServiceName: serviceName.String(),
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: endPoint.String(),
		},
	}
	tracer, closer, err = cfg.NewTracer(config.Logger(jaeger.StdLogger))
	return
}
