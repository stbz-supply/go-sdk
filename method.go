package stbz

// var (
// 	Method.
// )
type method uint

// UtilMethod UtilMethod
type UtilMethod struct {
	GET    method
	POST   method
	PUT    method
	PATCH  method
	DELETE method
}

// Method Method
var Method = new(UtilMethod)

func init() {
	Method.GET = 1
	Method.POST = 2
	Method.PUT = 3
	Method.PATCH = 4
	Method.DELETE = 5
}
