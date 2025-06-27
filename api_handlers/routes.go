package api_handlers

// @title Smarket API
// @version 1.0
// @description This is the Smarket API for managing sales and inventory
// @host localhost:8050
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
import (
	"Smarket/api_handlers/middleware"
	"Smarket/docs"
	"Smarket/internal/configs"
	"Smarket/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunServer() error {
	router := gin.Default()

	docs.SwaggerInfo.Title = "Smarket"
	docs.SwaggerInfo.Description = "This is a Smarket with Swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8050"
	docs.SwaggerInfo.BasePath = "/"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//http://localhost:8050/swagger/index.html

	router.GET("/ping", Ping)
	authG := router.Group("/auth")
	{
		authG.POST("/sign-up", SignUp)
		authG.POST("/sign-in", SignIn)
	}
	apiG := router.Group("/api", middleware.CheckUserAuthentication)

	profileG := apiG.Group("/profile")
	{
		profileG.PUT("")
	}

	category := apiG.Group("/categories")
	{
		category.GET("/", GetAllCategories)
		category.GET("/:id", GetCategoryByID)
		category.POST("/", CreateCategory)
		category.PUT("/:id", UpdateCategory)
		category.DELETE("/:id", DeleteCategory)
	}

	product := apiG.Group("/products")
	{
		product.GET("/", GetAllProducts)
		product.GET("/:id", GetProductByID)
		product.POST("/", CreateProduct)
		product.PUT("/:id", UpdateProduct)
		product.DELETE("/:id", DeleteProduct)
	}

	sale := apiG.Group("/sales")
	{
		sale.GET("/", GetAllSales)
		sale.GET("/:id", GetSaleByID)
		sale.POST("/", CreateSale) //with items
		sale.PUT("/:id", UpdateSale)
		sale.DELETE("/:id", DeleteSale)
	}

	saleItems := apiG.Group("/sale-items")
	{
		saleItems.GET("/", GetAllSaleItems)
		saleItems.GET("/:id", GetSaleItemByID)
		saleItems.PUT("/:id", UpdateSaleItem)
		saleItems.DELETE("/:id", DeleteSaleItem)
	}
	apiG.GET("/sales/:id/receipt", GetReceipt)
	apiG.GET("/report", GetSalesReport)

	if err := router.Run(configs.AppSettings.AppParams.PortRun); err != nil {
		logger.Error.Printf("[api_handlers] RunServer():  Error during running HTTP server: %s", err.Error())
		return err
	}

	return nil
}
