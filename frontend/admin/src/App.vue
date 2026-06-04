<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { isAuthenticated, fetchUser, logout, user } from './stores/auth'
import { setLocale } from './i18n'

const router = useRouter()
const { locale, t } = useI18n()

onMounted(() => {
  if (isAuthenticated.value) {
    fetchUser()
  }
})

function handleLogout() {
  logout()
  router.push({ name: 'login' })
}

function switchLocale(lang) {
  setLocale(lang)
}

function tierLabel(tier) {
  const key = tier || 'basic'
  return t(`tier.${key}`)
}
</script>

<template>
  <!-- Navbar for authenticated users -->
  <nav class="navbar" v-if="isAuthenticated">
    <div class="nav-inner">
      <router-link to="/" class="nav-brand">
        <svg class="nav-logo" viewBox="0 0 32 32" fill="none" stroke="currentColor" stroke-width="3.5" stroke-linecap="round" stroke-linejoin="round">
          <path d="M7 10l5 6-5 6"/>
          <path d="M13 10l5 6-5 6"/>
          <path d="M19 10l5 6-5 6"/>
        </svg>
        <span>{{ $t('nav.brand') }}</span>
      </router-link>
      <div class="nav-right">
        <router-link v-if="user?.role === 'admin'" to="/admin/analytics" class="nav-link" :title="$t('nav.analytics')">
          <svg class="nav-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 3v18h18"/><path d="M7 14l4-4 4 4 5-5"/></svg>
          <span class="nav-link-text">{{ $t('nav.analytics') }}</span>
        </router-link>
        <router-link v-if="user?.role === 'admin'" to="/admin/users" class="nav-link" :title="$t('nav.users')">
          <svg class="nav-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M23 21v-2a4 4 0 00-3-3.87"/><path d="M16 3.13a4 4 0 010 7.75"/></svg>
          <span class="nav-link-text">{{ $t('nav.users') }}</span>
        </router-link>
        <div class="lang-switch" :title="$t('common.language')">
          <button :class="{ active: locale === 'ru' }" @click="switchLocale('ru')">RU</button>
          <button :class="{ active: locale === 'en' }" @click="switchLocale('en')">EN</button>
        </div>
        <span class="nav-user" v-if="user">
          {{ user.name }}
          <span v-if="user.role !== 'admin'" class="tier-chip" :class="user.tier || 'basic'">{{ tierLabel(user.tier) }}</span>
        </span>
        <button class="nav-logout" @click="handleLogout" :title="$t('nav.logout')">
          <svg class="nav-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M9 21H5a2 2 0 01-2-2V5a2 2 0 012-2h4"/><path d="M16 17l5-5-5-5"/><path d="M21 12H9"/></svg>
          <span class="nav-link-text">{{ $t('nav.logout') }}</span>
        </button>
      </div>
    </div>
  </nav>

  <main :class="{ 'with-nav': isAuthenticated }">
    <!-- Floating language switch for guest pages (login) -->
    <div class="lang-switch floating" v-if="!isAuthenticated" :title="$t('common.language')">
      <button :class="{ active: locale === 'ru' }" @click="switchLocale('ru')">RU</button>
      <button :class="{ active: locale === 'en' }" @click="switchLocale('en')">EN</button>
    </div>
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
  color: var(--primary);
  display: block;
}

.nav-right {
  display: flex;
  align-items: center;
  gap: 16px;
}

.nav-link {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 14px;
  font-weight: 500;
  color: var(--primary);
  padding: 6px 12px;
  border-radius: var(--radius);
  transition: background 0.2s;
}

.nav-link:hover {
  background: var(--primary-light);
}

.nav-icon {
  flex-shrink: 0;
}

.lang-switch {
  display: inline-flex;
  background: var(--bg-secondary, #f3f4f6);
  border-radius: 8px;
  padding: 2px;
}

.lang-switch button {
  padding: 4px 10px;
  font-size: 12px;
  font-weight: 600;
  border-radius: 6px;
  color: var(--text-secondary);
  letter-spacing: 0.3px;
}

.lang-switch button.active {
  background: #fff;
  color: var(--text);
  box-shadow: 0 1px 2px rgba(0,0,0,0.08);
}

.lang-switch.floating {
  position: fixed;
  top: 16px;
  right: 16px;
  z-index: 200;
  background: #fff;
  border: 1px solid var(--border);
}

.nav-user {
  font-size: 14px;
  color: var(--text-secondary);
  display: inline-flex;
  align-items: center;
  gap: 8px;
}

.tier-chip {
  padding: 2px 8px;
  border-radius: 100px;
  font-size: 11px;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.3px;
}

.tier-chip.basic {
  background: #F1F5F9;
  color: #475569;
}

.tier-chip.business {
  background: #DBEAFE;
  color: #1D4ED8;
}

.tier-chip.business_max {
  background: #FEF3C7;
  color: #B45309;
}

.nav-logout {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 14px;
  color: var(--danger);
  font-weight: 500;
  padding: 6px 12px;
  border-radius: var(--radius);
  transition: background 0.2s;
}

.nav-logout:hover {
  background: #FEE2E2;
}

.with-nav {
  padding-top: 56px;
}

@media (max-width: 700px) {
  .nav-inner {
    padding: 0 12px;
    gap: 6px;
  }
  .nav-brand span {
    display: none;
  }
  .nav-right {
    gap: 4px;
  }
  .nav-link,
  .nav-logout {
    padding: 6px 8px;
  }
  .nav-link-text {
    display: none;
  }
  .nav-user {
    display: none;
  }
  .lang-switch button {
    padding: 3px 7px;
  }
}
</style>
