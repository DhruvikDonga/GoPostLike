package Postcontrollers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	dbs "github.com/DhruvikDonga/goshopcart/DBs"
	"github.com/DhruvikDonga/goshopcart/models"
	"github.com/gorilla/mux"
)

//--------------
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)

}
func Generateslug(length int) string {
	return StringWithCharset(length, charset)
}

//---------------

//get all the posts
func GetPosts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content- Type", "application/json")
	var posts []models.Posts
	dbs.DB.Find(&posts)
	json.NewEncoder(w).Encode(posts)

}

//get individual post
func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content- Type", "application/json")
	params := mux.Vars(r)
	var post models.Posts
	dbs.DB.First(&post, params["slug"])
	json.NewEncoder(w).Encode(post)

}

//create a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "user" {
		w.Write([]byte("Not authorized."))
		//log.Println(r.Header.Get("Email"))
		return
	}
	w.Header().Set("Content- Type", "application/json")
	var post models.Posts
	log.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&post)
	post.Postslug = Generateslug(10)
	//log.Println(post)
	dbs.DB.Create(&post)
	json.NewEncoder(w).Encode(post)
}

//update post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content- Type", "application/json")
	params := mux.Vars(r)
	var post models.Posts
	dbs.DB.First(&post, params["id"])
	json.NewDecoder(r.Body).Decode(&post)
	dbs.DB.Save(&post)
	json.NewEncoder(w).Encode(post)
}

//delete a post
func DeletePosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content- Type", "application/json")
	params := mux.Vars(r)
	var post models.Posts
	dbs.DB.Delete(&post, params["id"])

	//permanent delte the record
	//dbs.DB.Unscoped().Delete(&user, params["id"])

	json.NewEncoder(w).Encode("The post is deleted successfully!")
}
