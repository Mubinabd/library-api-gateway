package clients

import (
	"log/slog"

	pb "github.com/Mubinabd/library-api-gateway/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	BookClient     pb.BookServiceClient
	BorrowerClient pb.BorrowerServiceClient
	AuthorClient   pb.AuthorServiceClient
	GenreClient    pb.GenreServiceClient
}

func NewClients() *Clients {
	conn, err := grpc.NewClient("library:8085", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("error:", err)
	}
	authC := pb.NewAuthorServiceClient(conn)
	bookC := pb.NewBookServiceClient(conn)
	boorowerC := pb.NewBorrowerServiceClient(conn)
	genreC := pb.NewGenreServiceClient(conn)

	return &Clients{
		AuthorClient:   authC,
		BookClient:     bookC,
		BorrowerClient: boorowerC,
		GenreClient:    genreC,
	}
}
