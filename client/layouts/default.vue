<template>
    <a-layout id="components-layout-demo-top" class="layout">
      <a-layout-content class="main-wrapper">
        <div :style="{ background: '#fff', padding: '24px', height: '100%' }">
          <h1 class="page-title">What can I do?</h1>
          <p class="page-description">Discover methods which allow one person to has real impact to the World and our future!</p>
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

  .page-title {
    font-size: 33px;
    font-weight: bold;
    text-align: center;
  }

  .page-description {
    text-align: center;
  }
</style>
