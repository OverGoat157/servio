import { ref } from 'vue'
import { auth as authApi } from '../api/client'

export const user = ref(null)
export const isAuthenticated = ref(!!localStorage.getItem('token'))

export async function login(email, password) {
  const data = await authApi.login({ email, password })
  localStorage.setItem('token', data.token)
  user.value = data.user
  isAuthenticated.value = true
  return data
}

export async function register(email, password, name) {
  const data = await authApi.register({ email, password, name })
  localStorage.setItem('token', data.token)
  user.value = data.user
  isAuthenticated.value = true
  return data
}

export async function fetchUser() {
  try {
    user.value = await authApi.me()
    isAuthenticated.value = true
  } catch {
    logout()
  }
}

export function logout() {
  localStorage.removeItem('token')
  user.value = null
  isAuthenticated.value = false
}
