<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import homeImg from '../assets/demo/home.png'
import menuImg from '../assets/demo/menu.png'
import cartImg from '../assets/demo/cart.png'

const { tm } = useI18n()
const images = [homeImg, menuImg, cartImg]
const screens = computed(() => tm('demo.screens').map((label, i) => ({ src: images[i], label })))
</script>

<template>
  <section id="demo" class="demo">
    <div class="container">
      <div class="text-center">
        <h2 class="section-title">{{ $t('demo.title') }}</h2>
        <p class="section-subtitle">
          {{ $t('demo.subtitle') }}
        </p>
      </div>

      <div class="phones">
        <div class="phone-col" v-for="s in screens" :key="s.label">
          <div class="screenshot">
            <img :src="s.src" :alt="s.label" />
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
  width: 320px;
  max-width: 100%;
  border-radius: 24px;
  overflow: hidden;
  background: var(--off-white);
  border: 1px solid var(--light-gray);
  box-shadow: var(--shadow-lg);
  display: flex;
}

.screenshot img {
  width: 100%;
  height: auto;
  display: block;
  object-fit: contain;
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
    width: min(320px, 100%);
  }
}
</style>
