namespace go push_proxy
include "base.thrift"

struct PushMessageReq{
   1: required string clientId
   2: required string event
   3: required string type
   4: required string roomId
   5: required string message
}

struct PushMessageResp{
    1: required base.BaseResp baseResp
}
service PushProxyService {
   PushMessageResp PushMessage(PushMessageReq req)
}
