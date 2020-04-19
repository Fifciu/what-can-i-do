<template>
  <div class="profile">
    <h1>Hello {{ $store.state.auth.user.name }}</h1>
    <p>Here you can check information about your account and stats of posted by you ideas.</p>
    <a-divider orientation="left">Proposed problems</a-divider>
    <div v-if="!problems.length">
      You have not proposed any new problem to solve yet. You can do it on Home's page!
    </div>
    <a-list size="small" bordered :dataSource="problems" v-else>
      <a-list-item slot="renderItem" slot-scope="item, index">{{ item.name }}</a-list-item>
    </a-list>
    <a-divider orientation="left">Proposed ideas</a-divider>
    <div v-if="!ideas.length">
      You have not proposed any new idea yet. You can do it on any problem's page!
    </div>
    <a-list size="small" bordered :dataSource="ideas" v-else>
      <a-list-item slot="renderItem" slot-scope="item, index">
        <div class="ideas__item">
          <div>
            {{ item.action_description }} against <strong>{{ item.problem_name }}</strong>
          </div>
          <a-icon type="check" v-if="item.is_published"/>
          <a-icon type="clock-circle" v-else/>
        </div>
      </a-list-item>
    </a-list>
  </div>
</template>

<script>
    export default {
        name: "Profile",
        middleware: 'auth',

        async asyncData({ store, $axios }) {
            try {
                const token = store.getters['auth/token']
                if (!token) {
                    throw 'No token'
                }
              let [problemsResponse, ideasResponse] = await Promise.all([
                  $axios.post('problems/mine', {}, {
                      headers: {
                          Authorization: `Bearer ${token}`
                      }
                  }),
                  $axios.post('ideas/mine', {}, {
                      headers: {
                          Authorization: `Bearer ${token}`
                      }
                  })
              ])

                return {
                    problems: problemsResponse.data.problems,
                    ideas: ideasResponse.data.ideas
                }

            } catch (err) {
                return {
                    problems: [],
                    ideas: []
                }
            }
        }
    }
</script>

<style lang="scss" scoped>
  .profile {
    padding: $cardPadding;
  }

  .ideas {
    &__item {
      display: flex;
      justify-content: space-between;
      align-items: center;
      width: 100%;
    }
  }
</style>
