<template>
  <div class="profile">
    <h1>Hello {{ $store.state.auth.user.name }}</h1>
    <p>Here you can check information about your account and stats of posted by you ideas.</p>
  </div>
</template>

<script>
    export default {
        name: "Profile",
        middleware: 'auth'
    }
</script>

<style lang="scss" scoped>
  .profile {
    padding: $cardPadding;
  }
</style>
