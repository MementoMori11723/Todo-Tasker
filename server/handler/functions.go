package handler

import (
	// "fmt"
	// "go-server/database"
	"net/http"
)

var (
  Routes = make(map[string]func(http.ResponseWriter, *http.Request))
)

// need to work on these functions... 


