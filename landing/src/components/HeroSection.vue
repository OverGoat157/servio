<script setup>
const menu = [
  { name: 'Том Ям с креветками', desc: '350 мл', price: '890 ₽' },
  { name: 'Стейк Рибай', desc: '300 г, medium rare', price: '1 490 ₽' },
  { name: 'Тирамису', desc: '150 г', price: '490 ₽' },
]
</script>

<template>
  <section class="hero">
    <div class="container hero-grid">
      <div class="hero-text">
        <div class="badge">
          <span class="badge-dot" />
          Новый способ принимать заказы
        </div>

        <h1>Мини-приложение с&nbsp;меню для вашего <span class="accent">ресторана</span></h1>

        <p class="lead">
          Создайте цифровое меню за&nbsp;15&nbsp;минут. Клиенты делают заказ прямо
          со&nbsp;смартфона, а&nbsp;вы получаете его в&nbsp;WhatsApp, Telegram
          или любой другой мессенджер.
        </p>

        <div class="hero-btns">
          <a href="#contact" class="btn btn-primary">
            Создать меню бесплатно
            <svg width="20" height="20" viewBox="0 0 20 20" fill="none"><path d="M4 10h12m0 0l-4-4m4 4l-4 4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
          </a>
          <a href="#demo" class="btn btn-outline">Посмотреть демо</a>
        </div>

        <div class="stats">
          <div class="stat"><strong>500+</strong><span>Ресторанов</span></div>
          <div class="stat-line" />
          <div class="stat"><strong>50K+</strong><span>Заказов / мес</span></div>
          <div class="stat-line" />
          <div class="stat"><strong>15 мин</strong><span>На настройку</span></div>
        </div>
      </div>

      <!-- phone mockup -->
      <div class="hero-visual">
        <div class="phone">
          <div class="phone-notch" />
          <div class="screen">
            <div class="screen-head">
              <div class="rest-name">AB Team</div>
              <div class="rest-sub">Автоматизация вашего бизнеса</div>
            </div>
            <div class="screen-body">
              <div class="cat">Популярное</div>
              <div class="item" v-for="m in menu" :key="m.name">
                <div>
                  <div class="item-name">{{ m.name }}</div>
                  <div class="item-desc">{{ m.desc }}</div>
                </div>
                <div class="item-right">
                  <span class="item-price">{{ m.price }}</span>
                  <span class="item-add">+</span>
                </div>
              </div>
            </div>
            <div class="screen-cart">
              <div><small>3 позиции</small><strong>2 870 ₽</strong></div>
              <div class="cart-btn">
                <svg width="14" height="14" viewBox="0 0 16 16" fill="none"><path d="M14 2L7.5 14L5.5 8.5L2 6.5L14 2Z" fill="currentColor"/></svg>
                Отправить
              </div>
            </div>
          </div>
        </div>

      </div>
    </div>
  </section>
</template>

<style scoped>
.hero {
  padding: 120px 0 80px;
  overflow: hidden;
}

.hero-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 64px;
  align-items: center;
}

.badge {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 6px 16px;
  background: var(--primary-light);
  border: 1px solid #c5dbf5;
  border-radius: 100px;
  font-size: 14px;
  font-weight: 600;
  color: var(--primary);
  margin-bottom: 24px;
}

.badge-dot {
  width: 8px;
  height: 8px;
  background: var(--primary);
  border-radius: 50%;
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.3; }
}

h1 {
  font-size: 52px;
  font-weight: 700;
  line-height: 1.1;
  letter-spacing: -1px;
  margin-bottom: 24px;
}

.accent {
  color: var(--primary);
}

.lead {
  font-size: 18px;
  color: var(--gray);
  line-height: 1.7;
  margin-bottom: 32px;
  max-width: 520px;
}

.hero-btns {
  display: flex;
  gap: 16px;
  margin-bottom: 48px;
}

.stats {
  display: flex;
  align-items: center;
  gap: 24px;
}

.stat strong {
  display: block;
  font-size: 24px;
}

.stat span {
  font-size: 13px;
  color: var(--gray);
}

.stat-line {
  width: 1px;
  height: 40px;
  background: var(--light-gray);
}

/* phone */
.hero-visual {
  position: relative;
  display: flex;
  justify-content: center;
}

.phone {
  width: 300px;
  background: #1F2937;
  border-radius: 40px;
  padding: 12px;
  box-shadow: var(--shadow-lg);
  position: relative;
}

.phone-notch {
  width: 120px;
  height: 28px;
  background: #1F2937;
  border-radius: 0 0 20px 20px;
  position: absolute;
  top: 0;
  left: 50%;
  transform: translateX(-50%);
  z-index: 2;
}

.screen {
  background: #fff;
  border-radius: 30px;
  overflow: hidden;
  min-height: 520px;
  display: flex;
  flex-direction: column;
}

.screen-head {
  background: linear-gradient(135deg, #1F2937 0%, #374151 100%);
  color: #fff;
  padding: 48px 20px 20px;
  text-align: center;
}

.rest-name {
  font-size: 28px;
  font-weight: 700;
  letter-spacing: 2px;
}

.rest-sub {
  font-size: 12px;
  opacity: 0.8;
  margin-top: 4px;
}

.screen-body {
  flex: 1;
  padding: 16px;
}

.cat {
  font-size: 13px;
  font-weight: 600;
  color: var(--gray);
  text-transform: uppercase;
  letter-spacing: 1px;
  margin-bottom: 12px;
}

.item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 0;
  border-bottom: 1px solid var(--light-gray);
}

.item:last-child {
  border-bottom: none;
}

.item-name {
  font-size: 14px;
  font-weight: 600;
  margin-bottom: 2px;
}

.item-desc {
  font-size: 12px;
  color: var(--gray);
}

.item-right {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-shrink: 0;
}

.item-price {
  font-size: 14px;
  font-weight: 600;
}

.item-add {
  width: 28px;
  height: 28px;
  border-radius: 8px;
  background: var(--primary);
  color: #fff;
  font-size: 16px;
  font-weight: 600;
  display: flex;
  align-items: center;
  justify-content: center;
}

.screen-cart {
  background: var(--primary);
  margin: 8px 12px 12px;
  border-radius: 16px;
  padding: 14px 16px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: #fff;
}

.screen-cart small {
  display: block;
  font-size: 12px;
  opacity: 0.8;
}

.screen-cart strong {
  font-size: 16px;
}

.cart-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  font-weight: 600;
  background: rgba(255, 255, 255, 0.2);
  padding: 8px 12px;
  border-radius: 10px;
}


@media (max-width: 1024px) {
  .hero { padding: 100px 0 60px; }

  .hero-grid {
    grid-template-columns: 1fr;
    gap: 48px;
    text-align: center;
  }

  h1 { font-size: 36px; }

  .lead {
    margin-left: auto;
    margin-right: auto;
  }

  .hero-btns { justify-content: center; }
  .stats { justify-content: center; }
  .hero-visual { order: -1; }

  .phone {
    width: 260px;
  }

  .screen {
    min-height: 440px;
  }
}

@media (max-width: 480px) {
  h1 { font-size: 30px; }

  .hero-btns {
    flex-direction: column;
    align-items: center;
  }

  .stats {
    flex-direction: column;
    gap: 16px;
  }

  .stat-line {
    width: 40px;
    height: 1px;
  }

}
</style>
