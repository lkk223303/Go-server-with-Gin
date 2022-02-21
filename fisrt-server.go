// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/julienschmidt/httprouter"
// )

// func main() {

//     // ********  Chapter First Web app ************
//     // Listen to the root path of the web app
//     // http.HandleFunc("/", handler)

//     // Start a web server.
//     // http.ListenAndServeTLS(":8080","cert.pem", "key.pem", nil)
//     // http.ListenAndServe(":8080", nil)

//     // ********  Chapter 路由Route *****************
//     // Create a new HTTP request multiplexer
//     // mux := http.NewServeMux()

// 	// // Set routes to their handlers
//     // mux.HandleFunc("/",handler)

//     // server := http.Server{
//     //     Addr: "0.0.0.0:8080",
//     //     Handler: mux,
//     // }

//     // server.ListenAndServe()

//     // ********  Chapter 動態路由 使用第三方路由物件 *****************
//     // mux:=httprouter.New()

//     // mux.GET("/hello/:name",hello)

//     // server := http.Server{
//     //     Addr: "0.0.0.0:8080",
//     //     Handler: mux,
//     // }
//     // server.ListenAndServe()
// }

// The handler for the root path.
// func handler(w http.ResponseWriter, request *http.Request) {
//     fmt.Fprintf(w, "Hello, World!!!!!")
// }

// func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//     fmt.Fprintf(w, "Hello, %s!", p.ByName("name"))
// }

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexData struct {
	Title   string
	Content string
}

func test(c *gin.Context) {
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "我的第一個首頁"
	c.HTML(http.StatusOK, "index.html", gin.H{"data":"Hello this is data"})
}
func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/*.html")
	server.GET("/", test)
	server.Run(":8000")
}
