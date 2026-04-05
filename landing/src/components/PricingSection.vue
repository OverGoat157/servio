<script setup>
const plans = [
  {
    name: 'Бизнес',
    price: '10 000 ₽',
    period: '',
    desc: 'Полная настройка мини-приложения под ваш бизнес',
    features: ['Мини-приложение с меню', 'Безлимитное меню', 'Все мессенджеры', 'Индивидуальный дизайн', 'QR-код для печати', 'Техническая поддержка'],
    cta: 'Оставить заявку',
    featured: true,
  },
]
</script>

<template>
  <section id="pricing" class="pricing">
    <div class="container">
      <div class="text-center">
        <h2 class="section-title">Простые и понятные тарифы</h2>
        <p class="section-subtitle">Без скрытых комиссий и процентов с заказов</p>
      </div>

      <div class="plans">
        <div class="plan" v-for="p in plans" :key="p.name" :class="{ featured: p.featured }">
          <div class="plan-badge" v-if="p.badge">{{ p.badge }}</div>
          <h3>{{ p.name }}</h3>
          <div class="plan-price">
            <span class="price-val">{{ p.price }}</span>
            <span class="price-per" v-if="p.period">{{ p.period }}</span>
          </div>
          <p class="plan-desc">{{ p.desc }}</p>
          <ul>
            <li v-for="f in p.features" :key="f">
              <svg width="16" height="16" viewBox="0 0 20 20" fill="none"><path d="M7 10l2 2 4-4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
              {{ f }}
            </li>
          </ul>
          <a href="#contact" class="btn" :class="p.featured ? 'btn-primary' : 'btn-outline'">{{ p.cta }}</a>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.pricing {
  padding: 100px 0;
}

.text-center {
  text-align: center;
}

.plans {
  display: flex;
  justify-content: center;
  gap: 24px;
}

.plan {
  background: #fff;
  border: 2px solid var(--light-gray);
  border-radius: var(--radius-lg);
  padding: 36px 32px;
  text-align: center;
  position: relative;
  transition: all var(--ease);
}

.plan:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow);
}

.plan.featured {
  border-color: var(--primary);
  box-shadow: var(--shadow-lg);
  transform: scale(1.04);
}

.plan.featured:hover {
  transform: scale(1.04) translateY(-4px);
}

.plan-badge {
  position: absolute;
  top: -13px;
  left: 50%;
  transform: translateX(-50%);
  background: var(--primary);
  color: #fff;
  padding: 4px 16px;
  border-radius: 100px;
  font-size: 12px;
  font-weight: 600;
  white-space: nowrap;
}

.plan h3 {
  font-size: 20px;
  font-weight: 600;
  margin-bottom: 16px;
}

.price-val {
  font-size: 36px;
  font-weight: 700;
}

.price-per {
  font-size: 16px;
  color: var(--gray);
}

.plan-desc {
  font-size: 14px;
  color: var(--gray);
  margin: 8px 0 24px;
}

.plan ul {
  text-align: left;
  margin-bottom: 32px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.plan li {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 14px;
}

.plan li svg {
  flex-shrink: 0;
}

.plan .btn {
  width: 100%;
}

.plan {
  max-width: 400px;
  width: 100%;
}

@media (max-width: 768px) {
  .pricing {
    padding: 64px 0;
  }
}
</style>
