package triemap

import (
	"net/http"
	"path"
)

func HttpHandler(c Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lookup := path.Base(r.URL.Path)
		if len(lookup) > 64 {
			http.Error(w, "bad path", 400)
			return
		}
		sig := c.Hex(lookup)
		if sig == "" {
			http.Error(w, "not found", 404)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sig))
	}
}
