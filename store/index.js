/**
 * Created by joylee on 2018/6/15.
 */
import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    userName: '未知'
  },
  mutations: {
    updateUserName (state, payload) {
      state.userName = payload.userName
    }
  },
  modules: {
    vux: {
      state: {
        demoScrollTop: 0,
        isLoading: false
      },
      mutations: {
        updateDemoPosition (state, payload) {
          state.demoScrollTop = payload.top
        },
        updateLoadingStatus (state, payload) {
          state.isLoading = payload.isLoading
        }
      },
      actions: {
        updateDemoPosition ({ commit }, top) {
          commit({ type: 'updateDemoPosition', top: top })
        }
      }
    },
    demo: {
      namespaced: true,
      state: {
        name: 'xiao'
      },
      mutations: {
        updateName (state, payload) {
          state.name = payload.name
        }
      },
      actions: {
        updateName ({ commit }, name) {
          commit('updateName', { name: name })
        }
      }
    }
  }
})
