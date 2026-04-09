import { createRouter, createWebHistory } from 'vue-router'
import { isAuthenticated, user, fetchUser } from './stores/auth'

import LoginPage from './views/LoginPage.vue'
import DashboardPage from './views/DashboardPage.vue'
import RestaurantPage from './views/RestaurantPage.vue'
import MenuPage from './views/MenuPage.vue'
import OrdersPage from './views/OrdersPage.vue'
import MessengersPage from './views/MessengersPage.vue'
import CombosPage from './views/CombosPage.vue'
import AnalyticsPage from './views/AnalyticsPage.vue'
import AdminUsersPage from './views/AdminUsersPage.vue'

const routes = [
  { path: '/login', name: 'login', component: LoginPage, meta: { guest: true } },
  { path: '/', name: 'dashboard', component: DashboardPage, meta: { auth: true } },
  { path: '/restaurant/:id', name: 'restaurant', component: RestaurantPage, meta: { auth: true } },
  { path: '/restaurant/:id/menu', name: 'menu', component: MenuPage, meta: { auth: true } },
  { path: '/restaurant/:id/orders', name: 'orders', component: OrdersPage, meta: { auth: true } },
  { path: '/restaurant/:id/combos', name: 'combos', component: CombosPage, meta: { auth: true } },
  { path: '/restaurant/:id/messengers', name: 'messengers', component: MessengersPage, meta: { auth: true } },
  { path: '/restaurant/:id/analytics', name: 'analytics', component: AnalyticsPage, meta: { auth: true } },
  // Admin
  { path: '/admin/users', name: 'admin-users', component: AdminUsersPage, meta: { auth: true, admin: true } },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach(async (to) => {
  if (to.meta.auth && !isAuthenticated.value) {
    return { name: 'login' }
  }
  if (to.meta.guest && isAuthenticated.value) {
    return { name: 'dashboard' }
  }
  // Load user if authenticated but user not loaded
  if (to.meta.auth && isAuthenticated.value && !user.value) {
    await fetchUser()
  }
  // Admin check
  if (to.meta.admin && user.value?.role !== 'admin') {
    return { name: 'dashboard' }
  }
})

export default router
