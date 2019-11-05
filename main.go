package main

import (
	"fmt"
	"net/http"
	"os"

	driver "./driver"

	ph "./handler/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connection, err := driver.ConnectSQL(dbHost, dbPort, "admin", dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	organizerHandler := ph.InitOrganizerHandler(connection)
	userHandler := ph.InitUserHandler(connection)
	accountHandler := ph.InitAccountHandler(connection)
	rundownHandler := ph.InitRundownHandler(connection)
	rundownItemHandler := ph.InitRundownItemHandler(connection)

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/organizer", initOrganizerRoute(organizerHandler))
		rt.Mount("/user", initUserRoute(userHandler))
		rt.Mount("/account", initAccountRoute(accountHandler))
		rt.Mount("/rundown", initRundownRoute(rundownHandler))
		rt.Mount("/rundown_item", initRundownItemRoute(rundownItemHandler))
	})

	fmt.Println("Server listen at :8005")
	http.ListenAndServe(":3000", r)
}

func initOrganizerRoute(handler *ph.OrganizerHandler) http.Handler {
	route := chi.NewRouter()
	route.Post("/", handler.CreateOrganizer)
	route.Get("/", handler.GetAll)
	route.Get("/{id:[0-9]+}", handler.GetByID)
	route.Put("/{id:[0-9]+}", handler.Update)
	route.Delete("/{id:[0-9]+}", handler.Delete)

	return route
}

func initUserRoute(handler *ph.UserHandler) http.Handler {
	route := chi.NewRouter()
	route.Post("/", handler.Create)
	route.Get("/", handler.GetAll)
	route.Get("/{id:[0-9]+}", handler.GetByID)
	route.Put("/{id:[0-9]+}", handler.Update)
	route.Delete("/{id:[0-9]+}", handler.Delete)

	return route
}

func initAccountRoute(handler *ph.AccountHandler) http.Handler {
	route := chi.NewRouter()
	route.Post("/", handler.Create)
	route.Get("/", handler.GetAll)
	route.Get("/{id:[0-9]+}", handler.GetByID)
	route.Put("/{id:[0-9]+}", handler.Update)
	route.Delete("/{id:[0-9]+}", handler.Delete)

	return route
}

func initRundownRoute(handler *ph.RundownHandler) http.Handler {
	route := chi.NewRouter()
	route.Post("/", handler.Create)
	route.Get("/", handler.GetAll)
	route.Get("/{organizerId:[0-9]+}", handler.GetByOrganizerId)
	route.Get("/{organizerId:[0-9]+}/{startDate:[0-100]+}/{endDate:[0-100]+}", handler.GetByOrganizerIdAndDate)
	route.Get("/{organizerId[0-9]+}/{id:[0-9]+}", handler.GetByOrganizerAndId)
	route.Put("/{id:[0-9]+}", handler.Update)
	route.Delete("/{id:[0-9]+}", handler.Delete)

	return route
}

func initRundownItemRoute(handler *ph.RundownItemHandler) http.Handler {
	route := chi.NewRouter()
	route.Post("/", handler.Create)
	route.Get("/", handler.GetAll)
	route.Get("/{rundownId:[0-9]+}", handler.GetByRundownId)
	route.Put("/{id:[0-9]+}", handler.Update)
	route.Delete("/{id:[0-9]+}", handler.Delete)

	return route
}
