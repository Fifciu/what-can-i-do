import {ActionTree} from "~/node_modules/vuex";
import {BreadcrumbsState} from "~/store/breadcrumbs";
const cookieparser = process.server ? require('cookieparser') : undefined

export const state = () => ({

})

export type RootState = ReturnType<typeof state>

export const actions: ActionTree<BreadcrumbsState, RootState> = {
  nuxtServerInit ({ dispatch, commit }, { req }) {
    let token = null
    let expiresAt = null
    if (req.headers.cookie) {
      const parsed = cookieparser.parse(req.headers.cookie)
      try {
        token = parsed.token
        expiresAt = parsed.token_expires_at

        dispatch('auth/setCredentials', {
          token,
          expiresAt
        }, {
          root: true
        })
      } catch (err) {
        console.error(err)
        // No valid cookie found
      }
    }
  }
}
