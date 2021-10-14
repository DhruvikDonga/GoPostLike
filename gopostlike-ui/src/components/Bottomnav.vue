<template>
  <v-bottom-navigation
    color="teal"
      fixed
      horizontal
  >
    <v-btn link
         :to = "{ name:'home' }">
      <span>Home</span>

      <v-icon>mdi-home</v-icon>
    </v-btn>

    <v-btn link
              v-if="gettoken==null"

         :to = "{ name:'login' }">
      <span>Login</span>

      <v-icon>mdi-login</v-icon>
    </v-btn> 

    <v-btn link
              v-if="gettoken==null"

         :to = "{ name:'register' }">
      <span>Register</span>

      <v-icon>mdi-open-in-new</v-icon>
    </v-btn>
    <v-btn link
              v-if="gettoken!=null"

         :to = "{ name:'userboard' , params: { username:  username }}">
      <span>Profile</span>

      <v-icon>mdi-account</v-icon>
    </v-btn>

      <v-btn link
                v-if="gettoken!=null"

         @click="logout">
      <span>Logout</span>

      <v-icon>mdi-logout</v-icon>
    </v-btn>
  </v-bottom-navigation>
</template>

<script>

export default {
  
  data() {
    return {
  username:null,
    }
  },
  computed: {
gettoken() {

    return localStorage.getItem('jwt')

},


  },
  created(){

    this.username= localStorage.getItem('email').slice(1, -1)

  },
  methods: {
    
    logout(e) {
            e.preventDefault()

    localStorage.removeItem('jwt')

    this.$router.push('/')
}
  }
}
</script>