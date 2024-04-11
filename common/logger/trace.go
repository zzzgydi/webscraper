package logger

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type TraceLogger struct {
	ctx       *gin.Context
	IP        string
	Uid       string
	RequestId string
	RequestTs int64
	Request   any
	Response  any
	BizCode   int
	BizMsg    string
	Params    []any
	Logger    *slog.Logger
}

func NewTraceLogger(c *gin.Context) *TraceLogger {
	requestId := xid.New().String()

	return &TraceLogger{
		ctx:       c,
		RequestId: requestId,
		RequestTs: time.Now().UnixMilli(),
		Params:    make([]any, 0),
		Logger:    slog.With("reqid", requestId),
	}
}

func (t *TraceLogger) SetIp(ip string) {
	t.IP = ip
}

func (t *TraceLogger) SetUid(uid string) {
	t.Uid = uid
}

func (t *TraceLogger) SetBizRequest(req any) {
	t.Request = req
}

func (t *TraceLogger) SetBizResponse(res any) {
	t.Response = res
}

func (t *TraceLogger) Trace(key string, value any) {
	t.Params = append(t.Params, key, value)
}

func (t *TraceLogger) Tracef(key string, format string, args ...any) {
	t.Params = append(t.Params, key, fmt.Sprintf(format, args...))
}

func (t *TraceLogger) Write() {
	duration := time.Now().UnixMilli() - t.RequestTs
	msg := fmt.Sprintf("%s %s", t.ctx.Request.Method, t.ctx.Request.URL.Path)
	biz := fmt.Sprintf("%d:%s", t.BizCode, t.BizMsg)
	params := []any{"uid", t.Uid, "ip", t.IP, "biz", biz}

	if t.Request != nil {
		params = append(params, "biz_req", t.Request)
	}
	if t.Response != nil {
		params = append(params, "biz_res", t.Response)
	}

	params = append(params, "duration", duration)
	params = append(params, t.Params...)
	t.Logger.Info(msg, params...)
}
