<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { restaurants as restApi, orders as ordersApi } from '../api/client'

const route = useRoute()
const router = useRouter()
const { t, locale } = useI18n()
const id = route.params.id

const rest = ref(null)
const orderList = ref([])
const loading = ref(true)

const STATUS_KEYS = ['new', 'confirmed', 'preparing', 'ready', 'delivered', 'cancelled']
const statusLabels = computed(() => {
  const map = {}
  for (const key of STATUS_KEYS) map[key] = t(`orderStatus.${key}`)
  return map
})

const statusColors = {
  new: '#2167C7',
  confirmed: '#7C3AED',
  preparing: '#F59E0B',
  ready: '#22C55E',
  delivered: '#6B7280',
  cancelled: '#DC2626',
}

onMounted(async () => {
  try {
    rest.value = await restApi.get(id)
    orderList.value = await ordersApi.list(id) || []
  } catch {
    router.push({ name: 'dashboard' })
  }
  loading.value = false
})

function formatPrice(kopecks) {
  const tag = locale.value === 'en' ? 'en-US' : 'ru-RU'
  return (kopecks / 100).toLocaleString(tag) + ' ₽'
}

function formatDate(dateStr) {
  const tag = locale.value === 'en' ? 'en-US' : 'ru-RU'
  return new Date(dateStr).toLocaleString(tag, {
    day: 'numeric',
    month: 'short',
    hour: '2-digit',
    minute: '2-digit',
  })
}

function parseItems(itemsJson) {
  try {
    return JSON.parse(itemsJson)
  } catch {
    return []
  }
}

async function changeStatus(orderId, status) {
  await ordersApi.updateStatus(orderId, status)
  const order = orderList.value.find(o => o.id === orderId)
  if (order) order.status = status
}
</script>

<template>
  <div class="page">
    <div class="page-header">
      <div>
        <button class="back-link" @click="router.push({ name: 'dashboard' })">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M19 12H5M12 19l-7-7 7-7"/></svg>
          {{ $t('common.back') }}
        </button>
        <h1 v-if="rest">{{ $t('orders.title', { name: rest.name }) }}</h1>
      </div>
    </div>

    <div v-if="loading" class="loading">{{ $t('common.loading') }}</div>

    <div class="empty card" v-else-if="orderList.length === 0">
      <p>{{ $t('orders.empty') }}</p>
    </div>

    <div class="orders" v-else>
      <div class="order-card card" v-for="order in orderList" :key="order.id">
        <div class="order-header">
          <div>
            <span class="order-id">#{{ order.id }}</span>
            <span class="order-date">{{ formatDate(order.created_at) }}</span>
          </div>
          <span class="status-badge" :style="{ background: statusColors[order.status] + '20', color: statusColors[order.status] }">
            {{ statusLabels[order.status] || order.status }}
          </span>
        </div>

        <div class="order-items">
          <div class="oi-row" v-for="item in parseItems(order.items)" :key="item.menu_item_id">
            <span>{{ item.name }} x{{ item.quantity }}</span>
            <span class="oi-price">{{ formatPrice(item.price * item.quantity) }}</span>
          </div>
        </div>

        <div class="order-total">
          <span>{{ $t('orders.total') }}</span>
          <strong>{{ formatPrice(order.total) }}</strong>
        </div>

        <div class="order-meta">
          <span v-if="order.customer_name">{{ order.customer_name }}</span>
          <span v-if="order.customer_phone">{{ order.customer_phone }}</span>
          <span class="messenger-tag">{{ order.messenger }}</span>
        </div>

        <div class="order-actions">
          <select
            :value="order.status"
            @change="changeStatus(order.id, ($event.target).value)"
            class="input status-select"
          >
            <option v-for="(label, key) in statusLabels" :key="key" :value="key">{{ label }}</option>
          </select>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.page {
  max-width: 800px;
  margin: 0 auto;
  padding: 32px 24px;
}

.page-header {
  margin-bottom: 24px;
}

.back-link {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 14px;
  color: var(--text-secondary);
  margin-bottom: 8px;
}

.page-header h1 {
  font-size: 22px;
  font-weight: 700;
}

.loading {
  text-align: center;
  padding: 60px 0;
  color: var(--text-secondary);
}

.empty {
  text-align: center;
  padding: 48px 24px;
  color: var(--text-secondary);
}

.orders {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.order-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.order-id {
  font-weight: 700;
  font-size: 15px;
  margin-right: 10px;
}

.order-date {
  font-size: 13px;
  color: var(--text-secondary);
}

.status-badge {
  padding: 4px 12px;
  border-radius: 100px;
  font-size: 12px;
  font-weight: 600;
}

.order-items {
  margin-bottom: 12px;
}

.oi-row {
  display: flex;
  justify-content: space-between;
  padding: 6px 0;
  font-size: 14px;
  border-bottom: 1px solid var(--border);
}

.oi-price {
  font-weight: 600;
}

.order-total {
  display: flex;
  justify-content: space-between;
  padding: 10px 0;
  font-size: 15px;
  font-weight: 700;
}

.order-meta {
  display: flex;
  gap: 12px;
  font-size: 13px;
  color: var(--text-secondary);
  padding: 8px 0;
}

.messenger-tag {
  background: var(--primary-light);
  color: var(--primary);
  padding: 2px 10px;
  border-radius: 100px;
  font-weight: 600;
  font-size: 12px;
}

.order-actions {
  padding-top: 8px;
}

.status-select {
  max-width: 200px;
  padding: 8px 12px;
}
</style>
