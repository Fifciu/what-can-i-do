import { GetterTree, ActionTree, MutationTree } from 'vuex'
import { RootState } from '~/store'

export const state = () => ({
  currentRouteTitle: '',
})

export type BreadcrumbsState = ReturnType<typeof state>

export const mutations: MutationTree<BreadcrumbsState> = {
  CHANGE_TITLE: (state, newTitle: string) => (state.currentRouteTitle = newTitle),
}

export const actions: ActionTree<BreadcrumbsState, RootState> = {
  changeTitle({ commit }, newTitle: string) {
    commit('CHANGE_TITLE', newTitle)
  },
}
