package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/lordfarshad/simplebank/db/sqlc"
)

// Server: serves HTTP requests for our banking serrvice.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creaates a new HTTP and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//register our custom validator(util/validator) with gin
	// get current validator engin that gin uses
	// this Engine() function retruns general interface type
	// and we convert this validator output to *validator.Validator pointer
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address to listen to API request
func (server *Server) Start(address string) error {
	return server.router.Run(address)
	// router.Run is private and cannot be accessed out of this api pacckge
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
