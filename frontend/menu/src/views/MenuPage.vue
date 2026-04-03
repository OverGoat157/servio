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

const filteredItems = computed(() => {
  if (!activeCategory.value) {
    return categories.value.flatMap(c => c.items || [])
  }
  const cat = categories.value.find(c => c.id === activeCategory.value)
  return cat ? cat.items || [] : []
})

function selectCategory(id) {
  activeCategory.value = activeCategory.value === id ? null : id
}

function formatPrice(kopecks) {
  return Math.floor(kopecks / 100).toLocaleString('ru-RU') + ' ₽'
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
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M19 12H5M12 19l-7-7 7-7"/>
        </svg>
      </button>
      <h2>{{ restaurant.name }}</h2>
      <div class="spacer"></div>
    </div>

    <!-- Category tabs -->
    <div class="tabs" v-if="categories.length > 1">
      <button
        class="tab"
        :class="{ active: !activeCategory }"
        @click="activeCategory = null"
      >
        Все
      </button>
      <button
        class="tab"
        v-for="cat in categories"
        :key="cat.id"
        :class="{ active: activeCategory === cat.id }"
        @click="selectCategory(cat.id)"
      >
        {{ cat.name }}
      </button>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="loading">
      <div class="spinner"></div>
    </div>

    <!-- Items -->
    <div class="items" v-else>
      <template v-if="activeCategory === null">
        <div v-for="cat in categories" :key="cat.id" class="category-group">
          <div class="category-title" v-if="categories.length > 1">{{ cat.name }}</div>
          <div
            class="item-card"
            v-for="item in cat.items"
            :key="item.id"
          >
            <div class="item-img" v-if="item.image">
              <img :src="item.image" :alt="item.name" />
            </div>
            <div class="item-info">
              <div class="item-name">{{ item.name }}</div>
              <div class="item-desc" v-if="item.description">{{ item.description }}</div>
            </div>
            <div class="item-right">
              <div class="item-price">{{ formatPrice(item.price) }}</div>
              <button class="add-btn" @click="handleAdd(item)">+</button>
            </div>
          </div>
        </div>
      </template>

      <template v-else>
        <div
          class="item-card"
          v-for="item in filteredItems"
          :key="item.id"
        >
          <div class="item-img" v-if="item.image">
            <img :src="item.image" :alt="item.name" />
          </div>
          <div class="item-info">
            <div class="item-name">{{ item.name }}</div>
            <div class="item-desc" v-if="item.description">{{ item.description }}</div>
          </div>
          <div class="item-right">
            <div class="item-price">{{ formatPrice(item.price) }}</div>
            <button class="add-btn" @click="handleAdd(item)">+</button>
          </div>
        </div>
      </template>

      <div class="empty" v-if="filteredItems.length === 0 && !loading">
        <p>В этой категории пока нет блюд</p>
      </div>
    </div>

    <!-- Cart floating button -->
    <Transition name="slide-up">
      <button v-if="cartCount > 0" class="cart-fab" @click="goToCart">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="9" cy="21" r="1"/><circle cx="20" cy="21" r="1"/>
          <path d="M1 1h4l2.68 13.39a2 2 0 002 1.61h9.72a2 2 0 002-1.61L23 6H6"/>
        </svg>
        <span>Корзина</span>
        <span class="cart-badge">{{ cartCount }}</span>
      </button>
    </Transition>
  </div>
</template>

<style scoped>
.menu-page {
  padding-bottom: 100px;
}

.header {
  display: flex;
  align-items: center;
  padding: 16px;
  gap: 12px;
  position: sticky;
  top: 0;
  background: var(--bg);
  z-index: 10;
  border-bottom: 1px solid var(--border);
}

.header h2 {
  font-size: 18px;
  font-weight: 700;
  flex: 1;
  text-align: center;
}

.back-btn {
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  color: var(--text);
  transition: background var(--ease);
}

.back-btn:active {
  background: var(--bg-secondary);
}

.spacer {
  width: 36px;
}

/* Tabs */
.tabs {
  display: flex;
  gap: 8px;
  padding: 12px 16px;
  overflow-x: auto;
  -webkit-overflow-scrolling: touch;
  position: sticky;
  top: 69px;
  background: var(--bg);
  z-index: 9;
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
}

.tab.active {
  background: var(--primary);
  color: #fff;
  font-weight: 600;
}

/* Loading */
.loading {
  display: flex;
  justify-content: center;
  padding: 60px 0;
}

.spinner {
  width: 32px;
  height: 32px;
  border: 3px solid var(--border);
  border-top-color: var(--primary);
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

/* Items */
.items {
  padding: 8px 16px;
}

.category-group {
  margin-bottom: 8px;
}

.category-title {
  font-size: 13px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: var(--text-secondary);
  padding: 16px 0 8px;
}

.item-card {
  display: flex;
  align-items: center;
  gap: 12px;
  background: var(--bg-secondary);
  border-radius: var(--radius);
  padding: 14px;
  margin-bottom: 10px;
}

.item-img {
  width: 56px;
  height: 56px;
  border-radius: 10px;
  overflow: hidden;
  flex-shrink: 0;
}

.item-img img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.item-info {
  flex: 1;
  min-width: 0;
}

.item-name {
  font-size: 15px;
  font-weight: 600;
}

.item-desc {
  font-size: 12px;
  color: var(--text-secondary);
  margin-top: 2px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.item-right {
  text-align: right;
  flex-shrink: 0;
}

.item-price {
  font-size: 15px;
  font-weight: 700;
  white-space: nowrap;
}

.add-btn {
  width: 28px;
  height: 28px;
  background: var(--primary);
  color: #fff;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  font-weight: 500;
  margin-top: 6px;
  margin-left: auto;
  transition: opacity var(--ease);
}

.add-btn:active {
  opacity: 0.8;
}

.empty {
  text-align: center;
  padding: 60px 0;
  color: var(--text-secondary);
  font-size: 14px;
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
  color: #fff;
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

.cart-badge {
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
