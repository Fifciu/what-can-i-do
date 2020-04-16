<template>
  <div>
    Welciome
    {{ $store.state.auth }}

    {{ $store.getters['auth/isLoggedIn'] }}
  </div>
</template>

<script>
    export default {
        name: "Profile",
        middleware: 'auth'
    }
</script>

<style scoped>

</style>
