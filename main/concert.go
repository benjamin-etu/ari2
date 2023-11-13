package concert

import (
	"concert/main/artist"
	"errors"
	"time"
)

var concerts []Concert

type Concert struct {
	Id     int
	Name   string
	Date   time.Time
	Price  float64
	Artist artist.Artist
}

func ConcertInit() {
	concert := Concert{
		Id:     1,
		Name:   "Gazo Tour",
		Date:   time.Now(),
		Price:  5,
		Artist: artist.Artist{Id: 1, Name: "Artiste1", Note: "Note Artiste1"},
	}
	concert2 := Concert{
		Id:     2,
		Name:   "Damso Tour",
		Date:   time.Now(),
		Price:  66,
		Artist: artist.Artist{Id: 2, Name: "Artiste2", Note: "Note Artiste2"},
	}
	concerts = append(concerts, concert, concert2)
}

func AddConcert(concertAdd Concert) {
	concerts = append(concerts, concertAdd)
}

func GetConcert() []Concert {
	return concerts
}

func DeleteConcertById(id int) (Concert, error) {
	var indexToDelete int = -1
	var concertToDelete Concert

	for i, concert := range concerts {
		if concert.Id == id {
			indexToDelete = i
			concertToDelete = concert
			break
		}
	}
	if indexToDelete >= 0 {
		concerts = append(concerts[:indexToDelete], concerts[indexToDelete+1:]...)
		return concertToDelete, nil
	}

	return Concert{}, errors.New("Concert non trouvé")
}
