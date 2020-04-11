<template>
    <a-layout id="components-layout-demo-top" class="layout">
      <a-layout-header>
        <a-menu
          theme="dark"
          mode="horizontal"
          :selectedKeys="selectedNavItem"
          :style="{ lineHeight: '64px' }"
        >
          <a-menu-item key="1">
            <nuxt-link to="/">
              Home
            </nuxt-link>
          </a-menu-item>
          <a-menu-item key="2">About</a-menu-item>
          <a-menu-item key="3">Join us</a-menu-item>
          <a-menu-item key="4">Sign in</a-menu-item>
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
          <nuxt />
        </div>
      </a-layout-content>
      <a-layout-footer style="text-align: center">
        Whatcanido.club ©2020 Created by <a href="https://fifciuu.com" target="_blank">Fifciuu</a>
      </a-layout-footer>
    </a-layout>
</template>

<script>
  export default {
      computed: {
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
            if (this.$route.name == 'index') {
                return ['1']
            }
            return []
          }
      }
  }
</script>

<style>
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
    padding: 0 15px;
  }

</style>
