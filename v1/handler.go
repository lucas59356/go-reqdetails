package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler Agrupamento de funções
type Handler struct{}

// Handle Faz o trabalho sujo
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s FROM %s\n", r.Method, r.RequestURI, r.RemoteAddr)
	user, passwd, ok := r.BasicAuth()

	data := map[string]interface{}{
		"vars": mux.Vars(r),
		"auth": map[string]interface{}{
			"user":   user,
			"passwd": passwd,
			"ok":     ok,
		},
		"content_length": r.ContentLength,
		"cookies":        r.Cookies(),
		"form":           r.Form,
		"header":         r.Header,
		"host":           r.Host,
		"method":         r.Method,
		"proto": map[string]interface{}{
			"proto": r.Proto,
			"major": r.ProtoMajor,
			"minor": r.ProtoMinor,
		},
		"remote_addr":       r.RemoteAddr,
		"request_uri":       r.RequestURI,
		"transfer_encoding": r.TransferEncoding,
		"url":               r.URL,
		"user_agent":        r.UserAgent(),
	}
	err := r.ParseForm()
	if err != nil {
		println(err.Error())
	} else {
		data["post_form"] = r.PostForm
	}

	err = r.ParseMultipartForm(64 * 1024 * 1024)
	if err != nil {
		// println(err.Error())
	} else {
		data["multipart_form"] = r.MultipartForm.Value
	}

	d, err := json.MarshalIndent(data, "", "\t")
	w.Write(d)
	if err != nil {
		println(err.Error())
	}
}
