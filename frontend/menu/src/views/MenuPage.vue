<script setup>
import { ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { restaurant, loading } from '../stores/restaurant'
import { addToCart, cartCount } from '../stores/cart'

const route = useRoute()
const router = useRouter()
const slug = route.params.slug

const activeCategory = ref(null)
const categories = computed(() => restaurant.categories || [])

const displayCategories = computed(() => {
  if (!activeCategory.value) return categories.value
  const cat = categories.value.find(c => c.id === activeCategory.value)
  return cat ? [cat] : []
})

function selectCategory(id) {
  activeCategory.value = activeCategory.value === id ? null : id
}

function formatPrice(kopecks) {
  return Math.floor(kopecks / 100).toLocaleString('ru-RU') + ' \u20BD'
}

function handleAdd(item) {
  addToCart(item)
}

function goToCart() {
  router.push({ name: 'cart', params: { slug } })
}

function goBack() {
  router.push({ name: 'home', params: { slug } })
}
</script>

<template>
  <div class="menu-page">
    <!-- Header -->
    <div class="header">
      <button class="back-btn" @click="goBack">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
          <path d="M19 12H5M12 19l-7-7 7-7"/>
        </svg>
      </button>
      <h2>{{ restaurant.name }}</h2>
      <button class="cart-btn" @click="goToCart" v-if="cartCount > 0">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
          <circle cx="9" cy="21" r="1"/><circle cx="20" cy="21" r="1"/>
          <path d="M1 1h4l2.68 13.39a2 2 0 002 1.61h9.72a2 2 0 002-1.61L23 6H6"/>
        </svg>
        <span class="cart-count">{{ cartCount }}</span>
      </button>
      <div class="spacer" v-else></div>
    </div>

    <!-- Category tabs -->
    <div class="tabs-wrap" v-if="categories.length > 1">
      <div class="tabs">
        <button
          class="tab"
          :class="{ active: !activeCategory }"
          @click="activeCategory = null"
        >Все</button>
        <button
          class="tab"
          v-for="cat in categories"
          :key="cat.id"
          :class="{ active: activeCategory === cat.id }"
          @click="selectCategory(cat.id)"
        >{{ cat.name }}</button>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="loading">
      <div class="spinner"></div>
    </div>

    <!-- Content -->
    <div class="content" v-else>
      <div v-for="cat in displayCategories" :key="cat.id" class="category-section">
        <div class="category-name" v-if="!activeCategory && categories.length > 1">{{ cat.name }}</div>

        <div class="items-grid" v-if="cat.items?.length">
          <div class="item-card" v-for="item in cat.items" :key="item.id" @click="handleAdd(item)">
            <!-- Image -->
            <div class="item-img" v-if="item.image">
              <img :src="item.image" :alt="item.name" loading="lazy" />
            </div>
            <div class="item-img item-img-placeholder" v-else>
              <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="var(--text-muted)" stroke-width="1.5" stroke-linecap="round">
                <path d="M18 8h1a4 4 0 010 8h-1M2 8h16v9a4 4 0 01-4 4H6a4 4 0 01-4-4V8zM6 1v3M10 1v3M14 1v3"/>
              </svg>
            </div>

            <!-- Info -->
            <div class="item-body">
              <div class="item-name">{{ item.name }}</div>
              <div class="item-desc" v-if="item.description">{{ item.description }}</div>
              <div class="item-footer">
                <span class="item-price">{{ formatPrice(item.price) }}</span>
                <button class="add-btn" @click.stop="handleAdd(item)">
                  <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round">
                    <path d="M12 5v14M5 12h14"/>
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>

        <div class="empty-cat" v-else>
          <p>Пока пусто</p>
        </div>
      </div>

      <div class="empty" v-if="categories.length === 0">
        <p>Меню пока формируется</p>
      </div>
    </div>

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
  </div>
</template>

<style scoped>
.menu-page {
  padding-bottom: 100px;
  min-height: 100vh;
}

/* Header */
.header {
  display: flex;
  align-items: center;
  padding: 14px 16px;
  gap: 12px;
  position: sticky;
  top: 0;
  background: rgba(255,255,255,0.92);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  z-index: 20;
  border-bottom: 1px solid var(--border);
}

.header h2 {
  font-size: 17px;
  font-weight: 700;
  flex: 1;
  text-align: center;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.back-btn, .cart-btn {
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  color: var(--text);
  position: relative;
  flex-shrink: 0;
  transition: background var(--ease);
}

.back-btn:active, .cart-btn:active {
  background: var(--bg-secondary);
}

.cart-count {
  position: absolute;
  top: 2px;
  right: 2px;
  width: 16px;
  height: 16px;
  border-radius: 50%;
  background: var(--accent);
  color: #fff;
  font-size: 10px;
  font-weight: 700;
  display: flex;
  align-items: center;
  justify-content: center;
}

.spacer {
  width: 36px;
  flex-shrink: 0;
}

/* Tabs */
.tabs-wrap {
  position: sticky;
  top: 65px;
  z-index: 19;
  background: rgba(255,255,255,0.92);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  border-bottom: 1px solid var(--border);
}

.tabs {
  display: flex;
  gap: 8px;
  padding: 12px 16px;
  overflow-x: auto;
  -webkit-overflow-scrolling: touch;
  scrollbar-width: none;
}

.tabs::-webkit-scrollbar {
  display: none;
}

.tab {
  padding: 8px 18px;
  background: var(--bg-secondary);
  border-radius: 100px;
  font-size: 13px;
  font-weight: 500;
  color: var(--text-secondary);
  white-space: nowrap;
  transition: all var(--ease);
  flex-shrink: 0;
}

.tab.active {
  background: var(--primary);
  color: var(--primary-foreground);
  font-weight: 600;
}

/* Content */
.content {
  padding: 16px;
}

.category-section {
  margin-bottom: 24px;
}

.category-section:last-child {
  margin-bottom: 0;
}

.category-name {
  font-size: 18px;
  font-weight: 700;
  color: var(--text);
  padding: 4px 0 14px;
}

/* Grid layout */
.items-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
}

.item-card {
  background: var(--bg);
  border-radius: var(--radius);
  overflow: hidden;
  box-shadow: var(--shadow-sm);
  border: 1px solid var(--border);
  cursor: pointer;
  transition: transform 0.15s ease, box-shadow 0.15s ease;
}

.item-card:active {
  transform: scale(0.97);
}

/* Image */
.item-img {
  width: 100%;
  aspect-ratio: 1 / 1;
  overflow: hidden;
  background: var(--bg-secondary);
}

.item-img img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.item-img-placeholder {
  display: flex;
  align-items: center;
  justify-content: center;
}

/* Body */
.item-body {
  padding: 12px;
}

.item-name {
  font-size: 14px;
  font-weight: 600;
  line-height: 1.3;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.item-desc {
  font-size: 12px;
  color: var(--text-secondary);
  margin-top: 4px;
  line-height: 1.3;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.item-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 10px;
}

.item-price {
  font-size: 15px;
  font-weight: 700;
  font-variant-numeric: tabular-nums;
  color: var(--text);
}

.add-btn {
  width: 32px;
  height: 32px;
  border-radius: 10px;
  background: var(--primary);
  color: var(--primary-foreground);
  display: flex;
  align-items: center;
  justify-content: center;
  transition: opacity var(--ease);
  flex-shrink: 0;
}

.add-btn:active {
  opacity: 0.8;
}

/* Empty states */
.empty-cat {
  padding: 32px 0;
  text-align: center;
  color: var(--text-muted);
  font-size: 14px;
}

.empty {
  padding: 80px 0;
  text-align: center;
  color: var(--text-secondary);
  font-size: 15px;
}

/* Loading */
.loading {
  display: flex;
  justify-content: center;
  padding: 80px 0;
}

.spinner {
  width: 32px;
  height: 32px;
  border: 3px solid var(--border);
  border-top-color: var(--primary);
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}

@keyframes spin { to { transform: rotate(360deg); } }

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

/* Responsive */
@media (max-width: 340px) {
  .items-grid {
    grid-template-columns: 1fr;
  }
}
</style>
