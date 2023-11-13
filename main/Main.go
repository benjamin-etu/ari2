package main

import (
	"concert/main/artist"
	"concert/main/concert"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/* ------ CONCERT ------*/
func displayConcert(ctx *gin.Context) {

	var affichage []gin.H

	for _, concert := range concert.GetConcert() {
		concertData := gin.H{
			"id":     concert.Id,
			"nom":    concert.Name,
			"date":   concert.Date,
			"prix":   concert.Price,
			"artist": concert.Artist,
		}
		affichage = append(affichage, concertData)
	}

	ctx.JSON(http.StatusOK, affichage)
}

func deleteConcertById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide, veuillez mettre un nombre"})
		return
	}
	concertDelete, errConcert := concert.DeleteConcertById(id)

	if errConcert == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Concert supprimé avec succès", "concert supprimé : ": concertDelete})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Concert non trouvé", "Id du concert demandé ": id})
	}
}

func initConcert() {
	concert.ConcertInit()
}

/* ------ ARTIST ------*/

func initArtist() {
	artist.ArtistInit()
}

func displayArtist(ctx *gin.Context) {
	var affichage []gin.H

	for _, artist := range artist.GetArtist() {
		artistData := gin.H{
			"id":   artist.Id,
			"nom":  artist.Name,
			"note": artist.Note,
		}
		affichage = append(affichage, artistData)
	}

	ctx.JSON(http.StatusOK, affichage)
}

func deleteArtistById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide, veuillez mettre un nombre"})
		return
	}
	artistDelete, errConcert := artist.DeleteArtistById(id)

	if errConcert == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "artiste supprimé avec succès", "artiste supprimé : ": artistDelete})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "artiste non trouvé", "Id du artiste demandé ": id})
	}
}

func main() {
	initConcert()
	initArtist()
	router := gin.Default()
	router.GET("/concerts", displayConcert)
	router.GET("/deleteConcert/:id", deleteConcertById)
	router.GET("/artists", displayArtist)
	router.GET("/deleteArtist/:id", deleteArtistById)

	router.Run(":8080")
}
