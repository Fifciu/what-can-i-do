<template>
    <a-layout id="components-layout-demo-top" class="layout">
      <a-layout-header>
        <a-menu
          theme="dark"
          mode="horizontal"
          :defaultSelectedKeys="['0']"
          :style="{ lineHeight: '64px' }"
        >
          <a-menu-item key="1">Home</a-menu-item>
          <a-menu-item key="2">About</a-menu-item>
          <a-menu-item key="3">Join Us</a-menu-item>
        </a-menu>
      </a-layout-header>
      <a-layout-content class="main-wrapper">
        <div :style="{ background: '#fff', height: '100%' }">
          <nuxt />
        </div>
      </a-layout-content>
      <a-layout-footer style="text-align: center">
        Whatcanido ©2020 Created by <a href="https://fifciuu.com" target="_blank">Fifciuu</a>
      </a-layout-footer>
    </a-layout>
</template>

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
    padding: 0 35px;
  }

</style>
