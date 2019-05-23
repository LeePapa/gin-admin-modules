import Vue from 'vue'
import {getMyMenuList, getMenuTree} from '@/api/system'

const menu = {
  state: {
    menus: [],
    treeMenus: []
  },

  mutations: {
    SET_MENUS: (state, menus) => {
      state.menus = menus
    },
    SET_TREE_MENUS: (state, treeMenus) => {
      state.treeMenus = treeMenus
    }
  },

  actions: {
    ChangeMyMenu({commit, state}) {
      return new Promise((resolve) => {
        getMyMenuList().then(response => {
          commit('SET_MENUS', response.data)
          getMenuTree().then(res => {
            commit('SET_TREE_MENUS', res.data)
          })
        })
      })
    }
  }
}

export default menu
