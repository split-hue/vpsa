# Info

Ta projekt predstavlja implementacijo preproste redovalnice v programskem jeziku **Go**.  
Sestavljen je iz dveh delov:

- Paket `redovalnica`, ki vsebuje funkcije za upravljanje ocen študentov.
- CLI aplikacija `main.go`, ki uporablja paket `redovalnica` in omogoča njegovo uporabo preko ukazne vrstice.

Vse javne funkcije in ena skrita modula:
```go
// DodajOceno doda oceno, s podano vpisno in min/max.
// Če ni v obmnočju izpiše napako.
func DodajOceno(studenti map[int]Student, vpisnaStevilka int, ocena int, minOcena int, maxOcena int)
```
```go
// IzpisVsehOcen izpiše celo redovalnico ocen.
// Prav tako komu pripadajo z njegovo vpisno št.
func IzpisVsehOcen(studenti map[int]Student)
```
```go
// IzpisiKoncniUspeh izpiše končni učni uspeh.
// Poleg števila izpiše tudi komentar oz. ali je premalo ocen.
func IzpisiKoncniUspeh(studenti map[int]Student, stOcen int)
```
```go
func povprecje(studenti map[int]Student, vpisnaStevilka int, stOcen int) float64 
```
### Struktura
```
vpsa/
│
├── go.mod
├── go.sum
├── main.go
│
└── redovalnica/
    └── redovalnica.go
```
## Namestitev
Namestitev paketa `redovalnica`.
```bash
go get github.com/split-hue/vpsa/redovalnica
```
Namestitev CLI knjižnice za gradnjo ukaznovrstnega vmesnika.
```bash
go get github.com/urfave/cli/v2
```
### Uporaba
Uporaba paketa `redovalnica`.
```bash
import "github.com/split-hue/vpsa/redovalnica"
```





### Zagon
`stOcen` je minimum ocen potrebnih za pozitiven uspeh (privzeta vrednost: **6**)

`minOcena` je najmanjša dovoljena ocena (privzeta vrednost: **0**)

`maxOcena` je največja dovoljena ocena (privzeta vrednost: **10**)

Primer uporabe za zagon:
```sh
go run . --stOcen 6 --minOcena 5 --maxOcena 10
```


## Licenca
Projekt je odprtokoden.
