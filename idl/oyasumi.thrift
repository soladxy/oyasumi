// idl/oyasumi.thrift
namespace go oyasumi

struct HelloReq {
    1: string message (api.query="message")
}

struct HelloResp {
    1: required i32 st // 状态码，0: 成功
    2: required string msg
    3: HelloData data
}

struct HelloData {
    1: string now
    2: string db_test_data
}

struct CheckSSLReq {
    1: string domain (api.query="domain")
}

struct CheckSSLResp {
    1: required i32 st // 状态码，0: 成功
    2: required string msg
    3: CheckSSLData data
}

struct CheckSSLData {
    1: bool is_expired // 是否过期
}


service OyasumiService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/api/hello");

    CheckSSLResp CheckSSL(1: CheckSSLReq request) (api.get="/api/checkSsl");
}
