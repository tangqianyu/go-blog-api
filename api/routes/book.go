package routes

import (
	"blog/api/controller"
	"blog/infrastructure"
	"blog/middleware/jwt"
)

//BookRoute BookRoute -> Route for question module
type BookRoute struct {
	Controller controller.BookController
	Handler    infrastructure.GinRouter
}

//NewBookRoute -> initializes new choice routes
func NewBookRoute(
	controller controller.BookController,
	handler infrastructure.GinRouter,

) BookRoute {
	return BookRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (b BookRoute) Setup() {
	book := b.Handler.Gin.Group("/books")
	book.Use(jwt.JwtAuthMiddleware())
	{
		book.GET("/", b.Controller.GetBooks)
		book.POST("/", b.Controller.AddBook)
		book.GET("/:id", b.Controller.GetBook)
		book.PUT("/:id", b.Controller.UpdateBook)
	}
}
