<style scoped>
  .example {
    text-align: center;
    background: rgba(0, 0, 0, 0.05);
    border-radius: 4px;
    margin-bottom: 20px;
    padding: 30px 50px;
    margin: 20px 0;
  }
</style>
<template>
  <div class="example">
    <a-spin />
  </div>
</template>

<script>
    export default {
        name: "Provider",
        mounted () {
            this.$store.dispatch('auth/logout')
            this.$nuxt.$router.push('/')
        }
    }
</script>
