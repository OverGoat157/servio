<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { login } from '../stores/auth'

const router = useRouter()
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
    error.value = e.message === 'invalid credentials' ? 'Неверный email или пароль' : e.message
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="auth-page">
    <div class="auth-card card">
      <div class="auth-header">
        <div class="auth-logo">AB</div>
        <h1>Вход в AB Team</h1>
        <p>Управляйте ресторанами и меню</p>
      </div>

      <form @submit.prevent="handleSubmit">
        <div class="field">
          <label class="label">Email</label>
          <input v-model="email" type="email" class="input" placeholder="email@example.com" required />
        </div>
        <div class="field">
          <label class="label">Пароль</label>
          <input v-model="password" type="password" class="input" placeholder="Минимум 6 символов" required />
        </div>
        <div class="error-msg" v-if="error">{{ error }}</div>
        <button type="submit" class="btn btn-primary submit-btn" :disabled="loading">
          {{ loading ? 'Вход...' : 'Войти' }}
        </button>
      </form>

      <div class="auth-footer">
        Нет аккаунта? <router-link to="/register">Зарегистрироваться</router-link>
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
  width: 48px;
  height: 48px;
  background: var(--primary);
  color: #fff;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 22px;
  font-weight: 700;
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
