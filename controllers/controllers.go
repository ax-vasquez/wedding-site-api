package controllers

import (
	"os"
	"time"

	docs "github.com/ax-vasquez/wedding-site-api/docs"
	"github.com/ax-vasquez/wedding-site-api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@securityDefinitions.Bearer.type JWT
//@in header
//@name Authorization

// @BasePath /api/v1

func paveRoutes() *gin.Engine {
	r := gin.Default()

	corsOrigin := os.Getenv("CORS_ORIGIN")
	if gin.Mode() == "debug" && corsOrigin == "" {
		corsOrigin = "*"
	}
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{corsOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Disable trusting all proxies for now since there aren't any concerns around using a load balancer (app is small scale).
	r.SetTrustedProxies(nil)
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")

	// Routes without auth middleware (these are used to set/update the user's token, used by the auth middleware)
	{
		v1.POST("/signup", Signup)
		v1.POST("/login", Login)
	}

	// Routes for obtaining full or partial data sets for the base data types (admin-only)
	resourceRoutesV1 := v1.Group("")
	{
		resourceRoutesV1.Use(middleware.AuthenticateV1())
		resourceRoutesV1.GET("/entrees", GetEntrees)
		resourceRoutesV1.GET("/users", GetUsers)
		resourceRoutesV1.GET("/horsdoeuvres", GetHorsDoeuvres)
	}

	horsDoeuvresRoutesV1 := v1.Group("/horsdoeuvres")
	{
		horsDoeuvresRoutesV1.Use(middleware.AuthenticateV1())
		horsDoeuvresRoutesV1.GET("/:id", GetHorsDoeuvres)
		horsDoeuvresRoutesV1.POST("", middleware.IsAdmin(), CreateHorsDoeuvres)
		horsDoeuvresRoutesV1.DELETE("/:id", middleware.IsAdmin(), DeleteHorsDoeuvres)
	}

	entreeRoutesV1 := v1.Group("/entree")
	{
		entreeRoutesV1.Use(middleware.AuthenticateV1())
		entreeRoutesV1.GET("/:id", GetEntrees)
		entreeRoutesV1.POST("", middleware.IsAdmin(), CreateEntree)
		entreeRoutesV1.DELETE("/:id", middleware.IsAdmin(), DeleteEntree)
	}

	userRoutesV1 := v1.Group("/user")
	{
		userRoutesV1.Use(middleware.AuthenticateV1())
		userRoutesV1.GET("", middleware.IsAdminOrLoggedInUser(), GetLoggedInUser)
		userRoutesV1.GET("/invitees", middleware.IsAdminOrLoggedInUser(), GetInviteesForLoggedInUser)
		// TODO: I've fixed the API that this was using before - it's better to have a specific "EntreeForUser" controller since GetEntrees gets one or all entrees, now
		userRoutesV1.GET("/:id/entrees", middleware.IsAdminOrLoggedInUser(), GetEntrees)
		// TODO: Same note as for entrees - should use a different controller to get hors doeuvres for a user
		userRoutesV1.GET("/:id/horsdoeuvres", middleware.IsAdminOrLoggedInUser(), GetHorsDoeuvres)
		userRoutesV1.PATCH("", middleware.IsAdminOrLoggedInUser(), UpdateLoggedInUser)
		userRoutesV1.PATCH("/update-other", middleware.IsAdmin(), AdminUpdateUser)
		userRoutesV1.PATCH("/invitees/:id", middleware.IsAdminOrLoggedInUser(), UpdateInviteeForLoggedInUser)
		userRoutesV1.POST("", middleware.IsAdmin(), CreateUser)
		userRoutesV1.POST("/add-invitee", middleware.IsAdminOrLoggedInUser(), CreateUserInvitee)
		userRoutesV1.DELETE("/:id", middleware.IsAdmin(), DeleteUser)
		userRoutesV1.DELETE("/invitees/:id", middleware.IsAdminOrLoggedInUser(), DeleteInviteeForLoggedInUser)
	}

	inviteeRoutesV1 := v1.Group("/invitee")
	{
		inviteeRoutesV1.Use(middleware.AuthenticateV1())
		inviteeRoutesV1.DELETE("/:id", middleware.IsAdmin(), DeleteInvitee)
	}

	venueGroupV1 := v1.Group("/venue")
	{
		venueGroupV1.Use(middleware.AuthenticateV1())
		venueGroupV1.GET("/reservation-link", GetHotelRoomReservationBlockLink)
	}

	return r
}

func SetupRoutes() error {
	port := os.Getenv("PORT")
	if port == "" {
		// Set to 5000 since that's what EB listens to by default
		port = "5000"
	}
	r := paveRoutes()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r.Run(":" + port)
}
