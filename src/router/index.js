import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../components/login.vue'
import Home from '../components/home.vue'
import Index from '../components/articleIndex.vue'
import Article from '../components/article.vue'
import Write from '../components/wirte.vue'

Vue.use(VueRouter)

const routes = [
  { path: '/login', component: Login },
  { path: '/', redirect: '/login' },
  {
    path: '/home',
    component: Home,
    redirect: '/index',
    children: [{ path: '/index', component: Index },
      { path: '/article', name: 'article', component: Article },
      { path: '/article/write', component: Write }]
  }

]

const router = new VueRouter({
  routes
})

export default router
