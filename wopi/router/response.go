package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Json(w http.ResponseWriter, content interface{}) {
	raw, err := json.Marshal(content)
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}

	fmt.Fprintln(w, string(raw))
}
