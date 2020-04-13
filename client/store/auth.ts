import { GetterTree, ActionTree, MutationTree } from 'vuex'
import { RootState } from '~/store'
import Cookie from 'js-cookie'

export const state = () => ({
  token: '',
  expiresIn: '',
  user: {
    email: '',
    name: ''
  }
})

export type AuthState = ReturnType<typeof state>

export const getters: GetterTree<AuthState, RootState> = {
  isLoggedIn: state => !!state.token,
}

export const mutations: MutationTree<AuthState> = {
  SET_CREDENTIALS: (state, { token, expiresIn }) => {
    state.token = token
    state.expiresIn = expiresIn
  },
  SET_USERDATA: (state, { email, name }) => {
    state.user.email = email
    state.user.name = name
  }
}

export const actions: ActionTree<AuthState, RootState> = {
  setCredentials({ commit }, { token, expiresIn}) {
    commit('SET_CREDENTIALS', { token, expiresIn })
  },
  setUserdata({ commit }, { email, name }) {
    commit('SET_USERDATA', { email, name })
  },
  setCookieTokenFromState({ state }) {
    const jwtOffset = Number(process.env.jwt_offset) || 0
    const uselessTokenDate = new Date(new Date().getTime() + state.expiresIn + jwtOffset * 60 * 1000);
    Cookie.set('token', state.token, { expires: uselessTokenDate})
    Cookie.set('token_expires_in', state.expiresIn, { expires: uselessTokenDate})
  },
  setStateTokenFromCookie({ commit }) {
    const token = Cookie.get('token')
    const expiresIn = Cookie.get('token_expires_in')
    commit('SET_CREDENTIALS', { token, expiresIn })
  }
}
