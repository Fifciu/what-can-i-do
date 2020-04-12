<template>
  <div>{{ user }}</div>
</template>

<script>
    export default {
        name: "Provider",
        async asyncData({ $axios, route, redirect, env }) {
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
                // console.log(url)
                let response = await $axios.post(url)
                console.log(response.data)
                const user = response.data.user
                return {
                    user
                }
            } catch (err) {
                console.log(err.response)
                return {
                    user: 'lol'
                }
            }
        }
    }
</script>

<style scoped>

</style>
