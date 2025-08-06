package routes

import (
	"furious/iam-api/internal/configs"
	"furious/iam-api/internal/controllers"
	"furious/iam-api/internal/repositories"
	"furious/iam-api/internal/services"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	db := configs.ConnectDB()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	users := r.Group("/users")
	{
		userRepo := repositories.NewUserRepository(db)
		userService := services.NewUserService(userRepo)
		userController := controllers.NewUserController(userService)
		users.GET("", userController.Search)
		users.POST("", userController.Persist)
		users.DELETE("/:id", userController.Delete)
	}

	roles := r.Group("/roles")
	{
		roleRepo := repositories.NewRoleRepository(db)
		roleService := services.NewRoleService(roleRepo)
		roleController := controllers.NewRoleController(roleService)
		roles.GET("", roleController.Search)
		roles.POST("", roleController.Persist)
		roles.DELETE("/:id", roleController.Delete)
	}

	return r
}
