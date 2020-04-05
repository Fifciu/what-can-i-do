<template>
  <div class="main-page">
    <h1 class="page-title">What can I do?</h1>
    <p class="page-description">Discover methods which allow one person to has real impact to the World and our future!</p>
    <a-input-search
      placeholder="Find problem and ideas"
      v-model="searchQuery"
      @search="onSearch"
      enterButton
    />
    <br /><br />
      <TransitionFadeExpand>
        <a-list v-if="beenFinding && !!foundProblems.length" itemLayout="horizontal" :dataSource="foundProblems">
          <a-list-item slot="renderItem" slot-scope="item, index">
            <a
              slot="actions"
              @click="$router.push(`/problem/${item.id}`)"
            >Details</a>
            <a-list-item-meta
              :description="item.description.slice(0, 20) + '...'"
            >
              <a slot="title" href="https://www.antdv.com/">{{item.title}}</a>
            </a-list-item-meta>
          </a-list-item>
        </a-list>
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
          this.$router.push({
              query: {
                  ...this.$route.query,
                  searchQuery: this.searchQuery
              }
          })
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
    },

    async asyncData ({ query, error, $axios }) {
        try {
            if (!query || !query.searchQuery || !query.searchQuery.length) {
                return
            }
            let { problems } = await $axios.$get(`problems?searchQuery=${query.searchQuery}`)
            return {
                foundProblems: problems,
                beenFinding: true
            }
        } catch (err) {
            error({ statusCode: 404, message: err + 'Problem not found' })
        }
    },
})
</script>

<style>
  .page-title {
    font-size: 33px;
    font-weight: bold;
    text-align: center;
  }

  .page-description {
    text-align: center;
  }

  .main-page {
    padding: 24px;
  }
</style>
