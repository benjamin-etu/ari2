package concert

import (
	"concert/main/artist"
	"testing"
	"time"
)

func TestConcertInit(contextTest *testing.T) {
	ConcertInit()

	if len(concerts) == 0 {
		contextTest.Errorf("La liste de concerts n'a pas été initialisée correctement")
	}
}

func TestAddConcert(contextTest *testing.T) {
	newConcert := Concert{
		Id:     3,
		Name:   "TestConcert",
		Date:   time.Now(),
		Price:  10.99,
		Artist: artist.Artist{Id: 1, Name: "Artiste1", Note: "Note Artiste1"},
	}

	AddConcert(newConcert)

	if len(concerts) == 0 {
		contextTest.Errorf("Le concert n'a pas été ajouté correctement")
	}
}

func TestDeleteConcertById(contextTest *testing.T) {
	length := len(concerts)

	deletedConcert, err := DeleteConcertById(1)

	if err != nil {
		contextTest.Errorf("Erreur lors de la suppression du concert: %v", err)
	}

	if len(concerts) != length-1 {
		contextTest.Errorf("Le concert n'a pas été supprimé correctement")
	}

	expectedConcert := Concert{
		Id:     1,
		Name:   "Gazo Tour",
		Date:   time.Now(),
		Price:  5,
		Artist: artist.Artist{Id: 1, Name: "Artiste1", Note: "Note Artiste1"},
	}

	if !isConcertEqual(deletedConcert, expectedConcert) {
		contextTest.Errorf("Le concert supprimé ne correspond pas au concert attendu.")
	}

}

func isConcertEqual(concert1, concert2 Concert) bool {
	return concert1.Id == concert2.Id &&
		concert1.Name == concert2.Name &&
		concert1.Date.Equal(concert2.Date) &&
		concert1.Price == concert2.Price
}
