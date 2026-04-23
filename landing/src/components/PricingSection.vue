<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'

const { t, tm } = useI18n()

const plans = computed(() => {
  const items = tm('pricing.plans')
  return items.map((plan, i) => ({
    ...plan,
    featured: i === 1,
    badge: i === 1 ? t('pricing.popular') : null,
    cta: t('pricing.cta'),
  }))
})
</script>

<template>
  <section id="pricing" class="pricing">
    <div class="container">
      <div class="text-center">
        <h2 class="section-title">{{ $t('pricing.title') }}</h2>
        <p class="section-subtitle">{{ $t('pricing.subtitle') }}</p>
      </div>

      <div class="plans">
        <div class="plan" v-for="p in plans" :key="p.name" :class="{ featured: p.featured }">
          <div class="plan-badge" v-if="p.badge">{{ p.badge }}</div>
          <h3>{{ p.name }}</h3>
          <div class="plan-price">
            <span class="price-val">{{ p.price }}</span>
            <span class="price-per" v-if="p.period">{{ p.period }}</span>
          </div>
          <div class="price-hint" v-if="p.priceHint">{{ p.priceHint }}</div>
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
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24px;
  align-items: stretch;
  margin-top: 48px;
}

.plan {
  background: #fff;
  border: 2px solid var(--light-gray);
  border-radius: var(--radius-lg);
  padding: 36px 32px;
  text-align: center;
  position: relative;
  transition: all var(--ease);
  display: flex;
  flex-direction: column;
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

.price-hint {
  font-size: 13px;
  color: var(--gray);
  margin-top: 6px;
}

.plan-desc {
  font-size: 14px;
  color: var(--gray);
  margin: 12px 0 24px;
  min-height: 40px;
}

.plan ul {
  text-align: left;
  margin-bottom: 32px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  flex: 1;
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
  margin-top: auto;
}

@media (max-width: 960px) {
  .plans {
    grid-template-columns: 1fr;
    max-width: 420px;
    margin-left: auto;
    margin-right: auto;
  }

  .plan.featured {
    transform: none;
  }

  .plan.featured:hover {
    transform: translateY(-4px);
  }
}

@media (max-width: 768px) {
  .pricing {
    padding: 64px 0;
  }
}
</style>
