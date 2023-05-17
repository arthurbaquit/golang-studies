package main

import (
	"fmt"
	"sort"
)


type user struct{
	First string
	Last string
	Age int
	Sayings []string
}
type orderByAge []user
func (o orderByAge) Len() int { return len(o) }
func (o orderByAge) Swap(i, j int) { o[i], o[j] = o[j], o[i] }
func (o orderByAge) Less(i, j int) bool { return o[i].Age < o[j].Age }

type orderByLastName []user
func (o orderByLastName) Len() int { return len(o) }
func (o orderByLastName) Swap(i, j int) { o[i], o[j] = o[j], o[i] }
func (o orderByLastName) Less(i, j int) bool { return o[i].Last < o[j].Last }
func main(){
	user1 := user{
		First: "James",
		Last: "Bond",
		Age: 32, 
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}
	user2 := user{
		First: "Miss",
		Last: "Moneypenny",
		Age: 27, 
		Sayings: []string{"James, it is soo good to see you", "Would you like me to take care of that for you, James?", "I would really prefer to be a secret"},
	}
	user3 := user{
		First: "M",
		Last: "Hmmmm",
		Age: 54, 
		Sayings: []string{"Oh, James. You didn't", "Dear God, what has James done now?", "Can someone please tell me where James Bond is?"},
	}
	users := []user{user1, user2, user3}
	for _, v := range(users){
		sort.Strings(v.Sayings)
	}
	usersByAge := users
	usersByLastName := users
	
	for _, v := range(usersByAge){
		fmt.Println(v.First, v.Age)
	}
	sort.Sort(orderByAge(usersByAge))
	for _, v := range(usersByAge){
		fmt.Println(v.First, v.Age, v.Sayings)
	}
	for _, v := range(usersByLastName){
		fmt.Println(v.Last)
	}
	sort.Sort(orderByLastName(usersByLastName))
	for _, v := range(usersByLastName){
		fmt.Println(v.Last, v.Sayings)
	}
	

}