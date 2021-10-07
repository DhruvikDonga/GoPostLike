<template>
  <v-card
    class="mx-auto"
    max-width="344"
  >
  <h1>User Dashboard </h1>
<v-btn
      class="mr-4"
      @click="logout"
    >
      logout
    </v-btn>
  </v-card>
</template>

<script>

export default {
  
  methods: {
    logout(e) {
            e.preventDefault()

      console.log(localStorage.getItem('jwt'))
    localStorage.removeItem('jwt')

    this.$router.push('/')
}
  }


}
</script>