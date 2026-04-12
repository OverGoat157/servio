<script setup>
const features = [
  {
    icon: '<svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="20" x2="18" y2="10"/><line x1="12" y1="20" x2="12" y2="4"/><line x1="6" y1="20" x2="6" y2="14"/><line x1="3" y1="20" x2="21" y2="20"/></svg>',
    title: 'Удобная CRM и статистика',
    desc: 'Видите заказы, клиентов, выручку и популярные блюда. Удобная админка для управления меню и ценами.',
  },
  {
    icon: '<svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 01-2 2H7l-4 4V5a2 2 0 012-2h14a2 2 0 012 2z"/></svg>',
    title: 'Заказы в мессенджер',
    desc: 'Заказы приходят прямо в WhatsApp, Telegram или Viber. Не нужен POS-терминал.',
  },
  {
    icon: '<svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="5" y="2" width="14" height="20" rx="2" ry="2"/><line x1="12" y1="18" x2="12.01" y2="18"/></svg>',
    title: 'Работает на телефоне',
    desc: 'Адаптивный дизайн для любого устройства. Клиент открывает ссылку и видит меню.',
  },
  {
    icon: '<svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 12h-4l-3 9L9 3l-3 9H2"/></svg>',
    title: 'Поддержка 24/7',
    desc: 'Мы всегда на связи. Помогаем с настройкой, регулярно выпускаем обновления и новые возможности.',
  },
  {
    icon: '<svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><polyline points="12,6 12,12 16,14"/></svg>',
    title: 'Запуск за 1-2 дня',
    desc: 'Мы сами делаем первоначальный сайт, загружаем меню и подключаем мессенджер. Вы сразу принимаете заказы.',
  },
  {
    icon: '<svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/></svg>',
    title: 'Ваш бренд',
    desc: 'Логотип, цвета, фон — всё настраивается. Клиент видит ваш ресторан.',
  },
]
</script>

<template>
  <section id="features" class="features">
    <div class="container">
      <div class="text-center">
        <h2 class="section-title">Всё что нужно вашему ресторану</h2>
        <p class="section-subtitle">
          Готовое цифровое меню с&nbsp;поддержкой и&nbsp;обновлениями. Заказы с&nbsp;первого дня.
        </p>
      </div>

      <div class="grid">
        <div class="card" v-for="f in features" :key="f.title">
          <div class="card-icon" v-html="f.icon" />
          <h3>{{ f.title }}</h3>
          <p>{{ f.desc }}</p>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.features {
  padding: 100px 0;
  background: var(--section-bg);
}

.text-center {
  text-align: center;
  margin-bottom: 64px;
}

.grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24px;
}

.card {
  background: #fff;
  padding: 32px;
  border-radius: var(--radius-lg);
  border: 1px solid var(--light-gray);
  transition: all var(--ease);
}

.card:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow-lg);
}

.card-icon {
  width: 52px;
  height: 52px;
  border-radius: 14px;
  background: var(--primary-light);
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 20px;
  color: var(--primary);
}

.card h3 {
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 8px;
}

.card p {
  font-size: 15px;
  color: var(--gray);
  line-height: 1.6;
}

@media (max-width: 1024px) {
  .grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 640px) {
  .features {
    padding: 64px 0;
  }

  .grid {
    grid-template-columns: 1fr;
  }
}
</style>
