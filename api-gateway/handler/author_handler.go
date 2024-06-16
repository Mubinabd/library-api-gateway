package handler

import (
	"context"

	pb "github.com/Mubinabd/library-api-gateway/genproto"

	"github.com/gin-gonic/gin"
)

// @Router 				/author/create [POST]
// @Summary 			CREATE AUTHOR
// @Description		 	This api create author
// @Tags 				AUTHOR
// @Accept 				json
// @Produce 			json
// @Param data 			body pb.AuthorCreate true "Author"
// @Success 201 		{object} pb.Author
// @Failure 400 		string Error
func (h *HandlerStruct) CreateAuthor(c *gin.Context) {
	var req pb.AuthorCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	author, err := h.Clients.AuthorClient.CreateAuthor(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, author)
}

// GetAuthor handles GET requests to retrieve an author by ID.
// @Router /author/{id} [GET]
// @Summary GET AUTHOR
// @Description This API gets an author by ID
// @Tags AUTHOR
// @Accept json
// @Produce json
// @Param id path string true "AUTHOR ID"
// @Success 200 {object} pb.Author
// @Failure 400  string Error
// @Failure 404  string Error
func (h *HandlerStruct) GetAuthor(c *gin.Context) {
	var req pb.ById
	id := c.Param("id")
	req.Id = id

	author, err := h.Clients.AuthorClient.GetAuthor(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, author)
}

// @Router 				/author/update [PUT]
// @Summary 			UPDATES AUTHOR
// @Description		 	This api updatedes author
// @Tags 				AUTHOR
// @Accept 				json
// @Produce 			json
// @Param  author  body pb.Author true "Author"
// @Success 200			{object} string "author updated successfully"
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) UpdateAuthor(c *gin.Context) {
	var req pb.AuthorCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Clients.AuthorClient.UpdateAuthor(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"Error when updating author": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "author updated successfully"})
}

// @Router 				/author/{id} [DELETE]
// @Summary 			DELETE AUTHOR
// @Description		 	This api deletes author by id
// @Tags 				AUTHOR
// @Accept 				json
// @Produce 			json
// @Param 			    id path string true "AUTHOR ID"
// @Success 200			{object} string "author deleted successfully"
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) DeleteAuthor(c *gin.Context) {
	var req pb.ById
	id := c.Param("id")

	req.Id = id
	_, err := h.Clients.AuthorClient.DeleteAuthor(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "Author deleted")
}

// GetAllAuthors handles GET requests to retrieve authors based on a name filter.
// @Router /author [GET]
// @Summary Get All Authors
// @Description This API retrieves all authors based on a name filter
// @Tags AUTHOR
// @Accept json
// @Produce json
// @Param name query string false "Author Name"
// @Success 200 {object} pb.Authors
// @Failure 400  string Error
// @Failure 404  string Error
func (h *HandlerStruct) GetAllAuthors(c *gin.Context) {
	var namefilter pb.NameFilter
	name := c.Query("name")
	namefilter.Name = name

	authors, err := h.Clients.AuthorClient.GetAllAuthors(c.Request.Context(), &namefilter)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, authors)
}

// @Router 				/author/author/{id} [GET]
// @Summary 			GET AUTHOR BOOK
// @Description		 	This api logs book in
// @Tags 				AUTHOR
// @Accept 				json
// @Produce 			json
// @Param 			    id path string true "AUTHOR ID"
// @Success 201 		{object} pb.UserBook
// @Failure 400 		string Error
func (h *HandlerStruct)GetAuthorBooks(c *gin.Context) {
	authorID := c.Param("id")

	req := &pb.AuthorID{AuthorId: authorID}

	books, err := h.Clients.AuthorClient.GetAuthorBooks(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if len(books.Books) == 0 {
		c.JSON(404, gin.H{"error": "No books found for the specified author"})
		return
	}

	c.JSON(200, books)
}
