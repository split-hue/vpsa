package redovalnica //Ksenija Korošec, 63220163

import (
	"fmt"
)

// Student je definicija strukture študenta.
// Sestavljen je iz Imena, Priimka in ocene.
type Student struct { // definiramo strukturo študenta
	Ime     string //vse daš na VELIKE začetnice da bo JAVNO
	Priimek string
	Ocene   []int
}

// DodajOceno doda oceno, s podano vpisno in min/max.
// Če ni v obmnočju izpiše napako.
func DodajOceno(studenti map[int]Student, vpisnaStevilka int, ocena int, minOcena int, maxOcena int) { // vse imena fun daš na VELIKE začetnice da bo JAVNO
	if ocena < minOcena || ocena > maxOcena {
		fmt.Printf("Ocnea %d ni vrljavna, vnesi oceno med %d in %d.\n", ocena, minOcena, maxOcena)
		return
	}
	student, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		fmt.Printf("Študent pod vpisno št. %d ne obstaja.\n", vpisnaStevilka)
		return
	}

	student.Ocene = append(student.Ocene, ocena)
	studenti[vpisnaStevilka] = student
}

func povprecje(studenti map[int]Student, vpisnaStevilka int, stOcen int) float64 { //v navodilih: ta OSTANE skrita :D
	student, obstaja := studenti[vpisnaStevilka] //ne damo ji zato komentarjev
	if !obstaja {
		return -1.0
	}
	if len(student.Ocene) < stOcen {
		return 0.0
	}
	vsota := 0 //var vsota int = 0
	for _, ocena := range student.Ocene {
		vsota += ocena
	}
	return float64(vsota) / float64(len(student.Ocene))
}

// IzpisVsehOcen izpiše celo redovalnico ocen.
// Prav tako komu pripadajo z njegovo vpisno št.
func IzpisVsehOcen(studenti map[int]Student) {
	fmt.Print("REDOVALNICA:\n")
	for vpisna, student := range studenti {
		fmt.Printf("%d - %s %s: %v\n", vpisna, student.Ime, student.Priimek, student.Ocene)
	}
	fmt.Println()
}

// IzpisiKoncniUspeh izpiše končni učni uspeh.
// Poleg števila izpiše tudi komentar oz. ali je premalo ocen.
func IzpisiKoncniUspeh(studenti map[int]Student, stOcen int) {

	for vpisnaStevilka, student := range studenti { //Ana Novak: povprečna ocena 9.0 -> Odličen študent!
		uspeh_bes := povprecje(studenti, vpisnaStevilka, stOcen) //kličn funk zgori, k rača tud če ne ubstaja pa to
		var komentar string                                      //rabm prej da lah ta kumentar shranm

		if len(student.Ocene) < stOcen {
			fmt.Printf("%s %s: premalo ocen (%d/%d)\n",
				student.Ime, student.Priimek, len(student.Ocene), stOcen)
			continue
		}

		if uspeh_bes == -1.0 {
			komentar = "Študent ne obstaja."
		} else if uspeh_bes >= 9 {
			komentar = "Odličen študent!"
		} else if uspeh_bes >= 6 {
			komentar = "Povprečen študent"
		} else {
			komentar = "Neuspešen študent"
		}
		fmt.Printf("%s %s: povprečna ocena %.1f -> %s\n", student.Ime, student.Priimek,
			uspeh_bes, komentar)
	}

}
