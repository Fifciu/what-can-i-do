<template>
    <a-layout id="components-layout-demo-top" class="layout">
      <a-layout-header>
        <div class="logo">
          <nuxt-link to="/">Whatcanido</nuxt-link>
        </div>
        <div class="logo logo--mobile">
          <nuxt-link to="/">WCID</nuxt-link>
        </div>
        <a-menu
          theme="dark"
          mode="horizontal"
          :selectedKeys="selectedNavItem"
          :style="{ lineHeight: '64px' }"
        >
          <a-menu-item
            v-for="item in menuItems"
            :key="item.name"
          >
            <nuxt-link
              v-if="item.path"
              :to="item.path"
            >
              {{ item.name }}
            </nuxt-link>
            <span
              v-else-if="item.action"
              @click="item.action"
              class="clickable"
            >
              {{ item.name }}
            </span>
          </a-menu-item>
        </a-menu>
      </a-layout-header>

      <a-layout-content class="main-wrapper">
        <a-breadcrumb style="margin: 0px 0px 16px" v-if="showBreadrcrumbs">
          <a-breadcrumb-item
            v-for="item in breadcrumbs"
            :key="item.title"
            @click.native="'path' in item ? $router.push(item.path) : null"
          >
            {{ item.title }}
          </a-breadcrumb-item>
        </a-breadcrumb>
        <div :style="{ background: '#fff', height: '100%' }">
          <nuxt/>
        </div>
      </a-layout-content>
      <a-layout-footer style="text-align: center">
        Whatcanido.club ©2020 Created by <a href="https://fifciuu.com" target="_blank">Fifciuu</a>
      </a-layout-footer>

      <BaseModal
        v-if="modalVisibility"
        title="Review feedback"
        @handleOk="onSubmitFromModal"
      >
        <template #content>
          <a-textarea placeholder="Why do you want to discard this idea?" allow-clear v-model="modalMessage" />
        </template>
      </BaseModal>
    </a-layout>
</template>

<script>
import { mapState } from 'vuex'

  export default {
  data() {
    return {
      modalMessage: ''
    }
  },
  components: {
    BaseModal: () => import('~/components/Base/BaseModal/BaseModal.vue')
  },

    watch: {
      modalVisibility(value) {
        const el = document.querySelector('body')
        if (value) el.classList.add('no-scroll')
        else el.classList.remove('no-scroll')
      }
    },

      computed: {
        ...mapState({
          modalVisibility: state => state.ui.modalVisibility
        }),
          isLoggedIn () {
              return this.$store.getters['auth/isLoggedIn']
          },
          showBreadrcrumbs () {
              const allowedRoutes = [
                  'problem-id'
              ]
              return allowedRoutes.includes(this.$route.name)
          },
          breadcrumbs () {
              const breadcrumbs = [
                  {
                      title: 'Home',
                      path: '/'
                  }
              ]
              if (!!this.$store.state.breadcrumbs.currentRouteTitle) {
                  breadcrumbs.push({
                      title: this.$store.state.breadcrumbs.currentRouteTitle
                  })
              }
              return breadcrumbs
          },

          selectedNavItem () {
            const index = this.menuItems.findIndex(item => this.$route.path == item.path) + 1
              if (index > 0) {
                  return [String(index)]
              }
            return []
          },

          menuItems () {
              const items = [
                  {
                      name: 'About',
                      path: '/about'
                  }
              ]

              if (this.isLoggedIn) {
                  items.push(
                      {
                        name: 'My account',
                        path: '/account'
                      },
                      {
                          name: 'Logout',
                          action: () => {
                              this.$store.dispatch('auth/logout')
                              if (this.$route.name == 'account') {
                                  this.$router.push('/')
                              }
                          }
                      }
                  )
              } else {
                  items.push({
                      name: 'Sign up/in',
                      path: '/sign-in'
                  })
              }

              return items
          }
      },

    methods: {
      onSubmitFromModal () {
        if (!this.modalMessage) {
          this.$notification.error({
            message: 'Please provide a valid message',
          });
          return
        }
        this.$store.state.ui.onSubmitFromModal(this.modalMessage)
      }
    }
  }
</script>

<style lang="scss">
  #components-layout-demo-top {
    min-height: 100vh;
  }
  #components-layout-demo-top .logo {
    width: 120px;
    height: 31px;
    background: rgba(255, 255, 255, 0.2);
    margin: 16px 24px 16px 0;
    float: left;
  }

  .main-wrapper {
    padding: 20px;
  }

  .ant-layout-header {
    padding: 0 10px;
  }

  .ant-menu-item {
    padding: 0 10px;
    @media screen and (min-width: 600px) {
      padding: 0 15px;
    }
  }

  .clickable {
    cursor: pointer;
  }

  #components-layout-demo-top .logo {
    display: none;
    align-items: center;
    justify-content: center;
    font-weight: bold;
    font-family: 'Ubuntu', sans-serif;
    border-radius: 15px;
    margin: 16px 4px 16px 14px;
    font-size: 1.2em;

    a {
      color: #f0f0f0;
      &:hover {
        color: #fff;
      }
    }

    &--mobile {
      width: 65px;
      display: flex;
    }

    @media screen and (min-width: 500px) {
      display: flex;
      &--mobile {
        display: none;
      }
    }
  }

  * {
    box-sizing: border-box;
  }

</style>
