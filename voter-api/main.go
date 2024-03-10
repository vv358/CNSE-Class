package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"drexel.edu/todo/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	hostFlag string
	portFlag uint
)

func processCmdLineFlags() {

	flag.StringVar(&hostFlag, "h", "0.0.0.0", "Listen on all interfaces")
	flag.UintVar(&portFlag, "p", 1080, "Default Port")

	flag.Parse()
}

// main is the entry point for our todo API application.  It processes
// the command line flags and then uses the db package to perform the
// requested operation
func main() {
	processCmdLineFlags()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())

	apiHandler, err := api.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//HTTP Standards for "REST" APIS
	//GET - Read/Query
	//POST - Create
	//PUT - Update
	//DELETE - Delete

	app.Put("/voters/:id<int>", apiHandler.UpdateVoters)
	app.Put("/voters/:id<int>/polls/:pollid<int>", apiHandler.UpdateVotersPoll)
	app.Delete("/voters/:id<int>", apiHandler.DeleteVoters)
	app.Delete("/voters/:id<int>/polls/:pollid<int>", apiHandler.DeleteVotersPoll)
	app.Delete("/voters", apiHandler.DeleteAllVoters)
	app.Get("/voters", apiHandler.ListAllVoters)
	app.Get("/voters/:id<int>", apiHandler.GetVoters)
	app.Get("/voters/:id<int>/polls", apiHandler.GetVotersPoll)
	app.Get("/voters/:id<int>/polls/:pollid<int>", apiHandler.GetVotersPollId)
	app.Post("/voters", apiHandler.AddVoters)
	app.Post("/voters/:id<int>/polls", apiHandler.AddVotersPoll)

	app.Get("/crash", apiHandler.CrashSim)
	app.Get("/crash2", apiHandler.CrashSim2)
	app.Get("/crash3", apiHandler.CrashSim3)
	app.Get("/voters/health", apiHandler.HealthCheck)

	//We will now show a common way to version an API and add a new
	//version of an API handler under /v2.  This new API will support
	//a path parameter to search for todos based on a status
	// v2 := app.Group("/v2")
	// v2.Get("/todo", apiHandler.ListSelectTodos)

	serverPath := fmt.Sprintf("%s:%d", hostFlag, portFlag)
	log.Println("Starting server on ", serverPath)
	app.Listen(serverPath)
}
