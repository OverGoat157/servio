<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { login } from '../stores/auth'

const router = useRouter()
const { t } = useI18n()
const email = ref('')
const password = ref('')
const error = ref('')
const loading = ref(false)

async function handleSubmit() {
  error.value = ''
  loading.value = true
  try {
    await login(email.value, password.value)
    router.push({ name: 'dashboard' })
  } catch (e) {
    error.value = e.message === 'invalid credentials' ? t('auth.invalidCreds') : e.message
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="auth-page">
    <div class="auth-card card">
      <div class="auth-header">
        <svg class="auth-logo" viewBox="0 0 32 32" fill="none" stroke="currentColor" stroke-width="3.5" stroke-linecap="round" stroke-linejoin="round">
          <path d="M7 10l5 6-5 6"/>
          <path d="M13 10l5 6-5 6"/>
          <path d="M19 10l5 6-5 6"/>
        </svg>
        <h1>{{ $t('auth.loginTitle') }}</h1>
        <p>{{ $t('auth.loginSubtitle') }}</p>
      </div>

      <form @submit.prevent="handleSubmit">
        <div class="field">
          <label class="label">{{ $t('common.email') }}</label>
          <input v-model="email" type="email" class="input" placeholder="email@example.com" required />
        </div>
        <div class="field">
          <label class="label">{{ $t('common.password') }}</label>
          <input v-model="password" type="password" class="input" :placeholder="$t('auth.passwordPlaceholder')" required />
        </div>
        <div class="error-msg" v-if="error">{{ error }}</div>
        <button type="submit" class="btn btn-primary submit-btn" :disabled="loading">
          {{ loading ? $t('auth.submitLoading') : $t('auth.submit') }}
        </button>
      </form>

      <div class="auth-footer">
        {{ $t('auth.footer') }}
      </div>
    </div>
  </div>
</template>

<style scoped>
.auth-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
}

.auth-card {
  width: 100%;
  max-width: 400px;
}

.auth-header {
  text-align: center;
  margin-bottom: 32px;
}

.auth-logo {
  display: block;
  width: 56px;
  height: 56px;
  color: var(--primary);
  margin: 0 auto 16px;
}

.auth-header h1 {
  font-size: 22px;
  font-weight: 700;
  margin-bottom: 4px;
}

.auth-header p {
  font-size: 14px;
  color: var(--text-secondary);
}

.field {
  margin-bottom: 16px;
}

.error-msg {
  padding: 10px;
  background: #FEE2E2;
  color: var(--danger);
  border-radius: var(--radius);
  font-size: 13px;
  margin-bottom: 16px;
  text-align: center;
}

.submit-btn {
  width: 100%;
  padding: 12px;
}

.submit-btn:disabled {
  opacity: 0.6;
}

.auth-footer {
  text-align: center;
  margin-top: 20px;
  font-size: 14px;
  color: var(--text-secondary);
}

.auth-footer a {
  color: var(--primary);
  font-weight: 600;
}
</style>
