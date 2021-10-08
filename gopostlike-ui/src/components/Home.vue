<template>
  <v-card
    class="mx-auto"
    max-width="344"
  >
    <v-card-text>
      <div>Word of the Day</div>
      <p class="text-h4 text--primary">
        {{posts[0].description}}
      </p>
      <p>WTF</p>
      <div class="text--primary">
        {{posts[1].description}}
      </div>
    </v-card-text>
    <v-card-actions>
      <v-btn
        text
        color="teal accent-4"
      >
        Learn More
      </v-btn>
    </v-card-actions>

   
  </v-card>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      posts: [],
    }
  },

  // Fetches posts when the component is created.
  created() {
    
    axios.get(`http://localhost:9000/api/welcome`)
    .then(response => {
      // JSON responses are automatically parsed.
      this.posts = response.data
    })
    .catch(e => {
      this.errors.push(e)
    })
  }
}
</script>
