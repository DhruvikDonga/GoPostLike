<template>

  <form>
    
    <v-text-field
      v-model="email"
      label="E-mail"
       type="email"

      required
     
    ></v-text-field>
    <v-text-field
      v-model="password"
      :counter="10"
      label="Password"
       type="password"

      required
    
    ></v-text-field>

    

    <v-btn
      class="mr-4"
      @click="handlesubmit"
    >
      submit
    </v-btn>
    <v-alert
    v-if="error"
  color="red"
  elevation="24"
  type="warning"
>{{error}}</v-alert>
  </form>
  
</template>

<script>
import axios from 'axios';

export default {
  data () {
    return {
      email : "",
      password : "",
      error:"",
    }
  },
  methods: {

    handlesubmit(e) {
      e.preventDefault()
      if (this.password.length > 0) {
          
        axios.post(`http://localhost:9000/api/signin`, {
          email: this.email,
          password: this.password
        })
        .then(response => {
          let usertype = response.data.role

          if (response.data.token!=null){


          localStorage.setItem('email',JSON.stringify(response.data.email))
          localStorage.setItem('jwt',response.data.token)

          if (localStorage.getItem('jwt') != null) {
            this.$emit('loggedIn')
            if (this.$route.params.nextUrl != null) {
              this.$router.push(this.$route.params.nextUrl)
            }
            else {
              if (usertype == "admin") {
                this.$router.push('admin')
              }
              else {
                this.$router.push(localStorage.getItem("email").slice(1, -1))
              }
            }
          }
          }
          else {
            this.error = response.data
          }
        })
        .catch(function (error) {
          console.error(error.response);
        });
      }
    }
  }
}
</script>

