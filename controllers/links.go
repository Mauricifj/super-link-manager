package controllers

import (
	"github.com/super-link-manager/database"
	"github.com/super-link-manager/network"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	links := database.Links()

	message := network.Get(w, r)

	data := map[string]interface{}{
		"Message": message,
		"Links" : links,
	}

	templates.ExecuteTemplate(w, "Index", data)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		linkType := r.FormValue("linkType")
		name := r.FormValue("name")
		price := r.FormValue("price")

		convertedPrice, err := strconv.Atoi(price)
		if err != nil {
			log.Println("Error on converting price:", err)
		}

		response := network.CreateLink(linkType, name, convertedPrice)
		if response.Link != nil {
			if database.CreateLink(response.Link.Id, linkType, name, convertedPrice) {
				network.Set(w, r, "New Link Generated Successfully!", "success")
			} else {
				network.DeleteLink(response.Link.Id)
				network.Set(w, r, "Error on link creation...", "danger")
			}
			http.Redirect(w, r, "/", 301)
		} else if response.Errors != nil {
			templates.ExecuteTemplate(w, "New", response.Errors)
		}
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if network.DeleteLink(id) && database.DeleteLink(id) {
		network.Set(w, r, "Link Deleted Successfully!", "success")
	} else {
		network.Set(w, r, "Error on link deletion...", "danger")
	}

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	link := network.LinkById(id)

	if link != nil {
		templates.ExecuteTemplate(w, "Edit", link)
	} else {
		network.Set(w, r, "Link not found!", "danger")
		http.Redirect(w, r, "/", 301)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		linkType := r.FormValue("linkType")
		name := r.FormValue("name")
		price := r.FormValue("price")

		convertedPrice, err := strconv.Atoi(price)
		if err != nil {
			log.Println("Error on converting price:", err)
		}

		response := network.UpdateLink(id, linkType, name, convertedPrice)
		if response.Link != nil {
			if database.UpdateLink(response.Link.Id, linkType, name, convertedPrice) {
				network.Set(w, r, "Link Updated Successfully!", "success")
			} else {
				network.DeleteLink(response.Link.Id)
				network.Set(w, r, "Error on updating link...", "danger")
			}
			http.Redirect(w, r, "/", 301)
		} else if response.Errors != nil {
			templates.ExecuteTemplate(w, "Edit", response.Errors)
		}
	}
}