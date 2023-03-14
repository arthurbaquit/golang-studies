package solid

// Single Responsibility Principle
//
// The Single Responsibility Principle (SRP) states that a class should have one
// and only one reason to change, meaning that a class should have only one job.

import (
	"io/ioutil"
	"strings"
)

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	j.entries = append(j.entries, text)
	return len(j.entries) - 1
}

// remove entry based on index
func (j *Journal) RemoveEntry(index int) {
	j.entries = append(j.entries[:index], j.entries[index+1:]...)
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// func (j *Journal) Save(filename string) {
// 	// save to file txt
// 	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
//}

// This breaks the Single Responsibility Principle, once the Journal deals
// with the journal entries, not with persistence. We should create a new
// class to handle persistence, like DB.

// Then, we could remove the Save method from the Journal class, and create
// a new class to handle persistence

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	// save to file txt
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func mainSRP() {
	j := Journal{}
	j.AddEntry("I studied today")
	j.AddEntry("I took so many trains")
	//j.Save("journal.txt")
	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}

// Therefore, the Journal has the job of keeping track of entries, and the
// Persistence class has the job of saving the entries to a file. This is
// the Single Responsibility Principle in action.
