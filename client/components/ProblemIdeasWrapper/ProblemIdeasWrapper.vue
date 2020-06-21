<template>
  <a-tabs default-active-key="1" v-if="isModerator">
    <a-tab-pane key="1">
      <span slot="tab">
        <a-icon type="apple" />
        Public
      </span>
      <ProblemIdeas :ideas="ideas" />
    </a-tab-pane>
    <a-tab-pane key="2">
      <span slot="tab">
        <a-icon type="android" />
        Waiting for review
      </span>
      <ProblemIdeas :ideas="ideas" moderable/>
    </a-tab-pane>
  </a-tabs>
  <ProblemIdeas :ideas="ideas" v-else />
</template>

<script>
    export default {
        name: "ProblemIdeasWrapper",
        components: {
            ProblemIdeas: () => import('~/components/ProblemIdeas/ProblemIdeas.vue'),
        },
        props: {
            ideas: {
                type: Array,
                default: []
            }
        },
        computed: {
            isModerator () {
                return this.$store.getters['auth/isModerator']
            }
        }
    }
</script>

<style scoped>

</style>
