package routes 

import (
     "github.com/gorilla/mux"
     "github.com/andoriyaprashant/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("Post")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("Get")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("Get")
    router.HandleFunc("/book/{bookId}", controllers.UpdateBook) .Methods("Put")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook) .Methods("Delete")
} 