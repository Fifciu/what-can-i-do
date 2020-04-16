<template>
  <div>
    <a-alert
      message="Error"
      description="Could not authenticate"
      type="error"
      showIcon
      v-if="$route.query && 'error' in $route.query"
    />
    <a-button @click.native="googleInitAuth">Google auth</a-button>
  </div>
</template>

<script>
    export default {
        name: "JoinUs",
        methods: {
            async googleInitAuth () {
                try {
                    let response = await this.$axios.post('/auth/init/google')
                    const redirectUrl = response.data.redirectUrl
                    window.location.href = redirectUrl
                } catch (err) {
                    console.log(err)
                }
            }
        }
    }
</script>

<style scoped>

</style>
