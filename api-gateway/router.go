package api

import (
	"github.com/Mubinabd/library-api-gateway/api-gateway/handler"
	_ "github.com/Mubinabd/library-api-gateway/docs"
	// "github.com/Mubinabd/library-api-gateway/middleware"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title Online Voting System Swagger UI
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handler.HandlerStruct) *gin.Engine {
	r := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://localhost:8090"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}
	r.Use(cors.New(corsConfig))

	author := r.Group("/author")
	{
		author.POST("/create", h.CreateAuthor)
		author.GET("/:id", h.GetAuthor)
		author.GET("", h.GetAllAuthors)
		author.PUT("/update", h.UpdateAuthor)
		author.DELETE("/:id", h.DeleteAuthor)
		author.GET("/author/:id", h.GetAuthorBooks)
	}

	book := r.Group("/book")
	{
		book.POST("/create", h.CreateBook)
		book.GET("/:title", h.GetBook)
		book.GET("/all", h.GetBooks)
		book.PUT("/update", h.UpdateBook)
		book.DELETE("/del/:id", h.DeletesBook)
		book.GET("/search", h.SearchTitleAndAuthor)
		
	}

	borrower := r.Group("/borrower")
	{
		borrower.POST("/create", h.CreateBorrower)
		borrower.GET("/:id", h.GetBorrower)
		borrower.GET("/all", h.GetBorrowers)
		borrower.PUT("/update", h.UpdateBorrower)
		borrower.DELETE("/:id", h.DeleteBorrower)
		borrower.GET("/users/:id", h.GetBorrowerBooks)
		borrower.GET("/history/:id", h.HistoryUser)
		borrower.GET("/overdue", h.GetOverdueBooks)
	}

	genre := r.Group("/genre")
	{
		genre.POST("/create", h.CreateGenre)
		genre.GET("/:name", h.GetGenre)
		genre.GET("/all", h.GetAllGenres)
		genre.PUT("/update", h.UpdateGenre)
		genre.DELETE("/:id", h.DeleteGenre)

		genre.GET("/genre/:id", h.GetBooksByGenre)
	}
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
