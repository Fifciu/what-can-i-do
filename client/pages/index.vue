<template>
  <div class="home">
    <h1 class="home__title">What can I do?</h1>
    <p class="home__description">Discover methods which allow one person to has real impact to the World and our future!</p>
    <a-input-search
      placeholder="Find problems and ideas"
      v-model="searchQuery"
      @search="onSearch"
      enterButton
    />
    <br /><br />
      <TransitionFadeExpand>
        <ProblemsSearchResults
          v-if="beenFinding"
          :foundProblems="foundProblems"
        />
      </TransitionFadeExpand>
    <AddProblem
      v-if="beenFinding"
      class="problem__add"
    />
    </div>
</template>

<script>

export default {

    components: {
        ProblemsSearchResults: () => import('~/components/ProblemsSearchResults/ProblemsSearchResults.vue'),
        AddProblem: () => import('~/components/AddProblem/AddProblem.vue'),
    },

    head () {
        return {
            title: 'Home | WhatCanIdo.club'
        }
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
                beenFinding: true,
                searchQuery: query.searchQuery
            }
        } catch (err) {
            return {
                foundProblems: [],
                beenFinding: true,
                searchQuery: query.searchQuery
            }
        }
    },
}
</script>

<style lang="scss" scoped>
  .home {
    padding: $cardPadding;
    margin: 0 auto;
    padding-top: 50px;
    padding-bottom: 50px;
    max-width: 544px;

    &__title {
      font-size: 2.42em;
      font-weight: bold;
      text-align: center;
      margin-bottom: 2px;
    }

    &__description {
      text-align: center;
      font-size: 1.28em;
    }
  }

  .problem {
    &__add {
      margin-top: 15px;
    }
  }
</style>

