<script setup>
import { ref, onMounted } from 'vue'
import { admin as api } from '../api/client'

const users = ref([])
const loading = ref(true)

const showForm = ref(false)
const editingId = ref(null)
const form = ref({ email: '', password: '', name: '', role: 'user', max_restaurants: 3 })
const saving = ref(false)
const error = ref('')

onMounted(async () => {
  await loadUsers()
})

async function loadUsers() {
  loading.value = true
  try {
    users.value = await api.listUsers() || []
  } catch { /* empty */ }
  loading.value = false
}

function openCreate() {
  editingId.value = null
  form.value = { email: '', password: '', name: '', role: 'user', max_restaurants: 3 }
  error.value = ''
  showForm.value = true
}

function openEdit(u) {
  editingId.value = u.id
  form.value = {
    email: u.email,
    password: '',
    name: u.name,
    role: u.role,
    max_restaurants: u.max_restaurants,
  }
  error.value = ''
  showForm.value = true
}

async function saveUser() {
  saving.value = true
  error.value = ''
  try {
    if (editingId.value) {
      const data = { ...form.value }
      if (!data.password) delete data.password
      await api.updateUser(editingId.value, data)
    } else {
      await api.createUser(form.value)
    }
    showForm.value = false
    await loadUsers()
  } catch (e) {
    error.value = e.message
  }
  saving.value = false
}

async function deleteUser(u) {
  if (!confirm(`Удалить пользователя ${u.name} (${u.email})? Все его рестораны и данные будут удалены!`)) return
  try {
    await api.deleteUser(u.id)
    await loadUsers()
  } catch { /* empty */ }
}

function formatDate(d) {
  return new Date(d).toLocaleDateString('ru-RU', { day: 'numeric', month: 'short', year: 'numeric' })
}
</script>

<template>
  <div class="page">
    <div class="page-header">
      <h1>Пользователи</h1>
      <button class="btn btn-primary" @click="openCreate">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"><path d="M12 5v14M5 12h14"/></svg>
        Создать
      </button>
    </div>

    <div class="loading" v-if="loading">Загрузка...</div>

    <div class="table-wrap" v-else>
      <table class="table">
        <thead>
          <tr>
            <th>ID</th>
            <th>Имя</th>
            <th>Email</th>
            <th>Роль</th>
            <th>Лимит</th>
            <th>Дата</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="u in users" :key="u.id">
            <td class="td-id">{{ u.id }}</td>
            <td class="td-name">{{ u.name }}</td>
            <td>{{ u.email }}</td>
            <td><span class="role-badge" :class="u.role">{{ u.role }}</span></td>
            <td>{{ u.max_restaurants }}</td>
            <td class="td-date">{{ formatDate(u.created_at) }}</td>
            <td class="td-actions">
              <button class="icon-btn" title="Редактировать" @click="openEdit(u)">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M11 4H4a2 2 0 00-2 2v14a2 2 0 002 2h14a2 2 0 002-2v-7"/><path d="M18.5 2.5a2.121 2.121 0 013 3L12 15l-4 1 1-4 9.5-9.5z"/></svg>
              </button>
              <button class="icon-btn danger" title="Удалить" @click="deleteUser(u)" v-if="u.role !== 'admin'">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M3 6h18M19 6v14a2 2 0 01-2 2H7a2 2 0 01-2-2V6m3 0V4a2 2 0 012-2h4a2 2 0 012 2v2"/></svg>
              </button>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="empty" v-if="users.length === 0">Нет пользователей</div>
    </div>

    <!-- Modal -->
    <div class="modal-overlay" v-if="showForm" @click.self="showForm = false">
      <div class="modal card">
        <h2>{{ editingId ? 'Редактировать пользователя' : 'Новый пользователь' }}</h2>
        <form @submit.prevent="saveUser">
          <div class="field">
            <label class="label">Имя</label>
            <input v-model="form.name" class="input" required placeholder="Иван Петров" />
          </div>
          <div class="field">
            <label class="label">Email</label>
            <input v-model="form.email" type="email" class="input" required placeholder="user@example.com" />
          </div>
          <div class="field">
            <label class="label">{{ editingId ? 'Новый пароль (оставьте пустым)' : 'Пароль' }}</label>
            <input v-model="form.password" type="password" class="input" :required="!editingId" minlength="6" placeholder="Минимум 6 символов" />
          </div>
          <div class="row">
            <div class="field">
              <label class="label">Роль</label>
              <select v-model="form.role" class="input">
                <option value="user">user</option>
                <option value="admin">admin</option>
              </select>
            </div>
            <div class="field">
              <label class="label">Макс. ресторанов</label>
              <input v-model.number="form.max_restaurants" type="number" min="0" class="input" />
            </div>
          </div>
          <div class="error-msg" v-if="error">{{ error }}</div>
          <div class="modal-actions">
            <button type="button" class="btn btn-outline" @click="showForm = false">Отмена</button>
            <button type="submit" class="btn btn-primary" :disabled="saving">
              {{ saving ? 'Сохранение...' : 'Сохранить' }}
            </button>
          </div>
        </form>
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

.loading {
  text-align: center;
  padding: 60px 0;
  color: var(--text-secondary);
}

.table-wrap {
  overflow-x: auto;
}

.table {
  width: 100%;
  border-collapse: collapse;
  font-size: 14px;
}

.table th {
  text-align: left;
  padding: 10px 12px;
  font-size: 12px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: var(--text-secondary);
  border-bottom: 2px solid var(--border);
}

.table td {
  padding: 12px;
  border-bottom: 1px solid var(--border);
}

.td-id {
  color: var(--text-secondary);
  font-size: 13px;
}

.td-name {
  font-weight: 600;
}

.td-date {
  color: var(--text-secondary);
  font-size: 13px;
  white-space: nowrap;
}

.td-actions {
  display: flex;
  gap: 4px;
  justify-content: flex-end;
}

.role-badge {
  padding: 2px 10px;
  border-radius: 100px;
  font-size: 12px;
  font-weight: 600;
}

.role-badge.admin {
  background: #EDE9FE;
  color: #7C3AED;
}

.role-badge.user {
  background: var(--primary-light);
  color: var(--primary);
}

.icon-btn {
  width: 30px;
  height: 30px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text-secondary);
}

.icon-btn:hover {
  background: var(--bg);
}

.icon-btn.danger:hover {
  background: #FEE2E2;
  color: var(--danger);
}

.empty {
  text-align: center;
  padding: 40px;
  color: var(--text-secondary);
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
  max-width: 480px;
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
</style>
