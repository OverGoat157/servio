<script setup>
import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { setLocale } from '../i18n'

const { t, locale } = useI18n()

const menuOpen = ref(false)

const links = computed(() => [
  { label: t('nav.features'), href: '#features' },
  { label: t('nav.how'), href: '#how' },
  { label: t('nav.demo'), href: '#demo' },
  { label: t('nav.pricing'), href: '#pricing' },
])

function close() {
  menuOpen.value = false
}

function switchLocale(lang) {
  setLocale(lang)
}
</script>

<template>
  <header class="nav">
    <div class="container nav-inner">
      <a href="#" class="logo">
        <svg class="logo-icon" viewBox="0 0 32 32" fill="none" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round">
          <path d="M5 8l6 8-6 8"/>
          <path d="M13 8l6 8-6 8"/>
          <path d="M21 8l6 8-6 8"/>
        </svg>
        <span>AB Team</span>
      </a>

      <nav class="nav-links">
        <a v-for="l in links" :key="l.href" :href="l.href">{{ l.label }}</a>
      </nav>

      <div class="nav-right">
        <div class="lang-switch" :title="$t('common.language')">
          <button :class="{ active: locale === 'ru' }" @click="switchLocale('ru')">RU</button>
          <button :class="{ active: locale === 'en' }" @click="switchLocale('en')">EN</button>
        </div>
        <a href="#contact" class="btn btn-primary nav-cta">{{ $t('nav.cta') }}</a>
      </div>

      <button class="burger" :class="{ active: menuOpen }" @click="menuOpen = !menuOpen" :aria-label="$t('nav.menuAria')">
        <span /><span /><span />
      </button>
    </div>

    <!-- mobile -->
    <div class="overlay" :class="{ open: menuOpen }" @click="close" />
    <nav class="mobile-menu" :class="{ open: menuOpen }">
      <a v-for="l in links" :key="l.href" :href="l.href" @click="close">{{ l.label }}</a>
      <div class="lang-switch mobile-lang" :title="$t('common.language')">
        <button :class="{ active: locale === 'ru' }" @click="switchLocale('ru')">RU</button>
        <button :class="{ active: locale === 'en' }" @click="switchLocale('en')">EN</button>
      </div>
      <a href="#contact" class="btn btn-primary" @click="close">{{ $t('nav.cta') }}</a>
    </nav>
  </header>
</template>

<style scoped>
.nav {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 100;
  background: rgba(255, 255, 255, 0.92);
  backdrop-filter: blur(12px);
  border-bottom: 1px solid var(--light-gray);
}

.nav-inner {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 72px;
  gap: 32px;
}

.logo {
  display: flex;
  align-items: center;
  gap: 10px;
  font-weight: 700;
  font-size: 20px;
  flex-shrink: 0;
}

.logo-icon {
  width: 36px;
  height: 36px;
  color: var(--primary);
  flex-shrink: 0;
}

.nav-links {
  display: flex;
  gap: 32px;
}

.nav-links a {
  font-size: 15px;
  font-weight: 500;
  color: var(--gray);
  transition: color var(--ease);
}

.nav-links a:hover {
  color: var(--primary);
}

.nav-right {
  display: flex;
  align-items: center;
  gap: 16px;
}

.lang-switch {
  display: inline-flex;
  background: var(--off-white, #f3f4f6);
  border-radius: 100px;
  padding: 3px;
}

.lang-switch button {
  padding: 4px 10px;
  font-size: 12px;
  font-weight: 700;
  border-radius: 100px;
  color: var(--gray);
  border: none;
  background: transparent;
  cursor: pointer;
  letter-spacing: 0.3px;
}

.lang-switch button.active {
  background: #fff;
  color: var(--black);
  box-shadow: 0 1px 2px rgba(0,0,0,0.08);
}

.nav-cta {
  padding: 10px 24px;
  font-size: 14px;
}

/* burger */
.burger {
  display: none;
  flex-direction: column;
  gap: 5px;
  background: none;
  border: none;
  cursor: pointer;
  padding: 4px;
  z-index: 110;
}

.burger span {
  width: 24px;
  height: 2px;
  background: var(--black);
  border-radius: 2px;
  transition: all var(--ease);
}

.burger.active span:nth-child(1) {
  transform: rotate(45deg) translate(5px, 5px);
}

.burger.active span:nth-child(2) {
  opacity: 0;
}

.burger.active span:nth-child(3) {
  transform: rotate(-45deg) translate(5px, -5px);
}

.overlay {
  display: none;
}

.mobile-menu {
  display: none;
}

@media (max-width: 768px) {
  .nav-links,
  .nav-right {
    display: none;
  }

  .burger {
    display: flex;
  }

  .overlay {
    display: block;
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.4);
    opacity: 0;
    pointer-events: none;
    transition: opacity var(--ease);
    z-index: 90;
  }

  .overlay.open {
    opacity: 1;
    pointer-events: auto;
  }

  .mobile-menu {
    display: flex;
    flex-direction: column;
    position: fixed;
    top: 0;
    right: 0;
    width: 280px;
    height: 100dvh;
    background: #fff;
    padding: 88px 24px 32px;
    gap: 8px;
    transform: translateX(100%);
    transition: transform var(--ease);
    z-index: 95;
    box-shadow: var(--shadow-lg);
  }

  .mobile-menu.open {
    transform: translateX(0);
  }

  .mobile-menu a {
    padding: 12px 16px;
    font-size: 16px;
    font-weight: 500;
    border-radius: 8px;
    transition: background var(--ease);
  }

  .mobile-menu a:hover {
    background: var(--off-white);
  }

  .mobile-menu .btn {
    margin-top: 16px;
  }

  .mobile-lang {
    align-self: flex-start;
    margin: 8px 16px;
  }
}
</style>
