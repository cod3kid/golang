package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

var entryCount = 0
type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s",
		entryCount,
		text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	entryCount--

	j.entries = append(j.entries[:index],j.entries[index+1:]...)

}



// Below three receiver functions breaks the SRP

func (j *Journal) Save(filename string) {
	_ = os.WriteFile(filename,
		[]byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}



// Follows SRP as SaveToFile is not linked to Journal struct

func SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename,
		[]byte(strings.Join(j.entries, "\n")), 0644)
}




func main() {
	j := Journal{}
	j.AddEntry("I cried today.")
	j.AddEntry("I ate a bug")

	// j.RemoveEntry(0)
	fmt.Println(strings.Join(j.entries, "\n"))

	// separate function
	SaveToFile(&j, "journal.txt")
}