## Hypermedia as the engine of application state (HATEOAS)

HATEOAS enables the client to interact with the server with minimal knowledge about how to interact with the server. When a client makes a REST call, the server returns data that enables the client to take further actions and move from one application state to the next. This is possible through simple hypermedia links contained within the response.

So for our **VoterAPI**, **PollsAPI**, and **VotesAPI** to work together we could use hypermedia as below:

```
**VoterAPI**
{
  "voter_id": 1,
  "first_name": "John",
  "last_name": "Doe",
  "vote_history": [
    {
      "poll_id": 1,
      "vote_date": "2024-03-15T10:00:00Z"
    }
  ],
  "_links": {
    "self": {
      "href": "/voters/1"
    },
    "vote_history": {
      "href": "/voters/1/polls"
    },
    "vote_response": {
      "href": "/votes/voterid/1/pollid/1"
    },
    "voter_responses": {
      "href": "/votes/voterid/1/"
    },
    "make_vote": {
      "href": "/votes"
    },
    "poll_info": {
      "href": "/polls/1"
    },
    "all_polls": {
      "href": "/polls"
    }
  }
}
```
Explanation:

* The _links object contains hypermedia controls for navigating to different resources or actions related to the voter.
* The self link points to the voter's profile, allowing clients to retrieve more details about the voter.
* The vote_history link points to the voter's vote history, allowing clients to view the list of polls the voter participated in and the corresponding vote dates.
* The make_vote link points to the endpoint for making a new vote, allowing clients to initiate the voting process for the voter.
