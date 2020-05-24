<template functional>
  <a-list
    v-if="props.ideas.length"
    class="comment-list"
    itemLayout="horizontal"
    :dataSource="props.ideas"
  >
    <a-list-item slot="renderItem" slot-scope="item, index">
      <a-comment author="Han Solo">
        <template slot="actions">
          <span key="comment-basic-like">
            <a-tooltip title="Like">
              <a-icon type="like" :theme="item.my_vote == 1 ? 'filled' : 'outlined'" @click="like" />
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
                @click="dislike"
              />
            </a-tooltip>
          </span>
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
            }
        }
    }
</script>

<style scoped>

</style>
