<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { admin } from '../api/client'

const router = useRouter()
const { locale } = useI18n()
const data = ref(null)
const loading = ref(true)
const period = ref(30)

async function load() {
  loading.value = true
  try {
    data.value = await admin.analytics(period.value)
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
  return Math.floor(kopecks / 100).toLocaleString(locale.value === 'en' ? 'en-US' : 'ru-RU') + ' ₽'
}

function formatDate(dateStr) {
  const d = new Date(dateStr)
  const tag = locale.value === 'en' ? 'en-US' : 'ru-RU'
  return d.toLocaleDateString(tag, { day: 'numeric', month: 'short' })
}

const maxViews = computed(() => {
  if (!data.value?.day_stats) return 1
  return Math.max(...data.value.day_stats.map(d => d.views), 1)
})

const maxOrders = computed(() => {
  if (!data.value?.day_stats) return 1
  return Math.max(...data.value.day_stats.map(d => d.orders), 1)
})
</script>

<template>
  <div class="page">
    <div class="page-header">
      <div>
        <button class="back-link" @click="router.push({ name: 'dashboard' })">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M19 12H5M12 19l-7-7 7-7"/></svg>
          {{ $t('common.back') }}
        </button>
        <h1>{{ $t('analytics.platformTitle') }}</h1>
      </div>
      <div class="period-tabs">
        <button :class="{ active: period === 7 }" @click="setPeriod(7)">{{ $t('analytics.period7') }}</button>
        <button :class="{ active: period === 30 }" @click="setPeriod(30)">{{ $t('analytics.period30') }}</button>
        <button :class="{ active: period === 90 }" @click="setPeriod(90)">{{ $t('analytics.period90') }}</button>
      </div>
    </div>

    <div v-if="loading" class="loading">{{ $t('common.loading') }}</div>

    <template v-else-if="data">
      <!-- Summary cards -->
      <div class="stats-grid">
        <div class="stat-card">
          <div class="stat-icon users-icon">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M23 21v-2a4 4 0 00-3-3.87M16 3.13a4 4 0 010 7.75"/></svg>
          </div>
          <div class="stat-value">{{ data.total_users }}</div>
          <div class="stat-label">{{ $t('analytics.totalUsers') }}</div>
        </div>
        <div class="stat-card">
          <div class="stat-icon rest-icon">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M3 3h18v18H3z"/><path d="M12 8v8M8 12h8"/></svg>
          </div>
          <div class="stat-value">{{ data.total_restaurants }}</div>
          <div class="stat-label">{{ $t('analytics.totalRestaurants') }}</div>
        </div>
        <div class="stat-card">
          <div class="stat-icon views-icon">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></svg>
          </div>
          <div class="stat-value">{{ data.total_views.toLocaleString(locale === 'en' ? 'en-US' : 'ru-RU') }}</div>
          <div class="stat-label">{{ $t('analytics.menuViews') }}</div>
        </div>
        <div class="stat-card">
          <div class="stat-icon orders-icon">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M6 2L3 6v14a2 2 0 002 2h14a2 2 0 002-2V6l-3-4z"/><path d="M3 6h18M16 10a4 4 0 01-8 0"/></svg>
          </div>
          <div class="stat-value">{{ data.total_orders }}</div>
          <div class="stat-label">{{ $t('analytics.orders') }}</div>
        </div>
        <div class="stat-card">
          <div class="stat-icon revenue-icon">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M12 1v22M17 5H9.5a3.5 3.5 0 000 7h5a3.5 3.5 0 010 7H6"/></svg>
          </div>
          <div class="stat-value">{{ formatMoney(data.total_revenue) }}</div>
          <div class="stat-label">{{ $t('analytics.totalRevenue') }}</div>
        </div>
        <div class="stat-card">
          <div class="stat-icon clients-icon">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>
          </div>
          <div class="stat-value">{{ data.unique_clients }}</div>
          <div class="stat-label">{{ $t('analytics.uniqueClients') }}</div>
        </div>
      </div>

      <!-- Charts -->
      <div class="card chart-card">
        <h3>{{ $t('analytics.menuViewsByDay') }}</h3>
        <div class="chart">
          <div class="chart-bar-wrap" v-for="d in data.day_stats" :key="'v'+d.date">
            <div class="chart-bar-value" v-if="d.views">{{ d.views }}</div>
            <div class="chart-bar views-bar" :style="{ height: (d.views / maxViews * 120) + 'px' }"></div>
            <div class="chart-bar-label">{{ formatDate(d.date) }}</div>
          </div>
        </div>
      </div>

      <div class="card chart-card">
        <h3>{{ $t('analytics.ordersByDay') }}</h3>
        <div class="chart">
          <div class="chart-bar-wrap" v-for="d in data.day_stats" :key="'o'+d.date">
            <div class="chart-bar-value" v-if="d.orders">{{ d.orders }}</div>
            <div class="chart-bar orders-bar" :style="{ height: (d.orders / maxOrders * 120) + 'px' }"></div>
            <div class="chart-bar-label">{{ formatDate(d.date) }}</div>
          </div>
        </div>
      </div>

      <!-- Restaurant breakdown -->
      <div class="card table-card">
        <h3>{{ $t('analytics.byRestaurant') }}</h3>
        <div class="table-wrap">
          <table class="data-table">
            <thead>
              <tr>
                <th>{{ $t('analytics.restaurantCol') }}</th>
                <th>{{ $t('analytics.ownerCol') }}</th>
                <th>{{ $t('analytics.views') }}</th>
                <th>{{ $t('analytics.orders') }}</th>
                <th>{{ $t('analytics.revenue') }}</th>
                <th>{{ $t('analytics.clientsCol') }}</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="r in data.restaurant_stats" :key="r.id">
                <td class="cell-name">{{ r.name }}</td>
                <td class="cell-muted">{{ r.owner }}</td>
                <td>{{ r.views }}</td>
                <td>{{ r.orders }}</td>
                <td>{{ formatMoney(r.revenue) }}</td>
                <td>{{ r.clients }}</td>
              </tr>
              <tr v-if="!data.restaurant_stats?.length">
                <td colspan="6" class="cell-empty">{{ $t('common.empty') }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </template>
  </div>
</template>

<style scoped>
.page {
  max-width: 960px;
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

/* Stats */
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

.users-icon { background: #F5F3FF; color: #8B5CF6; }
.rest-icon { background: #FFF7ED; color: #F97316; }
.views-icon { background: #EFF6FF; color: #3B82F6; }
.orders-icon { background: #F0FDF4; color: #22C55E; }
.revenue-icon { background: #FFFBEB; color: #F59E0B; }
.clients-icon { background: #FDF2F8; color: #EC4899; }

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

.views-bar { background: #3B82F6; }
.orders-bar { background: #22C55E; }

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

/* Table */
.table-card {
  padding: 20px;
  margin-bottom: 16px;
}

.table-card h3 {
  font-size: 15px;
  font-weight: 700;
  margin-bottom: 16px;
}

.table-wrap {
  overflow-x: auto;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 14px;
}

.data-table th {
  text-align: left;
  padding: 10px 12px;
  font-size: 12px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: var(--text-muted, #9ca3af);
  border-bottom: 2px solid var(--border);
}

.data-table td {
  padding: 10px 12px;
  border-bottom: 1px solid var(--border);
  font-variant-numeric: tabular-nums;
}

.cell-name {
  font-weight: 600;
}

.cell-muted {
  color: var(--text-secondary);
  font-size: 13px;
}

.cell-empty {
  text-align: center;
  color: var(--text-muted);
  padding: 24px !important;
}

@media (max-width: 640px) {
  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  .chart-bar-label {
    display: none;
  }
  .data-table {
    font-size: 13px;
  }
  .data-table th, .data-table td {
    padding: 8px 8px;
  }
}
</style>
