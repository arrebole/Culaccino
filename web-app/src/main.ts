import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import { Pagination,Loading,Upload } from "element-ui"

import "normalize.css"
import "./assets/iconfont/iconfont.css"

Vue.use(Loading.directive)
Vue.component(Pagination.name,Pagination)
Vue.component(Upload.name,Upload)


new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
