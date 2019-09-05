import Vue from 'vue'
import Vuex from 'vuex'
import api from './api/index'
import { getCookie, setCookie } from './util'

Vue.use(Vuex)

export const types = {
  SET_USER_INFO: "SET_USER_INFO"
}

export default new Vuex.Store({
  state: {
    user: {
      uid: -1,
      domain: "",
      token: "",
      permission: 0
    }
  },
  getters: {
    // 判断登录状
    getUserStatus(state) {
      return state.user
    },
    islogin(state) {
      return (state.user.uid > 0)
    }
  },
  mutations: {
    [types.SET_USER_INFO](state, payload) {
      state.user.domain = payload.domain
      state.user.uid = payload.uid
      state.user.token = payload.token
      state.user.permission = payload.permission
    }
  },
  actions: {
    // 通过cookie获取用户信息
    async fetchUserInfoBycookie(context) {
      if (getCookie() == "" || context.getters.islogin) return
      const res = await api.sessionExist(getCookie())
      context.commit(types.SET_USER_INFO, res.data.user)
    },
    async fetchLogin(context, payload) {
      const res = await api.sessionLogin(payload)
      if (res.code < 0) return false
      setCookie(res.data.user.token);
      context.commit(types.SET_USER_INFO, res.data.user)
      return true
    }
  }
})
