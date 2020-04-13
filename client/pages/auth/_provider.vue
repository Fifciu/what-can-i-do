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
                    expiresIn: data.expires_in,
                    token: data.token
                })
                this.$store.dispatch('auth/setUserdata', {
                    email:  user.email,
                    name: user.name
                })

                this.$store.dispatch('auth/setCookieTokenFromState')
                this.$nuxt.$router.push('/profile')

            } catch (err) {
                console.log(err.response.data)
                this.$nuxt.$router.push('/join-us?error')
            }
        }
    }
</script>
