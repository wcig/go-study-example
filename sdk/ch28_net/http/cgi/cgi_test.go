package cgi

// net/http/cgi: 实现了RFC 3875标准的CGI接口

// 函数
// func Request() (*http.Request, error)
// func RequestFromMap(params map[string]string) (*http.Request, error)
// func Serve(handler http.Handler) error
// type Handler
//    func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request)
