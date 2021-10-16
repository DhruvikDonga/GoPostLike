<template>
   <form>
           <h3>Lets create new post {{user}}!</h3>
<v-text-field
      v-model="postname"
      label="postname"

      required
     
    ></v-text-field>
    <v-text-field
      v-model="description"
      label="post description"

      required
     
    ></v-text-field>

    <v-file-input accept="image/*" 
                label="Select files"
                   prepend-icon="mdi-camera"

                multiple chips color="pink"
                v-model="postimage"
                ></v-file-input>
          
    

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
      postname:"",
      description : "",
      lastname: "",
     postimage: [],
    readers: [],
      error:"",
      user:localStorage.getItem("email").slice(1,-1),
    }
  },
  methods: {
    submit(e) {
      e.preventDefault()
        console.log('files', this.postimage,localStorage.getItem("jwt"))
          const formData = new FormData()
          var postjson={
              postname:this.postname,
          description: this.description,
          imgno:this.postimage.length
          }
          
          const blob = new Blob([JSON.stringify(postjson)], {
                type: 'application/json'
            }); 
              formData.append('post', blob)
          for( var i = 0; i < this.postimage.length; i++ ){
              let file = this.postimage[i];
          formData.append('postimage', file);
        }



        axios.defaults.headers.common = {
          'Token': `${localStorage.getItem("jwt")}`,

        }
        axios.post(`http://localhost:9000/api/createpost`, 
            formData,
            )
          .then(response => {
                console.log(response.data)
            if ( response.data== 1) {
                this.$router.push(localStorage.getItem("email").slice(1,-1))

            } else {
                this.error = JSON.stringify(response.data)
            }
          })
          .catch(error => {
            console.error(error);
          });
     
    }
  }
}
</script>