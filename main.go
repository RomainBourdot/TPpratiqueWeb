package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func main() {

	fileServer := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	temp, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Println(fmt.Printf("ERREUR => %v", err.Error()))
		os.Exit(02)
	}

	type Etudiant struct {
		Nom    string
		Prenom string
		Age    int
		Sexe   string
	}

	type PagePromo struct {
		NomDeClasse  string
		Filiere      string
		Niveau       string
		NbrEtudiant  int
		ListEtudiant []Etudiant
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		etudiant1 := Etudiant{Nom: "Amumu", Prenom: "Tom", Age: 18, Sexe: "Homme"}
		etudiant2 := Etudiant{Nom: "Coquette", Prenom: "Tomyette", Age: 19, Sexe: "Femme"}

		classe := PagePromo{NomDeClasse: "B1 Informatique", Filiere: "Informatique", Niveau: "B1", NbrEtudiant: 2,
			ListEtudiant: []Etudiant{etudiant1, etudiant2}}

		temp.ExecuteTemplate(w, "promo", classe)
	})
	type PageChange struct {
		IsPair  bool
		Counter int
	}

	var Counter int

	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		Counter++
		data := PageChange{Counter: Counter, IsPair: Counter%2 == 0}

		temp.ExecuteTemplate(w, "change", data)
	})
	http.ListenAndServe("localhost:8000", nil)
}
