import { reactive, computed } from 'vue'

export const cart = reactive([])

export const cartTotal = computed(() =>
  cart.reduce((sum, item) => sum + item.price * item.quantity, 0)
)

export const cartCount = computed(() =>
  cart.reduce((sum, item) => sum + item.quantity, 0)
)

export function addToCart(menuItem) {
  const existing = cart.find(i => i.id === menuItem.id)
  if (existing) {
    existing.quantity++
  } else {
    cart.push({
      id: menuItem.id,
      name: menuItem.name,
      price: menuItem.price,
      quantity: 1,
    })
  }
}

export function removeFromCart(itemId) {
  const idx = cart.findIndex(i => i.id === itemId)
  if (idx !== -1) cart.splice(idx, 1)
}

export function updateQuantity(itemId, qty) {
  const item = cart.find(i => i.id === itemId)
  if (!item) return
  if (qty <= 0) {
    removeFromCart(itemId)
  } else {
    item.quantity = qty
  }
}

export function clearCart() {
  cart.splice(0, cart.length)
}
