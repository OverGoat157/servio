const API_BASE = import.meta.env.VITE_API_URL || 'http://localhost:8080'

export function imageUrl(path) {
  if (!path) return ''
  if (path.startsWith('http')) return path
  return API_BASE + path
}

function getToken() {
  return localStorage.getItem('token')
}

async function request(path, options = {}) {
  const headers = { 'Content-Type': 'application/json', ...options.headers }
  const token = getToken()
  if (token) {
    headers['Authorization'] = `Bearer ${token}`
  }
  const res = await fetch(`${API_BASE}${path}`, { ...options, headers })
  if (res.status === 401) {
    localStorage.removeItem('token')
    window.location.href = '/login'
    throw new Error('Unauthorized')
  }
  const data = await res.json()
  if (!res.ok) throw new Error(data.error || 'Request failed')
  return data
}

// Upload
export async function uploadFile(file) {
  const form = new FormData()
  form.append('file', file)
  const token = getToken()
  const res = await fetch(`${API_BASE}/api/upload`, {
    method: 'POST',
    headers: token ? { Authorization: `Bearer ${token}` } : {},
    body: form,
  })
  if (!res.ok) {
    const data = await res.json()
    throw new Error(data.error || 'Upload failed')
  }
  return res.json()
}

// Auth
export const auth = {
  login: (body) => request('/api/auth/login', { method: 'POST', body: JSON.stringify(body) }),
  me: () => request('/api/me'),
}

// Admin
export const admin = {
  listUsers: () => request('/api/admin/users'),
  getUser: (id) => request(`/api/admin/users/${id}`),
  createUser: (body) => request('/api/admin/users', { method: 'POST', body: JSON.stringify(body) }),
  updateUser: (id, body) => request(`/api/admin/users/${id}`, { method: 'PUT', body: JSON.stringify(body) }),
  deleteUser: (id) => request(`/api/admin/users/${id}`, { method: 'DELETE' }),
  listRestaurants: () => request('/api/admin/restaurants'),
  updateRestaurant: (id, body) => request(`/api/admin/restaurants/${id}`, { method: 'PUT', body: JSON.stringify(body) }),
}

// Restaurants
export const restaurants = {
  list: () => request('/api/restaurants'),
  get: (id) => request(`/api/restaurants/${id}`),
  create: (body) => request('/api/restaurants', { method: 'POST', body: JSON.stringify(body) }),
  update: (id, body) => request(`/api/restaurants/${id}`, { method: 'PUT', body: JSON.stringify(body) }),
  delete: (id) => request(`/api/restaurants/${id}`, { method: 'DELETE' }),
}

// Categories
export const categories = {
  list: (restaurantId) => request(`/api/restaurants/${restaurantId}/categories`),
  create: (restaurantId, body) => request(`/api/restaurants/${restaurantId}/categories`, { method: 'POST', body: JSON.stringify(body) }),
  update: (id, body) => request(`/api/categories/${id}`, { method: 'PUT', body: JSON.stringify(body) }),
  delete: (id) => request(`/api/categories/${id}`, { method: 'DELETE' }),
}

// Menu items
export const menuItems = {
  list: (categoryId) => request(`/api/categories/${categoryId}/items`),
  create: (categoryId, body) => request(`/api/categories/${categoryId}/items`, { method: 'POST', body: JSON.stringify(body) }),
  update: (id, body) => request(`/api/items/${id}`, { method: 'PUT', body: JSON.stringify(body) }),
  delete: (id) => request(`/api/items/${id}`, { method: 'DELETE' }),
}

// Orders
export const orders = {
  list: (restaurantId, limit = 50, offset = 0) => request(`/api/restaurants/${restaurantId}/orders?limit=${limit}&offset=${offset}`),
  updateStatus: (id, status) => request(`/api/orders/${id}/status`, { method: 'PATCH', body: JSON.stringify({ status }) }),
}

// Messengers
export const messengers = {
  list: (restaurantId) => request(`/api/restaurants/${restaurantId}/messengers`),
  upsert: (restaurantId, body) => request(`/api/restaurants/${restaurantId}/messengers`, { method: 'POST', body: JSON.stringify(body) }),
  delete: (restaurantId, type) => request(`/api/restaurants/${restaurantId}/messengers/${type}`, { method: 'DELETE' }),
}
