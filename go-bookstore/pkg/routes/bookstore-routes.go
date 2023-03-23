package routes

//  golang has absolute paths
import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/wickedbaba/go-projects/go-bookstore/pkg/controllers"
)

var pl = fmt.Println
var pf = fmt.Printf

var RegisterBookStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods("DELETE")

}
