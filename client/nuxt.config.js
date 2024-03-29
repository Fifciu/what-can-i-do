require('dotenv').config({path: '../server/.env'})

export default {
  mode: 'spa',
  /*
  ** Headers of the page
  */
  head: {
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: process.env.npm_package_description || '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css2?family=Ubuntu:wght@300;500;700&display=swap' }
    ]
  },

  router: {
    middleware: 'loginIfPossible'
  },

  server: {
    port: process.env.pwa_port,
    host: process.env.pwa_host
  },

  // pageTransition: 'fade',
  /*
  ** Customize the progress-bar color
  */
  loading: { color: '#fff' },
  /*
  ** Global CSS
  */
  css: [
    'ant-design-vue/dist/antd.css',
    '~/assets/scss/util.scss',
    '~/assets/scss/transitions.scss'
  ],
  /*
  ** Plugins to load before mounting the App
  */
  plugins: [
    '@/plugins/antd-ui',
    '@/plugins/global-transition'
  ],
  /*
  ** Nuxt.js dev-modules
  */
  buildModules: [
    '@nuxt/typescript-build',
      [
        '@nuxtjs/dotenv',
        {
          path: '../server'
        }
      ]
  ],
  /*
  ** Nuxt.js modules
  */
  modules: [
    // Doc: https://axios.nuxtjs.org/usage
    '@nuxtjs/axios',
    '@nuxtjs/pwa',
    '@nuxtjs/style-resources',
    '@nuxtjs/proxy'
  ],
  proxy: {
    '/api': {
      target: `http://${process.env.api_host}:${process.env.api_port}`,
      pathRewrite: {
        '^/api' : '/'
      }
    }
  },
  styleResources: {
    scss: [
      '~/assets/scss/variables.scss',
    ]
  },
  /*
  ** Axios module configuration
  ** See https://axios.nuxtjs.org/options
  */
  axios: {
    baseURL: `http://${process.env.pwa_host}:${process.env.pwa_port}/api/` // cuz of proxy
  },
  /*
  ** Build configuration
  */
  build: {
    /*
    ** You can extend webpack config here
    */
    extend (config, ctx) {
    }
  }
}
