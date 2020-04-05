<template>
  <div>
    <a-input-search
      placeholder="Find problem and ideas"
      v-model="searchQuery"
      @search="onSearch"
      enterButton
    />
    <br /><br />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
    data () {
        return {
            searchQuery: ''
        }
    },

    methods: {
      async onSearch() {
        await this.$axios.$get('/')
      }
    }
})
</script>

<style>

</style>
