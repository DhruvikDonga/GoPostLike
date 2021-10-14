<template>
   <form>
           <h3>Signup:-</h3>
<v-text-field
      v-model="username"
      label="Username"

      required
     
    ></v-text-field>
    <v-text-field
      v-model="firstname"
      label="First Name"

      required
     
    ></v-text-field>
    <v-text-field
      v-model="lastname"
      label="Last Name"

      required
     
    ></v-text-field>
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
    <v-text-field
      v-model="password_confirmation"
      :counter="10"
      label="Conform Password"
       type="password"

      required
    
    ></v-text-field>
    

    <v-btn
      class="mr-4"
      @click="submit"
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
      username:"",
      firstname : "",
      lastname: "",
      email : "",
      password : "",
      password_confirmation : "",
      error:"",
    }
  },
  methods: {
    submit(e) {
      e.preventDefault()

      if (this.password === this.password_confirmation && this.password.length > 0) {
      

        axios.post(`http://localhost:9000/api/signup`, {
          username:this.username,
          firstname: this.firstname,
        lastname: this.lastname,

          email: this.email,
          password: this.password,
        })
          .then(response => {
                console.log(response.data)
            if ( response.data== 1) {
                this.$router.push('login')

            } else {
                this.error = JSON.stringify(response.data)
            }
          })
          .catch(error => {
            console.error(error);
          });
      } else {
        this.password = ""
        this.passwordConfirm = ""

        return alert("Passwords do not match")
      }
    }
  }
}
</script>