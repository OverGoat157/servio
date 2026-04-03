<script setup>
import { ref, reactive } from 'vue'

const sent = ref(false)
const form = reactive({ name: '', phone: '', restaurant: '', messenger: '' })

function submit() {
  sent.value = true
  setTimeout(() => {
    sent.value = false
    Object.assign(form, { name: '', phone: '', restaurant: '', messenger: '' })
  }, 3000)
}
</script>

<template>
  <section id="contact" class="cta">
    <div class="container">
      <div class="cta-card">
        <div class="cta-text">
          <h2>Готовы запустить цифровое меню?</h2>
          <p>Оставьте заявку и мы настроим мини-приложение для вашего ресторана за 15 минут. Первый месяц — бесплатно.</p>
        </div>
        <form class="form" @submit.prevent="submit">
          <div class="form-row">
            <input v-model="form.name" type="text" placeholder="Ваше имя" required />
            <input v-model="form.phone" type="tel" placeholder="+7 (___) ___-__-__" required />
          </div>
          <div class="form-row">
            <input v-model="form.restaurant" type="text" placeholder="Название ресторана" />
            <select v-model="form.messenger">
              <option value="" disabled>Предпочитаемый мессенджер</option>
              <option value="whatsapp">WhatsApp</option>
              <option value="telegram">Telegram</option>
              <option value="viber">Viber</option>
              <option value="other">Другой</option>
            </select>
          </div>
          <button type="submit" class="btn btn-white" :disabled="sent">
            {{ sent ? 'Заявка отправлена!' : 'Оставить заявку' }}
            <svg v-if="!sent" width="20" height="20" viewBox="0 0 20 20" fill="none"><path d="M4 10h12m0 0l-4-4m4 4l-4 4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
          </button>
        </form>
      </div>
    </div>
  </section>
</template>

<style scoped>
.cta {
  padding: 100px 0;
  background: var(--section-bg);
}

.cta-card {
  background: linear-gradient(135deg, #374151, #1F2937);
  border-radius: var(--radius-lg);
  padding: 64px;
  color: #fff;
}

.cta-text {
  text-align: center;
  margin-bottom: 40px;
}

.cta-text h2 {
  font-size: 36px;
  font-weight: 700;
  margin-bottom: 16px;
  color: #fff;
}

.cta-text p {
  font-size: 18px;
  opacity: 0.75;
  max-width: 520px;
  margin: 0 auto;
  line-height: 1.6;
}

.form {
  max-width: 560px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.form-row {
  display: flex;
  gap: 16px;
}

.form input,
.form select {
  flex: 1;
  padding: 14px 18px;
  border: 2px solid rgba(255, 255, 255, 0.2);
  border-radius: var(--radius);
  background: rgba(255, 255, 255, 0.08);
  color: #fff;
  font-family: inherit;
  font-size: 15px;
  transition: border-color var(--ease);
  outline: none;
}

.form input::placeholder {
  color: rgba(255, 255, 255, 0.45);
}

.form input:focus,
.form select:focus {
  border-color: var(--primary);
}

.form select {
  appearance: none;
  cursor: pointer;
  color: rgba(255, 255, 255, 0.45);
}

.form select option {
  color: var(--black);
  background: #fff;
}

.btn-white {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  width: 100%;
  padding: 16px 32px;
  border-radius: var(--radius);
  font-family: inherit;
  font-size: 17px;
  font-weight: 600;
  border: none;
  cursor: pointer;
  background: #fff;
  color: var(--black);
  transition: all var(--ease);
}

.btn-white:hover {
  background: var(--off-white);
  transform: translateY(-2px);
  box-shadow: 0 8px 24px rgba(255, 255, 255, 0.15);
}

.btn-white:disabled {
  background: rgba(255, 255, 255, 0.2);
  color: #fff;
  cursor: default;
  transform: none;
  box-shadow: none;
}

@media (max-width: 768px) {
  .cta {
    padding: 64px 0;
  }

  .cta-card {
    padding: 40px 24px;
  }

  .cta-text h2 {
    font-size: 26px;
  }

  .form-row {
    flex-direction: column;
  }
}
</style>
