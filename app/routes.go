package app

import "dating/controllers"

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
}
