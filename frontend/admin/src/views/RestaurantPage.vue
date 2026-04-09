<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { restaurants as api, uploadFile, imageUrl } from '../api/client'

const API_BASE = import.meta.env.VITE_API_URL || 'http://localhost:8080'

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
  theme: '',
  cover_image: '',
  logo: '',
  promo_title: '',
  promo_description: '',
})

const dayNames = ['Пн', 'Вт', 'Ср', 'Чт', 'Пт', 'Сб', 'Вс']

function defaultSchedule() {
  return dayNames.map(day => ({
    day,
    open: '10:00',
    close: '22:00',
    break_start: '',
    break_end: '',
    day_off: false,
  }))
}

const schedule = ref(defaultSchedule())

function parseScheduleFromData(data) {
  if (!data) return defaultSchedule()
  const wh = typeof data === 'string' ? JSON.parse(data) : data

  // New format: array of day objects
  if (Array.isArray(wh)) return wh

  // Old format: { "Пн — Пт": "10:00 — 22:00" } — convert to defaults
  return defaultSchedule()
}

function scheduleToJSON() {
  return JSON.stringify(schedule.value)
}

const socialLinks = ref([])
const socialTypes = [
  { value: 'instagram', label: 'Instagram' },
  { value: 'vk', label: 'VK' },
  { value: 'telegram', label: 'Telegram' },
  { value: 'youtube', label: 'YouTube' },
  { value: 'tiktok', label: 'TikTok' },
  { value: 'facebook', label: 'Facebook' },
  { value: 'website', label: 'Сайт' },
]

function addSocialLink() {
  socialLinks.value.push({ type: 'instagram', url: '' })
}

function removeSocialLink(idx) {
  socialLinks.value.splice(idx, 1)
}

const uploadingLogo = ref(false)
const uploadingCover = ref(false)
const qrKey = ref(0)

const qrUrl = computed(() => {
  if (!rest.value) return ''
  return `${API_BASE}/api/restaurants/${id}/qrcode?size=512&_=${qrKey.value}`
})

function regenerateQR() {
  qrKey.value++
}

function downloadQR() {
  const token = localStorage.getItem('token')
  fetch(`${API_BASE}/api/restaurants/${id}/qrcode?size=1024`, {
    headers: { Authorization: `Bearer ${token}` }
  })
    .then(r => r.blob())
    .then(blob => {
      const url = URL.createObjectURL(blob)
      const a = document.createElement('a')
      a.href = url
      a.download = `qr-${rest.value?.slug || 'menu'}.png`
      a.click()
      URL.revokeObjectURL(url)
    })
}

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
      theme: data.theme,
      cover_image: data.cover_image || '',
      logo: data.logo || '',
      promo_title: data.promo_title || '',
      promo_description: data.promo_description || '',
    }
    try {
      schedule.value = parseScheduleFromData(data.working_hours)
    } catch { schedule.value = defaultSchedule() }
    try {
      const links = typeof data.social_links === 'string' ? JSON.parse(data.social_links) : (data.social_links || [])
      socialLinks.value = Array.isArray(links) ? links : []
    } catch { socialLinks.value = [] }
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
    const validLinks = socialLinks.value.filter(l => l.url.trim())
    await api.update(id, { ...form.value, working_hours: scheduleToJSON(), social_links: JSON.stringify(validLinks) })
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
        <h3 class="section-title">Соцсети и ссылки</h3>
        <div class="social-list">
          <div class="social-row" v-for="(link, idx) in socialLinks" :key="idx">
            <select v-model="link.type" class="input social-type">
              <option v-for="st in socialTypes" :key="st.value" :value="st.value">{{ st.label }}</option>
            </select>
            <input v-model="link.url" class="input" placeholder="https://..." />
            <button type="button" class="social-remove" @click="removeSocialLink(idx)">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M18 6L6 18M6 6l12 12"/></svg>
            </button>
          </div>
        </div>
        <button type="button" class="btn-add-social" @click="addSocialLink">
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"><path d="M12 5v14M5 12h14"/></svg>
          Добавить ссылку
        </button>

        <div class="divider"></div>
        <h3 class="section-title">Режим работы</h3>
        <div class="schedule-list">
          <div class="sched-row" v-for="s in schedule" :key="s.day" :class="{ 'sched-off': s.day_off }">
            <div class="sched-day">{{ s.day }}</div>
            <div class="sched-fields" v-if="!s.day_off">
              <div class="sched-time-group">
                <input type="time" v-model="s.open" class="sched-time" />
                <span class="sched-sep">—</span>
                <input type="time" v-model="s.close" class="sched-time" />
              </div>
              <div class="sched-break">
                <label class="sched-break-label" v-if="!s.break_start" @click="s.break_start = '13:00'; s.break_end = '14:00'">+ обед</label>
                <template v-else>
                  <span class="sched-break-tag">Обед:</span>
                  <input type="time" v-model="s.break_start" class="sched-time sched-time-sm" />
                  <span class="sched-sep">—</span>
                  <input type="time" v-model="s.break_end" class="sched-time sched-time-sm" />
                  <button type="button" class="sched-break-remove" @click="s.break_start = ''; s.break_end = ''">×</button>
                </template>
              </div>
            </div>
            <div class="sched-dayoff-label" v-else>Выходной</div>
            <label class="sched-toggle">
              <input type="checkbox" :checked="!s.day_off" @change="s.day_off = !s.day_off" />
              <span class="sched-toggle-slider"></span>
            </label>
          </div>
        </div>

        <div class="divider"></div>
        <h3 class="section-title">QR-код меню</h3>
        <div class="qr-section">
          <p class="qr-hint">Распечатайте и разместите на столах. Гости отсканируют и попадут в ваше меню.</p>
          <p class="qr-link">menu.ab-team.ru/{{ form.slug }}</p>
          <div class="qr-actions">
            <button type="button" class="btn btn-primary btn-sm" @click="downloadQR">Скачать PNG</button>
            <button type="button" class="btn btn-outline btn-sm" @click="regenerateQR">Перегенерировать</button>
          </div>
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

/* QR Code */
.qr-section {
  margin-bottom: 16px;
}

.qr-hint {
  font-size: 13px;
  color: var(--text-secondary);
  line-height: 1.5;
  margin-bottom: 8px;
}

.qr-link {
  font-size: 14px;
  font-weight: 600;
  color: var(--primary);
  margin-bottom: 12px;
}

.qr-actions {
  display: flex;
  gap: 8px;
}

/* Social links */
.social-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-bottom: 12px;
}

.social-row {
  display: flex;
  gap: 8px;
  align-items: center;
}

.social-type {
  width: 140px;
  flex-shrink: 0;
}

.social-remove {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text-secondary);
  flex-shrink: 0;
}

.social-remove:hover {
  background: #FEE2E2;
  color: #DC2626;
}

.btn-add-social {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 8px 16px;
  font-size: 13px;
  font-weight: 500;
  color: var(--primary);
  border: 1px dashed var(--primary);
  border-radius: var(--radius);
  margin-bottom: 8px;
}

.btn-add-social:hover {
  background: var(--bg);
}

/* Schedule */
.schedule-list {
  display: flex;
  flex-direction: column;
  gap: 6px;
  margin-bottom: 16px;
}

.sched-row {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 14px;
  background: var(--bg);
  border: 1px solid var(--border);
  border-radius: var(--radius);
  transition: opacity 0.2s;
}

.sched-row.sched-off {
  opacity: 0.5;
  background: var(--bg-secondary, #f9fafb);
}

.sched-day {
  width: 28px;
  font-size: 14px;
  font-weight: 700;
  flex-shrink: 0;
}

.sched-fields {
  flex: 1;
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 8px;
}

.sched-time-group {
  display: flex;
  align-items: center;
  gap: 6px;
}

.sched-time {
  width: 100px;
  padding: 6px 8px;
  border: 1px solid var(--border);
  border-radius: 6px;
  font-size: 13px;
  color: var(--text);
  background: var(--bg);
}

.sched-time-sm {
  width: 90px;
}

.sched-sep {
  font-size: 13px;
  color: var(--text-secondary);
}

.sched-break {
  display: flex;
  align-items: center;
  gap: 6px;
}

.sched-break-label {
  font-size: 12px;
  color: var(--primary);
  cursor: pointer;
  padding: 4px 10px;
  border: 1px dashed var(--primary);
  border-radius: 100px;
  white-space: nowrap;
}

.sched-break-label:hover {
  background: var(--primary-light, #f0f5ff);
}

.sched-break-tag {
  font-size: 12px;
  font-weight: 600;
  color: var(--text-secondary);
  white-space: nowrap;
}

.sched-break-remove {
  width: 22px;
  height: 22px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  color: var(--text-secondary);
  flex-shrink: 0;
}

.sched-break-remove:hover {
  background: #FEE2E2;
  color: #DC2626;
}

.sched-dayoff-label {
  flex: 1;
  font-size: 13px;
  color: var(--text-secondary);
  font-style: italic;
}

.sched-toggle {
  position: relative;
  display: inline-block;
  width: 36px;
  height: 20px;
  flex-shrink: 0;
}

.sched-toggle input {
  opacity: 0;
  width: 0;
  height: 0;
}

.sched-toggle-slider {
  position: absolute;
  inset: 0;
  background: var(--border);
  border-radius: 20px;
  cursor: pointer;
  transition: background 0.2s;
}

.sched-toggle-slider::before {
  content: '';
  position: absolute;
  width: 14px;
  height: 14px;
  left: 3px;
  bottom: 3px;
  background: #fff;
  border-radius: 50%;
  transition: transform 0.2s;
}

.sched-toggle input:checked + .sched-toggle-slider {
  background: var(--success, #22c55e);
}

.sched-toggle input:checked + .sched-toggle-slider::before {
  transform: translateX(16px);
}

@media (max-width: 640px) {
  .sched-row {
    flex-wrap: wrap;
  }
  .sched-fields {
    width: 100%;
  }
  .sched-time {
    width: 85px;
  }
  .sched-time-sm {
    width: 80px;
  }
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
