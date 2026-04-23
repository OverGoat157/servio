<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { restaurants as api } from '../api/client'
import { user } from '../stores/auth'

const router = useRouter()
const { t } = useI18n()
const list = ref([])
const loading = ref(true)

const limitReached = computed(() => {
  if (!user.value) return false
  if (user.value.role === 'admin') return false
  return list.value.length >= (user.value.max_restaurants ?? 0)
})
const showCreate = ref(false)
const form = ref({ name: '', slug: '', description: '', phone: '' })
const creating = ref(false)
const error = ref('')

onMounted(async () => {
  await loadList()
})

async function loadList() {
  loading.value = true
  try {
    list.value = await api.list() || []
  } catch { /* empty */ }
  loading.value = false
}

async function createRestaurant() {
  error.value = ''
  creating.value = true
  try {
    const rest = await api.create(form.value)
    list.value.unshift(rest)
    showCreate.value = false
    form.value = { name: '', slug: '', description: '', phone: '' }
    router.push({ name: 'menu', params: { id: rest.id } })
  } catch (e) {
    error.value = e.message
  }
  creating.value = false
}

function generateSlug() {
  form.value.slug = form.value.name
    .toLowerCase()
    .replace(/[а-яё]/g, (c) => {
      const map = { а:'a',б:'b',в:'v',г:'g',д:'d',е:'e',ё:'yo',ж:'zh',з:'z',и:'i',й:'y',к:'k',л:'l',м:'m',н:'n',о:'o',п:'p',р:'r',с:'s',т:'t',у:'u',ф:'f',х:'h',ц:'ts',ч:'ch',ш:'sh',щ:'sch',ъ:'',ы:'y',ь:'',э:'e',ю:'yu',я:'ya' }
      return map[c] || c
    })
    .replace(/[^a-z0-9]+/g, '-')
    .replace(/^-|-$/g, '')
}

async function deleteRestaurant(id) {
  if (!confirm(t('dashboard.confirmDelete'))) return
  await api.delete(id)
  list.value = list.value.filter(r => r.id !== id)
}
</script>

<template>
  <div class="page">
    <div class="page-header">
      <h1>{{ $t('dashboard.title') }}</h1>
      <button class="btn btn-primary" :disabled="limitReached" @click="!limitReached && (showCreate = true)">
        <svg v-if="!limitReached" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"><path d="M12 5v14M5 12h14"/></svg>
        {{ limitReached ? $t('dashboard.limitReached') : $t('dashboard.addRestaurant') }}
      </button>
    </div>

    <!-- Create modal -->
    <div class="modal-overlay" v-if="showCreate" @click.self="showCreate = false">
      <div class="modal card">
        <h2>{{ $t('dashboard.newRestaurant') }}</h2>
        <form @submit.prevent="createRestaurant">
          <div class="field">
            <label class="label">{{ $t('common.title') }}</label>
            <input v-model="form.name" class="input" :placeholder="$t('dashboard.namePlaceholder')" required @input="generateSlug" />
          </div>
          <div class="field">
            <label class="label">{{ $t('dashboard.slugLabel') }}</label>
            <input v-model="form.slug" class="input" :placeholder="$t('dashboard.slugPlaceholder')" required pattern="[a-z0-9\-]+" />
            <div class="hint">{{ $t('dashboard.slugHint', { slug: form.slug || '...' }) }}</div>
          </div>
          <div class="field">
            <label class="label">{{ $t('common.description') }}</label>
            <input v-model="form.description" class="input" :placeholder="$t('dashboard.descPlaceholder')" />
          </div>
          <div class="field">
            <label class="label">{{ $t('common.phone') }}</label>
            <input v-model="form.phone" class="input" :placeholder="$t('dashboard.phonePlaceholder')" />
          </div>
          <div class="error-msg" v-if="error">{{ error }}</div>
          <div class="modal-actions">
            <button type="button" class="btn btn-outline" @click="showCreate = false">{{ $t('common.cancel') }}</button>
            <button type="submit" class="btn btn-primary" :disabled="creating">
              {{ creating ? $t('common.creating') : $t('common.create') }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Loading -->
    <div class="loading" v-if="loading">{{ $t('common.loading') }}</div>

    <!-- Empty -->
    <div class="empty card" v-else-if="list.length === 0">
      <div class="empty-icon">
        <svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="var(--text-secondary)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
          <path d="M3 3h18v18H3z"/><path d="M12 8v8M8 12h8"/>
        </svg>
      </div>
      <p>{{ $t('dashboard.emptyList') }}</p>
      <button class="btn btn-primary" @click="showCreate = true">{{ $t('dashboard.createFirst') }}</button>
    </div>

    <!-- List -->
    <div class="rest-grid" v-else>
      <div class="rest-card card" v-for="r in list" :key="r.id">
        <div class="rest-top">
          <div class="rest-avatar">{{ r.name.charAt(0) }}</div>
          <div class="rest-info">
            <h3>{{ r.name }}</h3>
            <span class="rest-slug">/{{ r.slug }}</span>
          </div>
        </div>
        <p class="rest-desc" v-if="r.description">{{ r.description }}</p>
        <div class="rest-actions">
          <a :href="'https://menu.ab-team.ru/' + r.slug" target="_blank" class="btn btn-accent btn-sm">
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M18 13v6a2 2 0 01-2 2H5a2 2 0 01-2-2V8a2 2 0 012-2h6M15 3h6v6M10 14L21 3"/></svg>
            {{ $t('dashboard.actions.site') }}
          </a>
          <router-link :to="{ name: 'menu', params: { id: r.id } }" class="btn btn-primary btn-sm">{{ $t('dashboard.actions.menu') }}</router-link>
          <router-link :to="{ name: 'orders', params: { id: r.id } }" class="btn btn-outline btn-sm">{{ $t('dashboard.actions.orders') }}</router-link>
          <router-link :to="{ name: 'combos', params: { id: r.id } }" class="btn btn-outline btn-sm">{{ $t('dashboard.actions.combos') }}</router-link>
          <router-link :to="{ name: 'messengers', params: { id: r.id } }" class="btn btn-outline btn-sm">{{ $t('dashboard.actions.messengers') }}</router-link>
          <router-link :to="{ name: 'analytics', params: { id: r.id } }" class="btn btn-outline btn-sm">{{ $t('dashboard.actions.analytics') }}</router-link>
          <router-link :to="{ name: 'restaurant', params: { id: r.id } }" class="btn btn-outline btn-sm">{{ $t('dashboard.actions.settings') }}</router-link>
          <button class="btn btn-sm delete-btn" @click="deleteRestaurant(r.id)">
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M3 6h18M19 6v14a2 2 0 01-2 2H7a2 2 0 01-2-2V6m3 0V4a2 2 0 012-2h4a2 2 0 012 2v2"/></svg>
          </button>
        </div>
      </div>
    </div>
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
  align-items: center;
  margin-bottom: 24px;
}

.page-header h1 {
  font-size: 24px;
  font-weight: 700;
}

/* Modal */
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.4);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 200;
  padding: 24px;
}

.modal {
  width: 100%;
  max-width: 440px;
}

.modal h2 {
  font-size: 18px;
  font-weight: 700;
  margin-bottom: 20px;
}

.field {
  margin-bottom: 14px;
}

.hint {
  font-size: 12px;
  color: var(--text-secondary);
  margin-top: 4px;
}

.error-msg {
  padding: 10px;
  background: #FEE2E2;
  color: var(--danger);
  border-radius: var(--radius);
  font-size: 13px;
  margin-bottom: 14px;
  text-align: center;
}

.modal-actions {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
  margin-top: 20px;
}

/* Loading / Empty */
.loading {
  text-align: center;
  padding: 60px 0;
  color: var(--text-secondary);
}

.empty {
  text-align: center;
  padding: 60px 24px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
}

.empty p {
  color: var(--text-secondary);
  font-size: 15px;
}

/* Grid */
.rest-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 16px;
}

.rest-card {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.rest-top {
  display: flex;
  align-items: center;
  gap: 12px;
}

.rest-avatar {
  width: 44px;
  height: 44px;
  background: var(--primary-light);
  color: var(--primary);
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  font-weight: 700;
  flex-shrink: 0;
}

.rest-info h3 {
  font-size: 16px;
  font-weight: 700;
}

.rest-slug {
  font-size: 12px;
  color: var(--text-secondary);
}

.rest-desc {
  font-size: 13px;
  color: var(--text-secondary);
  line-height: 1.4;
}

.rest-actions {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.btn-accent {
  background: #EFF6FF;
  color: #2167C7;
  border: 1px solid #BFDBFE;
  display: inline-flex;
  align-items: center;
  gap: 5px;
}

.btn-accent:hover {
  background: #DBEAFE;
}

.delete-btn {
  color: var(--danger);
  margin-left: auto;
}

.delete-btn:hover {
  background: #FEE2E2;
  border-radius: var(--radius);
}
</style>
