<script setup>
const features = [
  {
    icon: '<svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2L2 7l10 5 10-5-10-5z"/><path d="M2 17l10 5 10-5"/><path d="M2 12l10 5 10-5"/></svg>',
    title: 'Готовые шаблоны',
    desc: 'Выберите шаблон, загрузите меню и логотип — мини-приложение готово. Никакого кода.',
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
    icon: '<svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4-4v2"/><circle cx="9" cy="7" r="4"/><path d="M23 21v-2a4 4 0 00-3-3.87"/><path d="M16 3.13a4 4 0 010 7.75"/></svg>',
    title: 'Без лишних расходов',
    desc: 'Никаких комиссий с заказов, дорогого оборудования или подписок на кассовые системы.',
  },
  {
    icon: '<svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><polyline points="12,6 12,12 16,14"/></svg>',
    title: 'Запуск за 15 минут',
    desc: 'Зарегистрируйтесь, заполните меню, получите ссылку. Всё.',
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
          Шаблонное мини-приложение, которое начнёт приносить заказы с&nbsp;первого дня.
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
