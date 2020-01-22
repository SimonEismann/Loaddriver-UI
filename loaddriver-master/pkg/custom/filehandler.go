package custom

import (
	"net/http"
)

type customFileServer struct {
	Handler http.Handler
}

func FileServer(root http.FileSystem) *customFileServer {
	return &customFileServer{Handler: http.FileServer(root)}
}

func (c *customFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.Handler.ServeHTTP(w, r)
	w.Header().Add("Cache-Control", "no-cache")
}
