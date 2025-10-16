package main

import (
	"Q_two/database"
	"Q_two/handle"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func main() {
	mux := defaultMux()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	zap.ReplaceGlobals(logger)
	db, err := database.Connect("identifier.sqlite")
	if err != nil {
		zap.S().Fatal(err)
	}

	pathsToUrls, err := database.GetURL(db)

	if err != nil {
		zap.S().Fatal("Error getting URLs", zap.Error(err))
	}

	mapHandler := handle.MapHandler(pathsToUrls, mux)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)

	/*

			mapHandler := MapHandler(pathsToUrls, mux)

			// Build the YAMLHandler using the mapHandler as the
			// fallback
			yaml := `- path: /urlshort
		  url: https://github.com/gophercises/urlshort
		- path: /urlshort-final
		  url: https://github.com/gophercises/urlshort/tree/solution`
			yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
			if err != nil {
				panic(err)
			}


		http.ListenAndServe(":8080", yamlHandler)
	*/
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
