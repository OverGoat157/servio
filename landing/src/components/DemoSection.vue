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
        <!-- Screen 1: Main page -->
        <div class="phone-col">
          <div class="phone-shell">
            <div class="phone-bar"><span>12:30</span></div>
            <div class="phone-body" style="background: linear-gradient(180deg, #1F2937 0%, #374151 50%, #fff 50%)">
              <div class="s1-header">
                <div class="s1-name">AB Team</div>
                <div class="s1-sub">Автоматизация вашего бизнеса</div>
              </div>
              <div class="s1-card">
                <div class="s1-label">РЕЖИМ РАБОТЫ</div>
                <div class="s1-row"><span>Пн — Чт</span><strong>11:00 — 00:00</strong></div>
                <div class="s1-row"><span>Пятница</span><strong>13:30 — 00:00</strong></div>
                <div class="s1-row last"><span>Сб — Вс</span><strong>11:00 — 00:00</strong></div>
              </div>
              <div class="s1-btn">Связаться с нами</div>
            </div>
          </div>
          <div class="phone-label">Главная страница</div>
        </div>

        <!-- Screen 2: Menu -->
        <div class="phone-col">
          <div class="phone-shell">
            <div class="phone-bar"><span>12:30</span></div>
            <div class="phone-body">
              <div class="s2-tabs">
                <span class="s2-tab active">Все</span>
                <span class="s2-tab">Супы</span>
                <span class="s2-tab">Стейки</span>
                <span class="s2-tab">Десерты</span>
              </div>
              <div class="s2-items">
                <div class="s2-item">
                  <div><div class="s2-name">Том Ям</div><div class="s2-desc">с креветками, 350 мл</div></div>
                  <div class="s2-right"><div class="s2-price">890 ₽</div><div class="s2-add">+</div></div>
                </div>
                <div class="s2-item">
                  <div><div class="s2-name">Стейк Рибай</div><div class="s2-desc">300 г, medium rare</div></div>
                  <div class="s2-right"><div class="s2-price">1 490 ₽</div><div class="s2-add">+</div></div>
                </div>
                <div class="s2-item">
                  <div><div class="s2-name">Тирамису</div><div class="s2-desc">классический, 150 г</div></div>
                  <div class="s2-right"><div class="s2-price">490 ₽</div><div class="s2-add">+</div></div>
                </div>
              </div>
            </div>
          </div>
          <div class="phone-label">Меню ресторана</div>
        </div>

        <!-- Screen 3: Order send -->
        <div class="phone-col">
          <div class="phone-shell">
            <div class="phone-bar"><span>12:30</span></div>
            <div class="phone-body">
              <div class="s3-wrap">
                <div class="s3-title">Ваш заказ</div>
                <div class="s3-line"><span>Том Ям x1</span><strong>890 ₽</strong></div>
                <div class="s3-line"><span>Стейк Рибай x1</span><strong>1 490 ₽</strong></div>
                <div class="s3-line"><span>Тирамису x1</span><strong>490 ₽</strong></div>
                <div class="s3-total"><span>Итого</span><strong>2 870 ₽</strong></div>
                <div class="s3-label">Куда отправить заказ?</div>
                <div class="s3-msg active">
                  <svg width="16" height="16" viewBox="0 0 24 24" fill="#333"><path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347z"/><path d="M12 0C5.373 0 0 5.373 0 12c0 2.127.556 4.12 1.525 5.855L0 24l6.335-1.652A11.94 11.94 0 0012 24c6.627 0 12-5.373 12-12S18.627 0 12 0z"/></svg>
                  WhatsApp
                </div>
                <div class="s3-msg">
                  <svg width="16" height="16" viewBox="0 0 24 24" fill="#666"><path d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12 12-5.373 12-12S18.627 0 12 0zm5.562 8.161c-.18 1.897-.962 6.502-1.359 8.627-.168.9-.5 1.201-.82 1.23-.697.064-1.226-.46-1.9-.903-1.056-.692-1.653-1.123-2.678-1.799-1.185-.781-.417-1.21.258-1.911.177-.184 3.247-2.977 3.307-3.23.007-.032.014-.15-.056-.212s-.174-.041-.249-.024c-.106.024-1.793 1.14-5.061 3.345-.479.33-.913.492-1.302.484-.428-.008-1.252-.241-1.865-.44-.752-.245-1.349-.374-1.297-.789.027-.216.325-.437.893-.663 3.498-1.524 5.831-2.529 6.998-3.014 3.332-1.386 4.025-1.627 4.476-1.635.099-.002.321.023.465.141a.506.506 0 01.171.325c.016.093.036.306.02.472z"/></svg>
                  Telegram
                </div>
              </div>
            </div>
          </div>
          <div class="phone-label">Отправка заказа</div>
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

.phone-shell {
  width: 240px;
  background: #1F2937;
  border-radius: 32px;
  padding: 8px;
  box-shadow: var(--shadow-lg);
  overflow: hidden;
}

.phone-bar {
  background: #374151;
  display: flex;
  justify-content: center;
  padding: 8px 16px 6px;
  border-radius: 24px 24px 0 0;
  font-size: 12px;
  color: #fff;
  font-weight: 600;
}

.phone-body {
  background: #fff;
  border-radius: 0 0 24px 24px;
  min-height: 420px;
  overflow: hidden;
}

.phone-label {
  font-size: 14px;
  font-weight: 600;
  color: var(--gray);
}

/* Screen 1 */
.s1-header {
  text-align: center;
  padding: 40px 16px 0;
  color: #fff;
}

.s1-name {
  font-size: 32px;
  font-weight: 700;
  letter-spacing: 3px;
}

.s1-sub {
  font-size: 11px;
  opacity: 0.7;
  margin-top: 4px;
}

.s1-card {
  background: #fff;
  margin: 24px 12px 12px;
  border-radius: 16px;
  padding: 16px;
}

.s1-label {
  font-size: 12px;
  color: var(--gray);
  font-weight: 600;
  margin-bottom: 8px;
}

.s1-row {
  display: flex;
  justify-content: space-between;
  font-size: 13px;
  padding: 6px 0;
  border-bottom: 1px solid var(--light-gray);
}

.s1-row span {
  color: var(--gray);
}

.s1-row.last {
  border-bottom: none;
}

.s1-btn {
  margin: 0 12px 12px;
  background: var(--primary);
  border-radius: 14px;
  padding: 14px;
  text-align: center;
  color: #fff;
  font-size: 14px;
  font-weight: 600;
}

/* Screen 2 */
.s2-tabs {
  display: flex;
  gap: 8px;
  padding: 16px 16px 0;
  overflow: hidden;
}

.s2-tab {
  padding: 6px 14px;
  background: var(--off-white);
  border-radius: 20px;
  font-size: 12px;
  font-weight: 500;
  color: var(--gray);
  white-space: nowrap;
}

.s2-tab.active {
  background: var(--primary);
  color: #fff;
  font-weight: 600;
}

.s2-items {
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.s2-item {
  background: var(--off-white);
  border-radius: 14px;
  padding: 14px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.s2-name {
  font-size: 14px;
  font-weight: 600;
}

.s2-desc {
  font-size: 11px;
  color: var(--gray);
  margin-top: 2px;
}

.s2-right {
  text-align: right;
}

.s2-price {
  font-size: 14px;
  font-weight: 700;
}

.s2-add {
  width: 24px;
  height: 24px;
  background: var(--primary);
  color: #fff;
  border-radius: 7px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  margin-top: 4px;
  margin-left: auto;
}

/* Screen 3 */
.s3-wrap {
  padding: 16px;
}

.s3-title {
  font-size: 16px;
  font-weight: 700;
  margin-bottom: 16px;
}

.s3-line {
  display: flex;
  justify-content: space-between;
  padding: 10px 0;
  border-bottom: 1px solid var(--light-gray);
  font-size: 13px;
}

.s3-total {
  display: flex;
  justify-content: space-between;
  padding: 14px 0;
  font-size: 15px;
  font-weight: 700;
}

.s3-label {
  font-size: 12px;
  color: var(--gray);
  margin-bottom: 12px;
}

.s3-msg {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 12px 14px;
  border: 2px solid var(--light-gray);
  border-radius: 12px;
  font-size: 13px;
  font-weight: 500;
  color: var(--gray);
  margin-bottom: 8px;
}

.s3-msg.active {
  border-color: var(--primary);
  color: var(--black);
  font-weight: 600;
}

@media (max-width: 768px) {
  .demo {
    padding: 64px 0;
  }

  .phones {
    gap: 32px;
  }

  .phone-shell {
    width: 220px;
  }

  .phone-body {
    min-height: 380px;
  }
}
</style>
