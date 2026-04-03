import { createRouter, createWebHistory } from 'vue-router'
import { isAuthenticated } from './stores/auth'

import LoginPage from './views/LoginPage.vue'
import RegisterPage from './views/RegisterPage.vue'
import DashboardPage from './views/DashboardPage.vue'
import RestaurantPage from './views/RestaurantPage.vue'
import MenuPage from './views/MenuPage.vue'
import OrdersPage from './views/OrdersPage.vue'
import MessengersPage from './views/MessengersPage.vue'

const routes = [
  { path: '/login', name: 'login', component: LoginPage, meta: { guest: true } },
  { path: '/register', name: 'register', component: RegisterPage, meta: { guest: true } },
  { path: '/', name: 'dashboard', component: DashboardPage, meta: { auth: true } },
  { path: '/restaurant/:id', name: 'restaurant', component: RestaurantPage, meta: { auth: true } },
  { path: '/restaurant/:id/menu', name: 'menu', component: MenuPage, meta: { auth: true } },
  { path: '/restaurant/:id/orders', name: 'orders', component: OrdersPage, meta: { auth: true } },
  { path: '/restaurant/:id/messengers', name: 'messengers', component: MessengersPage, meta: { auth: true } },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to) => {
  if (to.meta.auth && !isAuthenticated.value) {
    return { name: 'login' }
  }
  if (to.meta.guest && isAuthenticated.value) {
    return { name: 'dashboard' }
  }
})

export default router
