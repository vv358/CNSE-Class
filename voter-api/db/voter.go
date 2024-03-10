package db

import (
	"errors"
	"time"
)

type VoterHistory struct {
	PollId   uint      `json:"poll_id"`
	VoteId   uint      `json:"vote_id"`
	VoteDate time.Time `json:"vote_date"`
}

type Voter struct {
	VoterId     uint           `json:"voter_id"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	VoteHistory []VoterHistory `json:"vote_history"`
}

type VoterList struct {
	Voters map[uint]Voter //A map of VoterIDs as keys and Voter structs as values
}

func NewVoterList() (*VoterList, error) {

	voterList := &VoterList{
		Voters: make(map[uint]Voter),
	}

	return voterList, nil
}

func (v *VoterList) AddVoter(item Voter) error {

	//Before we add an item to the DB, lets make sure
	//it does not exist, if it does, return an error
	_, ok := v.Voters[item.VoterId]
	if ok {
		return errors.New("voter already exists")
	}

	//Now that we know the item doesn't exist, lets add it to our map
	v.Voters[item.VoterId] = item

	//If everything is ok, return nil for the error
	return nil
}

func (v *VoterList) AddVoterPoll(voterID uint, voterPoll VoterHistory) error {

	voter, ok := v.Voters[voterID]
	if !ok {
		return errors.New("voter does not exist")
	}
	//Now that we know the item doesn't exist, lets add it to our map
	voter.VoteHistory = append(voter.VoteHistory, voterPoll)
	v.Voters[voterID] = voter

	//If everything is ok, return nil for the error
	return nil
}

// DeleteItem accepts an item id and removes it from the DB.
// Preconditions:   (1) The database file must exist and be a valid
//
//					(2) The item must exist in the DB
//	    				because we use the item.Id as the key, this
//						function must check if the item already
//	    				exists in the DB, if not, return an error
//
// Postconditions:
//
//	 (1) The item will be removed from the DB
//		(2) The DB file will be saved with the item removed
//		(3) If there is an error, it will be returned
// func (t *ToDo) DeleteItem(id int) error {

// 	// we should if item exists before trying to delete it
// 	// this is a good practice, return an error if the
// 	// item does not exist

// 	//Now lets use the built-in go delete() function to remove
// 	//the item from our map
// 	delete(t.toDoMap, id)

// 	return nil
// }

// // DeleteAll removes all items from the DB.
// // It will be exposed via a DELETE /todo endpoint
// func (t *ToDo) DeleteAll() error {
// 	//To delete everything, we can just create a new map
// 	//and assign it to our existing map.  The garbage collector
// 	//will clean up the old map for us
// 	t.toDoMap = make(map[int]ToDoItem)

// 	return nil
// }

// // UpdateItem accepts a ToDoItem and updates it in the DB.
// // Preconditions:   (1) The database file must exist and be a valid
// //
// //					(2) The item must exist in the DB
// //	    				because we use the item.Id as the key, this
// //						function must check if the item already
// //	    				exists in the DB, if not, return an error
// //
// // Postconditions:
// //
// //	 (1) The item will be updated in the DB
// //		(2) The DB file will be saved with the item updated
// //		(3) If there is an error, it will be returned
// func (t *ToDo) UpdateItem(item ToDoItem) error {

// 	// Check if item exists before trying to update it
// 	// this is a good practice, return an error if the
// 	// item does not exist
// 	_, ok := t.toDoMap[item.Id]
// 	if !ok {
// 		return errors.New("item does not exist")
// 	}

// 	//Now that we know the item exists, lets update it
// 	t.toDoMap[item.Id] = item

// 	return nil
// }

// // GetItem accepts an item id and returns the item from the DB.
// // Preconditions:   (1) The database file must exist and be a valid
// //
// //					(2) The item must exist in the DB
// //	    				because we use the item.Id as the key, this
// //						function must check if the item already
// //	    				exists in the DB, if not, return an error
// //
// // Postconditions:
// //
// //	 (1) The item will be returned, if it exists
// //		(2) If there is an error, it will be returned
// //			along with an empty ToDoItem
// //		(3) The database file will not be modified
func (v *VoterList) GetVoter(id uint) (Voter, error) {

	voter, ok := v.Voters[id]
	if !ok {
		return Voter{}, errors.New("voter does not exist")
	}

	return voter, nil
}

func (v *VoterList) GetVoterPoll(id uint) ([]VoterHistory, error) {

	voter, ok := v.Voters[id]
	if !ok {
		return []VoterHistory{}, errors.New("voter does not exist")
	}

	return voter.VoteHistory, nil
}

func (v *VoterList) GetVoterPollId(id, pollId uint) (VoterHistory, error) {

	voter, ok := v.Voters[id]
	if !ok {
		return VoterHistory{}, errors.New("voter does not exist")
	}

	for _, poll := range voter.VoteHistory {
		if poll.PollId == pollId {
			return poll, nil
		}
	}

	return VoterHistory{}, errors.New("voter poll not found")
}

// ChangeItemDoneStatus accepts an item id and a boolean status.
// It returns an error if the status could not be updated for any
// reason.  For example, the item itself does not exist, or an
// IO error trying to save the updated status.

// Preconditions:   (1) The database file must exist and be a valid
//
//					(2) The item must exist in the DB
//	    				because we use the item.Id as the key, this
//						function must check if the item already
//	    				exists in the DB, if not, return an error
//
// Postconditions:
//
//	 (1) The items status in the database will be updated
//		(2) If there is an error, it will be returned.
//		(3) This function MUST use existing functionality for most of its
//			work.  For example, it should call GetItem() to get the item
//			from the DB, then it should call UpdateItem() to update the
//			item in the DB (after the status is changed).
// func (t *ToDo) ChangeItemDoneStatus(id int, value bool) error {

// 	//update was successful
// 	return errors.New("not implemented")
// }

// GetAllItems returns all items from the DB.  If successful it
// returns a slice of all of the items to the caller
// Preconditions:   (1) The database file must exist and be a valid
//
// Postconditions:
//
//	 (1) All items will be returned, if any exist
//		(2) If there is an error, it will be returned
//			along with an empty slice
//		(3) The database file will not be modified
func (v *VoterList) GetAllVoters() ([]Voter, error) {

	//Now that we have the DB loaded, lets crate a slice
	var voterList []Voter

	//Now lets iterate over our map and add each item to our slice
	for _, item := range v.Voters {
		voterList = append(voterList, item)
	}

	//Now that we have all of our items in a slice, return it
	return voterList, nil
}

func (v *VoterList) DeleteAll() error {
	//To delete everything, we can just create a new map
	//and assign it to our existing map.  The garbage collector
	//will clean up the old map for us
	v.Voters = make(map[uint]Voter)

	return nil
}

func (v *VoterList) DeleteVoter(id uint) error {

	if _, ok := v.Voters[id]; !ok {
		return errors.New("voter does not exist")
	}

	delete(v.Voters, id)

	return nil
}

func (v *VoterList) UpdateVoter(id uint, voter Voter) error {

	// Check if item exists before trying to update it
	// this is a good practice, return an error if the
	// item does not exist
	// _, ok := v.Voters[id]
	// if !ok {
	// 	return errors.New("Voter does not exist")
	// }

	//Now that we know the item exists, lets update it
	v.Voters[id] = voter

	return nil
}

// PrintItem accepts a ToDoItem and prints it to the console
// in a JSON pretty format. As some help, look at the
// json.MarshalIndent() function from our in class go tutorial.
// func (t *ToDo) PrintItem(item ToDoItem) {
// 	jsonBytes, _ := json.MarshalIndent(item, "", "  ")
// 	fmt.Println(string(jsonBytes))
// }

// // PrintAllItems accepts a slice of ToDoItems and prints them to the console
// // in a JSON pretty format.  It should call PrintItem() to print each item
// // versus repeating the code.
// func (t *ToDo) PrintAllItems(itemList []ToDoItem) {
// 	for _, item := range itemList {
// 		t.PrintItem(item)
// 	}
// }

// // JsonToItem accepts a json string and returns a ToDoItem
// // This is helpful because the CLI accepts todo items for insertion
// // and updates in JSON format.  We need to convert it to a ToDoItem
// // struct to perform any operations on it.
// func (t *ToDo) JsonToItem(jsonString string) (ToDoItem, error) {
// 	var item ToDoItem
// 	err := json.Unmarshal([]byte(jsonString), &item)
// 	if err != nil {
// 		return ToDoItem{}, err
// 	}

// 	return item, nil
// }
