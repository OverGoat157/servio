<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { restaurants as restApi, messengers as msgApi } from '../api/client'

const route = useRoute()
const router = useRouter()
const id = route.params.id

const rest = ref(null)
const configs = ref([])
const loading = ref(true)

// Telegram form
const tgForm = ref({ bot_token: '', chat_id: '' })
const tgEnabled = ref(false)
const tgSaving = ref(false)
const tgSuccess = ref(false)

// WhatsApp form
const waForm = ref({ phone: '' })
const waEnabled = ref(false)
const waSaving = ref(false)
const waSuccess = ref(false)

onMounted(async () => {
  try {
    rest.value = await restApi.get(id)
    configs.value = await msgApi.list(id)

    // Заполняем формы из существующих конфигов
    const tgConfig = configs.value.find(c => c.type === 'telegram')
    if (tgConfig) {
      try {
        const parsed = JSON.parse(tgConfig.config)
        tgForm.value = { bot_token: parsed.bot_token || '', chat_id: parsed.chat_id || '' }
        tgEnabled.value = tgConfig.enabled
      } catch { /* malformed config, use defaults */ }
    }

    const waConfig = configs.value.find(c => c.type === 'whatsapp')
    if (waConfig) {
      try {
        const parsed = JSON.parse(waConfig.config)
        waForm.value = { phone: parsed.phone || '' }
        waEnabled.value = waConfig.enabled
      } catch { /* malformed config, use defaults */ }
    }
  } catch {
    router.push({ name: 'dashboard' })
  }
  loading.value = false
})

async function saveTelegram() {
  tgSaving.value = true
  tgSuccess.value = false
  try {
    await msgApi.upsert(id, {
      type: 'telegram',
      config: JSON.stringify(tgForm.value),
      enabled: tgEnabled.value,
    })
    tgSuccess.value = true
    setTimeout(() => tgSuccess.value = false, 2000)
  } catch { /* empty */ }
  tgSaving.value = false
}

async function saveWhatsApp() {
  waSaving.value = true
  waSuccess.value = false
  try {
    await msgApi.upsert(id, {
      type: 'whatsapp',
      config: JSON.stringify(waForm.value),
      enabled: waEnabled.value,
    })
    waSuccess.value = true
    setTimeout(() => waSuccess.value = false, 2000)
  } catch { /* empty */ }
  waSaving.value = false
}

async function deleteConfig(type) {
  if (!confirm(`Удалить настройки ${type}?`)) return
  await msgApi.delete(id, type)
  if (type === 'telegram') {
    tgForm.value = { bot_token: '', chat_id: '' }
    tgEnabled.value = false
  } else {
    waForm.value = { phone: '' }
    waEnabled.value = false
  }
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
        <h1 v-if="rest">Мессенджеры: {{ rest.name }}</h1>
      </div>
    </div>

    <div v-if="loading" class="loading">Загрузка...</div>

    <div class="sections" v-else>
      <!-- Telegram -->
      <div class="card section">
        <div class="section-header">
          <div class="section-icon tg">
            <svg width="22" height="22" viewBox="0 0 24 24" fill="#fff"><path d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12 12-5.373 12-12S18.627 0 12 0zm5.562 8.161c-.18 1.897-.962 6.502-1.359 8.627-.168.9-.5 1.201-.82 1.23-.697.064-1.226-.46-1.9-.903-1.056-.692-1.653-1.123-2.678-1.799-1.185-.781-.417-1.21.258-1.911.177-.184 3.247-2.977 3.307-3.23.007-.032.014-.15-.056-.212s-.174-.041-.249-.024c-.106.024-1.793 1.14-5.061 3.345-.479.33-.913.492-1.302.484-.428-.008-1.252-.241-1.865-.44-.752-.245-1.349-.374-1.297-.789.027-.216.325-.437.893-.663 3.498-1.524 5.831-2.529 6.998-3.014 3.332-1.386 4.025-1.627 4.476-1.635.099-.002.321.023.465.141a.506.506 0 01.171.325c.016.093.036.306.02.472z"/></svg>
          </div>
          <div>
            <h3>Telegram</h3>
            <p>Заказы приходят сообщением от бота в чат</p>
          </div>
        </div>

        <form @submit.prevent="saveTelegram">
          <div class="field">
            <label class="label">Bot Token</label>
            <input v-model="tgForm.bot_token" class="input" placeholder="123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11" />
            <div class="hint">Получите у @BotFather в Telegram</div>
          </div>
          <div class="field">
            <label class="label">Chat ID</label>
            <input v-model="tgForm.chat_id" class="input" placeholder="-1001234567890" />
            <div class="hint">ID чата и��и группы, куда бот будет отправлять заказы</div>
          </div>
          <div class="toggle-row">
            <label class="toggle">
              <input type="checkbox" v-model="tgEnabled" />
              <span class="toggle-slider"></span>
            </label>
            <span>Включен��</span>
          </div>
          <div class="section-actions">
            <div class="success-msg" v-if="tgSuccess">Сохранено!</div>
            <button type="button" class="btn btn-sm" style="color: var(--danger)" @click="deleteConfig('telegram')" v-if="tgForm.bot_token">Удалить</button>
            <button type="submit" class="btn btn-primary btn-sm" :disabled="tgSaving">
              {{ tgSaving ? 'Сохранение...' : 'Сохранить' }}
            </button>
          </div>
        </form>
      </div>

      <!-- WhatsApp -->
      <div class="card section">
        <div class="section-header">
          <div class="section-icon wa">
            <svg width="22" height="22" viewBox="0 0 24 24" fill="#fff"><path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347z"/><path d="M12 0C5.373 0 0 5.373 0 12c0 2.127.556 4.12 1.525 5.855L0 24l6.335-1.652A11.94 11.94 0 0012 24c6.627 0 12-5.373 12-12S18.627 0 12 0z"/></svg>
          </div>
          <div>
            <h3>WhatsApp</h3>
            <p>Клиент отправляет заказ вам в WhatsApp по ссылке</p>
          </div>
        </div>

        <form @submit.prevent="saveWhatsApp">
          <div class="field">
            <label class="label">Номер телефона</label>
            <input v-model="waForm.phone" class="input" placeholder="79991234567" />
            <div class="hint">Без +, пробелов и скобок. Пример: 79991234567</div>
          </div>
          <div class="toggle-row">
            <label class="toggle">
              <input type="checkbox" v-model="waEnabled" />
              <span class="toggle-slider"></span>
            </label>
            <span>Включено</span>
          </div>
          <div class="section-actions">
            <div class="success-msg" v-if="waSuccess">Сохранено!</div>
            <button type="button" class="btn btn-sm" style="color: var(--danger)" @click="deleteConfig('whatsapp')" v-if="waForm.phone">Удалить</button>
            <button type="submit" class="btn btn-primary btn-sm" :disabled="waSaving">
              {{ waSaving ? 'Сохранение...' : 'Сохранить' }}
            </button>
          </div>
        </form>
      </div>

      <!-- How it works -->
      <div class="card info-card">
        <h3>Как это работает?</h3>
        <div class="info-item">
          <strong>Telegram:</strong> Создайт�� бота через @BotFather, добавьте его в группу. Заказы будут приходит�� автоматически от бота.
        </div>
        <div class="info-item">
          <strong>WhatsApp:</strong> Укажите номер. Когда клиент ��формляет заказ, он переходит в WhatsApp с готовым текстом заказа и отправляет его вам.
        </div>
      </div>
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

.sections {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.section {
  padding: 24px;
}

.section-header {
  display: flex;
  align-items: center;
  gap: 14px;
  margin-bottom: 20px;
}

.section-icon {
  width: 44px;
  height: 44px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.section-icon.tg {
  background: #229ED9;
}

.section-icon.wa {
  background: #25D366;
}

.section-header h3 {
  font-size: 16px;
  font-weight: 700;
}

.section-header p {
  font-size: 13px;
  color: var(--text-secondary);
  margin-top: 2px;
}

.field {
  margin-bottom: 14px;
}

.hint {
  font-size: 12px;
  color: var(--text-secondary);
  margin-top: 4px;
}

.toggle-row {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 16px;
  font-size: 14px;
}

.toggle {
  position: relative;
  display: inline-block;
  width: 44px;
  height: 24px;
}

.toggle input {
  opacity: 0;
  width: 0;
  height: 0;
}

.toggle-slider {
  position: absolute;
  inset: 0;
  background: var(--border);
  border-radius: 24px;
  cursor: pointer;
  transition: background var(--ease);
}

.toggle-slider::before {
  content: '';
  position: absolute;
  width: 18px;
  height: 18px;
  left: 3px;
  bottom: 3px;
  background: #fff;
  border-radius: 50%;
  transition: transform var(--ease);
}

.toggle input:checked + .toggle-slider {
  background: var(--primary);
}

.toggle input:checked + .toggle-slider::before {
  transform: translateX(20px);
}

.section-actions {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 10px;
}

.success-msg {
  font-size: 14px;
  color: var(--success);
  font-weight: 600;
}

.info-card {
  background: var(--primary-light);
  border-color: transparent;
}

.info-card h3 {
  font-size: 15px;
  font-weight: 700;
  margin-bottom: 12px;
}

.info-item {
  font-size: 13px;
  line-height: 1.6;
  color: var(--text-secondary);
  margin-bottom: 8px;
}

.info-item strong {
  color: var(--text);
}
</style>
