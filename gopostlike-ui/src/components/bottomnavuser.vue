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

   
  </v-bottom-navigation>
</template>

<script>

export default {
  
  data() {
    return {
  email:null,
    }
  },
  computed: {
gettoken() {

    return localStorage.getItem('jwt')

},


  },
  created(){

    this.email= localStorage.getItem('email').slice(1, -1)

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