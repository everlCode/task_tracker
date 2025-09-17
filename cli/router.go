package cli

type Router struct {
	
}

func(r Router) New() *Router {
	router := &Router{}
	router.registerCommands()

	return router
}

func(r Router) registerCommands() {
	
}