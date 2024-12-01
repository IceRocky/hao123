package api

import (
	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"hao123/db"
	"hao123/util"
	"os"
)

type Server struct {
	router *gin.Engine
	store  db.Store
	config util.Config
}

func NewServer(store db.Store, config util.Config) (*Server, error) {
	server := &Server{
		store:  store,
		config: config,
	}

	server.setUpRouter()

	return server, nil
}

// Start return the HTTP api on a specific route
func (s *Server) Start(address string) error {

	//return s.router.Run(address)
	err := s.router.Run(address)

	if err != nil {

		return err
	}
	return nil

}

// setUpRouter setup for different HTTP methods
func (s *Server) setUpRouter() {
	router := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Replace with your frontend's origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig))

	err := router.SetTrustedProxies([]string{"127.0.0.1"}) // Replace with actual trusted IPs or ranges
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to set trusted proxies")
	}
	if s.config.Environment == "development" {
		router.Use(ginzerolog.Logger("GIN"))
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	//router.POST(util.CreateUser, s.createUser)
	router.GET("/hotSearches", s.getAllHotSearch)
	router.GET("/categories", s.getAllCategories)
	router.GET("/items", s.getAllItem)
	router.GET("/itemCategories", s.getAllItemCategories)
	router.GET("/siteItem", s.getAllSiteItems)
	router.GET("/commonSiteItem", s.getAllCommonSiteItems)
	router.GET("/topNews", s.getAllTopNews)
	router.GET("/govSites", s.getAllGovSites)
	router.GET("/hotList", s.getAllHotList)
	router.GET("/hotTab", s.getAllHotTabs)
	router.GET("/topListItems", s.getAllTopListItems)
	router.GET("/topList", s.getAllTopList)
	s.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
