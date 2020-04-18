import { Middleware } from '@nuxt/types'

const authMiddleware: Middleware = async ({ store, redirect }) => {
  if (!store.state.auth.token) {
    return redirect('/sign-in')
  }

  if (!store.state.auth.user.email || !store.state.auth.user.name) {
    // request to post /me
    const success = await store.dispatch('auth/fetchAndSetUserdata', {
      token: store.state.auth.token
    })
    if (!success) {
      return redirect('/logout')
    }
  }
}


export default authMiddleware
