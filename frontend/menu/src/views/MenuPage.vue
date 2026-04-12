<script setup>
import { ref, computed, onMounted, onBeforeUnmount, watch, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { restaurant, loading } from '../stores/restaurant'
import { cart, addToCart, cartCount, updateQuantity, updateComment } from '../stores/cart'
import { imageUrl } from '../api/client'

const route = useRoute()
const router = useRouter()
const slug = route.params.slug

const searchQuery = ref('')
const categories = computed(() => restaurant.categories || [])
const combos = computed(() => (restaurant.combos || []).filter(c => c.available))
const selectedItem = ref(null)
const activeCategory = ref(null)

const displayCategories = computed(() => {
  const q = searchQuery.value.trim().toLowerCase()
  if (!q) return categories.value
  return categories.value.map(cat => ({
    ...cat,
    items: (cat.items || []).filter(item => item.name.toLowerCase().includes(q))
  })).filter(cat => cat.items.length > 0)
})

function scrollToCategory(id) {
  const el = document.getElementById('cat-' + id)
  if (!el) return
  const y = el.getBoundingClientRect().top + window.scrollY - 120
  window.scrollTo({ top: y, behavior: 'smooth' })
}

let observer = null

function setupObserver() {
  if (observer) observer.disconnect()
  const els = categories.value.map(c => document.getElementById('cat-' + c.id)).filter(Boolean)
  if (!els.length) return
  observer = new IntersectionObserver((entries) => {
    const visible = entries.filter(e => e.isIntersecting).sort((a, b) => a.boundingClientRect.top - b.boundingClientRect.top)
    if (visible.length) {
      const id = Number(visible[0].target.id.replace('cat-', ''))
      activeCategory.value = id
    }
  }, { rootMargin: '-130px 0px -60% 0px', threshold: 0 })
  els.forEach(el => observer.observe(el))
}

onMounted(() => {
  nextTick(setupObserver)
})
onBeforeUnmount(() => observer?.disconnect())
watch(categories, () => nextTick(setupObserver))
watch(activeCategory, (id) => {
  if (!id) return
  const tab = document.querySelector('.tab.active')
  if (tab && tab.scrollIntoView) {
    tab.scrollIntoView({ behavior: 'smooth', block: 'nearest', inline: 'center' })
  }
})

function formatPrice(kopecks) {
  return Math.floor(kopecks / 100).toLocaleString('ru-RU') + ' \u20BD'
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

const detailComment = ref('')

function openDetail(item) {
  selectedItem.value = item
  const ci = cart.find(i => i.id === item.id)
  detailComment.value = ci?.comment || ''
}

function addWithComment() {
  if (getQty(selectedItem.value.id)) {
    updateComment(selectedItem.value.id, detailComment.value)
  } else {
    addToCart(selectedItem.value, detailComment.value)
  }
}

function goToCart() {
  router.push({ name: 'cart', params: { slug } })
}

function addComboToCart(combo) {
  addToCart({ id: 'combo-' + combo.id, name: combo.name, price: combo.price, image: combo.image })
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

    <!-- Category tabs (scroll-to-section) -->
    <div class="tabs-wrap" v-if="categories.length > 1 && !searchQuery">
      <div class="tabs">
        <button
          class="tab"
          v-for="cat in categories"
          :key="cat.id"
          :class="{ active: activeCategory === cat.id }"
          @click="scrollToCategory(cat.id)"
        >{{ cat.name }}</button>
      </div>
    </div>

    <!-- Search -->
    <div class="search-bar">
      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="var(--text-muted)" stroke-width="2" stroke-linecap="round">
        <circle cx="11" cy="11" r="8"/><path d="M21 21l-4.35-4.35"/>
      </svg>
      <input v-model="searchQuery" type="text" placeholder="Поиск блюд..." class="search-input" />
      <button v-if="searchQuery" class="search-clear" @click="searchQuery = ''">
        <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M18 6L6 18M6 6l12 12"/></svg>
      </button>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="loading">
      <div class="spinner"></div>
    </div>

    <!-- Content -->
    <div class="content" v-else>
      <!-- Акция -->
      <div class="promo-card" v-if="restaurant.promo_title && !searchQuery">
        <div class="promo-badge">Акция</div>
        <h3 class="promo-title">{{ restaurant.promo_title }}</h3>
        <p class="promo-desc" v-if="restaurant.promo_description">{{ restaurant.promo_description }}</p>
      </div>

      <!-- Combos -->
      <div class="combos-section" v-if="combos.length && !searchQuery">
        <div class="category-name">Комбо-наборы</div>
        <div class="combos-list">
          <div class="combo-card" v-for="combo in combos" :key="combo.id">
            <div class="combo-img" v-if="combo.image">
              <img :src="imageUrl(combo.image)" :alt="combo.name" loading="lazy" />
            </div>
            <div class="combo-body">
              <div class="combo-name">{{ combo.name }}</div>
              <div class="combo-desc" v-if="combo.description">{{ combo.description }}</div>
              <div class="combo-contents" v-if="combo.items?.length">
                <span v-for="ci in combo.items" :key="ci.name" class="combo-chip">{{ ci.name }}<template v-if="ci.quantity > 1"> ×{{ ci.quantity }}</template></span>
              </div>
              <div class="combo-footer">
                <span class="combo-price">{{ formatPrice(combo.price) }}</span>
                <button class="add-btn" @click="addComboToCart(combo)">
                  <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"><path d="M12 5v14M5 12h14"/></svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div v-for="cat in displayCategories" :key="cat.id" :id="'cat-' + cat.id" class="category-section">
        <div class="category-name" v-if="categories.length > 1">{{ cat.name }}</div>

        <div class="items-grid" v-if="cat.items?.length">
          <div class="item-card" :class="{ 'item-stopped': !item.available }" v-for="item in cat.items" :key="item.id" @click="item.available && openDetail(item)">
            <!-- Image -->
            <div class="item-img" v-if="item.image">
              <img :src="imageUrl(item.image)" :alt="item.name" loading="lazy" />
            </div>
            <div class="item-img item-img-placeholder" v-else>
              <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="var(--text-muted)" stroke-width="1.5" stroke-linecap="round">
                <path d="M18 8h1a4 4 0 010 8h-1M2 8h16v9a4 4 0 01-4 4H6a4 4 0 01-4-4V8zM6 1v3M10 1v3M14 1v3"/>
              </svg>
            </div>
            <div class="stopped-badge" v-if="!item.available">Нет в наличии</div>

            <!-- Info -->
            <div class="item-body">
              <div class="item-name">{{ item.name }}</div>
              <div class="item-desc" v-if="item.description">{{ item.description }}</div>
              <div class="item-footer">
                <span class="item-price">{{ formatPrice(item.price) }}</span>
                <template v-if="item.available">
                  <!-- Quantity controls or add button -->
                  <div class="qty-inline" v-if="getQty(item.id)" @click.stop>
                    <button class="qty-btn-sm" @click="decrement(item.id)">
                      <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"><path d="M5 12h14"/></svg>
                    </button>
                    <span class="qty-val">{{ getQty(item.id) }}</span>
                    <button class="qty-btn-sm" @click="increment(item)">
                      <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"><path d="M12 5v14M5 12h14"/></svg>
                    </button>
                  </div>
                  <button v-else class="add-btn" @click.stop="increment(item)">
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round">
                      <path d="M12 5v14M5 12h14"/>
                    </svg>
                  </button>
                </template>
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

            <!-- Метаданные -->
            <div class="detail-meta" v-if="selectedItem.weight || selectedItem.cook_time">
              <div class="meta-chip" v-if="selectedItem.weight">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M12 3v18M3 12h18"/><circle cx="12" cy="12" r="9"/></svg>
                {{ selectedItem.weight }}
              </div>
              <div class="meta-chip" v-if="selectedItem.cook_time">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l4 2"/></svg>
                {{ selectedItem.cook_time }}
              </div>
            </div>

            <!-- КБЖУ -->
            <div class="kbzhu" v-if="selectedItem.calories || selectedItem.proteins || selectedItem.fats || selectedItem.carbs">
              <div class="kbzhu-label">Пищевая ценность</div>
              <div class="kbzhu-grid">
                <div class="kbzhu-item" v-if="selectedItem.calories">
                  <span class="kbzhu-val">{{ selectedItem.calories }}</span>
                  <span class="kbzhu-unit">ккал</span>
                </div>
                <div class="kbzhu-item" v-if="selectedItem.proteins">
                  <span class="kbzhu-val">{{ selectedItem.proteins }}</span>
                  <span class="kbzhu-unit">белки</span>
                </div>
                <div class="kbzhu-item" v-if="selectedItem.fats">
                  <span class="kbzhu-val">{{ selectedItem.fats }}</span>
                  <span class="kbzhu-unit">жиры</span>
                </div>
                <div class="kbzhu-item" v-if="selectedItem.carbs">
                  <span class="kbzhu-val">{{ selectedItem.carbs }}</span>
                  <span class="kbzhu-unit">углев.</span>
                </div>
              </div>
            </div>

            <!-- Ингредиенты -->
            <div class="detail-section" v-if="selectedItem.ingredients">
              <div class="detail-section-label">Состав</div>
              <p class="detail-section-text">{{ selectedItem.ingredients }}</p>
            </div>

            <!-- Комментарий к блюду -->
            <div class="detail-comment">
              <input v-model="detailComment" class="detail-comment-input" placeholder="Комментарий к блюду..." @change="getQty(selectedItem.id) && updateComment(selectedItem.id, detailComment)" />
            </div>

            <!-- Кнопка добавить / счётчик -->
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
              <button v-else class="detail-add-btn" @click="addWithComment()">
                Добавить в корзину
              </button>
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

/* Search */
.search-bar {
  display: flex;
  align-items: center;
  gap: 10px;
  margin: 12px 16px 0;
  padding: 10px 14px;
  background: var(--bg-secondary);
  border-radius: var(--radius);
  border: 1px solid var(--border);
}

.search-input {
  flex: 1;
  border: none;
  background: none;
  font-size: 14px;
  color: var(--text);
  outline: none;
}

.search-input::placeholder {
  color: var(--text-muted);
}

.search-clear {
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text-muted);
  border-radius: 50%;
}

.search-clear:active {
  background: var(--border);
}

/* Content */
.content {
  padding: 16px;
}

/* ===== Promo Card ===== */
.promo-card {
  margin-bottom: 16px;
  padding: 18px 20px;
  background: linear-gradient(135deg, #FFF7ED 0%, #FEF3C7 100%);
  border: 1px solid #FDE68A;
  border-radius: var(--radius);
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

/* Combos */
.combos-section {
  margin-bottom: 24px;
}

.combos-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.combo-card {
  background: var(--bg);
  border-radius: var(--radius);
  overflow: hidden;
  box-shadow: var(--shadow-sm);
  border: 1px solid var(--border);
  display: flex;
}

.combo-img {
  width: 100px;
  flex-shrink: 0;
  overflow: hidden;
}

.combo-img img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.combo-body {
  flex: 1;
  padding: 12px;
  min-width: 0;
}

.combo-name {
  font-size: 15px;
  font-weight: 600;
  line-height: 1.3;
}

.combo-desc {
  font-size: 12px;
  color: var(--text-secondary);
  margin-top: 2px;
}

.combo-contents {
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
  margin-top: 6px;
}

.combo-chip {
  padding: 2px 8px;
  background: var(--bg-secondary);
  border-radius: 100px;
  font-size: 11px;
  color: var(--text-secondary);
}

.combo-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 8px;
}

.combo-price {
  font-size: 16px;
  font-weight: 700;
  color: var(--text);
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

.item-stopped {
  opacity: 0.5;
  pointer-events: none;
  position: relative;
}

.item-stopped .add-btn,
.item-stopped .qty-inline {
  display: none;
}

.stopped-badge {
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

/* Quantity inline on card */
.qty-inline {
  display: flex;
  align-items: center;
  gap: 2px;
  background: var(--bg-secondary);
  border-radius: 10px;
  padding: 2px;
}

.qty-btn-sm {
  width: 28px;
  height: 28px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--primary);
  transition: background var(--ease);
}

.qty-btn-sm:active {
  background: var(--border);
}

.qty-val {
  min-width: 20px;
  text-align: center;
  font-size: 14px;
  font-weight: 700;
  font-variant-numeric: tabular-nums;
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
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 6px 12px;
  background: var(--bg-secondary);
  border-radius: 100px;
  font-size: 13px;
  color: var(--text-secondary);
  font-weight: 500;
}

/* КБЖУ */
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
  font-variant-numeric: tabular-nums;
}

.kbzhu-unit {
  display: block;
  font-size: 11px;
  color: var(--text-muted);
  margin-top: 2px;
}

/* Detail sections */
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

/* Detail comment */
.detail-comment {
  margin-top: 16px;
}

.detail-comment-input {
  width: 100%;
  padding: 12px 14px;
  border: 1.5px solid var(--border);
  border-radius: var(--radius-sm, 8px);
  font-size: 14px;
  color: var(--text);
  background: var(--bg);
  outline: none;
  transition: border-color var(--ease);
}

.detail-comment-input:focus {
  border-color: var(--primary);
}

.detail-comment-input::placeholder {
  color: var(--text-muted);
}

/* Detail actions */
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
  transition: opacity var(--ease);
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
  font-weight: 600;
  transition: background var(--ease);
}

.detail-qty-btn:active {
  background: var(--border);
}

.detail-qty-val {
  min-width: 40px;
  text-align: center;
  font-size: 20px;
  font-weight: 700;
  font-variant-numeric: tabular-nums;
}

/* Transitions */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.25s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
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
