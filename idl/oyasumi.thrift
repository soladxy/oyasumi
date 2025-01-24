// idl/oyasumi.thrift
namespace go oyasumi

struct HelloReq {
    1: string message (api.query="message")
}

struct HelloResp {
    1: required i32 st // 状态码，0: 成功
    2: required string msg
    3: Hello data
}

struct Hello {
    1: string message
}


service OyasumiService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
}
