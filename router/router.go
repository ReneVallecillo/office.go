package router

import (
	"time"

	"github.com/ReneVallecillo/office.go/auth" //remove dependecy
	"github.com/ReneVallecillo/office.go/handlers"
	"github.com/ReneVallecillo/office.go/mock"
	"github.com/ReneVallecillo/office.go/postgres"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/jmoiron/sqlx"
)

// InitRouter initializes the router
func InitRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	//TODO: check use of middleware after DDD
	//router.Use(postgres.Database(db))
	//Injects

	dbService := &postgres.UserService{DB: db}
	authService := &auth.AuthService{UserRepository: dbService}
	context := &AuthContext{AuthService: authService, Authorizer: authService}

	v1 := router.Group("/api/v1")
	{
		//Routes
		v1.POST("/login", context.AuthHandler)
		v1.GET("/", handlers.NotImplemented)
		v1.GET("/ping", Ping)
	}

	//Mock Routes
	test := router.Group("/test")
	{
		test.GET("/products/:slug/find", mock.MockHandler)
		test.GET("/products", mock.MockProductHandler)
	}

	authorized := router.Group("/", context.TokenAuthMiddleware())
	{
		authorized.GET("/profile", mock.MockProductHandler)
		// authorized.GET("/api/v1/users", handlers.UserListHandler)
	}

	return router
}
