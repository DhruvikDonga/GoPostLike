package Postcontrollers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
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
	var posts models.Posts
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
	w.Header().Set("Content- Type", "application/json")

	if r.Header.Get("Role") != "user" {
		w.Write([]byte("Not authorized."))
		//log.Println(r.Header.Get("Email"))
		return
	}
	mr, err := r.MultipartReader()
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	var post models.Posts
	var postimageid, no int
	part, err := mr.NextPart()
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	// JSON 'doc' part
	if part.FormName() == "post" {
		jsonDecoder := json.NewDecoder(part)
		jsonDecoder.Decode(&post)
		post.Postslug = Generateslug(20)
		dbs.DB.Create(&post)
		postimageid = int(post.ID)
		no = post.Images
		log.Println(postimageid)
	}

	//post images
	postimages := []models.PostImage{}
	for k := 0; k < no; k++ {
		part, err := mr.NextPart()
		if err != nil {
			json.NewEncoder(w).Encode(post)
			return
		}

		if part.FormName() == "postimage" {

			//fmt.Println("URL:", part.FileName())
			img := models.PostImage{
				PostsID: postimageid,
				Image:   part.FileName(),
			}

			outfile, err := os.Create("./storage/postimage/" + part.FileName())
			if err != nil {
				json.NewEncoder(w).Encode("Internal server error")
				return
			}
			defer outfile.Close()

			_, err = io.Copy(outfile, part)

			if err != nil {
				json.NewEncoder(w).Encode("Internal server error")
				return
			}
			postimages = append(postimages, img)

		}

	}
	log.Println("VVSVSVSFSVSSVS")
	fmt.Println("No shit:-", postimages)
	if postimages != nil {
		dbs.DB.Create(&postimages)
	}

	var wholepost models.PostComplete
	wholepost.Postmodel = post
	wholepost.Imagemodel = postimages

	json.NewEncoder(w).Encode(wholepost)

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
