import { GetterTree, ActionTree, MutationTree } from 'vuex'
import { RootState } from '~/store'

export const state = () => ({
  token: null,
  expiresIn: null,
  user: {
    email: null,
    name: null
  }
})

export type AuthState = ReturnType<typeof state>

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
  setCredentials({ commit }, { token, expiresIn }) {
    commit('SET_CREDENTIALS', { token, expiresIn })
  },
  setUserdata({ commit }, { email, name }) {
    commit('SET_USERDATA', { email, name })
  }
}
