<template>
  <div class="problem">

    <div class="problem__details">
      <div class="problem__description">
        <h2 class="problem__section-heading">{{ problem.name }}</h2>
        <p>{{ problem.description }}</p>
      </div>
    </div>

    <div class="problem__ideas">
      <h2 class="problem__section-heading">Ideas</h2>
      <ProblemIdeasWrapper
        v-if="problem.ideas && !!problem.ideas.length"
        :ideas="problem.ideas"
        @vote="afterVote"
      />
      <div
        v-else
        class="problem__no-ideas"
        message="No ideas"
        description="There are no ideas to solve this problem. Be the first who shares an brilliant idea!"
        type="warning"
        showIcon
      >
        There are no ideas to solve this problem. Be the first who shares an <strong>brilliant</strong> idea!
      </div>

      <TransitionFadeExpand>
        <AddIdea
          v-if="showAddIdea"
          @added="addedIdea = true"
        />
      </TransitionFadeExpand>

      <div class="problem__add-btn-wrapper">
        <div v-if="!isLoggedIn">
          <h3>Do you want to add own idea?</h3>
          <a-button type="primary" class="problem__add-btn" html-type="submit">
            <nuxt-link :to="`/sign-in?back-type=problem&back-slug=${problem.slug}`">
              Sign in at first
            </nuxt-link>
          </a-button>
        </div>
        <a-button type="primary" class="problem__add-btn" @click.native="showAddIdea = true" v-else-if="!showAddIdea">
          <template v-if="addedIdea">Add another idea</template>
          <template v-else>Add an idea</template>
        </a-button>
      </div>
    </div>
  </div>
</template>

<script>
    import moment from 'moment'

    export default {
        name: "ProblemView",
        components: {
            ProblemIdeasWrapper: () => import('~/components/ProblemIdeasWrapper/ProblemIdeasWrapper.vue'),
            AddIdea: () => import('~/components/AddIdea/AddIdea.vue'),
        },
        data () {
            return {
                showAddIdea: false,
                addedIdea: false
            }
        },
        computed: {
            isLoggedIn () {
                return this.$store.getters['auth/isLoggedIn']
            }
        },
        methods: {
            afterVote ({ ideaId, delta }) {
                const idea = this.problem.ideas.find(idea => idea.id == ideaId)
                if (idea) {
                    if (idea.my_vote != 0) {
                        idea.score -= idea.my_vote
                    }
                    idea.score += delta
                    idea.my_vote = delta
                }
            }
        },
        async asyncData ({ store, params, error, $axios }) {
            try {
                const token = store.getters['auth/token']
                let { data } = await $axios.get(`/problems/${params.slug}/ideas`, !token ? {} : {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                })
                store.dispatch('breadcrumbs/changeTitle', data.problem.name)
                return {
                     problem: {
                         ...data.problem,
                         ideas: data.problem.ideas.map(idea => ({
                             ...idea,
                             datetime: moment().subtract(1, 'days').fromNow()
                         }))
                     }
                }
            } catch (err) {
                error({ statusCode: 404, message: err + 'Problem not found' })
            }
        }
    }
</script>

<style lang="scss" scoped>
  .problem {
    &__details {

     }

    &__title {
      color: #fff;
      font-weight: bold;
      margin: 0;
    }

    &__ideas {
      padding: 10px 30px;
    }

    &__no-ideas {
      padding: 1em 1em 0;
      margin: 0;
    }

    &__add-btn {
      margin: 20px 0;
    }

    &__add-btn-wrapper {
      text-align: center;
    }

    &__add-ideas-form {
      padding: 18px;
    }

    &__heading {
      display: flex;
      align-items: center;
      margin-bottom: 5px;
      color: #f0f0f0;
      background: dodgerblue;
      padding: 24px;
    }

    &__back-icon {
      margin-right: 15px;
      font-size: 1.5em;
    }

    &__description {
      font-size: 18px;
      line-height: 1.6;
      padding: 10px 30px;
    }

    &__section-heading {
      font-size: 30px;
    }
  }
</style>
