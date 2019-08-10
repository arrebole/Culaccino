import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import mavonEditor from 'mavon-editor'

import { Pagination,Loading } from "element-ui"
import 'element-ui/lib/theme-chalk/index.css';

import "normalize.css"
import "./assets/iconfont/iconfont.css"
import 'mavon-editor/dist/css/index.css'
import "github-markdown-css"

Vue.use(mavonEditor)
Vue.use(Pagination)
Vue.use(Loading.directive)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
