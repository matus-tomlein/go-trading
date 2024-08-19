package main

import (
    "log"
    "net/http"
    "github.com/matus-tomlein/go-trading/internal/router"
)

func main() {
    r := router.SetupRouter()
    log.Fatal(http.ListenAndServe(":8080", r))
}
