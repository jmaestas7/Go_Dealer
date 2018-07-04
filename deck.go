package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new typeof deck which is a slice typeof strings
type deck []string

//create and return list of playing cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	//for the unused index number of the range of cardsuits as var suit
	for _, suit := range cardSuits {
		//for the unused index number of the range of cardValues as var value
		for _, value := range cardValues {
			//append to cards type deck slice of strings card suit + card value till loop complete
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//receiver d type of deck
//d = cards
func (d deck) print() {
	//for our index of cards type deck range of cards as card
	for i, card := range d {
		//println index and card at index
		fmt.Println(i, card)
	}
}

//d type of deck, integer handsize
func deal(d deck, handSize int) (deck, deck) {
	//return 2 decks seperated by handsize at index of slice
	return d[:handSize], d[handSize:]

}

//receiver d type of deck
func (d deck) toString() string {
	//join cards of deck with "," seperator as one string
	return strings.Join([]string(d), ",")
}

//receiver d type of deck
//save to filename specified as filename
func (d deck) saveToFile(filename string) error {
	//return byte slice of joinedstring, permisson 066(allow read/write to all)
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// create deck from fuilename specified
func newDeckFromFile(filename string) deck {
	//byte slice and error read file of filename
	bs, err := ioutil.ReadFile(filename)
	//if err if not null
	if err != nil {
		//print error and exit
		fmt.Println(err)
		os.Exit(1)
	}
	//s is byte slice(bs), to string and .Split into a slice of strings["1" , "2"]
	s := strings.Split(string(bs), ",")
	//str slice is now type deck
	return deck(s)
}

func (d deck) shuffle() {
	//source is rand from a int64 created by UnixNano of time now
	source := rand.NewSource(time.Now().UnixNano())
	//r is type *Rand created from new source created above
	r := rand.New(source)

	//for the klength of index of deck -1, newPoistion is random int to swap cards from current position to new position till index length is complete
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
