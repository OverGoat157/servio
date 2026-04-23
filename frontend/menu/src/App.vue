<script setup>
import { watch } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { loadRestaurant, restaurant } from './stores/restaurant'
import { setLocale } from './i18n'

const route = useRoute()
const { locale } = useI18n()

watch(
  () => route.params.slug,
  (slug) => {
    if (slug && !restaurant.id) {
      loadRestaurant(slug)
    }
  },
  { immediate: true }
)

function switchLocale(lang) {
  setLocale(lang)
}
</script>

<template>
  <div class="lang-switch" :title="$t('common.language')">
    <button :class="{ active: locale === 'ru' }" @click="switchLocale('ru')">RU</button>
    <button :class="{ active: locale === 'en' }" @click="switchLocale('en')">EN</button>
  </div>
  <router-view />
</template>

<style scoped>
.lang-switch {
  position: fixed;
  top: 12px;
  right: 12px;
  z-index: 50;
  display: inline-flex;
  background: rgba(255, 255, 255, 0.92);
  backdrop-filter: blur(8px);
  -webkit-backdrop-filter: blur(8px);
  border: 1px solid var(--border);
  border-radius: 100px;
  padding: 3px;
  box-shadow: 0 2px 6px rgba(0,0,0,0.05);
}

.lang-switch button {
  padding: 4px 10px;
  font-size: 11px;
  font-weight: 700;
  border-radius: 100px;
  color: var(--text-muted, #9ca3af);
  letter-spacing: 0.3px;
}

.lang-switch button.active {
  background: var(--primary);
  color: var(--primary-foreground, #fff);
}
</style>
