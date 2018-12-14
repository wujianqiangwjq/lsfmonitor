import Vue from 'vue'
import Router from 'vue-router'
import Login  from '../components/login'
import Main  from '../components/HelloWorld'
import Test  from '../components/Testvue'
Vue.use(Router)

export default new Router({
  routes: [
    {
        path: '/',
        component: Login
    },
    {
      path: '/main',
      component: Main
    }
    // {
    //   path: '/',
    //   component: Test
    // }
  ]
})
