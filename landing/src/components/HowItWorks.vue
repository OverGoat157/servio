<script setup>
const steps = [
  {
    title: 'Оставьте заявку',
    desc: 'Заполните короткую форму — мы свяжемся с вами в течение дня и обсудим детали.',
  },
  {
    title: 'Мы настраиваем всё за вас',
    desc: 'Делаем сайт, загружаем ваше меню с фотографиями и описаниями, оформляем обложку и карточку ресторана, подключаем Telegram или WhatsApp для заказов. Обычно за 1-2 дня.',
  },
  {
    title: 'Принимайте заказы',
    desc: 'Распечатайте QR-код на столики или добавьте ссылку в соцсети. Заказы приходят прямо в мессенджер, меню и цены редактируете сами в удобной админке.',
  },
]
</script>

<template>
  <section id="how" class="how">
    <div class="container">
      <div class="text-center">
        <h2 class="section-title">Как это работает</h2>
        <p class="section-subtitle">Мы всё делаем за вас. Вам остаётся только принимать заказы.</p>
      </div>

      <div class="steps">
        <div class="step" v-for="(s, i) in steps" :key="i">
          <div class="step-num">{{ i + 1 }}</div>
          <div class="step-body">
            <h3>{{ s.title }}</h3>
            <p>{{ s.desc }}</p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.how {
  padding: 100px 0;
}

.text-center {
  text-align: center;
  margin-bottom: 64px;
}

.steps {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.step {
  display: flex;
  align-items: center;
  gap: 32px;
  background: var(--off-white);
  border-radius: var(--radius-lg);
  padding: 40px;
  border: 1px solid var(--light-gray);
}

.step-num {
  width: 56px;
  height: 56px;
  min-width: 56px;
  border-radius: 16px;
  background: var(--primary);
  color: #fff;
  font-size: 24px;
  font-weight: 700;
  display: flex;
  align-items: center;
  justify-content: center;
}

.step-body h3 {
  font-size: 20px;
  font-weight: 600;
  margin-bottom: 8px;
}

.step-body p {
  font-size: 15px;
  color: var(--gray);
  line-height: 1.6;
}

@media (max-width: 768px) {
  .how {
    padding: 64px 0;
  }

  .step {
    flex-direction: column;
    text-align: center;
    padding: 24px;
    gap: 20px;
  }
}
</style>
