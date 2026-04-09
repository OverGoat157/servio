<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { restaurants as restApi, analytics as analyticsApi } from '../api/client'

const route = useRoute()
const router = useRouter()
const id = route.params.id

const rest = ref(null)
const data = ref(null)
const loading = ref(true)
const period = ref(30)

async function load() {
  loading.value = true
  try {
    if (!rest.value) rest.value = await restApi.get(id)
    data.value = await analyticsApi.summary(id, period.value)
  } catch {
    router.push({ name: 'dashboard' })
  }
  loading.value = false
}

onMounted(load)

function setPeriod(days) {
  period.value = days
  load()
}

function formatMoney(kopecks) {
  return Math.floor(kopecks / 100).toLocaleString('ru-RU') + ' \u20BD'
}

const maxBarValue = computed(() => {
  if (!data.value?.day_stats) return 1
  return Math.max(...data.value.day_stats.map(d => d.views), 1)
})

const maxOrderBar = computed(() => {
  if (!data.value?.day_stats) return 1
  return Math.max(...data.value.day_stats.map(d => d.orders), 1)
})

function formatDate(dateStr) {
  const d = new Date(dateStr)
  return d.toLocaleDateString('ru-RU', { day: 'numeric', month: 'short' })
}

const conversionRate = computed(() => {
  if (!data.value || !data.value.total_views) return '0'
  return ((data.value.total_orders / data.value.total_views) * 100).toFixed(1)
})
</script>

<template>
  <div class="page">
    <div class="page-header">
      <div>
        <button class="back-link" @click="router.push({ name: 'dashboard' })">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M19 12H5M12 19l-7-7 7-7"/></svg>
          Назад
        </button>
        <h1 v-if="rest">Аналитика: {{ rest.name }}</h1>
      </div>
      <div class="period-tabs">
        <button :class="{ active: period === 7 }" @click="setPeriod(7)">7 дней</button>
        <button :class="{ active: period === 30 }" @click="setPeriod(30)">30 дней</button>
        <button :class="{ active: period === 90 }" @click="setPeriod(90)">90 дней</button>
      </div>
    </div>

    <div v-if="loading" class="loading">Загрузка...</div>

    <template v-else-if="data">
      <!-- Summary cards -->
      <div class="stats-grid">
        <div class="stat-card">
          <div class="stat-icon views-icon">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></svg>
          </div>
          <div class="stat-value">{{ data.total_views.toLocaleString('ru-RU') }}</div>
          <div class="stat-label">Просмотров</div>
        </div>
        <div class="stat-card">
          <div class="stat-icon orders-icon">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M6 2L3 6v14a2 2 0 002 2h14a2 2 0 002-2V6l-3-4z"/><path d="M3 6h18M16 10a4 4 0 01-8 0"/></svg>
          </div>
          <div class="stat-value">{{ data.total_orders }}</div>
          <div class="stat-label">Заказов</div>
        </div>
        <div class="stat-card">
          <div class="stat-icon revenue-icon">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M12 1v22M17 5H9.5a3.5 3.5 0 000 7h5a3.5 3.5 0 010 7H6"/></svg>
          </div>
          <div class="stat-value">{{ formatMoney(data.total_revenue) }}</div>
          <div class="stat-label">Выручка</div>
        </div>
        <div class="stat-card">
          <div class="stat-icon clients-icon">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M23 21v-2a4 4 0 00-3-3.87M16 3.13a4 4 0 010 7.75"/></svg>
          </div>
          <div class="stat-value">{{ data.unique_clients }}</div>
          <div class="stat-label">Уникальных клиентов</div>
        </div>
        <div class="stat-card">
          <div class="stat-icon avg-icon">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><rect x="2" y="7" width="20" height="14" rx="2"/><path d="M16 7V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v3"/></svg>
          </div>
          <div class="stat-value">{{ data.avg_check ? formatMoney(data.avg_check) : '—' }}</div>
          <div class="stat-label">Средний чек</div>
        </div>
        <div class="stat-card">
          <div class="stat-icon conv-icon">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><polyline points="22 12 18 12 15 21 9 3 6 12 2 12"/></svg>
          </div>
          <div class="stat-value">{{ conversionRate }}%</div>
          <div class="stat-label">Конверсия</div>
        </div>
      </div>

      <!-- Views chart -->
      <div class="card chart-card">
        <h3>Просмотры по дням</h3>
        <div class="chart">
          <div class="chart-bar-wrap" v-for="d in data.day_stats" :key="d.date">
            <div class="chart-bar-value" v-if="d.views">{{ d.views }}</div>
            <div class="chart-bar views-bar" :style="{ height: (d.views / maxBarValue * 120) + 'px' }"></div>
            <div class="chart-bar-label">{{ formatDate(d.date) }}</div>
          </div>
        </div>
      </div>

      <!-- Orders chart -->
      <div class="card chart-card">
        <h3>Заказы по дням</h3>
        <div class="chart">
          <div class="chart-bar-wrap" v-for="d in data.day_stats" :key="d.date">
            <div class="chart-bar-value" v-if="d.orders">{{ d.orders }}</div>
            <div class="chart-bar orders-bar" :style="{ height: (d.orders / maxOrderBar * 120) + 'px' }"></div>
            <div class="chart-bar-label">{{ formatDate(d.date) }}</div>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<style scoped>
.page {
  max-width: 900px;
  margin: 0 auto;
  padding: 32px 24px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 24px;
  gap: 16px;
  flex-wrap: wrap;
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

.period-tabs {
  display: flex;
  gap: 4px;
  background: var(--bg-secondary, #f3f4f6);
  border-radius: 10px;
  padding: 3px;
}

.period-tabs button {
  padding: 7px 14px;
  font-size: 13px;
  font-weight: 500;
  border-radius: 8px;
  color: var(--text-secondary);
  transition: all 0.2s;
}

.period-tabs button.active {
  background: #fff;
  color: var(--text);
  box-shadow: 0 1px 3px rgba(0,0,0,0.1);
}

.loading {
  text-align: center;
  padding: 60px 0;
  color: var(--text-secondary);
}

/* Stats grid */
.stats-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 12px;
  margin-bottom: 20px;
}

.stat-card {
  background: var(--card-bg, #fff);
  border: 1px solid var(--border);
  border-radius: var(--radius);
  padding: 18px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.stat-icon {
  width: 40px;
  height: 40px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.views-icon { background: #EFF6FF; color: #3B82F6; }
.orders-icon { background: #F0FDF4; color: #22C55E; }
.revenue-icon { background: #FFFBEB; color: #F59E0B; }
.clients-icon { background: #FDF2F8; color: #EC4899; }
.avg-icon { background: #F5F3FF; color: #8B5CF6; }
.conv-icon { background: #ECFDF5; color: #10B981; }

.stat-value {
  font-size: 24px;
  font-weight: 700;
  font-variant-numeric: tabular-nums;
}

.stat-label {
  font-size: 13px;
  color: var(--text-secondary);
}

/* Charts */
.chart-card {
  padding: 20px;
  margin-bottom: 16px;
}

.chart-card h3 {
  font-size: 15px;
  font-weight: 700;
  margin-bottom: 16px;
}

.chart {
  display: flex;
  align-items: flex-end;
  gap: 3px;
  overflow-x: auto;
  padding-bottom: 4px;
  min-height: 160px;
}

.chart-bar-wrap {
  flex: 1;
  min-width: 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.chart-bar {
  width: 100%;
  max-width: 32px;
  min-height: 2px;
  border-radius: 4px 4px 0 0;
  transition: height 0.3s ease;
}

.views-bar {
  background: #3B82F6;
}

.orders-bar {
  background: #22C55E;
}

.chart-bar-value {
  font-size: 10px;
  font-weight: 600;
  color: var(--text-secondary);
}

.chart-bar-label {
  font-size: 9px;
  color: var(--text-muted, #9ca3af);
  white-space: nowrap;
  writing-mode: vertical-lr;
  transform: rotate(180deg);
  max-height: 40px;
}

@media (max-width: 640px) {
  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  .chart-bar-label {
    display: none;
  }
}
</style>
