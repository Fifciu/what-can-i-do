import { Middleware } from '@nuxt/types'

const authMiddleware: Middleware = async ({ store, redirect, isServer }) => {
  let token = store.state.auth.token
  let expiresAt = store.state.auth.expiresAt

  if (!store.state.auth.token || !store.state.auth.expiresAt) {
    token = window.localStorage.getItem(<string>process.env.ls_token_key)
    expiresAt = window.localStorage.getItem(<string>process.env.ls_expires_key)
    await store.dispatch('auth/setCredentials', {
      token,
      expiresAt
    })

    if (!token) {
      return redirect('/sign-in')
    }
  }

  if (!store.state.auth.user.email || !store.state.auth.user.name) {
    // request to post /me
    const success = await store.dispatch('auth/fetchAndSetUserdata', {
      token
    })
    if (!success) {
      return redirect('/logout')
    }
  }
}

// TODO Only MyAccount tries to refresh token, apply it everywhere

export default authMiddleware
