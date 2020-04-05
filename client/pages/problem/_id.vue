<template>
  <div class="problem">

    <div class="problem__details">
      <h2 class="problem__title">{{ problem.title }}</h2>
      <p class="problem__description">{{ problem.description }}</p>
    </div>

    <div class="problem__ideas">
      <a-list
        v-if="problem.ideas.length"
        class="comment-list"
        itemLayout="horizontal"
        :dataSource="problem.ideas"
      >
        <a-list-item slot="renderItem" slot-scope="item, index">
          <a-comment author="Han Solo">
            <template slot="actions">
      <span key="comment-basic-like">
        <a-tooltip title="Like">
          <a-icon type="like" :theme="action === 'liked' ? 'filled' : 'outlined'" @click="like" />
        </a-tooltip>
        <span style="padding-left: '8px';cursor: 'auto'">
          {{likes}}
        </span>
      </span>
              <span key="comment-basic-dislike">
        <a-tooltip title="Dislike">
          <a-icon
            type="dislike"
            :theme="action === 'disliked' ? 'filled' : 'outlined'"
            @click="dislike"
          />
        </a-tooltip>
        <span style="padding-left: '8px';cursor: 'auto'">
          {{dislikes}}
        </span>
      </span>
              <span key="comment-basic-reply-to">Fake news</span>
            </template>
            <p slot="content">{{item.description}}</p>
            <a-tooltip slot="datetime">
              <span>{{item.datetime}}</span>
            </a-tooltip>
          </a-comment>
        </a-list-item>
      </a-list>
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

      <div class="problem__add-btn-wrapper">
        <a-button type="primary" class="problem__add-btn">Add an idea</a-button>
      </div>
    </div>
  </div>
</template>

<script>
  import moment from 'moment'
    export default {
        name: "ProblemView",
        async asyncData ({ params, error, $axios }) {
            try {
                let { data } = await $axios.get(`/problems/${params.id}/ideas`)
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
      color: #f0f0f0;
      background: dodgerblue;
      padding: 24px;
     }

    &__title {
      color: #fff;
      font-weight: bold;
    }

    &__ideas {
      padding: 0 6px 6px;
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
  }
</style>
