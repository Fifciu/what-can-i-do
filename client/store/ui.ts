import { GetterTree, ActionTree, MutationTree } from 'vuex'
import { RootState } from '~/store'

export const state = () => ({
  modalVisibility: false
})

export type UiState = ReturnType<typeof state>

export const mutations: MutationTree<UiState> = {
  CHANGE_MODAL_VISIBILITY: (state, value: boolean) => (state.modalVisibility = value)
}

export const actions: ActionTree<UiState, RootState> = {
  changeModalVisibility({ commit }, value: boolean) {
    commit('CHANGE_MODAL_VISIBILITY', value)
  }
}
