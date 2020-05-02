<template functional>
  <a-list
    v-if="!!props.foundProblems.length"
    itemLayout="horizontal"
    :dataSource="props.foundProblems"
  >
    <a-list-item slot="renderItem" slot-scope="item, index">
      <a
        slot="actions"
        @click="parent.$router.push(`/problem/${item.slug}`)"
      >
        Details
      </a>
      <a-list-item-meta
        :description="item.description.slice(0, 20) + '...'"
      >
        <a slot="title" href="https://www.antdv.com/">
          {{ item.name }}
        </a>
      </a-list-item-meta>
    </a-list-item>
  </a-list>
  <a-alert
    v-else
    message="No results"
    description="Query returned 0 results"
    type="error"
    showIcon
  />
</template>

<script>
    export default {
        name: "ProblemsSearchResults",
        props: {
            foundProblems: {
                type: Array,
                default: []
            }
        }
    }
</script>

<style scoped>

</style>
