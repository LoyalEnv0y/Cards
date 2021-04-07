package deck

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func NewDeck() deck { // we are creating the whole deck of cards
	cards := deck{}
	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _,v := range cardSuits{
		for _,y := range cardValues{
			cards = append(cards,y+" of "+v)
		}
	}
	return cards
}

func(d deck) Print(){ // we are printing the givin slice one by one
	for _,v := range d{
		fmt.Println(v)
	}
}

func Deal(d deck, handsize int) (deck, deck) { // we are dealing said amount of cards in a deck
	return d[:handsize], d[handsize:]
}

func (d deck) ToString() string { // we are converting the slice of deck to a slice of bytes so that we can save it to our harddrive
	return strings.Join([]string(d), ",")
}
func (d deck) SaveToFile(filename string) error{ // we are saving the deck to our harddrive
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0666)
}
func NewDeckFromFile(filename string) deck { // we are reading from a file
	bs, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println("Error: ",err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

/*func (d deck) deal() {
	dealt := make([]string, 0)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 5; i++{
		dealt = append(dealt, d[rand.Intn(len(d))] )
	}
	for _,v := range dealt{
		fmt.Println(v)
	}
}*/ //don't use this until you learn how to make func it without duplacates