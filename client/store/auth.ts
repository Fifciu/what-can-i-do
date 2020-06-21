import { GetterTree, ActionTree, MutationTree } from 'vuex'
import { RootState } from '~/store'

const FLAGS = {
  MODERATOR: 1<<0
}

export const state = () => ({
  token: '',
  expiresAt: '',
  refreshTokenInterval: 0,
  flags: 0,
  user: {
    email: '',
    name: ''
  }
})

export type AuthState = ReturnType<typeof state>

export const getters: GetterTree<AuthState, RootState> = {
  isLoggedIn: state => !!state.token,
  token: state => state.token,
  isModerator: (state): Boolean => !!(state.flags & FLAGS.MODERATOR)
}

export const mutations: MutationTree<AuthState> = {
  SET_CREDENTIALS: (state, { token, expiresAt }) => {
    state.token = token
    state.expiresAt = expiresAt
  },
  SET_FLAGS: (state, { token }) => {
    try {
      const tokenParts = token.split('.')
      const userObject = JSON.parse(window.atob(tokenParts[1]))
      state.flags = userObject.flags
    } catch (err) {
      console.log("Couldn't set flags from token. Bad token! ", err)
    }
  },
  SET_USERDATA: (state, { email, name }) => {
    state.user.email = email
    state.user.name = name
  },
  SET_REFRESH_INTERVAL: (state, timeoutIdentifier: number) => {
    state.refreshTokenInterval = timeoutIdentifier
  },
  UNSET_REFRESH_INTERVAL: (state) => {
    if (state.refreshTokenInterval) {
      clearTimeout(state.refreshTokenInterval)
      state.refreshTokenInterval = 0
    }
  }
}

export const actions: ActionTree<AuthState, RootState> = {

  async setCredentials({ state, commit, dispatch }, { token, expiresAt }): Promise<void> {
    commit('SET_CREDENTIALS', { token, expiresAt })
    commit('SET_FLAGS', { token })

    if (!state.refreshTokenInterval) {

      const diff = new Date(expiresAt).getTime() - new Date().getTime()
      if (diff < 1) {
        return
      }

      const refreshLogicFunc = async () => {

        if (state.refreshTokenInterval) {
          return
        }

        let refreshOffset = Number(process.env.jwt_refresh_offset) * 1000
        if (!refreshOffset) {
          refreshOffset = 10000
        }

        const refresh = async () => {
          await dispatch('refresh', { token })
          refreshLogicFunc()
        }
        if (refreshOffset >= diff) {
          return await refresh()
        }
        const timeoutIdentifier = setTimeout(refresh, diff - refreshOffset)
        commit('SET_REFRESH_INTERVAL', timeoutIdentifier)
      }

      await refreshLogicFunc()
    }
  },

  setUserdata({ commit }, { email, name }) {
    commit('SET_USERDATA', { email, name })
  },

  logout ({ commit }) {
    commit('SET_CREDENTIALS', { token: '', expiresAt: '' })
    commit('SET_USERDATA', { email: '', name: '' })
    window.localStorage.removeItem(<string>process.env.ls_token_key)
    window.localStorage.removeItem(<string>process.env.ls_expires_key)
    commit('UNSET_REFRESH_INTERVAL')
  },

  async fetchAndSetUserdata({ dispatch, commit }, { token }): Promise<Boolean> {
    try {
      let response = await this.$axios.post('me', {}, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      })
      const { email, fullname } = response.data.user
      commit('SET_USERDATA', { email, name: fullname })
      return true
    } catch (err) {
      console.log(err)
      if (err.response.status == 401) {
        // Try to refresh
        return await dispatch('refresh', { token })
      } else {
        return false
      }
    }
  },

  async refresh({ commit }, { token }): Promise<Boolean> {
    try {
      let response = await this.$axios.post('auth/refresh', {}, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      })
      commit('SET_CREDENTIALS', { token: response.data.token, expiresAt: response.data.expires_at })
      return true
    } catch (err) {
      return false
    }
  }

}
