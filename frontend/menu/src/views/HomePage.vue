<script setup>
import { ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { restaurant, loading, error } from '../stores/restaurant'
import { cart, addToCart, cartCount, updateQuantity } from '../stores/cart'
import { imageUrl } from '../api/client'

const route = useRoute()
const router = useRouter()
const slug = route.params.slug

const popularItems = computed(() => {
  const cats = restaurant.categories || []
  const all = cats.flatMap(c => c.items || [])
  return all.slice(0, 4)
})

function goToMenu() {
  router.push({ name: 'menu', params: { slug } })
}

function goToCart() {
  router.push({ name: 'cart', params: { slug } })
}

function formatPrice(kopecks) {
  return Math.floor(kopecks / 100).toLocaleString('ru-RU') + ' \u20BD'
}

const selectedItem = ref(null)

function handleAdd(item) {
  addToCart(item)
}

function getQty(itemId) {
  const ci = cart.find(i => i.id === itemId)
  return ci ? ci.quantity : 0
}

function increment(item) {
  addToCart(item)
}

function decrement(itemId) {
  const ci = cart.find(i => i.id === itemId)
  if (ci) updateQuantity(itemId, ci.quantity - 1)
}

function openDetail(item) {
  selectedItem.value = item
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
      <!-- Hero с фото фоном -->
      <div class="hero" :class="{ 'hero-with-image': restaurant.cover_image }">
        <div class="hero-bg">
          <img v-if="restaurant.cover_image" :src="imageUrl(restaurant.cover_image)" alt="" class="hero-img" />
        </div>
        <div class="hero-overlay"></div>
        <div class="hero-content">
          <div class="logo-circle" v-if="restaurant.logo">
            <img :src="imageUrl(restaurant.logo)" :alt="restaurant.name" />
          </div>
          <div class="logo-circle logo-placeholder" v-else>
            {{ restaurant.name.charAt(0) }}
          </div>
          <h1>{{ restaurant.name }}</h1>
          <p class="subtitle" v-if="restaurant.description">{{ restaurant.description }}</p>
        </div>
      </div>

      <!-- Акция -->
      <div class="promo-card" v-if="restaurant.promo_title">
        <div class="promo-badge">Акция</div>
        <h3 class="promo-title">{{ restaurant.promo_title }}</h3>
        <p class="promo-desc" v-if="restaurant.promo_description">{{ restaurant.promo_description }}</p>
      </div>

      <!-- Популярные блюда -->
      <div class="section" v-if="popularItems.length">
        <div class="section-header">
          <h2>Популярные блюда</h2>
        </div>
        <div class="popular-grid">
          <div class="popular-card" :class="{ 'pop-stopped': !item.available }" v-for="item in popularItems" :key="item.id" @click="item.available && openDetail(item)">
            <div class="pop-img" v-if="item.image">
              <img :src="imageUrl(item.image)" :alt="item.name" loading="lazy" />
            </div>
            <div class="pop-img pop-img-empty" v-else>
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="var(--text-muted)" stroke-width="1.5" stroke-linecap="round">
                <path d="M18 8h1a4 4 0 010 8h-1M2 8h16v9a4 4 0 01-4 4H6a4 4 0 01-4-4V8zM6 1v3M10 1v3M14 1v3"/>
              </svg>
            </div>
            <div class="pop-stopped-badge" v-if="!item.available">Нет в наличии</div>
            <div class="pop-info">
              <div class="pop-name">{{ item.name }}</div>
              <div class="pop-price">{{ formatPrice(item.price) }}</div>
            </div>
            <template v-if="item.available">
              <div class="pop-qty" v-if="getQty(item.id)" @click.stop>
                <button class="pop-qty-btn" @click="decrement(item.id)">-</button>
                <span>{{ getQty(item.id) }}</span>
                <button class="pop-qty-btn" @click="increment(item)">+</button>
              </div>
              <button v-else class="pop-add" @click.stop="increment(item)">+</button>
            </template>
          </div>
        </div>
      </div>

      <!-- Кнопки -->
      <div class="actions">
        <button class="btn-primary" @click="goToMenu">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
            <path d="M4 6h16M4 12h16M4 18h16"/>
          </svg>
          Полное меню
        </button>

        <a v-if="restaurant.phone" :href="'tel:' + restaurant.phone" class="btn-outline">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
            <path d="M22 16.92v3a2 2 0 01-2.18 2 19.79 19.79 0 01-8.63-3.07 19.5 19.5 0 01-6-6 19.79 19.79 0 01-3.07-8.67A2 2 0 014.11 2h3a2 2 0 012 1.72c.127.96.361 1.903.7 2.81a2 2 0 01-.45 2.11L8.09 9.91a16 16 0 006 6l1.27-1.27a2 2 0 012.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0122 16.92z"/>
          </svg>
          Связаться с нами
        </a>
      </div>

      <!-- Часы работы -->
      <div class="section" v-if="restaurant.working_hours">
        <div class="info-card">
          <div class="info-icon">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="var(--text-secondary)" stroke-width="2" stroke-linecap="round">
              <circle cx="12" cy="12" r="10"/><path d="M12 6v6l4 2"/>
            </svg>
          </div>
          <div class="info-body">
            <div class="info-label">Режим работы</div>
            <div class="hours-grid">
              <div class="hours-row" v-for="(time, day) in restaurant.working_hours" :key="day">
                <span class="day">{{ day }}</span>
                <span class="time">{{ time }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Адрес -->
      <div class="section" v-if="restaurant.address">
        <div class="info-card">
          <div class="info-icon">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="var(--text-secondary)" stroke-width="2" stroke-linecap="round">
              <path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0118 0z"/><circle cx="12" cy="10" r="3"/>
            </svg>
          </div>
          <div class="info-body">
            <div class="info-label">Адрес</div>
            <div class="info-value">{{ restaurant.address }}</div>
          </div>
        </div>
      </div>

      <!-- Соцсети -->
      <div class="section" v-if="restaurant.social_links?.length">
        <div class="section-header">
          <h2>Мы в соцсетях</h2>
        </div>
        <div class="social-list">
          <a v-for="link in restaurant.social_links" :key="link.type + link.url" :href="link.url" target="_blank" rel="noopener" class="social-chip">
            <svg v-if="link.type === 'instagram'" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><rect x="2" y="2" width="20" height="20" rx="5"/><path d="M16 11.37A4 4 0 1112.63 8 4 4 0 0116 11.37z"/><path d="M17.5 6.5h.01"/></svg>
            <svg v-else-if="link.type === 'telegram'" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M22 2L11 13M22 2l-7 20-4-9-9-4z"/></svg>
            <svg v-else-if="link.type === 'vk'" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M4 4h16v16H4z"/><path d="M7 12c1 4 3 5 5 5s3-1 3-3c0-1-1-2-2-2"/></svg>
            <svg v-else-if="link.type === 'youtube'" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><rect x="2" y="4" width="20" height="16" rx="4"/><path d="M10 9l5 3-5 3V9z"/></svg>
            <svg v-else width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M10 13a5 5 0 007.54.54l3-3a5 5 0 00-7.07-7.07l-1.72 1.71"/><path d="M14 11a5 5 0 00-7.54-.54l-3 3a5 5 0 007.07 7.07l1.71-1.71"/></svg>
            {{ link.type === 'instagram' ? 'Instagram' : link.type === 'vk' ? 'VK' : link.type === 'telegram' ? 'Telegram' : link.type === 'youtube' ? 'YouTube' : link.type === 'tiktok' ? 'TikTok' : link.type === 'facebook' ? 'Facebook' : 'Сайт' }}
          </a>
        </div>
      </div>

      <!-- Мессенджеры -->
      <div class="section" v-if="restaurant.messengers?.length">
        <div class="section-header">
          <h2>Мы в мессенджерах</h2>
        </div>
        <div class="msg-list">
          <div class="msg-chip" v-for="m in restaurant.messengers" :key="m">
            <svg v-if="m === 'whatsapp'" width="20" height="20" viewBox="0 0 24 24" fill="#25D366"><path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347z"/><path d="M12 0C5.373 0 0 5.373 0 12c0 2.127.556 4.12 1.525 5.855L0 24l6.335-1.652A11.94 11.94 0 0012 24c6.627 0 12-5.373 12-12S18.627 0 12 0z"/></svg>
            <svg v-else-if="m === 'telegram'" width="20" height="20" viewBox="0 0 24 24" fill="#229ED9"><path d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12 12-5.373 12-12S18.627 0 12 0zm5.562 8.161c-.18 1.897-.962 6.502-1.359 8.627-.168.9-.5 1.201-.82 1.23-.697.064-1.226-.46-1.9-.903-1.056-.692-1.653-1.123-2.678-1.799-1.185-.781-.417-1.21.258-1.911.177-.184 3.247-2.977 3.307-3.23.007-.032.014-.15-.056-.212s-.174-.041-.249-.024c-.106.024-1.793 1.14-5.061 3.345-.479.33-.913.492-1.302.484-.428-.008-1.252-.241-1.865-.44-.752-.245-1.349-.374-1.297-.789.027-.216.325-.437.893-.663 3.498-1.524 5.831-2.529 6.998-3.014 3.332-1.386 4.025-1.627 4.476-1.635.099-.002.321.023.465.141a.506.506 0 01.171.325c.016.093.036.306.02.472z"/></svg>
            {{ m === 'whatsapp' ? 'WhatsApp' : m === 'telegram' ? 'Telegram' : m === 'viber' ? 'Viber' : m }}
          </div>
        </div>
      </div>

      <!-- Footer -->
      <footer class="footer">
        <p>Сделано на <a href="https://ab-team.ru" target="_blank">AB Team</a></p>
      </footer>

      <!-- Item detail modal -->
      <Transition name="fade">
        <div class="modal-overlay" v-if="selectedItem" @click.self="selectedItem = null">
          <div class="detail-modal">
            <button class="modal-close" @click="selectedItem = null">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M18 6L6 18M6 6l12 12"/></svg>
            </button>
            <div class="detail-img" v-if="selectedItem.image">
              <img :src="imageUrl(selectedItem.image)" :alt="selectedItem.name" />
            </div>
            <div class="detail-body">
              <h3 class="detail-name">{{ selectedItem.name }}</h3>
              <div class="detail-price">{{ formatPrice(selectedItem.price) }}</div>
              <p class="detail-desc" v-if="selectedItem.description">{{ selectedItem.description }}</p>
              <div class="detail-meta" v-if="selectedItem.weight || selectedItem.cook_time">
                <div class="meta-chip" v-if="selectedItem.weight">{{ selectedItem.weight }}</div>
                <div class="meta-chip" v-if="selectedItem.cook_time">{{ selectedItem.cook_time }}</div>
              </div>
              <div class="kbzhu" v-if="selectedItem.calories || selectedItem.proteins || selectedItem.fats || selectedItem.carbs">
                <div class="kbzhu-label">Пищевая ценность</div>
                <div class="kbzhu-grid">
                  <div class="kbzhu-item" v-if="selectedItem.calories"><span class="kbzhu-val">{{ selectedItem.calories }}</span><span class="kbzhu-unit">ккал</span></div>
                  <div class="kbzhu-item" v-if="selectedItem.proteins"><span class="kbzhu-val">{{ selectedItem.proteins }}</span><span class="kbzhu-unit">белки</span></div>
                  <div class="kbzhu-item" v-if="selectedItem.fats"><span class="kbzhu-val">{{ selectedItem.fats }}</span><span class="kbzhu-unit">жиры</span></div>
                  <div class="kbzhu-item" v-if="selectedItem.carbs"><span class="kbzhu-val">{{ selectedItem.carbs }}</span><span class="kbzhu-unit">углев.</span></div>
                </div>
              </div>
              <div class="detail-section" v-if="selectedItem.ingredients">
                <div class="detail-section-label">Состав</div>
                <p class="detail-section-text">{{ selectedItem.ingredients }}</p>
              </div>
              <div class="detail-actions">
                <div class="detail-qty" v-if="getQty(selectedItem.id)">
                  <button class="detail-qty-btn" @click="decrement(selectedItem.id)">
                    <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"><path d="M5 12h14"/></svg>
                  </button>
                  <span class="detail-qty-val">{{ getQty(selectedItem.id) }}</span>
                  <button class="detail-qty-btn" @click="increment(selectedItem)">
                    <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"><path d="M12 5v14M5 12h14"/></svg>
                  </button>
                </div>
                <button v-else class="detail-add-btn" @click="increment(selectedItem)">Добавить в корзину</button>
              </div>
            </div>
          </div>
        </div>
      </Transition>

      <!-- Cart floating button -->
      <Transition name="slide-up">
        <button v-if="cartCount > 0" class="cart-fab" @click="goToCart">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
            <circle cx="9" cy="21" r="1"/><circle cx="20" cy="21" r="1"/>
            <path d="M1 1h4l2.68 13.39a2 2 0 002 1.61h9.72a2 2 0 002-1.61L23 6H6"/>
          </svg>
          <span>Корзина</span>
          <span class="fab-badge">{{ cartCount }}</span>
        </button>
      </Transition>
    </template>
  </div>
</template>

<style scoped>
.home {
  padding-bottom: 0;
}

/* Loading & Error */
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

@keyframes spin { to { transform: rotate(360deg); } }

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

/* ===== Hero ===== */
.hero {
  position: relative;
  padding: 72px 24px 40px;
  text-align: center;
  overflow: hidden;
  min-height: 280px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.hero-bg {
  position: absolute;
  inset: 0;
  background: linear-gradient(180deg, #1F2937 0%, #374151 100%);
  z-index: 0;
}

.hero-with-image .hero-bg {
  background: #000;
}

.hero-bg .hero-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  opacity: 0.55;
}

.hero-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(180deg, rgba(0,0,0,0.1) 0%, rgba(0,0,0,0.5) 100%);
  z-index: 1;
}

.hero-content {
  position: relative;
  z-index: 2;
}

.logo-circle {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  margin: 0 auto 16px;
  overflow: hidden;
  border: 3px solid rgba(255,255,255,0.25);
  box-shadow: 0 4px 20px rgba(0,0,0,0.3);
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
  letter-spacing: 0.5px;
  text-shadow: 0 2px 8px rgba(0,0,0,0.3);
}

.subtitle {
  color: rgba(255,255,255,0.75);
  font-size: 14px;
  line-height: 1.5;
  max-width: 320px;
  margin: 0 auto;
}

/* ===== Promo Card ===== */
.promo-card {
  margin: 16px 16px 0;
  padding: 18px 20px;
  background: linear-gradient(135deg, #FFF7ED 0%, #FEF3C7 100%);
  border: 1px solid #FDE68A;
  border-radius: var(--radius);
  position: relative;
}

.promo-badge {
  display: inline-block;
  padding: 3px 10px;
  background: #F59E0B;
  color: #fff;
  font-size: 11px;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  border-radius: 100px;
  margin-bottom: 10px;
}

.promo-title {
  font-size: 17px;
  font-weight: 700;
  color: #92400E;
  line-height: 1.3;
}

.promo-desc {
  font-size: 13px;
  color: #A16207;
  margin-top: 6px;
  line-height: 1.4;
}

/* ===== Sections ===== */
.section {
  padding: 20px 16px 0;
}

.section-header {
  margin-bottom: 14px;
}

.section-header h2 {
  font-size: 18px;
  font-weight: 700;
  color: var(--text);
}

/* ===== Popular Grid ===== */
.popular-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 10px;
}

.popular-card {
  background: var(--bg);
  border: 1px solid var(--border);
  border-radius: var(--radius);
  overflow: hidden;
  cursor: pointer;
  position: relative;
  transition: transform 0.15s ease;
}

.popular-card:active {
  transform: scale(0.97);
}

.pop-img {
  width: 100%;
  aspect-ratio: 4 / 3;
  overflow: hidden;
  background: var(--bg-secondary);
}

.pop-img img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.pop-img-empty {
  display: flex;
  align-items: center;
  justify-content: center;
}

.pop-info {
  padding: 10px 12px;
}

.pop-name {
  font-size: 14px;
  font-weight: 600;
  line-height: 1.3;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.pop-price {
  font-size: 15px;
  font-weight: 700;
  margin-top: 4px;
  font-variant-numeric: tabular-nums;
}

.pop-add {
  position: absolute;
  bottom: 10px;
  right: 10px;
  width: 30px;
  height: 30px;
  border-radius: 8px;
  background: var(--primary);
  color: var(--primary-foreground);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  font-weight: 500;
  box-shadow: 0 2px 8px rgba(0,0,0,0.15);
}

.pop-add:active {
  opacity: 0.8;
}

.pop-stopped {
  opacity: 0.5;
  pointer-events: none;
  position: relative;
}

.pop-stopped-badge {
  position: absolute;
  top: 8px;
  left: 8px;
  padding: 3px 10px;
  background: rgba(0,0,0,0.6);
  color: #fff;
  font-size: 11px;
  font-weight: 600;
  border-radius: 100px;
  z-index: 2;
}

.pop-qty {
  position: absolute;
  bottom: 8px;
  right: 8px;
  display: flex;
  align-items: center;
  gap: 2px;
  background: var(--primary);
  border-radius: 8px;
  padding: 2px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.15);
}

.pop-qty-btn {
  width: 26px;
  height: 26px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--primary-foreground);
  font-size: 16px;
  font-weight: 600;
}

.pop-qty-btn:active {
  opacity: 0.7;
}

.pop-qty span {
  min-width: 18px;
  text-align: center;
  font-size: 13px;
  font-weight: 700;
  color: var(--primary-foreground);
}

/* ===== Actions ===== */
.actions {
  padding: 24px 16px 0;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.btn-primary {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  width: 100%;
  padding: 16px;
  background: var(--primary);
  color: var(--primary-foreground);
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
  border: 1.5px solid var(--border);
  border-radius: var(--radius);
  font-size: 16px;
  font-weight: 600;
  transition: border-color var(--ease);
}

.btn-outline:active {
  border-color: var(--primary);
}

/* ===== Info Cards ===== */
.info-card {
  display: flex;
  gap: 14px;
  padding: 16px;
  background: var(--bg);
  border: 1px solid var(--border);
  border-radius: var(--radius);
}

.info-icon {
  width: 38px;
  height: 38px;
  border-radius: 10px;
  background: var(--bg-secondary);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.info-body {
  flex: 1;
  min-width: 0;
}

.info-label {
  font-size: 11px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.8px;
  color: var(--text-muted);
  margin-bottom: 8px;
}

.info-value {
  font-size: 14px;
  color: var(--text);
  line-height: 1.5;
}

.hours-grid {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.hours-row {
  display: flex;
  justify-content: space-between;
  font-size: 14px;
  padding: 4px 0;
}

.day {
  color: var(--text-secondary);
}

.time {
  font-weight: 600;
}

/* ===== Social links ===== */
.social-list {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.social-chip {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 18px;
  background: var(--bg-secondary);
  border-radius: 100px;
  font-size: 14px;
  font-weight: 500;
  color: var(--text);
  text-decoration: none;
  transition: background var(--ease);
}

.social-chip:active {
  background: var(--border);
}

/* ===== Messengers ===== */
.msg-list {
  display: flex;
  gap: 10px;
}

.msg-chip {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 18px;
  background: var(--bg-secondary);
  border-radius: 100px;
  font-size: 14px;
  font-weight: 500;
}

/* ===== Footer ===== */
.footer {
  padding: 32px 16px 40px;
  text-align: center;
}

.footer p {
  font-size: 12px;
  color: var(--text-muted);
}

.footer a {
  color: var(--text-secondary);
  font-weight: 500;
}

/* Modal */
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.5);
  z-index: 100;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
}

.detail-modal {
  background: var(--bg);
  border-radius: 16px;
  width: 100%;
  max-width: 500px;
  max-height: 85vh;
  overflow-y: auto;
  position: relative;
}

.modal-close {
  position: absolute;
  top: 12px;
  right: 12px;
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: rgba(0,0,0,0.4);
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 2;
}

.detail-img {
  width: 100%;
  aspect-ratio: 16 / 10;
  overflow: hidden;
}

.detail-img img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.detail-body {
  padding: 20px 20px 24px;
}

.detail-name {
  font-size: 22px;
  font-weight: 700;
  line-height: 1.3;
}

.detail-price {
  font-size: 20px;
  font-weight: 700;
  color: var(--primary);
  margin-top: 6px;
}

.detail-desc {
  font-size: 14px;
  color: var(--text-secondary);
  line-height: 1.5;
  margin-top: 12px;
}

.detail-meta {
  display: flex;
  gap: 10px;
  margin-top: 14px;
}

.meta-chip {
  padding: 6px 12px;
  background: var(--bg-secondary);
  border-radius: 100px;
  font-size: 13px;
  color: var(--text-secondary);
  font-weight: 500;
}

.kbzhu {
  margin-top: 16px;
  padding: 14px;
  background: var(--bg-secondary);
  border-radius: var(--radius);
}

.kbzhu-label {
  font-size: 11px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: var(--text-muted);
  margin-bottom: 10px;
}

.kbzhu-grid {
  display: flex;
  gap: 6px;
}

.kbzhu-item {
  flex: 1;
  text-align: center;
  padding: 8px 4px;
  background: var(--bg);
  border-radius: 8px;
}

.kbzhu-val {
  display: block;
  font-size: 16px;
  font-weight: 700;
}

.kbzhu-unit {
  display: block;
  font-size: 11px;
  color: var(--text-muted);
  margin-top: 2px;
}

.detail-section {
  margin-top: 16px;
}

.detail-section-label {
  font-size: 11px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: var(--text-muted);
  margin-bottom: 6px;
}

.detail-section-text {
  font-size: 14px;
  color: var(--text-secondary);
  line-height: 1.5;
}

.detail-actions {
  margin-top: 20px;
}

.detail-add-btn {
  width: 100%;
  padding: 16px;
  background: var(--primary);
  color: var(--primary-foreground);
  border-radius: var(--radius);
  font-size: 16px;
  font-weight: 600;
}

.detail-add-btn:active {
  opacity: 0.85;
}

.detail-qty {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 4px;
  background: var(--bg-secondary);
  border-radius: var(--radius);
  padding: 6px;
}

.detail-qty-btn {
  width: 44px;
  height: 44px;
  border-radius: 10px;
  background: var(--bg);
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--primary);
}

.detail-qty-btn:active {
  background: var(--border);
}

.detail-qty-val {
  min-width: 40px;
  text-align: center;
  font-size: 20px;
  font-weight: 700;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.25s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* Cart FAB */
.cart-fab {
  position: fixed;
  bottom: 24px;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 14px 28px;
  background: var(--primary);
  color: var(--primary-foreground);
  border-radius: 100px;
  font-size: 15px;
  font-weight: 600;
  box-shadow: var(--shadow-lg);
  z-index: 50;
  transition: opacity var(--ease);
}

.cart-fab:active {
  opacity: 0.9;
}

.fab-badge {
  background: #fff;
  color: var(--primary);
  width: 22px;
  height: 22px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 700;
}

.slide-up-enter-active,
.slide-up-leave-active {
  transition: all 0.3s ease;
}
.slide-up-enter-from,
.slide-up-leave-to {
  transform: translateX(-50%) translateY(80px);
  opacity: 0;
}
</style>
