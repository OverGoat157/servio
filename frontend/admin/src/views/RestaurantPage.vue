<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { restaurants as api, uploadFile, imageUrl } from '../api/client'

const route = useRoute()
const router = useRouter()
const id = route.params.id

const rest = ref(null)
const loading = ref(true)
const saving = ref(false)
const success = ref(false)
const error = ref('')

const form = ref({
  name: '',
  slug: '',
  description: '',
  phone: '',
  address: '',
  working_hours: '',
  theme: '',
  cover_image: '',
  logo: '',
  promo_title: '',
  promo_description: '',
})

const uploadingLogo = ref(false)
const uploadingCover = ref(false)

onMounted(async () => {
  try {
    const data = await api.get(id)
    rest.value = data
    form.value = {
      name: data.name,
      slug: data.slug || '',
      description: data.description || '',
      phone: data.phone || '',
      address: data.address || '',
      working_hours: data.working_hours ? JSON.stringify(data.working_hours) : '',
      theme: data.theme,
      cover_image: data.cover_image || '',
      logo: data.logo || '',
      promo_title: data.promo_title || '',
      promo_description: data.promo_description || '',
    }
  } catch {
    router.push({ name: 'dashboard' })
  }
  loading.value = false
})

async function handleUpload(field, event) {
  const file = event.target.files?.[0]
  if (!file) return

  const uploading = field === 'logo' ? uploadingLogo : uploadingCover
  uploading.value = true
  try {
    const { url } = await uploadFile(file)
    form.value[field] = url
  } catch (e) {
    error.value = 'Ошибка загрузки: ' + e.message
  }
  uploading.value = false
}

function clearImage(field) {
  form.value[field] = ''
}

async function save() {
  saving.value = true
  success.value = false
  error.value = ''
  try {
    await api.update(id, form.value)
    success.value = true
    setTimeout(() => success.value = false, 2000)
  } catch (e) {
    error.value = 'Ошибка сохранения: ' + e.message
  }
  saving.value = false
}
</script>

<template>
  <div class="page">
    <div class="page-header">
      <button class="back-link" @click="router.push({ name: 'dashboard' })">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M19 12H5M12 19l-7-7 7-7"/></svg>
        Назад
      </button>
      <h1 v-if="rest">Настройки: {{ rest.name }}</h1>
    </div>

    <div v-if="loading" class="loading">Загрузка...</div>

    <div class="card form-card" v-else>
      <form @submit.prevent="save">
        <div class="field">
          <label class="label">Название ресторана</label>
          <input v-model="form.name" class="input" required />
        </div>
        <div class="field">
          <label class="label">Slug (URL)</label>
          <input v-model="form.slug" class="input" required placeholder="my-restaurant" />
          <div class="hint">Адрес меню: menu.ab-team.ru/<strong>{{ form.slug || 'slug' }}</strong></div>
        </div>
        <div class="field">
          <label class="label">Описание</label>
          <textarea v-model="form.description" class="input textarea" rows="3" placeholder="Ресторан изысканной кухни"></textarea>
        </div>
        <div class="row">
          <div class="field">
            <label class="label">Телефон</label>
            <input v-model="form.phone" class="input" placeholder="+7 (999) 123-45-67" />
          </div>
          <div class="field">
            <label class="label">Тема</label>
            <select v-model="form.theme" class="input">
              <option value="default">По умолчанию</option>
              <option value="dark">Тёмная</option>
              <option value="warm">Тёплая</option>
            </select>
          </div>
        </div>
        <div class="field">
          <label class="label">Адрес</label>
          <input v-model="form.address" class="input" placeholder="г. Москва, ул. Тверская, 1" />
        </div>

        <div class="divider"></div>

        <!-- Логотип -->
        <div class="field">
          <label class="label">Логотип</label>
          <div class="upload-area">
            <div class="preview" v-if="form.logo">
              <img :src="imageUrl(form.logo)" alt="Логотип" />
              <button type="button" class="preview-remove" @click="clearImage('logo')">×</button>
            </div>
            <label class="upload-btn" v-else>
              <input type="file" accept="image/*" hidden @change="handleUpload('logo', $event)" />
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M21 15v4a2 2 0 01-2 2H5a2 2 0 01-2-2v-4M17 8l-5-5-5 5M12 3v12"/></svg>
              {{ uploadingLogo ? 'Загрузка...' : 'Загрузить логотип' }}
            </label>
          </div>
        </div>

        <!-- Обложка -->
        <div class="field">
          <label class="label">Фото обложки</label>
          <div class="upload-area">
            <div class="preview preview-wide" v-if="form.cover_image">
              <img :src="imageUrl(form.cover_image)" alt="Обложка" />
              <button type="button" class="preview-remove" @click="clearImage('cover_image')">×</button>
            </div>
            <label class="upload-btn" v-else>
              <input type="file" accept="image/*" hidden @change="handleUpload('cover_image', $event)" />
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M21 15v4a2 2 0 01-2 2H5a2 2 0 01-2-2v-4M17 8l-5-5-5 5M12 3v12"/></svg>
              {{ uploadingCover ? 'Загрузка...' : 'Загрузить обложку' }}
            </label>
          </div>
          <div class="hint">Фон за названием на главной странице ресторана</div>
        </div>

        <div class="divider"></div>
        <h3 class="section-title">Акция / Спецпредложение</h3>
        <div class="field">
          <label class="label">Заголовок акции</label>
          <input v-model="form.promo_title" class="input" placeholder="Бизнес-ланч за 350 ₽" />
        </div>
        <div class="field">
          <label class="label">Описание акции</label>
          <textarea v-model="form.promo_description" class="input textarea" rows="2" placeholder="Каждый будний день с 12:00 до 15:00"></textarea>
        </div>

        <div class="divider"></div>
        <div class="field">
          <label class="label">Часы работы (JSON)</label>
          <textarea v-model="form.working_hours" class="input textarea" rows="4" placeholder='{"Пн — Пт": "10:00 — 22:00", "Сб — Вс": "11:00 — 23:00"}'></textarea>
          <div class="hint">Формат JSON: {"день": "время", ...}</div>
        </div>

        <div class="error-msg" v-if="error">{{ error }}</div>

        <div class="form-footer">
          <div class="success-msg" v-if="success">Сохранено!</div>
          <button type="submit" class="btn btn-primary" :disabled="saving">
            {{ saving ? 'Сохранение...' : 'Сохранить' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>
.page {
  max-width: 700px;
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
  margin-bottom: 12px;
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

.form-card {
  padding: 28px;
}

.field {
  margin-bottom: 16px;
}

.textarea {
  resize: vertical;
}

.row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}

.hint {
  font-size: 12px;
  color: var(--text-secondary);
  margin-top: 4px;
}

.divider {
  height: 1px;
  background: var(--border);
  margin: 8px 0;
}

.section-title {
  font-size: 15px;
  font-weight: 700;
  margin-bottom: 4px;
}

/* Upload */
.upload-area {
  display: flex;
  gap: 12px;
}

.upload-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 20px;
  border: 2px dashed var(--border);
  border-radius: var(--radius);
  cursor: pointer;
  font-size: 14px;
  color: var(--text-secondary);
  transition: border-color 0.2s, color 0.2s;
}

.upload-btn:hover {
  border-color: var(--primary);
  color: var(--primary);
}

.preview {
  position: relative;
  width: 80px;
  height: 80px;
  border-radius: var(--radius);
  overflow: hidden;
  border: 1px solid var(--border);
}

.preview img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.preview-wide {
  width: 160px;
  height: 90px;
}

.preview-remove {
  position: absolute;
  top: 2px;
  right: 2px;
  width: 22px;
  height: 22px;
  border-radius: 50%;
  background: rgba(0,0,0,0.6);
  color: #fff;
  font-size: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  border: none;
  line-height: 1;
}

.preview-remove:hover {
  background: rgba(220,38,38,0.8);
}

/* Messages */
.error-msg {
  padding: 10px 14px;
  background: #FEE2E2;
  color: #DC2626;
  border-radius: var(--radius);
  font-size: 14px;
  margin-bottom: 12px;
}

.form-footer {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 8px;
}

.success-msg {
  font-size: 14px;
  color: var(--success);
  font-weight: 600;
}

@media (max-width: 640px) {
  .row {
    grid-template-columns: 1fr;
  }
}
</style>
