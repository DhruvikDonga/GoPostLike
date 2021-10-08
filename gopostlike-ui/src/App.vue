<template>
  <v-app id="inspire">
    <v-navigation-drawer
      v-model="drawer"
      app
    >
     <v-list-item v-if="gettoken!=null">
        

        <v-list-item-content>
          <v-list-item-title><b>Welcome </b> <br><i>{{email}}</i></v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <v-divider></v-divider>

      <v-list dense>
        <v-list-item
         link
        >
          <v-list-item-icon>
            <v-icon>mdi-home</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>Home</v-list-item-title>
          </v-list-item-content>
        </v-list-item>


      <v-list-item
         link
         v-if="gettoken==null"
        >
          <v-list-item-icon>
            <v-icon>mdi-login</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>Login</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

       <v-list-item
         link
         v-if="gettoken!=null"
         @click="logout"
        >
          <v-list-item-icon>
            <v-icon>mdi-logout</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>Logout</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-list-item
         link
          v-if="gettoken==null"

        >
          <v-list-item-icon>
            <v-icon> mdi-open-in-new</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>Register</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      
      
      
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>Go post Like 👍🏻</v-toolbar-title>
    </v-app-bar>

    <v-main
    center
    >
    <br>
      <!--  -->
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
    drawer: null,
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
