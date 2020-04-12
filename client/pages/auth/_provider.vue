<template>
  <div>{{ user }}</div>
</template>

<script>
    export default {
        name: "Provider",
        async asyncData(context) {
            const { $axios, route, redirect, env, store } = context
            const supportedProviders = env.supported_oauth_providers.split(',')
            if (!supportedProviders.includes(route.params.provider)) {
                redirect('/')
            }

            try {
                let url = `/auth/complete/${route.params.provider}`
                let first = true
                for (const [key, value] of Object.entries(route.query)) {
                    if (first) {
                        url += '?'
                        first = false
                    } else {
                        url += '&'
                    }
                    url += `${key}=${value}`
                }
                let response = await $axios.post(url)
                const respData = response.data
                const user = response.data.user
                store.dispatch('auth/setCredentials', {
                    expiresIn: respData.expires_in,
                    token: respData.token
                })
                store.dispatch('auth/setUserdata', {
                    email:  user.email,
                    user: user.name
                })

                redirect('/profile')
            } catch (err) {
                console.log(err.response)
                redirect('/join-us?error')
            }
        }
    }
</script>

<style scoped>

</style>
