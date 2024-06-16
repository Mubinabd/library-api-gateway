package handler

import (
	"context"
	pb "github.com/Mubinabd/library-api-gateway/genproto"
	"github.com/gin-gonic/gin"
)

// @Router 				/book/create [POST]
// @Summary 			CREATE BOOK
// @Description		 	This api create book
// @Tags 				BOOK
// @Accept 				json
// @Produce 			json
// @Param data 			body pb.BookCreate true "Book"
// @Success 201 		{object} pb.Book
// @Failure 400 		string Error
func (h *HandlerStruct) CreateBook(c *gin.Context) {

	var req pb.BookCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	book, err := h.Clients.BookClient.CreateBook(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, book)
}

// @Router 				/book/{title} [GET]
// @Summary 			GET BOOK
// @Description		 	This api get book by title
// @Tags 				BOOK
// @Accept 				json
// @Produce 			json
// @Param 			    title path string true "BOOK TITLE"
// @Success 200			{object} pb.Book
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) GetBook(c *gin.Context) {
	var req pb.ByTitle
	title := c.Param("title")
	req.Title = title
	book, err := h.Clients.BookClient.GetBook(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, book)
}

// @Router 				/book/all [GET]
// @Summary 			GET ALL BOOKS
// @Description		 	This api get all books
// @Tags 				BOOK
// @Accept 				json
// @Produce 			json
// @Param 			    title query string false "Book Title"
// @Success 200			{object} pb.TitleFilter
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) GetBooks(c *gin.Context) {
	var titleFilter pb.TitleFilter
	title := c.Query("title")
	titleFilter.Title = title
	book, err := h.Clients.BookClient.GetAllBooks(c.Request.Context(), &titleFilter)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, book)
}

// @Router 				/book/update [PUT]
// @Summary 			UPDATES BOOK
// @Description		 	This api updatedes book
// @Tags 				BOOK
// @Accept 				json
// @Produce 			json
// @Param  book  body pb.BookCreate true "Book"
// @Success 200			{object} string "book updated successfully"
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) UpdateBook(c *gin.Context) {
	var req pb.BookCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Clients.BookClient.UpdateBook(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"Error when updating book": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "book successfully updated"})
}

// @Router 				/book/del/{id} [DELETE]
// @Summary 			DELETE BOOK
// @Description		 	This api logs book in
// @Tags 				BOOK
// @Accept 				json
// @Produce 			json
// @Param 			    id path string true "BOOK ID"
// @Success 201 		{object} pb.Void
// @Failure 400 		string Error
func (h *HandlerStruct) DeletesBook(c *gin.Context) {
	var req pb.ById
	id := c.Param("id")
	req.Id = id
	_, err := h.Clients.BookClient.DeleteBook(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "Book deleted")
}

//@Router /book/search [GET]
// @Summary Search Books
// @Description This API searches for books by title or author.
// @Tags BOOK
// @Accept json
// @Produce json
// @Param title query string false "Book Title"
// @Param author query string false "Author Name"
// @Success 200 {object} pb.Books
// @Failure 400 string Error
// @Failure 404 string Error
func(h *HandlerStruct)SearchTitleAndAuthor(c *gin.Context){
	var req pb.Search
	title := c.Query("title")
	req.Title = title
	author := c.Query("author")
	req.Author = author
	book, err := h.Clients.BookClient.SearchTitleAndAuthor(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, book)
}