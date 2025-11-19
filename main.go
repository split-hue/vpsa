package main

import (
	"fmt"
	"log"
	redovalnica "naloga05/redovalnica"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Redovalnica",
		Usage: "Program za upravljanje ocen študentov.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "Najmanjše število ocen za pozitivno oceno",
				Value: 6,
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "Najnižja dovoljena ocena",
				Value: 0,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "Najvišja dovoljena ocena",
				Value: 10,
			},
		},
		Action: func(ctx *cli.Context) error {

			minOcena := ctx.Int("minOcena")
			maxOcena := ctx.Int("maxOcena")
			stOcen := ctx.Int("stOcen")

			//drug ime, če ne prog meša
			studenti_3 := map[int]redovalnica.Student{ //popravljen na VELIKE začetnice
				63220001: {Ime: "Metka", Priimek: "Keber", Ocene: []int{10, 9, 8, 10, 9, 9}},
				63220002: {Ime: "Bobi", Priimek: "Kralj", Ocene: []int{6, 7, 5, 8, 6, 6}},
				63220003: {Ime: "Matija", Priimek: "Kos", Ocene: []int{4, 5, 3, 5}},
			}

			//package.Bla(...)
			redovalnica.DodajOceno(studenti_3, 63220001, 10, minOcena, maxOcena) // pravilni vnos
			redovalnica.DodajOceno(studenti_3, 63220001, -3, minOcena, maxOcena) //nevejlavna ocena
			redovalnica.DodajOceno(studenti_3, 11111111, 8, minOcena, maxOcena)  //ne obstaja
			redovalnica.IzpisVsehOcen(studenti_3)
			redovalnica.IzpisiKoncniUspeh(studenti_3, stOcen)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println()
		log.Fatal(err)
	}

}
