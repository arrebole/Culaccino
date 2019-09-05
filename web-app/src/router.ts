import Vue from 'vue'
import Router from 'vue-router'

import { notAllowedHadLogin, notAllowedNoLogin } from "./middleware/auth"

import Login from './views/account/login.vue'
import Repos from './views/domain/repository/index.vue'
import Dashboard from './views/dashboard/index.vue'
//把组件按组分块
const ManageRepos = () => import(/* webpackChunkName: "admin" */ './views/domain/manage.vue')
const NewRepo = () => import(/* webpackChunkName: "newRepo" */ './views/domain/new.vue')
const Commit = () => import(/* webpackChunkName: "update" */ './views/domain/repository/commit.vue')

Vue.use(Router)


export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: "Home",
      component: Dashboard
    },
    {
      path: '/login',
      name: 'Login',
      component: Login,
      beforeEnter: notAllowedHadLogin
    },
    {
      path: '/new',
      name: "New",
      component: NewRepo,
      beforeEnter: notAllowedNoLogin
    },
    {
      path: '/:domain',
      name: "ManageRepos",
      component: ManageRepos,
      beforeEnter: notAllowedNoLogin
    },
    {
      path: '/:domain/:repo',
      name: "Repo",
      component: Repos,
    },
    {
      path: '/:domain/:repo/commit',
      name: "Commit",
      component: Commit,
      beforeEnter: notAllowedNoLogin
    },
    {
      path: '/*',
      redirect: "/"
    }
  ]
})
