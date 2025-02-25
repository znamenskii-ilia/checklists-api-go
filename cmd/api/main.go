package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/znamenskii-ilia/checklists-api-go/internal/db"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/infrastructure/repositories"
	checklistsRouter "github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/interfaces/http"
)

func main() {
	port := "3001"
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db, err := db.New("./checklists.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	repo := repositories.NewSqliteChecklistsRepository(db)

	r.Mount("/api/checklists", checklistsRouter.New(repo))

	fmt.Printf("Server is running on port %s\n", port)

	err = http.ListenAndServe(":"+port, r)

	if err != nil {
		fmt.Println("Error starting server:", err)
		log.Fatal(err)
	}

}
