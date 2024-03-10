package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"drexel.edu/todo/db"
	fake "github.com/brianvoe/gofakeit/v6" //aliasing package name
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

var (
	BASE_API = "http://localhost:1080"

	cli = resty.New()
)

func TestMain(m *testing.M) {

	//SETUP GOES FIRST
	rsp, err := cli.R().Delete(BASE_API + "/voters")

	if rsp.StatusCode() != 200 {
		log.Printf("error clearing database, %v", err)
		os.Exit(1)
	}

	code := m.Run()

	//CLEANUP

	//Now Exit
	os.Exit(code)
}

func newRandVoter(id uint) db.Voter {
	return db.Voter{
		VoterId: id,
		Name:    fake.FirstName(),
		Email:   fake.Email(),
		VoteHistory: []db.VoterHistory{
			{
				PollId:   uint(id + 1),
				VoteId:   uint(id + 1),
				VoteDate: time.Now(),
			},
		},
	}
}

func Test_LoadDB(t *testing.T) {
	numLoad := 3
	for i := 0; i < numLoad; i++ {
		item := newRandVoter(uint(i))
		rsp, err := cli.R().
			SetBody(item).
			Post(BASE_API + "/voters")

		assert.Nil(t, err)
		assert.Equal(t, 200, rsp.StatusCode())
	}
}

func Test_GetAllVoters(t *testing.T) {
	var items []db.Voter

	rsp, err := cli.R().SetResult(&items).Get(BASE_API + "/voters")

	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())

	assert.Equal(t, 3, len(items))
}

func Test_GetVoterByID(t *testing.T) {
	rsp, err := cli.R().Get(BASE_API + "/voters/1")
	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())

	var individualVoter db.Voter
	err = json.Unmarshal(rsp.Body(), &individualVoter)
	assert.Nil(t, err)

	assert.NotNil(t, individualVoter.VoterId)
	assert.NotNil(t, individualVoter.Name)
	assert.NotNil(t, individualVoter.Email)
	assert.NotNil(t, individualVoter.VoteHistory)
	assert.True(t, len(individualVoter.VoteHistory) > 0)
}

func Test_GetVoterPolls(t *testing.T) {
	rsp, err := cli.R().Get(BASE_API + "/voters/1/polls")
	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())

	var pollsHistory []db.VoterHistory
	err = json.Unmarshal(rsp.Body(), &pollsHistory)
	assert.Nil(t, err)

	assert.True(t, len(pollsHistory) > 0)

	for _, pollsHistory := range pollsHistory {
		assert.NotNil(t, pollsHistory.PollId)
		assert.NotNil(t, pollsHistory.VoteId)
		assert.NotNil(t, pollsHistory.VoteDate)
	}
}

func Test_GetVoterPollByID(t *testing.T) {
	rsp, err := cli.R().Get(BASE_API + "/voters/1/polls")
	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())

	var pollsHistory []db.VoterHistory
	err = json.Unmarshal(rsp.Body(), &pollsHistory)
	assert.Nil(t, err)

	assert.True(t, len(pollsHistory) > 0)

	randIndex := rand.Intn(len(pollsHistory))
	selectedPoll := pollsHistory[randIndex]

	rsp, err = cli.R().Get(fmt.Sprintf("%s/voters/1/polls/%d", BASE_API, selectedPoll.PollId))
	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())

	var poll db.VoterHistory
	err = json.Unmarshal(rsp.Body(), &poll)
	assert.Nil(t, err)

	assert.NotNil(t, poll.PollId)
	assert.NotNil(t, poll.VoteId)
	assert.NotNil(t, poll.VoteDate)
}

func Test_DeleteVoter(t *testing.T) {
	var item db.VoterList

	rsp, err := cli.R().SetResult(&item).Get(BASE_API + "/voters/2")
	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode(), "voters #2 expected")

	rsp, err = cli.R().Delete(BASE_API + "/voters/2")
	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode(), "voters deleted expected")

	rsp, err = cli.R().SetResult(item).Get(BASE_API + "/voters/2")
	assert.Nil(t, err)
	assert.Equal(t, 404, rsp.StatusCode(), "expected not found error code")
}

func Test_DeleteAllVoters(t *testing.T) {
	//var item db.VoterList

	rsp, err := cli.R().Get(BASE_API + "/voters")
	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode(), "voters expected")

	rsp, err = cli.R().Delete(BASE_API + "/voters")
	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode(), "all voters  deleted expected")

	rsp, err = cli.R().Get(BASE_API + "/voters")
	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode(), "expected not found error code")
}
