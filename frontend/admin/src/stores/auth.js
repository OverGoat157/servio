import { ref } from 'vue'
import { auth as authApi } from '../api/client'

export const user = ref(null)
export const isAuthenticated = ref(!!localStorage.getItem('token'))

const TIER_RANK = { basic: 0, business: 1, business_max: 2 }

export const TIER_LABELS = {
  basic: 'Базовый',
  business: 'Бизнес',
  business_max: 'Бизнес Макс',
}

export function tierLabel(tier) {
  return TIER_LABELS[tier] || 'Базовый'
}

// Returns true if the current user's tier is >= minTier.
// Admins always pass — they see everything.
export function canUseFeature(minTier) {
  if (!user.value) return false
  if (user.value.role === 'admin') return true
  const current = TIER_RANK[user.value.tier] ?? 0
  const required = TIER_RANK[minTier] ?? 0
  return current >= required
}

export async function login(email, password) {
  const data = await authApi.login({ email, password })
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
