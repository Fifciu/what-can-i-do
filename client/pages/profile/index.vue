<template>
  <div>
    Welciome
    {{ $store.state.auth }}
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
