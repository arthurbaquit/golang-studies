package main

import (
	"encoding/json"
	"fmt"
)

type user struct{
	First string
	Last string
	Sayings []string
}

func main(){
	s := `[{"First":"James","Last":"Bond", "Sayings":["Shaken, not stirred", "Any last wishes?", "Never say never"]},{"First":"Miss","Last":"Moneypenny", "Sayings":["James, it is soo good to see you", "Would you like me to take care of that for you, James?", "I would really prefer to be a secret"]},{"First":"M","Last":"Hmmmm", "Sayings":["Oh, James. You didn't", "Dear God, what has James done now?", "Can someone please tell me where James Bond is?"]}]`
	var users []user
	err := json.Unmarshal([]byte(s), &users)
	if err != nil { 
		fmt.Println(err)
	}
	fmt.Println(users)
}