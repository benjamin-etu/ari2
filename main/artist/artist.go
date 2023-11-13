package artist

import (
	"errors"
)

var artists []Artist

type Artist struct {
	Id   int
	Name string
	Note string
}

func ArtistInit() {
	artist1 := Artist{
		Id:   1,
		Name: "Gazo",
		Note: "La mala est toujours gangs",
	}
	artist2 := Artist{
		Id:   2,
		Name: "Damso",
		Note: "Nouvel album Geya",
	}
	artists = append(artists, artist1, artist2)
}

func AddArtist(artistAdd Artist) {
	artists = append(artists, artistAdd)
}

func GetArtist() []Artist {
	return artists
}

func DeleteArtistById(id int) (Artist, error) {
	var indexToDelete int = -1
	var artistToDelete Artist

	for i, artist := range artists {
		if artist.Id == id {
			indexToDelete = i
			artistToDelete = artist
			break
		}
	}

	if indexToDelete >= 0 {
		artists = append(artists[:indexToDelete], artists[indexToDelete+1:]...)
		return artistToDelete, nil
	}

	return Artist{}, errors.New("artiste non supprimé")
}
