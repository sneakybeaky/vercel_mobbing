package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func Drivers(w http.ResponseWriter, _ *http.Request) {

	drivers := []string{"Jon", "Ciaran", "Marano", "Dorothy", "Muharrem"}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(drivers), func(i, j int) { drivers[i], drivers[j] = drivers[j], drivers[i] })

	fmt.Fprintf(w, strings.Join(drivers, ", "))

}
