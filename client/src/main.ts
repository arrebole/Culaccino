import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import mavonEditor from 'mavon-editor'

import "normalize.css"
import "./assets/iconfont/iconfont.css"
import 'mavon-editor/dist/css/index.css'

Vue.use(mavonEditor)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
