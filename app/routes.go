package app

import "dating/controllers"

var (
	parentGroupAuth    = "auth/"
	parentGroupProfile = "profile/"
	swipeGroupProfile  = "swipe/"
)

type IRoutes interface {
	registerRoutes()
}

type Routes struct {
	controller *controllers.Controller
}

func initRoutes(controller *controllers.Controller) *Routes {
	return &Routes{
		controller: controller,
	}

}
func (r *Routes) registerRoutes() {
	r.authRouter()
	r.profileRoutes()
	r.swipeRoutes()
}

func (r *Routes) authRouter() {
	authRoute := router.Group(parentGroupAuth)
	{
		authRoute.POST("register", r.controller.AuthController.Register)
		authRoute.POST("login", r.controller.AuthController.Login)
	}
}

func (r *Routes) profileRoutes() {
	profileRoute := router.Group(parentGroupProfile).Use(r.controller.AuthMiddleware.JWTVerifyToken)
	{
		profileRoute.GET("", r.controller.ProfileController.ListProfile)
	}
}

func (r *Routes) swipeRoutes() {
	swipeRoute := router.Group(swipeGroupProfile).Use(r.controller.AuthMiddleware.JWTVerifyToken)
	{
		swipeRoute.GET("left/:id", r.controller.SwipeController.SwipeLeft)
		swipeRoute.GET("right/:id", r.controller.SwipeController.SwipeRight)
	}
}
