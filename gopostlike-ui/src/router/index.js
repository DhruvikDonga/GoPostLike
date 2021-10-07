import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../components/Home.vue'
import Admin from '../components/Admin.vue'
import Userboard from '../components/Userboard.vue'
import Login from '../components/Login.vue'
import Register from '../components/Register.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
    meta: {
      guest: true
    }
  },
  {
    path: '/register',
    name: 'register',
    component: Register,
    meta: {
      guest: true
    }
  },
  {
    path: '/dashboard',
    name: 'userboard',
    component: Userboard,
    meta: {
      requiresAuth: true,
      role : "user"

    }
  },
  
  {
    path: '/admindasboard',
    name: 'admin',
    component: Admin,
    meta: {
      requiresAuth: true,
      role : "admin"
    }
  },
]
const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})
// Meta Handling
router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (localStorage.getItem('jwt') == null) {
      next({
        path: '/login',
        params: { nextUrl: to.fullPath }
      })
    } else {
      let user = JSON.parse(localStorage.getItem('user'))
      if (to.matched.some(record => record.meta.user)) {
        if (user.user == "admin") {
          next()
        } else {
          next({ name: 'userboard' })
        }
      } else {
        next()
      }
    }
  } else if (to.matched.some(record => record.meta.guest)) {
    if (localStorage.getItem('jwt') == null) {
      next()
    } else {
      next({ name: 'userboard' })
    }
  } else {
    next()
  }
})



export default router
