package backend

import (
	"net/http"
	"fmt"
	"Groupie/controller"
	"Groupie/structure"
	"Groupie/temp"
)

func Angele(w http.ResponseWriter, r *http.Request) {
	apiToken := controller.GetAccessToken()
	apiUrl := "https://api.spotify.com/v1/artists/3QVolfxko2UyCOtexhVTli"
	data := structure.Artist{}
	controller.RequestApi(apiUrl, apiToken, &data)
	apiUrl = "https://api.spotify.com/v1/artists/3QVolfxko2UyCOtexhVTli/top-tracks"
	data2 := structure.TopTracks{}
	controller.RequestApi(apiUrl, apiToken, &data2)
	fmt.Println(data, data2)
	data3 := structure.Bigdata{data, data2}
	templates.Temp.ExecuteTemplate(w, "artist", data3)
}

func DamsoArtist(w http.ResponseWriter, r *http.Request) {
	apiToken := controller.GetAccessToken()
	apiUrl := "https://api.spotify.com/v1/artists/2UwqpfQtNuhBwviIC0f2ie"
	data := structure.Artist{}
	controller.RequestApi(apiUrl, apiToken, &data)
	apiUrl = "https://api.spotify.com/v1/artists/2UwqpfQtNuhBwviIC0f2ie/top-tracks"
	data2 := structure.TopTracks{}
	controller.RequestApi(apiUrl, apiToken, &data2)
	fmt.Println("Pp : ", data.Images[0].URL)
	fmt.Println("Nom : ", data.Name)
	fmt.Println("Follower : ", data.Followers.Total)
	data3 := structure.Bigdata{data, data2}
	templates.Temp.ExecuteTemplate(w, "artist", data3)
}



func LafeveArtist(w http.ResponseWriter, r *http.Request) {
	apiToken := controller.GetAccessToken()
	apiUrl := "https://api.spotify.com/v1/artists/2sBKOwN0fSjx39VtL2WpjJ"
	data := structure.Artist{}
	controller.RequestApi(apiUrl, apiToken, &data)
	apiUrl = "https://api.spotify.com/v1/artists/2sBKOwN0fSjx39VtL2WpjJ/top-tracks"
	data2 := structure.TopTracks{}
	controller.RequestApi(apiUrl, apiToken, &data2)
	fmt.Println("Pp : ", data.Images[0].URL)
	fmt.Println("Nom : ", data.Name)
	fmt.Println("Follower : ", data.Followers.Total)
	data3 := structure.Bigdata{data, data2}
	templates.Temp.ExecuteTemplate(w, "artist", data3)
}

func Index (w http.ResponseWriter, r *http.Request) {
	apiToken := controller.GetAccessToken()
	apiUrl := "https://api.spotify.com/v1/artists/2UwqpfQtNuhBwviIC0f2ie"
	listdata := []structure.Artist{}
	data := structure.Artist{}
	controller.RequestApi(apiUrl, apiToken, &data)
	listdata = append(listdata, data)
	fmt.Println(data)
	apiUrl = "https://api.spotify.com/v1/artists/3QVolfxko2UyCOtexhVTli"
	data = structure.Artist{}
	controller.RequestApi(apiUrl, apiToken, &data)
	listdata = append(listdata, data)
	fmt.Println(data)
	apiUrl = "https://api.spotify.com/v1/artists/2sBKOwN0fSjx39VtL2WpjJ"
	data = structure.Artist{}
	controller.RequestApi(apiUrl, apiToken, &data)
	listdata = append(listdata, data)
	fmt.Println(data)
	fmt.Println(listdata)
	for _, data := range listdata {
		fmt.Println(data.Name)
		fmt.Println(data.Images[0].URL)
	}
	templates.Temp.ExecuteTemplate(w, "index", listdata)
	
}