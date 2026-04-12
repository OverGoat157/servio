<script setup>
import { ref, computed, onMounted, onBeforeUnmount, watch, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { restaurant, loading, error } from '../stores/restaurant'
import { cart, addToCart, cartCount, updateQuantity, updateComment } from '../stores/cart'
import { imageUrl } from '../api/client'

const route = useRoute()
const router = useRouter()
const slug = route.params.slug

const categories = computed(() => restaurant.categories || [])
const combos = computed(() => (restaurant.combos || []).filter(c => c.available))

const popularItems = computed(() => {
  const all = categories.value.flatMap(c => (c.items || []).filter(i => i.available))
  return all.slice(0, 8)
})

const activeCategory = ref(null)

function scrollToCategory(id) {
  const el = document.getElementById('cat-' + id)
  if (!el) return
  const y = el.getBoundingClientRect().top + window.scrollY - 80
  window.scrollTo({ top: y, behavior: 'smooth' })
}

let observer = null
function setupObserver() {
  if (observer) observer.disconnect()
  const ids = []
  if (popularItems.value.length) ids.push('popular')
  if (combos.value.length) ids.push('combos')
  ids.push(...categories.value.map(c => c.id))
  const els = ids.map(id => document.getElementById('cat-' + id)).filter(Boolean)
  if (!els.length) return
  observer = new IntersectionObserver((entries) => {
    const visible = entries.filter(e => e.isIntersecting).sort((a, b) => a.boundingClientRect.top - b.boundingClientRect.top)
    if (visible.length) {
      const raw = visible[0].target.id.replace('cat-', '')
      activeCategory.value = (raw === 'popular' || raw === 'combos') ? raw : Number(raw)
    }
  }, { rootMargin: '-60px 0px -60% 0px', threshold: 0 })
  els.forEach(el => observer.observe(el))
}

onMounted(() => { nextTick(setupObserver) })
onBeforeUnmount(() => observer?.disconnect())
watch(categories, () => nextTick(setupObserver))
watch(combos, () => nextTick(setupObserver))
watch(activeCategory, () => {
  nextTick(() => {
    const tab = document.querySelector('.tab.active')
    const container = document.querySelector('.tabs')
    if (!tab || !container) return
    const tabRect = tab.getBoundingClientRect()
    const containerRect = container.getBoundingClientRect()
    const target = container.scrollLeft + (tabRect.left - containerRect.left) - (containerRect.width - tabRect.width) / 2
    container.scrollTo({ left: target, behavior: 'smooth' })
  })
})

function goToCart() {
  router.push({ name: 'cart', params: { slug } })
}

function addComboToCart(combo) {
  addToCart({ id: 'combo-' + combo.id, name: combo.name, price: combo.price, image: combo.image })
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

const parsedSchedule = computed(() => {
  const wh = restaurant.working_hours
  if (!wh) return null

  // New format: array of day objects
  if (Array.isArray(wh)) return wh

  // Old format: { "Пн — Пт": "10:00 — 22:00" } — show as key/value
  if (typeof wh === 'object') {
    return Object.entries(wh).map(([day, time]) => ({ day, time, legacy: true }))
  }

  return null
})

const todayDayName = computed(() => {
  const days = ['Вс', 'Пн', 'Вт', 'Ср', 'Чт', 'Пт', 'Сб']
  return days[new Date().getDay()]
})

const todaySchedule = computed(() => {
  const sched = parsedSchedule.value
  if (!sched?.length) return null
  if (sched[0]?.legacy) return null
  const today = sched.find(s => s.day === todayDayName.value)
  if (!today) return null
  if (today.day_off) return { label: 'Сегодня выходной', off: true }
  return { label: `Сегодня ${today.open} — ${today.close}`, off: false }
})
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
      <!-- Cover banner -->
      <div class="cover-banner" v-if="restaurant.cover_image">
        <img :src="imageUrl(restaurant.cover_image)" :alt="restaurant.name" />
      </div>

      <!-- Compact info card -->
      <div class="info-card" :class="{ 'info-card-overlap': restaurant.cover_image }">
        <div class="info-logo" v-if="restaurant.logo">
          <img :src="imageUrl(restaurant.logo)" :alt="restaurant.name" />
        </div>
        <div class="info-logo info-logo-placeholder" v-else>
          {{ restaurant.name.charAt(0) }}
        </div>
        <div class="info-body">
          <h1 class="info-name">{{ restaurant.name }}</h1>
          <p class="info-desc" v-if="restaurant.description">{{ restaurant.description }}</p>
          <div class="info-address" v-if="restaurant.address">
            <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0118 0z"/><circle cx="12" cy="10" r="3"/>
            </svg>
            {{ restaurant.address }}
          </div>
          <div class="info-hours" v-if="todaySchedule" :class="{ 'info-hours-off': todaySchedule.off }">
            <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l4 2"/></svg>
            {{ todaySchedule.label }}
          </div>
        </div>
        <a v-if="restaurant.phone" :href="'tel:' + restaurant.phone" class="info-contact">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
            <path d="M22 16.92v3a2 2 0 01-2.18 2 19.79 19.79 0 01-8.63-3.07 19.5 19.5 0 01-6-6 19.79 19.79 0 01-3.07-8.67A2 2 0 014.11 2h3a2 2 0 012 1.72c.127.96.361 1.903.7 2.81a2 2 0 01-.45 2.11L8.09 9.91a16 16 0 006 6l1.27-1.27a2 2 0 012.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0122 16.92z"/>
          </svg>
          Связаться
        </a>
      </div>

      <!-- Closed / Closing soon -->
      <div class="status-banner closed" v-if="!restaurant.is_open">
        Сейчас закрыто. Заказы не принимаются.
      </div>
      <div class="status-banner closing" v-else-if="restaurant.closing_soon">
        Закрываемся в {{ restaurant.close_time }}. Заказы больше не принимаются.
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

      <!-- Category tabs (scroll-to-section) -->
      <div class="tabs-wrap" v-if="popularItems.length || combos.length || categories.length > 1">
        <div class="tabs">
          <button
            v-if="popularItems.length"
            class="tab"
            :class="{ active: activeCategory === 'popular' }"
            @click="scrollToCategory('popular')"
          >Популярное</button>
          <button
            v-if="combos.length"
            class="tab"
            :class="{ active: activeCategory === 'combos' }"
            @click="scrollToCategory('combos')"
          >Комбо</button>
          <button
            v-for="cat in categories"
            :key="cat.id"
            class="tab"
            :class="{ active: activeCategory === cat.id }"
            @click="scrollToCategory(cat.id)"
          >{{ cat.name }}</button>
        </div>
      </div>

      <!-- Menu (popular + combos + categories) -->
      <div class="menu-content" v-if="popularItems.length || combos.length || categories.length">
        <!-- Популярное (горизонтальный скролл, карточки как у основного меню) -->
        <div id="cat-popular" class="category-section popular-section" v-if="popularItems.length">
          <div class="category-name">Популярное</div>
          <div class="popular-scroll">
            <div class="item-card" v-for="item in popularItems" :key="item.id" @click="openDetail(item)">
              <div class="item-img" v-if="item.image">
                <img :src="imageUrl(item.image)" :alt="item.name" loading="lazy" />
              </div>
              <div class="item-img item-img-placeholder" v-else>
                <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="var(--text-muted)" stroke-width="1.5" stroke-linecap="round">
                  <path d="M18 8h1a4 4 0 010 8h-1M2 8h16v9a4 4 0 01-4 4H6a4 4 0 01-4-4V8zM6 1v3M10 1v3M14 1v3"/>
                </svg>
              </div>
              <div class="item-body">
                <div class="item-name">{{ item.name }}</div>
                <div class="item-desc" v-if="item.description">{{ item.description }}</div>
                <div class="item-footer">
                  <span class="item-price">{{ formatPrice(item.price) }}</span>
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
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"><path d="M12 5v14M5 12h14"/></svg>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div id="cat-combos" class="category-section combos-section" v-if="combos.length">
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

        <!-- Categories with items -->
        <div v-for="cat in categories" :key="cat.id" :id="'cat-' + cat.id" class="category-section">
          <div class="category-name">{{ cat.name }}</div>
          <div class="items-grid" v-if="cat.items?.length">
            <div class="item-card" :class="{ 'item-stopped': !item.available }" v-for="item in cat.items" :key="item.id" @click="item.available && openDetail(item)">
              <div class="item-img" v-if="item.image">
                <img :src="imageUrl(item.image)" :alt="item.name" loading="lazy" />
              </div>
              <div class="item-img item-img-placeholder" v-else>
                <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="var(--text-muted)" stroke-width="1.5" stroke-linecap="round">
                  <path d="M18 8h1a4 4 0 010 8h-1M2 8h16v9a4 4 0 01-4 4H6a4 4 0 01-4-4V8zM6 1v3M10 1v3M14 1v3"/>
                </svg>
              </div>
              <div class="stopped-badge" v-if="!item.available">Нет в наличии</div>
              <div class="item-body">
                <div class="item-name">{{ item.name }}</div>
                <div class="item-desc" v-if="item.description">{{ item.description }}</div>
                <div class="item-footer">
                  <span class="item-price">{{ formatPrice(item.price) }}</span>
                  <template v-if="item.available">
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
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"><path d="M12 5v14M5 12h14"/></svg>
                    </button>
                  </template>
                </div>
              </div>
            </div>
          </div>
          <div class="empty-cat" v-else><p>Пока пусто</p></div>
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
              <div class="detail-comment">
                <input v-model="detailComment" class="detail-comment-input" placeholder="Комментарий к блюду..." @change="getQty(selectedItem.id) && updateComment(selectedItem.id, detailComment)" />
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
                <button v-else class="detail-add-btn" @click="addWithComment()">Добавить в корзину</button>
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

/* ===== Cover banner ===== */
.cover-banner {
  width: 100%;
  height: 180px;
  overflow: hidden;
  background: var(--bg-secondary);
}

.cover-banner img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

/* ===== Compact info card ===== */
.info-card {
  display: flex;
  align-items: center;
  gap: 14px;
  margin: 16px;
  padding: 14px 16px;
  background: var(--bg);
  border: 1px solid var(--border);
  border-radius: 16px;
  box-shadow: 0 4px 16px rgba(0,0,0,0.06);
}

.info-card-overlap {
  margin-top: -40px;
  position: relative;
  z-index: 2;
}

.info-logo {
  width: 56px;
  height: 56px;
  border-radius: 50%;
  overflow: hidden;
  flex-shrink: 0;
  background: var(--bg-secondary);
  border: 2px solid var(--bg);
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.info-logo img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.info-logo-placeholder {
  background: var(--primary);
  color: var(--primary-foreground);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 22px;
  font-weight: 700;
}

.info-body {
  flex: 1;
  min-width: 0;
}

.info-name {
  font-size: 17px;
  font-weight: 700;
  color: var(--text);
  line-height: 1.2;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.info-desc {
  font-size: 12px;
  color: var(--text-secondary);
  margin-top: 2px;
  line-height: 1.3;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.info-address {
  display: flex;
  align-items: flex-start;
  gap: 5px;
  font-size: 12px;
  color: var(--text-secondary);
  line-height: 1.4;
  margin-top: 6px;
}

.info-address svg {
  flex-shrink: 0;
  margin-top: 2px;
}

.info-hours {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  font-size: 12px;
  color: var(--text-secondary);
  margin-top: 6px;
}

.info-hours-off {
  color: #DC2626;
}

.info-contact {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  padding: 8px 12px;
  background: var(--primary);
  color: var(--primary-foreground);
  border-radius: 10px;
  font-size: 13px;
  font-weight: 600;
  flex-shrink: 0;
  text-decoration: none;
}

.info-contact:active {
  opacity: 0.85;
}

/* ===== Status banner ===== */
.status-banner {
  margin: 12px 16px 0;
  padding: 12px 16px;
  border-radius: var(--radius);
  font-size: 14px;
  font-weight: 600;
  text-align: center;
}

.status-banner.closed {
  background: #FEE2E2;
  color: #DC2626;
}

.status-banner.closing {
  background: #FEF3C7;
  color: #92400E;
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

/* ===== Popular horizontal scroll (inside menu-content) ===== */
.popular-section {
  /* break out of .menu-content's 16px padding so the scroll reaches screen edges */
  margin-left: -16px;
  margin-right: -16px;
}

.popular-section .category-name {
  padding-left: 16px;
  padding-right: 16px;
}

.popular-scroll {
  display: flex;
  gap: 12px;
  overflow-x: auto;
  padding: 2px 16px 12px;
  scroll-snap-type: x proximity;
  -webkit-overflow-scrolling: touch;
  scrollbar-width: thin;
  scrollbar-color: var(--border) transparent;
}

.popular-scroll::-webkit-scrollbar {
  height: 6px;
}

.popular-scroll::-webkit-scrollbar-track {
  background: transparent;
  margin: 0 16px;
}

.popular-scroll::-webkit-scrollbar-thumb {
  background: var(--border);
  border-radius: 3px;
}

.popular-scroll::-webkit-scrollbar-thumb:hover {
  background: var(--text-muted);
}

/* Items inside horizontal scroll use the same card style as the grid but fixed width */
.popular-scroll .item-card {
  flex: 0 0 170px;
  scroll-snap-align: start;
}

/* ===== Category tabs (sticky) ===== */
.tabs-wrap {
  position: sticky;
  top: 0;
  z-index: 19;
  background: rgba(255,255,255,0.92);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  border-bottom: 1px solid var(--border);
  margin-top: 16px;
}

.tabs {
  display: flex;
  gap: 8px;
  padding: 12px 16px;
  overflow-x: auto;
  -webkit-overflow-scrolling: touch;
  scrollbar-width: none;
}

.tabs::-webkit-scrollbar { display: none; }

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

/* ===== Menu content ===== */
.menu-content {
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

/* Item grid */
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
  position: relative;
}

.item-card:active { transform: scale(0.97); }

.item-stopped {
  opacity: 0.5;
  pointer-events: none;
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

.item-body { padding: 12px; }

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

.add-btn:active { opacity: 0.8; }

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

.qty-btn-sm:active { background: var(--border); }

.qty-val {
  min-width: 20px;
  text-align: center;
  font-size: 14px;
  font-weight: 700;
  font-variant-numeric: tabular-nums;
}

.empty-cat {
  text-align: center;
  padding: 24px;
  color: var(--text-muted);
  font-size: 14px;
}

/* Legacy: old .actions/.btn-primary styles kept in case they are referenced elsewhere */
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

.schedule-card {
  background: var(--bg-secondary);
  border-radius: var(--radius);
  padding: 12px 16px;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.hours-row {
  display: flex;
  align-items: center;
  font-size: 14px;
  padding: 6px 8px;
  border-radius: 8px;
  gap: 8px;
}

.hours-row.hours-today {
  background: var(--primary-light, rgba(59,130,246,0.08));
  font-weight: 600;
}

.hours-row.hours-off {
  opacity: 0.5;
}

.sday {
  width: 26px;
  font-weight: 600;
  flex-shrink: 0;
}

.sdots {
  flex: 1;
  border-bottom: 1px dotted var(--border);
  margin: 0 4px;
  min-width: 20px;
  align-self: flex-end;
  margin-bottom: 4px;
}

.stime {
  white-space: nowrap;
}

.stime-off {
  color: var(--text-muted);
  font-style: italic;
}

.sbreak {
  font-size: 12px;
  color: var(--text-muted);
  white-space: nowrap;
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
