<template>
  <v-app id="inspire">
   
    <v-app-bar app>

      <v-toolbar-title>Go post Like 👍🏻</v-toolbar-title>
    </v-app-bar>

    <v-main
    center
    >
    <br>
      <!--  -->
      <router-view/>
    </v-main>
    <keep-alive>

    <component v-bind:is="component"/>
    </keep-alive>
  </v-app>
</template>

<script>
import Bottomnav from './components/Bottomnav.vue'

export default {
  components:{
    Bottomnav
  },
  data() {
    return {
    drawer: null,
  email:null,
        component:"Bottomnav"

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
