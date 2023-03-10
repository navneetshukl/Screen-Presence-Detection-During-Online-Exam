package routes

import (
	"encoding/base64"
	"fmt"
	"go_modules/database"
	"go_modules/render"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)
var email string
func Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,"form.html")
}
func GetData(w http.ResponseWriter, r *http.Request){
	name:=r.FormValue("name")
	email=r.FormValue("email")
	invitation_code:=r.FormValue("invitation_code")
	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(invitation_code)
	database.InsertRow(name,email,invitation_code)
	http.Redirect(w, r, "/save", http.StatusSeeOther)
}
func WebCam(w http.ResponseWriter, r *http.Request) {
   render.RenderTemplate(w,"index.html")
}

func Save(w http.ResponseWriter, r *http.Request){
	image := r.FormValue("image")
	
		data, err := base64.StdEncoding.DecodeString(image[strings.IndexByte(image, ',')+1:])
		fmt.Println(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(data)
		timestamp:=time.Now().UTC()
		database.InsertCameraData(timestamp.String(),image,email)
		err = ioutil.WriteFile("image.jpeg", data, 0666)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Image saved!")
	}

	func Display(w http.ResponseWriter, r *http.Request){
		image, err := ioutil.ReadFile("image.jpeg")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(image)))
		w.Header().Set("Access-Control-Allow-Origin", "*")    // Allow CORS from any domain
		w.Header().Set("Access-Control-Allow-Methods", "GET") // Allow GET requests only
		w.Write(image)
	}
	func GetAllData(w http.ResponseWriter, r *http.Request){
		database.Getdata()
	}

















	