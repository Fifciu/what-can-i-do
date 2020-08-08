import { GetterTree, ActionTree, MutationTree } from 'vuex'
import { RootState } from '~/store'

export const state = () => ({
  modalVisibility: false,
  modalButtonLoader: false,
  onSubmitFromModal: (message: string) => {}
})

export type UiState = ReturnType<typeof state>

export const mutations: MutationTree<UiState> = {
  CHANGE_MODAL_VISIBILITY: (state, value: boolean) => (state.modalVisibility = value),

  CHANGE_BUTTON_LOADER: (state, value: boolean) => (state.modalButtonLoader = value),

  SET_SUBMIT_FROM_MODAL: (state, handler: () => {}) => (state.onSubmitFromModal = handler)
}

export const actions: ActionTree<UiState, RootState> = {
  changeModalVisibility({ commit }, value: boolean) {
    commit('CHANGE_MODAL_VISIBILITY', value)
  },

  changeButtonLoader({ commit }, value: boolean) {
    commit("CHANGE_BUTTON_LOADER", value)
  },

  onSubmitFromModal({ commit }, handler: () => {}) {
    commit("SET_SUBMIT_FROM_MODAL", handler)
  }
}
