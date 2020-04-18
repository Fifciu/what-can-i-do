<template>
  <div class="sign">
    <h2>Sign up/in</h2>
    <p>We only accept authenticate via external platforms. Joining to our society requires having an account in one of the services listed below. </p>
    <a-alert
      message="Error"
      description="Could not authenticate"
      type="error"
      class="my-14"
      showIcon
      v-if="$route.query && 'error' in $route.query"
    />
    <div class="sign-options">
      <div class="sign-option">
        <a-button type="primary" @click.native="googleInitAuth">Sign in via Google</a-button>
      </div>
    </div>
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

<style lang="scss" scoped>
.sign {
  padding: $cardPadding;
}
</style>
