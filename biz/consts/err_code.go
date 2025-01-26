package consts

import "fmt"

const (
	MinErrCode = 10000
)

var (
	ParamErr      = NewRespError(MinErrCode+1, "参数错误")
	DownstreamErr = NewRespError(MinErrCode+2, "下游异常")
)

type RespError struct {
	St  int32  `thrift:"st,1,required" form:"st,required" json:"st,required" query:"st,required"`
	Msg string `thrift:"msg,2,required" form:"msg,required" json:"msg,required" query:"msg,required"`
}

func (e RespError) Error() string {
	return fmt.Sprintf("status: %d, message: %s", e.St, e.Msg)
}

func (e RespError) AppendMsg(s string) *RespError {
	return &RespError{e.St, fmt.Sprintf("%s, %s", e.Msg, s)}
}

func NewRespError(code int32, msg string) *RespError {
	return &RespError{code, msg}
}
