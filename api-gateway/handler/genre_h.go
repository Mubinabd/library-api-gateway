package handler

import (
	"context"
	pb "github.com/Mubinabd/library-api-gateway/genproto"
	"github.com/gin-gonic/gin"
)

// @Router 				/genre/create [POST]
// @Summary 			CREATE GENRE
// @Description		 	This api create genre
// @Tags 				GENRE
// @Accept 				json
// @Produce 			json
// @Param data 			body pb.GenreCreate true "Genre"
// @Success 201 		{object} pb.Genre
// @Failure 400 		string Error
func (h *HandlerStruct) CreateGenre(c *gin.Context) {
	var req pb.GenreCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	genre, err := h.Clients.GenreClient.CreateGenre(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, genre)
}

// @Router 				/genre/{name} [GET]
// @Summary 			GET GENRE
// @Description		 	This api get genre by name
// @Tags 				GENRE
// @Accept 				json
// @Produce 			json
// @Param 			    name path string true "GENRE NAME"
// @Success 200			{object} pb.Genre
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) GetGenre(c *gin.Context) {
	var req pb.ByName
	name := c.Param("name")
	req.Name = name

	genre, err := h.Clients.GenreClient.GetGenre(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, genre)
}

// @Router 				/genre/update [PUT]
// @Summary 			UPDATES GENRE
// @Description		 	This api updatedes genre
// @Tags 				GENRE
// @Accept 				json
// @Produce 			json
// @Param  genre  body pb.GenreCreate true "Genre"
// @Success 200			{object} string "genre updated successfully"
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) UpdateGenre(c *gin.Context) {
	var req pb.GenreCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := h.Clients.GenreClient.UpdateGenre(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "genre updated successfully")
}

// @Router 				/genre/{id} [DELETE]
// @Summary 			DELETES GENRE
// @Description		 	This api delete genre by id
// @Tags 				GENRE
// @Accept 				json
// @Produce 			json
// @Param 			    id path string true "GENRE ID"
// @Success 200			{object} string "genre deleted successfully"
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) DeleteGenre(c *gin.Context) {
	var req pb.ById
	id := c.Param("id")
	req.Id = id
	_, err := h.Clients.GenreClient.DeleteGenre(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "Genre deleted")
}

// @Router 				/genre/all [GET]
// @Summary 			GET ALL GENRES
// @Description		 	This api get all books by id
// @Tags 				GENRE
// @Accept 				json
// @Produce 			json
// @Param 			    name query string false "Genre Name"
// @Success 200			{object} pb.Genres
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) GetAllGenres(c *gin.Context) {
	var nameFilter pb.NameFilter
	name := c.Query("name")
	nameFilter.Name = name
	genres, err := h.Clients.GenreClient.GetAllGenres(c.Request.Context(), &nameFilter)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, genres)
}

// @Router 				/genre/genre/{id} [GET]
// @Summary 			GET GENRE BOOK
// @Description		 	This api logs book in
// @Tags 				GENRE
// @Accept 				json
// @Produce 			json
// @Param 			    id path string true "GENRE ID"
// @Success 201 		{object} pb.GenreBooks
// @Failure 400 		string Error
func(h *HandlerStruct)GetBooksByGenre(c *gin.Context){
	var req pb.GenreId
	id := c.Param("id")
	req.GenreId = id
	book, err := h.Clients.GenreClient.GetBooksByGenre(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, book)
}