package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"drexel.edu/todo/db"
	"github.com/gofiber/fiber/v2"
)

// The api package creates and maintains a reference to the data handler
// this is a good design practice
type VoterAPI struct {
	db *db.VoterList
}

func New() (*VoterAPI, error) {
	dbHandler, err := db.NewVoterList()
	if err != nil {
		return nil, err
	}

	return &VoterAPI{db: dbHandler}, nil
}

func (vt *VoterAPI) ListAllVoters(c *fiber.Ctx) error {

	voterList, err := vt.db.GetAllVoters()
	if err != nil {
		log.Println("Error Getting All Items: ", err)
		return fiber.NewError(http.StatusNotFound,
			"Error Getting All Items")
	}
	//Note that the database returns a nil slice if there are no items
	//in the database.  We need to convert this to an empty slice
	//so that the JSON marshalling works correctly.  We want to return
	//an empty slice, not a nil slice. This will result in the json being []
	if voterList == nil {
		voterList = make([]db.Voter, 0)
	}

	return c.JSON(voterList)
}

func (vt *VoterAPI) GetVoters(c *fiber.Ctx) error {

	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	voter, err := vt.db.GetVoter(uint(id))
	if err != nil {
		log.Println("Item not found: ", err)
		return fiber.NewError(http.StatusNotFound)
	}

	return c.JSON(voter)
}

func (vt *VoterAPI) GetVotersPoll(c *fiber.Ctx) error {

	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	voter, err := vt.db.GetVoterPoll(uint(id))
	if err != nil {
		log.Println("Item not found: ", err)
		return fiber.NewError(http.StatusNotFound)
	}

	return c.JSON(voter)
}

func (vt *VoterAPI) GetVotersPollId(c *fiber.Ctx) error {

	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	pollIdStr := c.Params("pollid")
	pollId, err := strconv.ParseUint(pollIdStr, 10, 64)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	voter, err := vt.db.GetVoterPollId(uint(id), uint(pollId))
	if err != nil {
		log.Println("Item not found: ", err)
		return fiber.NewError(http.StatusNotFound)
	}

	return c.JSON(voter)
}

// implementation for POST /todo
// adds a new todo
func (vt *VoterAPI) AddVoters(c *fiber.Ctx) error {
	var voter db.Voter

	if err := c.BodyParser(&voter); err != nil {
		log.Println("Error binding JSON: ", err)
		return fiber.NewError(http.StatusBadRequest)
	}

	if err := vt.db.AddVoter(voter); err != nil {
		log.Println("Error adding item: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	}

	return c.JSON(voter)
}

func (vt *VoterAPI) AddVotersPoll(c *fiber.Ctx) error {
	voterIDStr := c.Params("id")

	voterID, err := strconv.ParseUint(voterIDStr, 10, 64)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	var voterPoll db.VoterHistory

	if err := c.BodyParser(&voterPoll); err != nil {
		log.Println("Error binding JSON: ", err)
		return fiber.NewError(http.StatusBadRequest)
	}

	if err := vt.db.AddVoterPoll(uint(voterID), voterPoll); err != nil {
		log.Println("Error adding item: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	}

	return c.JSON(voterPoll)
}

func (td *VoterAPI) CrashSim(c *fiber.Ctx) error {
	//panic() is go's version of throwing an exception
	//note with recover middleware this will not end program
	panic("Simulating an unexpected crash")
}

func (td *VoterAPI) CrashSim2(c *fiber.Ctx) error {
	//A stupid crash simulation example
	i := 0
	j := 1 / i
	jStr := fmt.Sprintf("%d", j)
	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"val_j": jStr,
		})
}

func (td *VoterAPI) CrashSim3(c *fiber.Ctx) error {
	//A stupid crash simulation example
	os.Exit(10)
	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"error": "will never get here, nothing you can do about this",
		})
}

// implementation of GET /health. It is a good practice to build in a
// health check for your API.  Below the results are just hard coded
// but in a real API you can provide detailed information about the
// health of your API with a Health Check
func (td *VoterAPI) HealthCheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"status":             "ok",
			"version":            "1.0.0",
			"uptime":             100,
			"users_processed":    1000,
			"errors_encountered": 10,
		})
}
