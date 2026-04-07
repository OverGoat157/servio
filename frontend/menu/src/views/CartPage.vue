<script setup>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { restaurant } from '../stores/restaurant'
import { cart, cartTotal, updateQuantity, removeFromCart, clearCart } from '../stores/cart'
import { createOrder } from '../api/client'

const route = useRoute()
const router = useRouter()
const slug = route.params.slug

const customerName = ref('')
const customerPhone = ref('')
const selectedMessenger = ref('')
const sending = ref(false)
const sent = ref(false)
const orderError = ref(null)

function formatPrice(kopecks) {
  return Math.floor(kopecks / 100).toLocaleString('ru-RU') + ' \u20BD'
}

function goBack() {
  router.push({ name: 'menu', params: { slug } })
}

async function submitOrder() {
  if (!selectedMessenger.value || cart.length === 0) return

  sending.value = true
  orderError.value = null

  try {
    const result = await createOrder(slug, {
      items: cart.map(i => ({
        menu_item_id: i.id,
        quantity: i.quantity,
      })),
      messenger: selectedMessenger.value,
      customer_name: customerName.value || undefined,
      customer_phone: customerPhone.value || undefined,
    })

    sent.value = true
    clearCart()

    if (result.whatsapp_url) {
      window.open(result.whatsapp_url, '_blank')
    }
  } catch {
    orderError.value = 'Не удалось отправить заказ. Попробуйте ещё раз.'
  } finally {
    sending.value = false
  }
}

function goHome() {
  router.push({ name: 'home', params: { slug } })
}
</script>

<template>
  <div class="cart-page">
    <!-- Header -->
    <div class="header">
      <button class="back-btn" @click="goBack">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
          <path d="M19 12H5M12 19l-7-7 7-7"/>
        </svg>
      </button>
      <h2>Ваш заказ</h2>
      <div class="spacer"></div>
    </div>

    <!-- Success -->
    <div v-if="sent" class="success-screen">
      <div class="success-icon">
        <svg width="36" height="36" viewBox="0 0 24 24" fill="none" stroke="#fff" stroke-width="2.5" stroke-linecap="round">
          <path d="M20 6L9 17l-5-5"/>
        </svg>
      </div>
      <h3>Заказ оформлен!</h3>
      <p v-if="selectedMessenger === 'telegram'">Заказ отправлен в Telegram ресторана</p>
      <p v-else-if="selectedMessenger === 'whatsapp'">Отправьте сообщение в открывшемся WhatsApp</p>
      <p v-else>Ресторан получил ваш заказ</p>
      <button class="btn-primary" @click="goHome">На главную</button>
    </div>

    <!-- Empty cart -->
    <div v-else-if="cart.length === 0" class="empty-cart">
      <div class="empty-icon">
        <svg width="36" height="36" viewBox="0 0 24 24" fill="none" stroke="var(--text-muted)" stroke-width="1.5" stroke-linecap="round">
          <circle cx="9" cy="21" r="1"/><circle cx="20" cy="21" r="1"/>
          <path d="M1 1h4l2.68 13.39a2 2 0 002 1.61h9.72a2 2 0 002-1.61L23 6H6"/>
        </svg>
      </div>
      <p class="empty-title">Корзина пуста</p>
      <button class="btn-outline" @click="goBack">Перейти в меню</button>
    </div>

    <!-- Cart items -->
    <div v-else class="cart-content">
      <!-- Items -->
      <div class="items-section">
        <div class="cart-item" v-for="item in cart" :key="item.id">
          <div class="ci-top">
            <div class="ci-name">{{ item.name }}</div>
            <button class="ci-remove" @click="removeFromCart(item.id)">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
                <path d="M18 6L6 18M6 6l12 12"/>
              </svg>
            </button>
          </div>
          <div class="ci-bottom">
            <div class="qty-controls">
              <button class="qty-btn" @click="updateQuantity(item.id, item.quantity - 1)">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"><path d="M5 12h14"/></svg>
              </button>
              <span class="qty">{{ item.quantity }}</span>
              <button class="qty-btn" @click="updateQuantity(item.id, item.quantity + 1)">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"><path d="M12 5v14M5 12h14"/></svg>
              </button>
            </div>
            <div class="ci-price">{{ formatPrice(item.price * item.quantity) }}</div>
          </div>
        </div>
      </div>

      <!-- Total -->
      <div class="total-row">
        <span>Итого</span>
        <strong>{{ formatPrice(cartTotal) }}</strong>
      </div>

      <!-- Customer info -->
      <div class="form-section">
        <div class="section-label">Ваши данные (необязательно)</div>
        <div class="form-fields">
          <input
            v-model="customerName"
            class="field-input"
            placeholder="Имя"
          />
          <input
            v-model="customerPhone"
            class="field-input"
            type="tel"
            placeholder="Телефон"
          />
        </div>
      </div>

      <!-- Messenger selection -->
      <div class="form-section">
        <div class="section-label">Куда отправить заказ?</div>
        <div class="msg-options">
          <button
            v-for="m in (restaurant.messengers?.length ? restaurant.messengers : ['whatsapp', 'telegram'])"
            :key="m"
            class="msg-option"
            :class="{ active: selectedMessenger === m }"
            @click="selectedMessenger = m"
          >
            <svg v-if="m === 'whatsapp'" width="22" height="22" viewBox="0 0 24 24" :fill="selectedMessenger === m ? '#25D366' : '#c4c4c4'"><path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347z"/><path d="M12 0C5.373 0 0 5.373 0 12c0 2.127.556 4.12 1.525 5.855L0 24l6.335-1.652A11.94 11.94 0 0012 24c6.627 0 12-5.373 12-12S18.627 0 12 0z"/></svg>
            <svg v-else-if="m === 'telegram'" width="22" height="22" viewBox="0 0 24 24" :fill="selectedMessenger === m ? '#229ED9' : '#c4c4c4'"><path d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12 12-5.373 12-12S18.627 0 12 0zm5.562 8.161c-.18 1.897-.962 6.502-1.359 8.627-.168.9-.5 1.201-.82 1.23-.697.064-1.226-.46-1.9-.903-1.056-.692-1.653-1.123-2.678-1.799-1.185-.781-.417-1.21.258-1.911.177-.184 3.247-2.977 3.307-3.23.007-.032.014-.15-.056-.212s-.174-.041-.249-.024c-.106.024-1.793 1.14-5.061 3.345-.479.33-.913.492-1.302.484-.428-.008-1.252-.241-1.865-.44-.752-.245-1.349-.374-1.297-.789.027-.216.325-.437.893-.663 3.498-1.524 5.831-2.529 6.998-3.014 3.332-1.386 4.025-1.627 4.476-1.635.099-.002.321.023.465.141a.506.506 0 01.171.325c.016.093.036.306.02.472z"/></svg>
            <span>{{ m === 'whatsapp' ? 'WhatsApp' : m === 'telegram' ? 'Telegram' : m === 'viber' ? 'Viber' : m }}</span>
          </button>
        </div>
      </div>

      <!-- Error -->
      <div class="order-error" v-if="orderError">{{ orderError }}</div>

      <!-- Submit -->
      <button
        class="submit-btn"
        :disabled="!selectedMessenger || sending"
        @click="submitOrder"
      >
        {{ sending ? 'Отправка...' : 'Отправить заказ' }}
      </button>
    </div>
  </div>
</template>

<style scoped>
.cart-page {
  min-height: 100vh;
}

/* Header */
.header {
  display: flex;
  align-items: center;
  padding: 14px 16px;
  gap: 12px;
  position: sticky;
  top: 0;
  background: rgba(255,255,255,0.92);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  z-index: 20;
  border-bottom: 1px solid var(--border);
}

.header h2 {
  font-size: 17px;
  font-weight: 700;
  flex: 1;
  text-align: center;
}

.back-btn {
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  color: var(--text);
  transition: background var(--ease);
}

.back-btn:active {
  background: var(--bg-secondary);
}

.spacer {
  width: 36px;
}

/* Success */
.success-screen {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 70vh;
  padding: 24px;
  gap: 14px;
  text-align: center;
}

.success-icon {
  width: 72px;
  height: 72px;
  border-radius: 50%;
  background: var(--success);
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 8px;
}

.success-screen h3 {
  font-size: 22px;
  font-weight: 700;
}

.success-screen p {
  color: var(--text-secondary);
  font-size: 14px;
  line-height: 1.5;
}

.success-screen .btn-primary {
  margin-top: 16px;
  padding: 14px 36px;
  background: var(--primary);
  color: var(--primary-foreground);
  border-radius: var(--radius);
  font-size: 15px;
  font-weight: 600;
}

/* Empty */
.empty-cart {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 60vh;
  gap: 12px;
  padding: 24px;
}

.empty-icon {
  width: 64px;
  height: 64px;
  border-radius: 50%;
  background: var(--bg-secondary);
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 8px;
}

.empty-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-secondary);
}

.empty-cart .btn-outline {
  margin-top: 8px;
  padding: 12px 28px;
  border: 1.5px solid var(--border);
  border-radius: var(--radius);
  font-size: 14px;
  font-weight: 600;
  color: var(--text);
  transition: border-color var(--ease);
}

.empty-cart .btn-outline:active {
  border-color: var(--primary);
}

/* Cart content */
.cart-content {
  padding: 16px;
}

/* Items */
.items-section {
  display: flex;
  flex-direction: column;
  gap: 2px;
  background: var(--bg-secondary);
  border-radius: var(--radius);
  overflow: hidden;
  margin-bottom: 16px;
}

.cart-item {
  padding: 14px 16px;
  background: var(--bg);
}

.ci-top {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 12px;
  margin-bottom: 10px;
}

.ci-name {
  font-size: 15px;
  font-weight: 600;
  line-height: 1.3;
}

.ci-remove {
  width: 28px;
  height: 28px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text-muted);
  flex-shrink: 0;
  transition: all var(--ease);
}

.ci-remove:active {
  background: #fee2e2;
  color: var(--danger);
}

.ci-bottom {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.qty-controls {
  display: flex;
  align-items: center;
  gap: 2px;
  background: var(--bg-secondary);
  border-radius: 10px;
  padding: 2px;
}

.qty-btn {
  width: 34px;
  height: 34px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text);
  transition: background var(--ease);
}

.qty-btn:active {
  background: var(--border);
}

.qty {
  font-size: 15px;
  font-weight: 600;
  min-width: 28px;
  text-align: center;
  font-variant-numeric: tabular-nums;
}

.ci-price {
  font-size: 16px;
  font-weight: 700;
  font-variant-numeric: tabular-nums;
}

/* Total */
.total-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 18px 0;
  font-size: 17px;
  border-bottom: 1px solid var(--border);
}

.total-row span {
  color: var(--text-secondary);
  font-weight: 500;
}

.total-row strong {
  font-size: 20px;
  font-weight: 700;
  font-variant-numeric: tabular-nums;
}

/* Form sections */
.form-section {
  padding: 20px 0;
  border-bottom: 1px solid var(--border);
}

.section-label {
  font-size: 11px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.8px;
  color: var(--text-muted);
  margin-bottom: 12px;
}

.form-fields {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.field-input {
  width: 100%;
  padding: 14px 16px;
  border: 1.5px solid var(--border);
  border-radius: var(--radius-sm);
  font-size: 15px;
  color: var(--text);
  background: var(--bg);
  transition: border-color var(--ease);
  outline: none;
}

.field-input:focus {
  border-color: var(--primary);
}

.field-input::placeholder {
  color: var(--text-muted);
}

/* Messenger options */
.msg-options {
  display: flex;
  gap: 10px;
}

.msg-option {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  padding: 16px 12px;
  border: 1.5px solid var(--border);
  border-radius: var(--radius);
  font-size: 13px;
  font-weight: 500;
  color: var(--text-secondary);
  transition: all var(--ease);
}

.msg-option.active {
  border-color: var(--primary);
  color: var(--text);
  background: var(--bg-secondary);
}

/* Error */
.order-error {
  padding: 12px 16px;
  background: #fee2e2;
  color: var(--danger);
  border-radius: var(--radius-sm);
  font-size: 14px;
  margin-top: 16px;
  text-align: center;
}

/* Submit */
.submit-btn {
  width: 100%;
  padding: 16px;
  margin-top: 20px;
  background: var(--primary);
  color: var(--primary-foreground);
  border-radius: var(--radius);
  font-size: 16px;
  font-weight: 600;
  transition: opacity var(--ease);
}

.submit-btn:disabled {
  opacity: 0.35;
  cursor: not-allowed;
}

.submit-btn:not(:disabled):active {
  opacity: 0.85;
}
</style>
