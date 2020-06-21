<template>
    <a-list
      v-if="ideas.length"
      class="comment-list"
      itemLayout="horizontal"
      :dataSource="ideas"
    >
      <a-list-item slot="renderItem" slot-scope="item, index">
        <a-comment author="Han Solo">
          <template slot="actions">
            <template v-if="moderable">
              <a-button-group>
                <a-button type="primary">Accept</a-button>
                <a-button type="danger">Discard</a-button>
              </a-button-group>
            </template>
            <template v-else>
              <span key="comment-basic-like">
                <a-tooltip title="Like">
                  <a-icon type="like" :theme="item.my_vote == 1 ? 'filled' : 'outlined'" @click="item.my_vote != 1 ? like(item.id) : null" />
                </a-tooltip>
              </span>
              <span style="padding: 0 18px 0 8px; cursor: auto">
                {{ item.score }}
              </span>
              <span key="comment-basic-dislike">
                <a-tooltip title="Dislike">
                  <a-icon
                    type="dislike"
                    :theme="item.my_vote == -1 ? 'filled' : 'outlined'"
                    @click="item.my_vote != -1 ? dislike(item.id) : null"
                  />
                </a-tooltip>
              </span>
            </template>
          </template>
          <div slot="content">
            <h5>What can I do?</h5>
            <p>{{item.action_description}}</p>
            <h5>What impact would my action have?</h5>
            <p>{{item.results_description}}</p>
            <h5>Budget</h5>
            <p>${{item.money_price}}</p>
            <h5>Needed time</h5>
            <p>
              <template v-if="item.time_price == -1">As much as possible</template>
              <template v-else>{{item.time_price}} minute{{item.time_price != 1 ? 's' : ''}}</template>
            </p>
          </div>
          <a-tooltip slot="datetime">
            <span>{{item.datetime}}</span>
          </a-tooltip>
        </a-comment>
      </a-list-item>
    </a-list>
</template>

<script>
    export default {
        name: "ProblemIdea",

        props: {
            ideas: {
                type: Array,
                default: []
            },
            moderable: {
                type: Boolean,
                default: false
            }
        },

        computed: {
            isLoggedIn () {
                return this.$store.getters['user/isLoggedIn']
            }
        },

        methods: {
            async like (idea_id) {
                await this.vote(idea_id, 1)
            },
            async dislike (idea_id) {
                await this.vote(idea_id, -1)
            },
            async vote (ideaId, delta) {
              try {
                  let { data } = this.$axios.post('/vote', {
                      idea_id: ideaId,
                      delta
                  }, {
                      headers: {
                          'Authorization': `Bearer ${this.$store.state.auth.token}`
                      }
                  })
                  // TODO Update in vuex state
                  this.$emit('vote', {
                      ideaId,
                      delta
                  })

              } catch (err) {
                  this.$notification.error({
                      message: 'Error',
                      description: err.response.data.message
                  });
              }
            }
        }
    }
</script>

<style scoped>

</style>
