package handler

import (
	"fmt"
	"net/http"
	"time"
)

func Date(w http.ResponseWriter, _ *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime)
}
