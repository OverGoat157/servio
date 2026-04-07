<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { restaurants as api } from '../api/client'

const route = useRoute()
const router = useRouter()
const id = route.params.id

const rest = ref(null)
const loading = ref(true)
const saving = ref(false)
const success = ref(false)

const form = ref({
  name: '',
  description: '',
  phone: '',
  address: '',
  working_hours: '',
  theme: '',
  cover_image: '',
  promo_title: '',
  promo_description: '',
})

onMounted(async () => {
  try {
    const data = await api.get(id)
    rest.value = data
    form.value = {
      name: data.name,
      description: data.description || '',
      phone: data.phone || '',
      address: data.address || '',
      working_hours: data.working_hours ? JSON.stringify(data.working_hours) : '',
      theme: data.theme,
      cover_image: data.cover_image || '',
      promo_title: data.promo_title || '',
      promo_description: data.promo_description || '',
    }
  } catch {
    router.push({ name: 'dashboard' })
  }
  loading.value = false
})

async function save() {
  saving.value = true
  success.value = false
  try {
    await api.update(id, form.value)
    success.value = true
    setTimeout(() => success.value = false, 2000)
  } catch { /* empty */ }
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
        <div class="field">
          <label class="label">Фото обложки (URL)</label>
          <input v-model="form.cover_image" class="input" placeholder="https://example.com/photo.jpg" />
          <div class="hint">Ссылка на фото ресторана для фона на главной странице</div>
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

@media (max-width: 640px) {
  .row {
    grid-template-columns: 1fr;
  }
}
</style>
