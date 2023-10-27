package json

import (
	"net/http"

	"github.com/jbonadiman/webscrapers"
)

//goland:noinspection GoUnusedExportedFunction
func Handler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	if !queryParams.Has("url") || !queryParams.Has("browserlessToken") {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("the query params 'url' and 'browserlessToken' are required"))
		return
	}

	url := queryParams.Get("url")
	browserlessToken := queryParams.Get("browserlessToken")

	bundle, err := webscrapers.GetBundleData(
		browserlessToken,
		url,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	jsonBundle, err := bundle.ToJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Cache-Control", "max-age=0, s-maxage=86400")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonBundle)
}
