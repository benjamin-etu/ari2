package artist

import (
	"testing"
)

func TestArtistInit(contextTest *testing.T) {
	ArtistInit()

	if len(artists) == 0 {
		contextTest.Errorf("La liste d'artistes n'a pas été initialisée correctement")
	}
}

func TestAddArtist(contextTest *testing.T) {
	newArtist := Artist{
		Id:   3,
		Name: "TestArtist",
		Note: "Ceci est un artiste de test",
	}

	AddArtist(newArtist)

	if len(artists) == 0 {
		contextTest.Errorf("L'artiste n'a pas été ajouté correctement")
	}
}

func TestDeleteArtistById(contextTest *testing.T) {
	length := len(artists)
	deletedArtist, err := DeleteArtistById(1)
	if err != nil {
		contextTest.Errorf("Erreur lors de la suppression de l'artiste: %v", err)
	}

	if len(artists) != length-1 {
		contextTest.Errorf("L'artiste n'a pas été supprimé correctement")
	}
	expectedArtist := Artist{
		Id:   1,
		Name: "Gazo",
		Note: "La mala est toujours gangs",
	}

	if !isArtistEqual(deletedArtist, expectedArtist) {
		contextTest.Errorf("L'artiste supprimé ne correspond pas à l'artiste attendu.")
	}
}

func isArtistEqual(artist1, artist2 Artist) bool {
	return artist1.Id == artist2.Id && artist1.Name == artist2.Name && artist1.Note == artist2.Note
}
