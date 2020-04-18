import { Middleware } from '@nuxt/types'

const authMiddleware: Middleware = ({ store, redirect, req }) => {
  if (!store.state.auth.token) {
    return redirect('/sign-in')
  }

  if (!store.state.auth.user.email || !store.state.auth.user.name) {
    // request to post /me
  }
}


export default authMiddleware
