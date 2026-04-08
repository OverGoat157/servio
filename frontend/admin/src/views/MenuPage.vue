<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { restaurants as restApi, categories as catApi, menuItems as itemApi, uploadFile, imageUrl } from '../api/client'

const route = useRoute()
const router = useRouter()
const id = route.params.id

const rest = ref(null)
const cats = ref([])
const loading = ref(true)

// Category form
const showCatForm = ref(false)
const catForm = ref({ name: '', sort_order: 0 })
const editingCatId = ref(null)

// Item form
const showItemForm = ref(false)
const itemForm = ref({ name: '', description: '', price: '', sort_order: 0, image: '', ingredients: '', weight: '', calories: '', proteins: '', fats: '', carbs: '', cook_time: '' })
const editingItemId = ref(null)
const targetCatId = ref(null)
const uploadingItemImage = ref(false)

onMounted(async () => {
  try {
    rest.value = await restApi.get(id)
    await loadCategories()
  } catch {
    router.push({ name: 'dashboard' })
  }
  loading.value = false
})

async function loadCategories() {
  const catList = await catApi.list(id) || []
  // Load items for each category
  for (const cat of catList) {
    cat.items = await itemApi.list(cat.id) || []
  }
  cats.value = catList
}

// Category CRUD
function openCatForm(cat = null) {
  if (cat) {
    editingCatId.value = cat.id
    catForm.value = { name: cat.name, sort_order: cat.sort_order }
  } else {
    editingCatId.value = null
    catForm.value = { name: '', sort_order: cats.value.length }
  }
  showCatForm.value = true
}

async function saveCat() {
  if (editingCatId.value) {
    await catApi.update(editingCatId.value, catForm.value)
  } else {
    await catApi.create(id, catForm.value)
  }
  showCatForm.value = false
  await loadCategories()
}

async function deleteCat(catId) {
  if (!confirm('Удалить категорию и все её позиции?')) return
  await catApi.delete(catId)
  await loadCategories()
}

// Item CRUD
function openItemForm(catId, item = null) {
  targetCatId.value = catId
  if (item) {
    editingItemId.value = item.id
    itemForm.value = {
      name: item.name,
      description: item.description || '',
      price: String(item.price / 100),
      sort_order: item.sort_order,
      image: item.image || '',
      ingredients: item.ingredients || '',
      weight: item.weight || '',
      calories: item.calories ?? '',
      proteins: item.proteins ?? '',
      fats: item.fats ?? '',
      carbs: item.carbs ?? '',
      cook_time: item.cook_time || '',
    }
  } else {
    editingItemId.value = null
    itemForm.value = { name: '', description: '', price: '', sort_order: 0, image: '', ingredients: '', weight: '', calories: '', proteins: '', fats: '', carbs: '', cook_time: '' }
  }
  showItemForm.value = true
}

async function saveItem() {
  const data = {
    ...itemForm.value,
    price: Math.round(parseFloat(itemForm.value.price) * 100),
    calories: itemForm.value.calories !== '' ? Number(itemForm.value.calories) : null,
    proteins: itemForm.value.proteins !== '' ? Number(itemForm.value.proteins) : null,
    fats: itemForm.value.fats !== '' ? Number(itemForm.value.fats) : null,
    carbs: itemForm.value.carbs !== '' ? Number(itemForm.value.carbs) : null,
  }
  if (editingItemId.value) {
    await itemApi.update(editingItemId.value, data)
  } else {
    await itemApi.create(targetCatId.value, data)
  }
  showItemForm.value = false
  await loadCategories()
}

async function deleteItem(itemId) {
  if (!confirm('Удалить позицию?')) return
  await itemApi.delete(itemId)
  await loadCategories()
}

async function handleItemImage(event) {
  const file = event.target.files?.[0]
  if (!file) return
  uploadingItemImage.value = true
  try {
    const { url } = await uploadFile(file)
    itemForm.value.image = url
  } catch { /* empty */ }
  uploadingItemImage.value = false
}

function formatPrice(kopecks) {
  return (kopecks / 100).toLocaleString('ru-RU') + ' ₽'
}
</script>

<template>
  <div class="page">
    <div class="page-header">
      <div>
        <button class="back-link" @click="router.push({ name: 'dashboard' })">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M19 12H5M12 19l-7-7 7-7"/></svg>
          Назад
        </button>
        <h1 v-if="rest">Меню: {{ rest.name }}</h1>
      </div>
      <button class="btn btn-primary" @click="openCatForm()">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"><path d="M12 5v14M5 12h14"/></svg>
        Добавить категорию
      </button>
    </div>

    <div v-if="loading" class="loading">Загрузка...</div>

    <!-- Empty -->
    <div class="empty card" v-else-if="cats.length === 0">
      <p>Добавьте первую категорию, например «Супы» или «Горячее»</p>
      <button class="btn btn-primary" @click="openCatForm()">Добавить категорию</button>
    </div>

    <!-- Categories -->
    <div class="categories" v-else>
      <div class="cat-section card" v-for="cat in cats" :key="cat.id">
        <div class="cat-header">
          <h3>{{ cat.name }}</h3>
          <div class="cat-actions">
            <button class="icon-btn" title="Добавить позицию" @click="openItemForm(cat.id)">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="var(--primary)" stroke-width="2.5" stroke-linecap="round"><path d="M12 5v14M5 12h14"/></svg>
            </button>
            <button class="icon-btn" title="Редактировать" @click="openCatForm(cat)">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M11 4H4a2 2 0 00-2 2v14a2 2 0 002 2h14a2 2 0 002-2v-7"/><path d="M18.5 2.5a2.121 2.121 0 013 3L12 15l-4 1 1-4 9.5-9.5z"/></svg>
            </button>
            <button class="icon-btn danger" title="Удалить" @click="deleteCat(cat.id)">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M3 6h18M19 6v14a2 2 0 01-2 2H7a2 2 0 01-2-2V6m3 0V4a2 2 0 012-2h4a2 2 0 012 2v2"/></svg>
            </button>
          </div>
        </div>

        <div class="items-list" v-if="cat.items?.length">
          <div class="item-row" v-for="item in cat.items" :key="item.id">
            <div class="item-thumb" v-if="item.image">
              <img :src="imageUrl(item.image)" :alt="item.name" />
            </div>
            <div class="item-info">
              <span class="item-name">{{ item.name }}</span>
              <span class="item-desc" v-if="item.description">{{ item.description }}</span>
            </div>
            <div class="item-price">{{ formatPrice(item.price) }}</div>
            <div class="item-actions">
              <button class="icon-btn-sm" @click="openItemForm(cat.id, item)">
                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M11 4H4a2 2 0 00-2 2v14a2 2 0 002 2h14a2 2 0 002-2v-7"/><path d="M18.5 2.5a2.121 2.121 0 013 3L12 15l-4 1 1-4 9.5-9.5z"/></svg>
              </button>
              <button class="icon-btn-sm danger" @click="deleteItem(item.id)">
                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M18 6L6 18M6 6l12 12"/></svg>
              </button>
            </div>
          </div>
        </div>

        <div class="empty-items" v-else>
          <button class="btn btn-outline btn-sm" @click="openItemForm(cat.id)">
            Добавить первую позицию
          </button>
        </div>
      </div>
    </div>

    <!-- Category modal -->
    <div class="modal-overlay" v-if="showCatForm" @click.self="showCatForm = false">
      <div class="modal card">
        <h2>{{ editingCatId ? 'Редактировать категорию' : 'Новая категория' }}</h2>
        <form @submit.prevent="saveCat">
          <div class="field">
            <label class="label">Название</label>
            <input v-model="catForm.name" class="input" placeholder="Супы" required />
          </div>
          <div class="field">
            <label class="label">Порядок сортировки</label>
            <input v-model.number="catForm.sort_order" type="number" class="input" />
          </div>
          <div class="modal-actions">
            <button type="button" class="btn btn-outline" @click="showCatForm = false">Отмена</button>
            <button type="submit" class="btn btn-primary">Сохранить</button>
          </div>
        </form>
      </div>
    </div>

    <!-- Item modal -->
    <div class="modal-overlay" v-if="showItemForm" @click.self="showItemForm = false">
      <div class="modal card">
        <h2>{{ editingItemId ? 'Редактировать позицию' : 'Новая позиция' }}</h2>
        <form @submit.prevent="saveItem">
          <div class="field">
            <label class="label">Фото блюда</label>
            <div class="upload-area">
              <div class="img-preview" v-if="itemForm.image">
                <img :src="imageUrl(itemForm.image)" alt="" />
                <button type="button" class="img-remove" @click="itemForm.image = ''">×</button>
              </div>
              <label class="upload-btn" v-else>
                <input type="file" accept="image/*" hidden @change="handleItemImage" />
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M21 15v4a2 2 0 01-2 2H5a2 2 0 01-2-2v-4M17 8l-5-5-5 5M12 3v12"/></svg>
                {{ uploadingItemImage ? 'Загрузка...' : 'Загрузить фото' }}
              </label>
            </div>
          </div>
          <div class="field">
            <label class="label">Название</label>
            <input v-model="itemForm.name" class="input" placeholder="Том Ям" required />
          </div>
          <div class="field">
            <label class="label">Описание</label>
            <input v-model="itemForm.description" class="input" placeholder="с креветками, 350 мл" />
          </div>
          <div class="row">
            <div class="field">
              <label class="label">Цена (₽)</label>
              <input v-model="itemForm.price" type="number" step="0.01" min="0" class="input" placeholder="890" required />
            </div>
            <div class="field">
              <label class="label">Порядок</label>
              <input v-model.number="itemForm.sort_order" type="number" class="input" />
            </div>
          </div>

          <div class="divider"></div>
          <div class="section-label">Подробности (необязательно)</div>

          <div class="field">
            <label class="label">Ингредиенты</label>
            <textarea v-model="itemForm.ingredients" class="input textarea" rows="2" placeholder="Креветки, кокосовое молоко, лемонграсс..."></textarea>
          </div>
          <div class="row">
            <div class="field">
              <label class="label">Граммовка</label>
              <input v-model="itemForm.weight" class="input" placeholder="350 г" />
            </div>
            <div class="field">
              <label class="label">Время готовки</label>
              <input v-model="itemForm.cook_time" class="input" placeholder="15 мин" />
            </div>
          </div>
          <div class="row row-4">
            <div class="field">
              <label class="label">Ккал</label>
              <input v-model="itemForm.calories" type="number" min="0" class="input" placeholder="320" />
            </div>
            <div class="field">
              <label class="label">Белки</label>
              <input v-model="itemForm.proteins" type="number" step="0.1" min="0" class="input" placeholder="12" />
            </div>
            <div class="field">
              <label class="label">Жиры</label>
              <input v-model="itemForm.fats" type="number" step="0.1" min="0" class="input" placeholder="8" />
            </div>
            <div class="field">
              <label class="label">Углев.</label>
              <input v-model="itemForm.carbs" type="number" step="0.1" min="0" class="input" placeholder="45" />
            </div>
          </div>
          <div class="modal-actions">
            <button type="button" class="btn btn-outline" @click="showItemForm = false">Отмена</button>
            <button type="submit" class="btn btn-primary">Сохранить</button>
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

/* Categories */
.categories {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.cat-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.cat-header h3 {
  font-size: 17px;
  font-weight: 700;
}

.cat-actions {
  display: flex;
  gap: 6px;
}

.icon-btn {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text-secondary);
  transition: background var(--ease);
}

.icon-btn:hover {
  background: var(--bg);
}

.icon-btn.danger:hover {
  background: #FEE2E2;
  color: var(--danger);
}

/* Items */
.items-list {
  display: flex;
  flex-direction: column;
}

.item-row {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 0;
  border-top: 1px solid var(--border);
}

.item-info {
  flex: 1;
  min-width: 0;
}

.item-name {
  font-size: 14px;
  font-weight: 600;
  display: block;
}

.item-desc {
  font-size: 12px;
  color: var(--text-secondary);
}

.item-price {
  font-size: 14px;
  font-weight: 700;
  white-space: nowrap;
}

.item-actions {
  display: flex;
  gap: 4px;
}

.icon-btn-sm {
  width: 26px;
  height: 26px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text-secondary);
}

.icon-btn-sm:hover {
  background: var(--bg);
}

.icon-btn-sm.danger:hover {
  background: #FEE2E2;
  color: var(--danger);
}

.empty-items {
  text-align: center;
  padding: 16px 0;
}

/* Modals */
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

.row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 14px;
}

.row-4 {
  grid-template-columns: repeat(4, 1fr);
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
  margin-bottom: 12px;
}

.textarea {
  resize: vertical;
}

.modal-actions {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
  margin-top: 20px;
}

/* Thumbnails */
.item-thumb {
  width: 44px;
  height: 44px;
  border-radius: 8px;
  overflow: hidden;
  flex-shrink: 0;
}

.item-thumb img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* Upload in modal */
.upload-area {
  display: flex;
  gap: 10px;
}

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
  transition: border-color 0.2s, color 0.2s;
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
  line-height: 1;
}

.img-remove:hover {
  background: rgba(220,38,38,0.8);
}
</style>
