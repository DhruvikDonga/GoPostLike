<template>
 <v-app>
  <h2 align="center">User :-{{userdetails.firstname}} </h2>


    <v-card
      class="d-flex justify-center mb-6"
      flat
      tile
    >
      <v-card
       
        class="pa-2"
        tile
      >
        <h3>Followers:-{{userdetails.totalfollowers}}</h3>
      </v-card>
      <v-card
       
        class="pa-2"
        tile
      >
      <h3>Posts:-{{userdetails.totalposts}}</h3>
      </v-card>
      
    </v-card>
    <v-main
    center
    >
    <br>
    <br>
    <v-card
         v-if="userdetails.totalposts==0"
         class="mx-auto"
    max-width="344"
    align="center"
    max-height="344"
        >
        <br>
        <h1>No posts yet</h1>
      <br>
    </v-card>  

   <v-container  v-if="userdetails.totalposts>0">
      <v-flex v-for="(post) in posts" :key="post.id">
        
  <v-card
    class="mx-auto"
    max-width="344"
    v-if="post.postdata!=null"
  >
    <v-carousel>
    <v-carousel-item
      v-for="(image,i) in post.postimages"
      :key="i"
      :src="`http://localhost:9000/`+image.image"
      reverse-transition="fade-transition"
      transition="fade-transition"
    ></v-carousel-item>
  </v-carousel>

    <v-card-title>
      {{post.postdata.postname}}
      
    </v-card-title>

    <v-card-subtitle>
      Likes:-{{post.postdata.likes}}
    </v-card-subtitle>
      <v-card-text>
          {{post.postdata.description}}
        </v-card-text>
    


  </v-card>
      </v-flex>
   </v-container>

    </v-main>
 </v-app>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      userdetails:{},
      show: false,

      posts:null,
     
    }
  },

  // Fetches posts when the component is created.
  created() {
    //console.log(localStorage.getItem("email"))

    axios.get(`http://localhost:9000/api/user/`+this.$route.params.username,
    
    {
    headers: {
   Token: 'Bearer ' + localStorage.getItem("jwt") //the token is a variable which holds the token
  }
    }
    )
    .then(response => {
      // JSON responses are automatically parsed.
      this.userdetails = response.data
    // console.log(this.userdetails.ID)
    })
    .catch(e => {
      this.errors.push(e)
    })

    


  },
  mounted() {
    this.getuserpost();

},

methods: {
    getuserpost() {
      //get user posts
    axios.get(`http://localhost:9000/api/getuserposts/`+this.$route.params.username)
    .then(res => {
      // JSON responses are automatically parsed.
      
     const post=JSON.parse(JSON.stringify(res.data))
     //console.log(JSON.parse(JSON.stringify(res.data))[0])
     this.posts = post
          console.log(this.posts)

    })
    .catch(e => {
      this.errors.push(e)
    })
    }

}
}
</script>
