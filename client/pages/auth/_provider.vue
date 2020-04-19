<style scoped>
  .example {
    text-align: center;
    background: rgba(0, 0, 0, 0.05);
    border-radius: 4px;
    margin-bottom: 20px;
    padding: 30px 50px;
    margin: 20px 0;
  }
</style>
<template>
  <div class="example">
    <a-spin />
  </div>
</template>

<script>
    export default {
        name: "Provider",
        async mounted () {
            const supportedProviders = process.env.supported_oauth_providers.split(',')
            if (!supportedProviders.includes(this.$route.params.provider)) {
                this.$nuxt.$router.go('/')
            }

            try {
                let url = `/auth/complete/${this.$route.params.provider}`
                let first = true
                for (const [key, value] of Object.entries(this.$route.query)) {
                    if (first) {
                        url += '?'
                        first = false
                    } else {
                        url += '&'
                    }
                    url += `${key}=${value}`
                }
                let { data } = await this.$axios.post(url)
                const user = data.user
                this.$store.dispatch('auth/setCredentials', {
                    expiresAt: data.expires_at,
                    token: data.token
                })
                this.$store.dispatch('auth/setUserdata', {
                    email:  user.email,
                    name: user.name
                })

                this.$store.dispatch('auth/setCookieTokenFromState')
                const savedRedirectType = localStorage.getItem('back-type')
                const savedRedirectSlug = localStorage.getItem('back-slug')
                if (savedRedirectSlug && savedRedirectType) {
                    this.$nuxt.$router.push(`/${savedRedirectType}/${savedRedirectSlug}`)
                } else {
                    this.$nuxt.$router.push('/account')
                }

                // Autorefresh/autoremove cookie from jwt
                // POST /me endpoint

            } catch (err) {
                console.log(err.response.data)
                this.$nuxt.$router.push('/sign-in?error')
            }
        }
    }
</script>
