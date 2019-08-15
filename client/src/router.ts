import Vue from 'vue'
import Router from 'vue-router'

import Home from './views/Home.vue'
import Login from './views/Login.vue'
import Archives from './views/Archive.vue'
import auth from "./middleware/auth"
//把组件按组分块
const Admin = () => import(/* webpackChunkName: "admin" */ './views/Admin.vue')
const Create = () => import(/* webpackChunkName: "create" */ './views/Create.vue')
const Update = () => import(/* webpackChunkName: "update" */ './views/Update.vue')

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
      path: '/archives/:id',
      name: "Archives",
      component: Archives,
    },
    {
      path: '/login',
      name: 'Login',
      component: Login,
      beforeEnter: auth.hadPower
    },
    {
      path: '/admin',
      name: "Admin",
      component: Admin,
      beforeEnter: auth.noPower
    },
    {
      path: '/admin/create',
      name: "Create",
      component: Create,
      beforeEnter: auth.noPower
    },

    {
      path: '/admin/update/:id',
      name: "Update",
      component: Update,
      beforeEnter: auth.noPower
    },
    {
      path: '/*',
      redirect: "/"
    }
  ]
})
