package main

import (
	// "github.com/xmufeng/fibnacii"
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
	"testing"
)

func TestHttpServer(t *testing.T) {
	globalMux := http.NewServeMux()

	apiRouter := mux.NewRouter()
	apiRouter.HandleFunc("/test/{id}/{name}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		w.Write([]byte(fmt.Sprintf("id=%s; name=%s", vars["id"], vars["name"])))
		fmt.Println("Test")
	})

	globalMux.Handle("/", http.FileServer(http.Dir("static")))
	globalMux.Handle("/test/", apiRouter)

	http.ListenAndServe(":8080", context.ClearHandler(globalMux))

}
