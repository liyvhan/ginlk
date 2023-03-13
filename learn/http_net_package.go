package learn

import "net/http"

func http_net_01() {
	http.ListenAndServe(":80", nil)
}
