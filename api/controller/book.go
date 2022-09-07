package controller

import (
	"blog/api/service"
	"blog/models"
	"blog/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//BookController -> BookController
type BookController struct {
	service service.BookService
}

//NewBookController : NewBookController
func NewBookController(s service.BookService) BookController {
	return BookController{
		service: s,
	}
}

// GetBooks : GetBooks controller
func (b BookController) GetBooks(ctx *gin.Context) {
	var books models.Book
	data, total, err := b.service.FindAll(books)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
		return
	}
	respArr := make([]map[string]interface{}, 0, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Book result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

// AddBook : AddBook controller
func (b *BookController) AddBook(ctx *gin.Context) {
	var book models.Book
	ctx.ShouldBindJSON(&book)

	if book.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	err := b.service.Save(book)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create book")
		return
	}

	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created book")
}

//GetBook : get book by id
func (b *BookController) GetBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	var book models.Book
	book.ID = id
	foundBook, err := b.service.Find(book)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Error Finding Book")
		return
	}

	response := foundBook.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Book",
		Data:    &response})

}

//UpdateBook : get update by id
func (b BookController) UpdateBook(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var book models.Book
	book.ID = id

	bookRecord, err := b.service.Find(book)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Book with given id not found")
		return
	}
	ctx.ShouldBindJSON(&bookRecord)

	if bookRecord.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	if err := b.service.Update(bookRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Book")
		return
	}
	response := bookRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Book",
		Data:    response,
	})
}
