import { createRouter, createWebHistory } from 'vue-router'
import HomePage from './views/HomePage.vue'
import MenuPage from './views/MenuPage.vue'
import CartPage from './views/CartPage.vue'

const routes = [
  { path: '/:slug', name: 'home', component: HomePage },
  { path: '/:slug/menu', name: 'menu', component: MenuPage },
  { path: '/:slug/cart', name: 'cart', component: CartPage },
]

export default createRouter({
  history: createWebHistory(),
  routes,
})
