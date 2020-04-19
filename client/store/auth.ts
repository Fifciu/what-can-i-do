import { GetterTree, ActionTree, MutationTree } from 'vuex'
import { RootState } from '~/store'
import Cookie from 'js-cookie'

export const state = () => ({
  token: '',
  expiresAt: '',
  user: {
    email: '',
    name: ''
  }
})

export type AuthState = ReturnType<typeof state>

export const getters: GetterTree<AuthState, RootState> = {
  isLoggedIn: state => !!state.token,
  token: state => state.token
}

export const mutations: MutationTree<AuthState> = {
  SET_CREDENTIALS: (state, { token, expiresAt }) => {
    state.token = token
    state.expiresAt = expiresAt
  },
  SET_USERDATA: (state, { email, name }) => {
    state.user.email = email
    state.user.name = name
  }
}

export const actions: ActionTree<AuthState, RootState> = {

  setCredentials({ commit }, { token, expiresAt}) {
    commit('SET_CREDENTIALS', { token, expiresAt })
  },

  setUserdata({ commit }, { email, name }) {
    commit('SET_USERDATA', { email, name })
  },

  setCookieTokenFromState({ state }) {
    const jwtOffset = Number(process.env.jwt_offset) || 0
    const uselessTokenDate = new Date(new Date().getTime() + state.expiresAt + jwtOffset * 60 * 1000);
    Cookie.set('token', state.token, { expires: uselessTokenDate})
    Cookie.set('token_expires_at', state.expiresAt, { expires: uselessTokenDate})
  },

  setStateTokenFromCookie({ commit }) {
    const token = Cookie.get('token')
    const expiresAt = Cookie.get('token_expires_at')
    commit('SET_CREDENTIALS', { token, expiresAt })
  },

  logout ({ commit }) {
    commit('SET_CREDENTIALS', { token: '', expiresAt: '' })
    commit('SET_USERDATA', { email: '', name: '' })
    Cookie.remove('token')
    Cookie.remove('token_expires_at')
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
      let response = await this.$axios.post('auth/refresh', {
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
