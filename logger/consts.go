package logger

// Logger variables
const (
	TagPid               = "pid"
	TagTime              = "time"
	TagReferer           = "referer"
	TagProtocol          = "protocol"
	TagPort              = "port"
	TagIP                = "ip"
	TagIPs               = "ips"
	TagHost              = "host"
	TagMethod            = "method"
	TagPath              = "path"
	TagURL               = "url"
	TagUA                = "ua"
	TagLatency           = "latency"
	TagStatus            = "status"
	TagResBody           = "resBody"
	TagReqHeaders        = "reqHeaders"
	TagQueryStringParams = "queryParams"
	TagBody              = "body"
	TagBytesSent         = "bytesSent"
	TagBytesReceived     = "bytesReceived"
	TagRoute             = "route"

	TagReqHeader  = "reqHeader:"
	TagRespHeader = "respHeader:"
	TagContext    = "context:"
	TagQuery      = "query:"
	TagForm       = "form:"
	TagCookie     = "cookie:"

	defaultFormat = "${time} ${status} ${latency} ${method} ${path}"
)
