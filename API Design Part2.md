## Hypermedia as the engine of application state (HATEOAS)

HATEOAS enables the client to interact with the server with minimal knowledge about how to interact with the server. When a client makes a REST call, the server returns data that enables the client to take further actions and move from one application state to the next. This is possible through simple hypermedia links contained within the response.

So for our **VoterAPI**, **PollsAPI**, and **VotesAPI** to work together we could use hypermedia as below:

```
{
  "voter_id": 1,
  "first_name": "John",
  "last_name": "Doe",
  "vote_history": [
    {
      "poll_id": 1,
      "vote_date": "2024-03-16T12:00:00Z"
    }
  ],
  "_links": {
    "self": {
      "href": "/voters/1"
    },
    "vote_history": {
      "href": "/voters/1/vote_history"
    },
    "make_vote": {
      "href": "/votes"
    }
  }
}

```

