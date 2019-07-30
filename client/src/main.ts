import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import "normalize.css"
import "./assets/iconfont/iconfont.css"

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
