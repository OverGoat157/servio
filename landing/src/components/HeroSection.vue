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

        <div class="float-card fc-wa">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="#333"><path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347z"/><path d="M12 0C5.373 0 0 5.373 0 12c0 2.127.556 4.12 1.525 5.855L0 24l6.335-1.652A11.94 11.94 0 0012 24c6.627 0 12-5.373 12-12S18.627 0 12 0z"/></svg>
          Новый заказ!
        </div>

        <div class="float-card fc-tg">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="#555"><path d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12 12-5.373 12-12S18.627 0 12 0zm5.562 8.161c-.18 1.897-.962 6.502-1.359 8.627-.168.9-.5 1.201-.82 1.23-.697.064-1.226-.46-1.9-.903-1.056-.692-1.653-1.123-2.678-1.799-1.185-.781-.417-1.21.258-1.911.177-.184 3.247-2.977 3.307-3.23.007-.032.014-.15-.056-.212s-.174-.041-.249-.024c-.106.024-1.793 1.14-5.061 3.345-.479.33-.913.492-1.302.484-.428-.008-1.252-.241-1.865-.44-.752-.245-1.349-.374-1.297-.789.027-.216.325-.437.893-.663 3.498-1.524 5.831-2.529 6.998-3.014 3.332-1.386 4.025-1.627 4.476-1.635.099-.002.321.023.465.141a.506.506 0 01.171.325c.016.093.036.306.02.472z"/></svg>
          +1 280 ₽
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

/* floating cards */
.float-card {
  position: absolute;
  display: flex;
  align-items: center;
  gap: 10px;
  background: #fff;
  padding: 12px 18px;
  border-radius: 14px;
  box-shadow: var(--shadow-lg);
  font-size: 14px;
  font-weight: 600;
  animation: float 3s ease-in-out infinite;
}

.fc-wa {
  top: 15%;
  right: -20px;
}

.fc-tg {
  bottom: 20%;
  left: -10px;
  animation-delay: 1.5s;
}

@keyframes float {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-8px); }
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

  .fc-wa { right: 5%; }
  .fc-tg { left: 5%; }
}
</style>
