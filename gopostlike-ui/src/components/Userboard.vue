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
    </v-main>
 </v-app>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      userdetails:[],
    }
  },

  // Fetches posts when the component is created.
  created() {
    
    axios.get(`http://localhost:9000/api/user/`+localStorage.getItem("email"),
    
    {
    headers: {
   Token: 'Bearer ' + localStorage.getItem("jwt") //the token is a variable which holds the token
  }
    }
    )
    .then(response => {
      // JSON responses are automatically parsed.
      this.userdetails = response.data
      console.log(this.userdetails,localStorage.getItem("jwt"))
    })
    .catch(e => {
      this.errors.push(e)
    })
  }
}
</script>
