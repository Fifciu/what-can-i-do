<template>
  <div>
    <a-input-search
      placeholder="Find problem and ideas"
      v-model="searchQuery"
      @search="onSearch"
      enterButton
    />
    <br /><br />
      <TransitionFadeExpand>
        <div v-if="beenFinding && !!foundProblems.length">
          Found
        </div>
        <a-alert
          v-else-if="beenFinding"
          message="No results"
          description="Query returned 0 results"
          type="error"
          showIcon
        />
      </TransitionFadeExpand>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
import TransitionFadeExpand from '~/components/TransitionFadeExpand.vue'

export default Vue.extend({

    components: {
        TransitionFadeExpand
    },

    data () {
        return {
            searchQuery: '',
            foundProblems: [],
            beenFinding: false
        }
    },

    methods: {
      async onSearch() {
          if (!this.beenFinding) {
              this.beenFinding = true
          }

          try {
            let { problems } = await this.$axios.$get(`problems?searchQuery=${this.searchQuery.trim()}`)
            this.foundProblems = problems
        } catch (err) {
            this.foundProblems = []
            console.log(err)
        }
      }
    }
})
</script>

<style>

</style>
