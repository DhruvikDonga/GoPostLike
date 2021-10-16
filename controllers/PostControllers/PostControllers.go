package Postcontrollers

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
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

//get  posts by a user
func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content- Type", "application/json")

	var getuser models.User
	params := mux.Vars(r)
	userposts := []models.PostComplete{}
	var post []models.Posts
	postimage := []models.PostImage{}
	dbs.DB.Model(&getuser).Where("username=?", params["username"]).First(&getuser)

	dbs.DB.Model(&post).Where("user_id=?", getuser.ID).Find(&post)

	for _, v := range post {
		dbs.DB.Model(&postimage).Where("posts_id=?", v.ID).Find(&postimage)

		var temp models.PostComplete
		temp.Postmodel = v
		temp.Imagemodel = postimage
		userposts = append(userposts, temp)
	}
	json.NewEncoder(w).Encode(userposts)

}

//create a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content- Type", "application/json")
	log.Println("Create post function:-", r.Header.Get("Email"))

	if r.Header.Get("Role") != "user" {
		w.Write([]byte("Not authorized."))
		return
	}
	var getuser models.User
	dbs.DB.Model(&getuser).Where("username= ?", r.Header.Get("email")).First(&getuser)

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
		post.UserID = int(getuser.ID)
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
			filename := part.FileName()
			filename = strconv.Itoa(postimageid) + "-" + filename
			log.Println("URL:", part.FileName())
			img := models.PostImage{
				PostsID: postimageid,
				Image:   "/" + filename,
			}

			outfile, err := os.Create("./storage/postimage/" + filename)
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

	if postimages != nil {
		dbs.DB.Create(&postimages)
	}

	var wholepost models.PostComplete
	wholepost.Postmodel = post
	wholepost.Imagemodel = postimages

	//update post no
	getuser.TotalPosts += 1
	dbs.DB.Model(&getuser).Update("total_posts", getuser.TotalPosts)

	json.NewEncoder(w).Encode("post created successfully")

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
