const API_BASE = import.meta.env.VITE_API_URL || 'http://localhost:8080'

export async function fetchMenu(slug) {
  const res = await fetch(`${API_BASE}/api/public/menu/${slug}`)
  if (!res.ok) throw new Error('Restaurant not found')
  return res.json()
}

export async function createOrder(slug, order) {
  const res = await fetch(`${API_BASE}/api/public/menu/${slug}/order`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(order),
  })
  if (!res.ok) throw new Error('Failed to create order')
  return res.json()
}
