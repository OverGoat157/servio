import { createRouter, createWebHistory } from 'vue-router'
import HomePage from './views/HomePage.vue'
import CartPage from './views/CartPage.vue'

const routes = [
  { path: '/:slug', name: 'home', component: HomePage },
  { path: '/:slug/menu', redirect: to => ({ name: 'home', params: { slug: to.params.slug } }) },
  { path: '/:slug/cart', name: 'cart', component: CartPage },
]

export default createRouter({
  history: createWebHistory(),
  routes,
})
