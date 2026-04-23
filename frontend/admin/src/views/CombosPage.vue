<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { restaurants as restApi, combos as comboApi, menuItems as itemApi, categories as catApi, uploadFile, imageUrl } from '../api/client'

const route = useRoute()
const router = useRouter()
const { t, locale } = useI18n()
const id = route.params.id

const rest = ref(null)
const comboList = ref([])
const loading = ref(true)
const menuItemsList = ref([])

// Form
const showForm = ref(false)
const editingId = ref(null)
const form = ref({ name: '', name_en: '', description: '', description_en: '', price: '', image: '', sort_order: 0 })
const formItems = ref([])
const uploadingImage = ref(false)
const saving = ref(false)

onMounted(async () => {
  try {
    rest.value = await restApi.get(id)
    await loadCombos()
    await loadMenuItems()
  } catch {
    router.push({ name: 'dashboard' })
  }
  loading.value = false
})

async function loadCombos() {
  comboList.value = await comboApi.list(id) || []
}

async function loadMenuItems() {
  const cats = await catApi.list(id) || []
  const all = []
  for (const cat of cats) {
    const items = await itemApi.list(cat.id) || []
    for (const item of items) {
      all.push({ id: item.id, name: item.name, name_en: item.name_en || '', category: cat.name })
    }
  }
  menuItemsList.value = all
}

function openForm(combo = null) {
  if (combo) {
    editingId.value = combo.id
    form.value = {
      name: combo.name,
      name_en: combo.name_en || '',
      description: combo.description || '',
      description_en: combo.description_en || '',
      price: String(combo.price / 100),
      image: combo.image || '',
      sort_order: combo.sort_order,
    }
    formItems.value = (combo.items || []).map(ci => ({
      menu_item_id: ci.menu_item_id,
      name: ci.name,
      name_en: ci.name_en || '',
      quantity: ci.quantity,
    }))
  } else {
    editingId.value = null
    form.value = { name: '', name_en: '', description: '', description_en: '', price: '', image: '', sort_order: 0 }
    formItems.value = []
  }
  showForm.value = true
}

function addFormItem() {
  formItems.value.push({ menu_item_id: null, name: '', name_en: '', quantity: 1 })
}

function removeFormItem(idx) {
  formItems.value.splice(idx, 1)
}

function onSelectMenuItem(idx) {
  const item = formItems.value[idx]
  if (item.menu_item_id) {
    const mi = menuItemsList.value.find(m => m.id === item.menu_item_id)
    if (mi) {
      item.name = mi.name
      item.name_en = mi.name_en || ''
    }
  }
}

async function saveCombo() {
  if (saving.value) return
  saving.value = true
  try {
    const data = {
      name: form.value.name,
      name_en: form.value.name_en,
      description: form.value.description,
      description_en: form.value.description_en,
      price: Math.round(parseFloat(form.value.price) * 100),
      image: form.value.image,
      sort_order: form.value.sort_order,
      items: formItems.value.filter(i => i.name.trim()).map(i => ({
        menu_item_id: i.menu_item_id || null,
        name: i.name,
        name_en: i.name_en,
        quantity: i.quantity || 1,
      })),
    }
    if (editingId.value) {
      await comboApi.update(editingId.value, data)
    } else {
      await comboApi.create(id, data)
    }
    showForm.value = false
    await loadCombos()
  } finally {
    saving.value = false
  }
}

async function toggleAvailable(combo) {
  await comboApi.update(combo.id, { available: !combo.available })
  await loadCombos()
}

async function deleteCombo(comboId) {
  if (!confirm(t('combos.confirmDelete'))) return
  await comboApi.delete(comboId)
  await loadCombos()
}

async function handleImage(event) {
  const file = event.target.files?.[0]
  if (!file) return
  uploadingImage.value = true
  try {
    const { url } = await uploadFile(file)
    form.value.image = url
  } catch { /* empty */ }
  uploadingImage.value = false
}

function formatPrice(kopecks) {
  const tag = locale.value === 'en' ? 'en-US' : 'ru-RU'
  return (kopecks / 100).toLocaleString(tag) + ' ₽'
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
        <h1 v-if="rest">{{ $t('combos.title', { name: rest.name }) }}</h1>
      </div>
      <button class="btn btn-primary" @click="openForm()">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"><path d="M12 5v14M5 12h14"/></svg>
        {{ $t('combos.addCombo') }}
      </button>
    </div>

    <div v-if="loading" class="loading">{{ $t('common.loading') }}</div>

    <div class="empty card" v-else-if="comboList.length === 0">
      <p>{{ $t('combos.emptyHint') }}</p>
      <button class="btn btn-primary" @click="openForm()">{{ $t('combos.addCombo') }}</button>
    </div>

    <div class="combo-grid" v-else>
      <div class="combo-card card" :class="{ unavailable: !combo.available }" v-for="combo in comboList" :key="combo.id">
        <div class="combo-top">
          <div class="combo-thumb" v-if="combo.image">
            <img :src="imageUrl(combo.image)" :alt="combo.name" />
          </div>
          <div class="combo-info">
            <h3>{{ combo.name }}</h3>
            <span class="combo-price">{{ formatPrice(combo.price) }}</span>
            <span class="badge-off" v-if="!combo.available">{{ $t('combos.notAvailable') }}</span>
          </div>
        </div>
        <p class="combo-desc" v-if="combo.description">{{ combo.description }}</p>
        <div class="combo-items-list" v-if="combo.items?.length">
          <span class="combo-item-chip" v-for="ci in combo.items" :key="ci.id">
            {{ ci.name }}<template v-if="ci.quantity > 1"> ×{{ ci.quantity }}</template>
          </span>
        </div>
        <div class="combo-actions">
          <button class="icon-btn-sm" :class="combo.available ? 'available' : 'off'" :title="combo.available ? $t('combos.hide') : $t('combos.show')" @click="toggleAvailable(combo)">
            <svg v-if="combo.available" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"><path d="M20 6L9 17l-5-5"/></svg>
            <svg v-else width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"><path d="M18 6L6 18M6 6l12 12"/></svg>
          </button>
          <button class="icon-btn-sm" @click="openForm(combo)">
            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M11 4H4a2 2 0 00-2 2v14a2 2 0 002 2h14a2 2 0 002-2v-7"/><path d="M18.5 2.5a2.121 2.121 0 013 3L12 15l-4 1 1-4 9.5-9.5z"/></svg>
          </button>
          <button class="icon-btn-sm danger" @click="deleteCombo(combo.id)">
            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M18 6L6 18M6 6l12 12"/></svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Form modal -->
    <div class="modal-overlay" v-if="showForm" @click.self="showForm = false">
      <div class="modal card">
        <h2>{{ editingId ? $t('combos.editCombo') : $t('combos.newCombo') }}</h2>
        <form @submit.prevent="saveCombo">
          <div class="field">
            <label class="label">{{ $t('combos.imageLabel') }}</label>
            <div class="upload-area">
              <div class="img-preview" v-if="form.image">
                <img :src="imageUrl(form.image)" alt="" />
                <button type="button" class="img-remove" @click="form.image = ''">×</button>
              </div>
              <label class="upload-btn" v-else>
                <input type="file" accept="image/*" hidden @change="handleImage" />
                {{ uploadingImage ? $t('restaurant.uploading') : $t('combos.imageUpload') }}
              </label>
            </div>
          </div>
          <div class="field">
            <label class="label">{{ $t('common.title') }} <span class="lang-tag">RU</span></label>
            <input v-model="form.name" class="input" :placeholder="$t('combos.namePlaceholder')" required />
          </div>
          <div class="field">
            <label class="label">{{ $t('common.title') }} <span class="lang-tag">EN</span></label>
            <input v-model="form.name_en" class="input" placeholder="Business lunch" />
          </div>
          <div class="field">
            <label class="label">{{ $t('common.description') }} <span class="lang-tag">RU</span></label>
            <input v-model="form.description" class="input" :placeholder="$t('combos.descPlaceholder')" />
          </div>
          <div class="field">
            <label class="label">{{ $t('common.description') }} <span class="lang-tag">EN</span></label>
            <input v-model="form.description_en" class="input" placeholder="Soup + main + drink" />
          </div>
          <div class="row">
            <div class="field">
              <label class="label">{{ $t('menu.priceLabel') }}</label>
              <input v-model="form.price" type="number" step="0.01" min="0" class="input" required />
            </div>
            <div class="field">
              <label class="label">{{ $t('common.order') }}</label>
              <input v-model.number="form.sort_order" type="number" class="input" />
            </div>
          </div>

          <div class="divider"></div>
          <div class="section-label">{{ $t('combos.compositionSection') }}</div>

          <div class="combo-form-items">
            <div class="combo-form-row" v-for="(ci, idx) in formItems" :key="idx">
              <select v-model="ci.menu_item_id" class="input combo-select" @change="onSelectMenuItem(idx)">
                <option :value="null">{{ $t('combos.ownDish') }}</option>
                <option v-for="mi in menuItemsList" :key="mi.id" :value="mi.id">{{ mi.category }}: {{ mi.name }}</option>
              </select>
              <input v-model="ci.name" class="input combo-name" :placeholder="$t('combos.dishNamePlaceholder') + ' (RU)'" />
              <input v-model="ci.name_en" class="input combo-name" :placeholder="$t('combos.dishNamePlaceholder') + ' (EN)'" />
              <input v-model.number="ci.quantity" type="number" min="1" class="input combo-qty" />
              <button type="button" class="combo-remove" @click="removeFormItem(idx)">×</button>
            </div>
          </div>
          <button type="button" class="btn-add-item" @click="addFormItem">{{ $t('combos.addPosition') }}</button>

          <div class="modal-actions">
            <button type="button" class="btn btn-outline" @click="showForm = false">{{ $t('common.cancel') }}</button>
            <button type="submit" class="btn btn-primary" :disabled="saving">{{ saving ? $t('common.saving') : $t('common.save') }}</button>
          </div>
        </form>
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
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
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
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
  color: var(--text-secondary);
}

.combo-grid {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.combo-card {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.combo-card.unavailable {
  opacity: 0.5;
}

.combo-top {
  display: flex;
  gap: 12px;
  align-items: center;
}

.combo-thumb {
  width: 56px;
  height: 56px;
  border-radius: 10px;
  overflow: hidden;
  flex-shrink: 0;
  border: 1px solid var(--border);
}

.combo-thumb img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.combo-info h3 {
  font-size: 16px;
  font-weight: 700;
}

.combo-price {
  font-size: 14px;
  font-weight: 700;
  color: var(--primary);
}

.badge-off {
  display: inline-block;
  font-size: 11px;
  font-weight: 600;
  color: #DC2626;
  background: #FEE2E2;
  padding: 1px 8px;
  border-radius: 100px;
  margin-left: 8px;
}

.combo-desc {
  font-size: 13px;
  color: var(--text-secondary);
}

.combo-items-list {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
}

.combo-item-chip {
  padding: 4px 10px;
  background: var(--bg);
  border: 1px solid var(--border);
  border-radius: 100px;
  font-size: 12px;
  color: var(--text-secondary);
}

.combo-actions {
  display: flex;
  gap: 6px;
}

.icon-btn-sm {
  width: 28px;
  height: 28px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text-secondary);
}

.icon-btn-sm:hover { background: var(--bg); }
.icon-btn-sm.available { color: #16A34A; }
.icon-btn-sm.available:hover { background: #DCFCE7; }
.icon-btn-sm.off { color: #DC2626; }
.icon-btn-sm.off:hover { background: #FEE2E2; }
.icon-btn-sm.danger:hover { background: #FEE2E2; color: #DC2626; }

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
  max-width: 520px;
  max-height: 90vh;
  overflow-y: auto;
}

.modal h2 {
  font-size: 18px;
  font-weight: 700;
  margin-bottom: 20px;
}

.field { margin-bottom: 14px; }

.row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 14px;
}

.divider {
  height: 1px;
  background: var(--border);
  margin: 6px 0 14px;
}

.section-label {
  font-size: 12px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: var(--text-secondary);
  margin-bottom: 10px;
}

.lang-tag {
  display: inline-block;
  padding: 1px 6px;
  margin-left: 6px;
  font-size: 10px;
  font-weight: 700;
  color: var(--text-secondary);
  background: var(--bg);
  border: 1px solid var(--border);
  border-radius: 4px;
  vertical-align: middle;
  letter-spacing: 0.3px;
}

.combo-form-items {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-bottom: 10px;
}

.combo-form-row {
  display: flex;
  gap: 6px;
  align-items: center;
}

.combo-select { flex: 1; min-width: 0; font-size: 13px; }
.combo-name { flex: 1; min-width: 0; font-size: 13px; }
.combo-qty { width: 50px; flex-shrink: 0; font-size: 13px; text-align: center; }

.combo-remove {
  width: 26px;
  height: 26px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text-secondary);
  font-size: 16px;
  flex-shrink: 0;
}

.combo-remove:hover {
  background: #FEE2E2;
  color: #DC2626;
}

.btn-add-item {
  font-size: 13px;
  color: var(--primary);
  font-weight: 500;
  padding: 6px 0;
  margin-bottom: 10px;
}

.modal-actions {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
  margin-top: 20px;
}

/* Upload */
.upload-area { display: flex; gap: 10px; }

.upload-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 10px 16px;
  border: 2px dashed var(--border);
  border-radius: var(--radius);
  cursor: pointer;
  font-size: 13px;
  color: var(--text-secondary);
}

.upload-btn:hover {
  border-color: var(--primary);
  color: var(--primary);
}

.img-preview {
  position: relative;
  width: 72px;
  height: 72px;
  border-radius: var(--radius);
  overflow: hidden;
  border: 1px solid var(--border);
}

.img-preview img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.img-remove {
  position: absolute;
  top: 2px;
  right: 2px;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: rgba(0,0,0,0.6);
  color: #fff;
  font-size: 13px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  border: none;
}
</style>
