import Vue from 'vue'
import Router from 'vue-router'

import Admin from './views/Admin.vue'
import Create from './views/Create.vue'
import Update from './views/Update.vue'
import Home from './views/Home.vue'
import Login from './views/Login.vue'
import Detail from './views/Detail.vue'
import middware from "./middleware/isAdmin"
Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: "Home",
      component: Home
    },
    {
      path: '/article/:id',
      name: "Detail",
      component: Detail,
    },
    {
      path: '/login',
      name: 'Login',
      component: Login,
      beforeEnter: middware.hadPower
    },
    {
      path: '/admin',
      name: "Admin",
      component: Admin,
      beforeEnter: middware.noPower
    },
    {
      path: '/admin/create',
      name: "Create",
      component: Create,
      beforeEnter: middware.noPower
    },

    {
      path: '/admin/update/:articleID',
      name: "Update",
      component: Update,
      beforeEnter: middware.noPower
    },
    {
      path: '/*',
      redirect: "/"
    }
  ]
})
