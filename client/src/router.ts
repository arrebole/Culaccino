import Vue from 'vue'
import Router from 'vue-router'

import Admin from './views/Admin.vue'
import Create from './views/Create.vue'
import Update from './views/Update.vue'
import Home from './views/Home.vue'
import Login from './views/Login.vue'

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
      path:'/login',
      name:'Login',
      component: Login
    },
    {
      path: '/admin',
      name: "Admin",
      component: Admin
    },
    {
      path: '/admin/create',
      name: "Create",
      component: Create
    },
    {
      path: '/admin/update/:id',
      name: "Update",
      component: Update
    },
    {
      path: '/*',
      redirect: "/"
    }
  ]
})
