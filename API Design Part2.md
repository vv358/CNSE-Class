## Hypermedia as the engine of application state (HATEOAS)

HATEOAS enables the client to interact with the server with minimal knowledge about how to interact with the server. When a client makes a REST call, the server returns data that enables the client to take further actions and move from one application state to the next. This is possible through simple hypermedia links contained within the response.

So for our **VoterAPI**, **PollsAPI**, and **VotesAPI** to work together we could use hypermedia as below:

**VoterAPI**

```
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
**Explanation:**

* The **_links** object contains hypermedia controls for navigating to different resources or actions related to the voter.
* The **self** link points to the voter's profile, allowing clients to retrieve more details about the voter.
* The **vote_history** link points to the voter's vote history, allowing clients to view the list of polls the voter participated in and the corresponding vote dates.
* The **vote_response** link points to the voter's response to the particular poll.
* The **voter_responses** link points to the voter's response to all the polls.
* The **make_vote** link points to the endpoint for making a new vote, allowing clients to initiate the voting process for the voter.
* The **poll_info** link points to the particular poll info, allowing the clients to view the poll details.
* The **all_polls** link points to all the polls available, allowing the clients to view all the available polls.

**PollsAPI**

```
{
  "poll_id": 1,
  "poll_title": "Favorite Pet",
  "poll_question": "What type of pet do you like best?",
  "poll_options": [
    {
      "poll_option_id": 1,
      "poll_option_value": "Dog"
    },
    {
      "poll_option_id": 2,
      "poll_option_value": "Cat"
    }
  ],
  "_links": {
    "self": {
      "href": "/polls/1"
    },
    "poll_options": {
      "href": "/polls/1/options"
    },
    "all_polls": {
      "href": "/polls"
    },
    "make_vote": {
      "href": "/votes"
    }
  }
}

```
**Explanation:**

* The **_links** object contains hypermedia controls for navigating to different resources or actions related to the poll.
* The **self** link points to the poll's details, allowing clients to retrieve more information about the poll.
* The **poll_options** link points to the list of options for the poll, allowing clients to view and select the available options.
* The **all_polls** link points to all the polls available, allowing the clients to view all the available polls.
* The **make_vote** link points to the endpoint for making a new vote on the poll, allowing clients to initiate the voting process for the poll.

**VotesAPI**

```
{
  "vote_id": 1,
  "voter_id": 1,
  "poll_id": 1,
  "vote_value": 1,
  "_links": {
    "voter": {
      "href": "/voters/1"
    },
    "poll": {
      "href": "/polls/1"
    }
  }
}

```
**Explanation:**

* The **_links** object contains hypermedia controls for navigating to different resources or actions related to the votes.
* The **voter** link points to the voter details, allowing clients to view the details of the voter who made the vote.
* The **poll** link points to the poll details, allowing clients to view the poll details for which the vote was made.

This way, the hypermedia links enable the clients to dynamically discover and navigate the available resources and actions. 

**References**

> 1. https://youtu.be/ybwo_70jpGc?si=Bq83QYAJ4WJT8Riw
> 2. https://www.geeksforgeeks.org/hateoas-and-why-its-needed-in-restful-api/
> 3. I also referred to Google and ChatGPT.



