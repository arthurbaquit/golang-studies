package main

import (
	"encoding/json"
	"os"
)

type user struct{
	First string
	Last string
	Sayings []string
}

func main(){
	user1 := user{
		First: "James",
		Last: "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}
	user2 := user{
		First: "Miss",
		Last: "Moneypenny",
		Sayings: []string{"James, it is soo good to see you", "Would you like me to take care of that for you, James?", "I would really prefer to be a secret"},
	}
	user3 := user{
		First: "M",
		Last: "Hmmmm",
		Sayings: []string{"Oh, James. You didn't", "Dear God, what has James done now?", "Can someone please tell me where James Bond is?"},
	}
	users := []user{user1, user2, user3}
	json.NewEncoder(os.Stdout).Encode(users)
}