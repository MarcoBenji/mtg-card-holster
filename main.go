package main

import (
	"encoding/json"
	"log"
	"fmt"
    "net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
)




// main runner
func main() {
	fmt.Println("starting the app")
	router := mux.NewRouter()

	addCardsToDB()

	router.HandleFunc("/cards", GetCards).Methods("GET")
    router.HandleFunc("/cards/{id}", GetCard).Methods("GET")
    router.HandleFunc("/cards/{id}", CreateCard).Methods("POST")
	router.HandleFunc("/cards/{id}", DeleteCard).Methods("DELETE")
	router.HandleFunc("/cardGet", MtgCardGet).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", router))
}
func GetCards(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(cards)
}
func GetCard(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range cards {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Card{})
}
func CreateCard(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	var card Card 
	json.NewDecoder(r.Body).Decode(&card)
	
	cards = append(cards, card)

	json.NewEncoder(w).Encode(cards)
}
func DeleteCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range cards {
		if item.ID == params["id"]{
			cards = append(cards[:index], cards[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(cards)
}
func MtgCardGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var queryString = "https://api.magicthegathering.io/v1/cards?name=" + params["name"]
	fmt.Println(queryString)
	response, err := http.Get(queryString)
	if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
		data, _ := ioutil.ReadAll(response.Body)
		// fmt.Println(data)

		mtgCards := make(map[string][]Card)
		if (json.Unmarshal([]byte(data), &mtgCards) != nil){
			fmt.Printf("The parse of the cards failed %s\n", err)
		}
		
		// if (len(mtgCards) > 0){
		// 	cards = append(cards, mtgCards.get(0))
		// }
		// fmt.Println(response)
		// fmt.Println(response.Body)
		// // fmt.Println(data)

		// json.NewDecoder(response.Body).Decode(&card)
		
		fmt.Println(mtgCards)
		
		// for _, element := range mtgCards {
		// 	cards = append(cards, element)
		// }

        json.NewEncoder(w).Encode(mtgCards)
    }
}

var numOfCards int
var cards []Card
func addCardsToDB(){
	cards = append(cards, Card{localID: "1", Name: "Brainstorm", ManaCost: "{B}", Text: "Draw three cards, then put two cards from your hand on top of your library in any order.",
	 	Flavor: "'I reeled from the blow, and then suddenly, I knew exactly what to do. Within moments, victory was mine.' —Gustha Ebbasdotter, Kjeldoran Royal Mage", Type: "Instant", Set: "IceAge"})
	cards = append(cards, Card{localID: "2", Name: "Lightning Bolt", ManaCost: "{R}", Text: "Lightning Bolt deals 3 damage to any target.",
		Flavor: "The sparkmage shrieked, calling on the rage of the storms of his youth. To his surprise, the sky responded with a fierce energy he'd never thought to see again.", Type: "Instant", Set: "M10"})
	cards = append(cards, Card{localID: "3", Name: "Swords to Plowshares", ManaCost: "{W}", Text: "Exile target creature. Its controller gains life equal to its power.",
		Flavor: "'The so-called Barbarians will not respect us for our military might—they will respect us for our honor.' —Lucilde Fiksdotter, Leader of the Order of the White Shield", Type: "Instant", Set: "Ice Age"})
	numOfCards = 3
}
//testing cards construct
type Card struct {
    localID       string   		`json:"id,omitempty"`
    Name          string        `json:"name,omitempty"`
	ManaCost      string        `json:"manaCost,omitempty"`
	Cmc           int           `json:"cmc,omitempty"`
	Colors        []string      `json:"colors,omitempty"`
	ColorIdentity []string      `json:"colorIdentity,omitempty"`
	Type          string        `json:"type,omitempty"`
	Supertypes    []interface{} `json:"supertypes,omitempty"`
	Types         []string      `json:"types,omitempty"`
	Subtypes      []interface{} `json:"subtypes,omitempty"`
	Rarity        string        `json:"rarity,omitempty"`
	Set           string        `json:"set,omitempty"`
	SetName       string        `json:"setName,omitempty"`
	Text          string        `json:"text,omitempty"`
	Flavor        string        `json:"flavor,omitempty"`
	Artist        string        `json:"artist,omitempty"`
	Number        string        `json:"number,omitempty"`
	Layout        string        `json:"layout,omitempty"`
	Multiverseid  int           `json:"multiverseid,omitempty"`
	ImageURL      string        `json:"imageUrl,omitempty"`
	Rulings       []struct {
		Date string 			`json:"date,omitempty"`
		Text string 			`json:"text,omitempty"`
	} 							`json:"rulings,omitempty"`
	ForeignNames []struct {
		Name         string      `json:"name,omitempty"`
		Text         string      `json:"text,omitempty"`
		Flavor       interface{} `json:"flavor,omitempty"`
		ImageURL     interface{} `json:"imageUrl,omitempty"`
		Language     string      `json:"language,omitempty"`
		Multiverseid interface{} `json:"multiverseid,omitempty"`
	} 							 `json:"foreignNames,omitempty"`
	Printings    []string 		 `json:"printings,omitempty"`
	OriginalText string   		 `json:"originalText,omitempty"`
	OriginalType string   		 `json:"originalType,omitempty"`
	Legalities   []struct {
		Format   string 		 `json:"format,omitempty"`
		Legality string 		 `json:"legality,omitempty"`
	} 							 `json:"legalities,omitempty"`
	ID string 					 `json:"id,omitempty"`
}