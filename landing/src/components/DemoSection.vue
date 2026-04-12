<script setup>
// TODO: положить реальные скриншоты в landing/public/demo/
const screens = [
  { src: '/demo/home.png', label: 'Главная страница' },
  { src: '/demo/menu.png', label: 'Меню ресторана' },
  { src: '/demo/cart.png', label: 'Отправка заказа' },
]
</script>

<template>
  <section id="demo" class="demo">
    <div class="container">
      <div class="text-center">
        <h2 class="section-title">Посмотрите как это выглядит</h2>
        <p class="section-subtitle">
          Так видит ваше меню клиент. Минималистичный дизайн и заказ в пару нажатий.
        </p>
      </div>

      <div class="phones">
        <div class="phone-col" v-for="s in screens" :key="s.label">
          <div class="screenshot">
            <img :src="s.src" :alt="s.label" @error="$event.target.classList.add('missing')" />
            <div class="screenshot-placeholder">
              <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/><circle cx="8.5" cy="8.5" r="1.5"/><polyline points="21 15 16 10 5 21"/>
              </svg>
              <span>Скриншот</span>
            </div>
          </div>
          <div class="phone-label">{{ s.label }}</div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.demo {
  padding: 100px 0;
  background: var(--section-bg);
}

.text-center {
  text-align: center;
}

.phones {
  display: flex;
  justify-content: center;
  gap: 40px;
  flex-wrap: wrap;
}

.phone-col {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
}

.screenshot {
  position: relative;
  width: 260px;
  aspect-ratio: 9 / 18;
  border-radius: 24px;
  overflow: hidden;
  background: var(--off-white);
  border: 1px solid var(--light-gray);
  box-shadow: var(--shadow-lg);
}

.screenshot img {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  z-index: 1;
}

.screenshot img.missing {
  display: none;
}

.screenshot-placeholder {
  position: absolute;
  inset: 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 10px;
  color: var(--gray);
  font-size: 13px;
  font-weight: 500;
  z-index: 0;
}

.phone-label {
  font-size: 14px;
  font-weight: 600;
  color: var(--gray);
}

@media (max-width: 768px) {
  .demo {
    padding: 64px 0;
  }

  .phones {
    gap: 32px;
  }

  .screenshot {
    width: 220px;
  }
}
</style>
