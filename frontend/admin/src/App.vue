<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { isAuthenticated, fetchUser, logout, user } from './stores/auth'

const router = useRouter()

onMounted(() => {
  if (isAuthenticated.value) {
    fetchUser()
  }
})

function handleLogout() {
  logout()
  router.push({ name: 'login' })
}
</script>

<template>
  <!-- Navbar for authenticated users -->
  <nav class="navbar" v-if="isAuthenticated">
    <div class="nav-inner">
      <router-link to="/" class="nav-brand">
        <span class="nav-logo">M</span>
        <span>MenuLink</span>
      </router-link>
      <div class="nav-right">
        <span class="nav-user" v-if="user">{{ user.name }}</span>
        <button class="nav-logout" @click="handleLogout">Выйти</button>
      </div>
    </div>
  </nav>

  <main :class="{ 'with-nav': isAuthenticated }">
    <router-view />
  </main>
</template>

<style scoped>
.navbar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  height: 56px;
  background: #fff;
  border-bottom: 1px solid var(--border);
  z-index: 100;
}

.nav-inner {
  max-width: 1100px;
  margin: 0 auto;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 24px;
}

.nav-brand {
  display: flex;
  align-items: center;
  gap: 10px;
  font-weight: 700;
  font-size: 17px;
}

.nav-logo {
  width: 32px;
  height: 32px;
  background: var(--primary);
  color: #fff;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 15px;
  font-weight: 700;
}

.nav-right {
  display: flex;
  align-items: center;
  gap: 16px;
}

.nav-user {
  font-size: 14px;
  color: var(--text-secondary);
}

.nav-logout {
  font-size: 14px;
  color: var(--danger);
  font-weight: 500;
}

.with-nav {
  padding-top: 56px;
}
</style>
