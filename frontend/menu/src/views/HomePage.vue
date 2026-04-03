<script setup>
import { useRoute, useRouter } from 'vue-router'
import { restaurant, loading, error } from '../stores/restaurant'

const route = useRoute()
const router = useRouter()
const slug = route.params.slug

function goToMenu() {
  router.push({ name: 'menu', params: { slug } })
}
</script>

<template>
  <div class="home">
    <!-- Loading -->
    <div v-if="loading" class="loading-screen">
      <div class="spinner"></div>
    </div>

    <!-- Error -->
    <div v-else-if="error" class="error-screen">
      <div class="error-icon">!</div>
      <p>Ресторан не найден</p>
    </div>

    <!-- Content -->
    <template v-else>
      <!-- Header -->
      <div class="hero">
        <div class="hero-bg"></div>
        <div class="hero-content">
          <div class="logo-circle" v-if="restaurant.logo">
            <img :src="restaurant.logo" :alt="restaurant.name" />
          </div>
          <div class="logo-circle logo-placeholder" v-else>
            {{ restaurant.name.charAt(0) }}
          </div>
          <h1>{{ restaurant.name }}</h1>
          <p class="subtitle" v-if="restaurant.description">{{ restaurant.description }}</p>
        </div>
      </div>

      <!-- Working hours -->
      <div class="card" v-if="restaurant.working_hours">
        <div class="card-label">Режим работы</div>
        <div
          class="hours-row"
          v-for="(time, day) in restaurant.working_hours"
          :key="day"
        >
          <span class="day">{{ day }}</span>
          <span class="time">{{ time }}</span>
        </div>
      </div>

      <!-- Address -->
      <div class="card" v-if="restaurant.address">
        <div class="card-label">Адрес</div>
        <p class="card-text">{{ restaurant.address }}</p>
      </div>

      <!-- Actions -->
      <div class="actions">
        <button class="btn-primary" @click="goToMenu">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M3 6h18M3 12h18M3 18h18"/>
          </svg>
          Открыть меню
        </button>

        <a
          v-if="restaurant.phone"
          :href="'tel:' + restaurant.phone"
          class="btn-outline"
        >
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M22 16.92v3a2 2 0 01-2.18 2 19.79 19.79 0 01-8.63-3.07 19.5 19.5 0 01-6-6 19.79 19.79 0 01-3.07-8.67A2 2 0 014.11 2h3a2 2 0 012 1.72c.127.96.361 1.903.7 2.81a2 2 0 01-.45 2.11L8.09 9.91a16 16 0 006 6l1.27-1.27a2 2 0 012.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0122 16.92z"/>
          </svg>
          Связаться с нами
        </a>
      </div>

      <!-- Messengers -->
      <div class="messengers" v-if="restaurant.messengers?.length">
        <div class="card-label">Мы в мессенджерах</div>
        <div class="msg-list">
          <div class="msg-chip" v-for="m in restaurant.messengers" :key="m">
            <svg v-if="m === 'whatsapp'" width="18" height="18" viewBox="0 0 24 24" fill="#25D366"><path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347z"/><path d="M12 0C5.373 0 0 5.373 0 12c0 2.127.556 4.12 1.525 5.855L0 24l6.335-1.652A11.94 11.94 0 0012 24c6.627 0 12-5.373 12-12S18.627 0 12 0z"/></svg>
            <svg v-else-if="m === 'telegram'" width="18" height="18" viewBox="0 0 24 24" fill="#229ED9"><path d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12 12-5.373 12-12S18.627 0 12 0zm5.562 8.161c-.18 1.897-.962 6.502-1.359 8.627-.168.9-.5 1.201-.82 1.23-.697.064-1.226-.46-1.9-.903-1.056-.692-1.653-1.123-2.678-1.799-1.185-.781-.417-1.21.258-1.911.177-.184 3.247-2.977 3.307-3.23.007-.032.014-.15-.056-.212s-.174-.041-.249-.024c-.106.024-1.793 1.14-5.061 3.345-.479.33-.913.492-1.302.484-.428-.008-1.252-.241-1.865-.44-.752-.245-1.349-.374-1.297-.789.027-.216.325-.437.893-.663 3.498-1.524 5.831-2.529 6.998-3.014 3.332-1.386 4.025-1.627 4.476-1.635.099-.002.321.023.465.141a.506.506 0 01.171.325c.016.093.036.306.02.472z"/></svg>
            <svg v-else width="18" height="18" viewBox="0 0 24 24" fill="#6B7280"><circle cx="12" cy="12" r="10"/><path d="M8 14s1.5 2 4 2 4-2 4-2" stroke="#fff" stroke-width="1.5" fill="none"/><circle cx="9" cy="10" r="1" fill="#fff"/><circle cx="15" cy="10" r="1" fill="#fff"/></svg>
            {{ m === 'whatsapp' ? 'WhatsApp' : m === 'telegram' ? 'Telegram' : m === 'viber' ? 'Viber' : m }}
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<style scoped>
.home {
  padding-bottom: 32px;
}

.loading-screen {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
}

.spinner {
  width: 40px;
  height: 40px;
  border: 3px solid var(--border);
  border-top-color: var(--primary);
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.error-screen {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  gap: 16px;
  color: var(--text-secondary);
}

.error-icon {
  width: 56px;
  height: 56px;
  border-radius: 50%;
  background: #FEE2E2;
  color: #DC2626;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  font-weight: 700;
}

/* Hero */
.hero {
  position: relative;
  padding: 60px 24px 32px;
  text-align: center;
  overflow: hidden;
}

.hero-bg {
  position: absolute;
  inset: 0;
  background: linear-gradient(180deg, #1F2937 0%, #374151 100%);
  z-index: 0;
}

.hero-content {
  position: relative;
  z-index: 1;
}

.logo-circle {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  margin: 0 auto 16px;
  overflow: hidden;
  border: 3px solid rgba(255,255,255,0.2);
}

.logo-circle img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.logo-placeholder {
  background: var(--primary);
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32px;
  font-weight: 700;
}

.hero h1 {
  font-size: 28px;
  font-weight: 700;
  color: #fff;
  margin-bottom: 8px;
  letter-spacing: 1px;
}

.subtitle {
  color: rgba(255,255,255,0.7);
  font-size: 14px;
  line-height: 1.5;
}

/* Cards */
.card {
  margin: 16px 16px 0;
  background: var(--bg);
  border: 1px solid var(--border);
  border-radius: var(--radius);
  padding: 16px;
}

.card-label {
  font-size: 11px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 1px;
  color: var(--text-secondary);
  margin-bottom: 12px;
}

.card-text {
  font-size: 14px;
  color: var(--text);
  line-height: 1.5;
}

.hours-row {
  display: flex;
  justify-content: space-between;
  padding: 8px 0;
  border-bottom: 1px solid var(--border);
  font-size: 14px;
}

.hours-row:last-child {
  border-bottom: none;
  padding-bottom: 0;
}

.day {
  color: var(--text-secondary);
}

.time {
  font-weight: 600;
}

/* Actions */
.actions {
  padding: 24px 16px 0;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.btn-primary {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  width: 100%;
  padding: 16px;
  background: var(--primary);
  color: #fff;
  border-radius: var(--radius);
  font-size: 16px;
  font-weight: 600;
  transition: opacity var(--ease);
}

.btn-primary:active {
  opacity: 0.85;
}

.btn-outline {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  width: 100%;
  padding: 16px;
  background: var(--bg);
  color: var(--text);
  border: 2px solid var(--border);
  border-radius: var(--radius);
  font-size: 16px;
  font-weight: 600;
  transition: border-color var(--ease);
}

.btn-outline:active {
  border-color: var(--primary);
}

/* Messengers */
.messengers {
  padding: 24px 16px 0;
}

.msg-list {
  display: flex;
  gap: 10px;
}

.msg-chip {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 16px;
  background: var(--bg-secondary);
  border-radius: 100px;
  font-size: 13px;
  font-weight: 500;
}
</style>
