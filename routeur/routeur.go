package routeur

import (
	"fmt"
	"log"
	"net/http"
	"Groupie/backend"
)

func InitServ() {


	http.HandleFunc("/Angèle", backend.Angele)
	http.HandleFunc("/AngeleSon", backend.Angele)
	http.HandleFunc("/Damso", backend.DamsoArtist)
	http.HandleFunc("/La%20Fève", backend.LafeveArtist)
	http.HandleFunc("/index", backend.Index)

	css := http.FileServer(http.Dir("./css"))
	http.Handle("/static/", http.StripPrefix("/static/",css))


	log.Println("serveur lancé")
	fmt.Println("http://localhost:8080/index")
	log.Fatal(http.ListenAndServe(":8080", nil))
}