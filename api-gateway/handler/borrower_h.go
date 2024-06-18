package handler

import (
	pb "github.com/Mubinabd/library-api-gateway/genproto"
	"github.com/gin-gonic/gin"
)

// @Router 				/admin/borrower/create [POST]
// @Summary 			CREATE BORROWER
// @Description		 	This api create borrower
// @Tags 				BORROWER
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param data 			body pb.BorrowerCreate true "Borrower"
// @Success 201 		{object} pb.Borrower
// @Failure 400 		string Error
func (h *HandlerStruct) CreateBorrower(c *gin.Context) {

	var req pb.BorrowerCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	borrower, err := h.Clients.BorrowerClient.CreateBorrower(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, borrower)
}

// @Router 				/borrower/{id} [GET]
// @Summary 			GET BORROWER
// @Description		 	This api get borrower by id
// @Tags 				BORROWER
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    id path string true "BORROWER ID"
// @Success 200			{object} pb.Borrower
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) GetBorrower(c *gin.Context) {
	var req pb.ById
	id := c.Param("id")
	req.Id = id
	borrower, err := h.Clients.BorrowerClient.GetBorrower(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, borrower)
}

// @Router 				/borrower/all [GET]
// @Summary 			GET ALL BORROWERS
// @Description		 	This api get all borrowers
// @Tags 				BORROWER
// @Accept 				json
// @Produce 			json
// @Success 200			{object} pb.Borrowers
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) GetBorrowers(c *gin.Context) {
	var req pb.Void
	borrowers, err := h.Clients.BorrowerClient.GetAllBorrowers(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, borrowers)
}

// @Router 				/admin/borrower/update [PUT]
// @Summary 			UPDATES BORROWER
// @Description		 	This api updates borrower
// @Tags 				BORROWER
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param  borrower  body pb.BorrowerCreate true "Borrower"
// @Success 200			{object} string "borrower updated successfully"
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) UpdateBorrower(c *gin.Context) {
	var req pb.BorrowerCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	borrower, err := h.Clients.BorrowerClient.UpdateBorrower(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, borrower)
}

// @Router 				/admin/borrower/{id} [DELETE]
// @Summary 			DELETES BORROWER
// @Description		 	This api deletes borrower by id
// @Tags 				BORROWER
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    id path string true "BORROWER ID"
// @Success 200			{object} string "borrower deleted successfully"
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) DeleteBorrower(c *gin.Context) {
	var req pb.ById
	id := c.Param("id")
	req.Id = id
	_, err := h.Clients.BorrowerClient.DeleteBorrower(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "Borrower deleted")
}

// @Router 				/borrower/users/{id} [GET]
// @Summary 			GET USERS BORROWER
// @Description		 	This api GETS borrower by id
// @Tags 				BORROWER
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    id path string true "USER ID"
// @Success 200			{object} pb.BorrowedBooks
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) GetBorrowerBooks(c *gin.Context) {
	var req pb.UserId
	id := c.Param("id")
	req.UserId = id
	book, err := h.Clients.BorrowerClient.BorrowerBooks(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, book)
}

// @Router /borrower/overdue [GET]
// @Summary Get Overdue Books
// @Description This API retrieves a list of books that are overdue.
// @Tags BORROWER
// @Accept json
// @Produce json
// @Security            BearerAuth
// @Success 200 {object} pb.BorrowedBooks
// @Failure 400 string Error
// @Failure 404 string Error
func (h *HandlerStruct) GetOverdueBooks(c *gin.Context) {
	req := &pb.OverdueRequest{}

	overdue, err := h.Clients.BorrowerClient.GetOverdueBooks(c.Request.Context(), req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, overdue)
}

// @Router 				/borrower/history/{id} [GET]
// @Summary 			GET USER BORROWER HISTORY
// @Description		 	This api GETS borrower by id
// @Tags 				BORROWER
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    id path string true "USER ID"
// @Success 200			{object} pb.BorrowingHistory
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) HistoryUser(c *gin.Context) {
	var req pb.UserId
	id := c.Param("id")
	req.UserId = id
	history, err := h.Clients.BorrowerClient.HistoryUser(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, history)
}
