<script setup>
import heroImg from '../assets/demo/hero.png'
</script>

<template>
  <section class="hero">
    <div class="container hero-grid">
      <div class="hero-text">
        <div class="badge">
          <span class="badge-dot" />
          Новый способ принимать заказы
        </div>

        <h1>Сайт-приложение с&nbsp;меню для вашего <span class="accent">ресторана</span></h1>

        <p class="lead">
          Мы создадим для вас цифровое меню за&nbsp;1-2&nbsp;дня. Клиенты делают заказ прямо
          со&nbsp;смартфона, а&nbsp;вы получаете его в&nbsp;WhatsApp, Telegram
          или любой другой мессенджер.
        </p>

        <div class="hero-btns">
          <a href="#contact" class="btn btn-primary">
            Оставить заявку
            <svg width="20" height="20" viewBox="0 0 20 20" fill="none"><path d="M4 10h12m0 0l-4-4m4 4l-4 4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
          </a>
          <a href="#demo" class="btn btn-outline">Посмотреть демо</a>
        </div>

        <div class="stats">
          <div class="stat"><strong>1-2 дня</strong><span>На запуск</span></div>
          <div class="stat-line" />
          <div class="stat"><strong>24/7</strong><span>Поддержка</span></div>
          <div class="stat-line" />
          <div class="stat"><strong>0%</strong><span>Комиссий</span></div>
        </div>
      </div>

      <div class="hero-visual">
        <div class="video-wrap">
          <img class="hero-image" :src="heroImg" alt="Пример сайт-приложения" />
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

.hero-visual {
  position: relative;
  display: flex;
  justify-content: center;
}

.video-wrap {
  position: relative;
  width: 100%;
  max-width: 540px;
  aspect-ratio: 4 / 5;
  border-radius: 24px;
  overflow: hidden;
  background: var(--off-white);
  box-shadow: var(--shadow-lg);
}

.hero-image {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
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

  .video-wrap {
    max-width: 420px;
    aspect-ratio: 16 / 10;
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
