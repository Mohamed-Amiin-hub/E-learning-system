package main

import (
	"database/sql"
	"e-learning-system/internal/api/controller"
	"e-learning-system/internal/api/gateway"
	"e-learning-system/internal/api/routes"
	"e-learning-system/internal/config"
	"e-learning-system/internal/domain/service"
	"fmt"

	// utils "kaabe-app/pkg/config"

	"log"
	// "net/http"
	// "os"
	// "time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file not found, relying on system environment variables")
	}
}

func main() {
	// Load configuration
	config.LoadEnv()
	appCfg, err := config.LoadAppConfig()
	if err != nil {
		log.Fatalf("Failed to load config.yaml: %v", err)
	}

	dbCfg := config.LoadDBConfig()
	db := config.InitDB(dbCfg)
	if db == nil {
		log.Fatal("Failed to initialize the database")
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing DB: %v", err)
		}
	}()

	fmt.Printf("Server running on port %s in %s mode\n", appCfg.App.Port, appCfg.App.Env)
	var dbConn *sql.DB = db

	userRepo := gateway.NewUserRepositry(dbConn)
	tokenRepo := gateway.NewTokenRepository(dbConn)
	organizationRepo := gateway.NewOrganizationRepository(dbConn)
	organizationAdminRepo := gateway.NewOrganizationAdminRepository(dbConn)
	organizationTutorRepo := gateway.NewOrganizationTutorRepository(dbConn)
	organizationBrandingRepo := gateway.NewOrganizationBrandingRepository(dbConn)
	organizationBillingRepo := gateway.NewOrganizationBillingRepository(dbConn)

	// Initialize Services
	userService := service.NewUserService(userRepo, tokenRepo)
	organizationService := service.NewOrganizationService(organizationRepo)
	organizationAdminService := service.NewOrganizationAdminService(organizationAdminRepo)
	organizationTutorService := service.NewOrganizationTutorService(organizationTutorRepo)
	organizationBrandingService := service.NewOrganizationBrandingService(organizationBrandingRepo)
	organizationBillingService := service.NewOrganizationBillingService(organizationBillingRepo)

	// Initialize Controllers
	userController := controller.NewUserController(userService)
	organizationController := controller.NewOrganizationController(organizationService)
	organizationAdminController := controller.NewOrganizationAdminController(organizationAdminService)
	organizationTotorController := controller.NewOrganizationTutorController(organizationTutorService)
	organizationBrandingController := controller.NewOrganizationBrandingController(organizationBrandingService)
	organizationBillingController := controller.NewOrganizationBillingController(organizationBillingService)
	// Setup Gin HTTP Server
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Register API Routes
	routes.RegisterUserRoutes(r, userController, tokenRepo)
	routes.RegisterOrganizationRoutes(r, organizationController, tokenRepo)
	routes.RegisterOrganizationAdminRoutes(r, organizationAdminController, tokenRepo)
	routes.RegisterOrganizationTutorRoutes(r, organizationTotorController, tokenRepo)
	routes.RegisterOrganizationBrandingRoutes(r, organizationBrandingController, tokenRepo)
	routes.RegisterOrganizationBillingRoutes(r, organizationBillingController, tokenRepo)

	// Start Gin server (blocks here, keeps container alive)
	if err := r.Run(fmt.Sprintf(":%s", appCfg.App.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
