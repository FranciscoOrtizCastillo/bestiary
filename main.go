package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	port := os.Getenv("PORT")

	var sandboxHTML = []byte(`
		<!DOCTYPE html>
		<html lang="en">
		<body style="margin: 0; overflow-x: hidden; overflow-y: hidden">
		<div id="sandbox" style="height:100vh; width:100vw;"></div>
		<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
		<script>
		new window.EmbeddedSandbox({
		target: "#sandbox",
		// Pass through your server href if you are embedding on an endpoint.
		// Otherwise, you can pass whatever endpoint you want Sandbox to start up with here.
		initialEndpoint: "http://localhost:` + os.Getenv("PORT") + `/graphql",
		});
		// advanced options: https://www.apollographql.com/docs/studio/explorer/sandbox#embedding-sandbox
		</script>
		</body>

		</html>`)

	h := handler.New(&handler.Config{
		Schema:   &BeastSchema,
		Pretty:   true,
		GraphiQL: false,
	})

	http.Handle("/graphql", h)

	http.Handle("/sandbox", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(sandboxHTML)
	}))

	//log.Println(string(sandboxHTML[:]))
	log.Println("Listening on :" + port)

	if err := http.ListenAndServe("localhost:"+port, nil); err != nil {
		log.Fatalln(err)
		panic(err)
	}

}
