package api

import (
	"github.com/Mubinabd/library-api-gateway/api-gateway/handler"
	_ "github.com/Mubinabd/library-api-gateway/docs"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title Library Swagger UI
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

	authorAdmin := r.Group("/admin/author")
	{
		authorAdmin.POST("/create", h.CreateAuthor)
		authorAdmin.PUT("/update", h.UpdateAuthor)
		authorAdmin.DELETE("/:id", h.DeleteAuthor)
	}

	author := r.Group("/author")
	{
		author.GET("/:id", h.GetAuthor)
		author.GET("", h.GetAllAuthors)
		author.GET("/author/:id", h.GetAuthorBooks)
	}

	bookAdmin := r.Group("/admin/book")
	{
		bookAdmin.POST("/create", h.CreateBook)
		bookAdmin.PUT("/update", h.UpdateBook)
		bookAdmin.DELETE("/del/:id", h.DeletesBook)

	}
	book := r.Group("/book")
	{
		book.GET("/:title", h.GetBook)
		book.GET("/all", h.GetBooks)
		book.GET("/search", h.SearchTitleAndAuthor)

	}

	borrowerAdmin := r.Group("/admin/borrower")
	{
		borrowerAdmin.POST("/create", h.CreateBorrower)
		borrowerAdmin.PUT("/update", h.UpdateBorrower)
		borrowerAdmin.DELETE("/:id", h.DeleteBorrower)
	}

	borrower := r.Group("/borrower")
	{
		borrower.GET("/:id", h.GetBorrower)
		borrower.GET("/all", h.GetAllBorrowers)
		borrower.GET("/users/:id", h.GetBorrowerBooks)
		borrower.GET("/history/:id", h.HistoryUser)
		borrower.GET("/overdue", h.GetOverdueBooks)
	}

	genreAdmin := r.Group("/admin/genre")
	{
		genreAdmin.POST("/create", h.CreateGenre)
		genreAdmin.PUT("/update", h.UpdateGenre)
		genreAdmin.DELETE("/:id", h.DeleteGenre)

	}
	genre := r.Group("/genre")
	{
		genre.GET("/:name", h.GetGenre)
		genre.GET("/all", h.GetAllGenres)
		genre.GET("/genre/:id", h.GetBooksByGenre)
	}
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
