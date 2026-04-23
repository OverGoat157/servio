<script setup>
import heroImg from '../assets/demo/hero.png'
</script>

<template>
  <section class="hero">
    <div class="container hero-grid">
      <div class="hero-text">
        <div class="badge">
          <span class="badge-dot" />
          {{ $t('hero.badge') }}
        </div>

        <h1>{{ $t('hero.titleLead') }}&nbsp;<span class="accent">{{ $t('hero.titleAccent') }}</span></h1>

        <p class="lead">
          {{ $t('hero.lead') }}
        </p>

        <div class="hero-btns">
          <a href="#contact" class="btn btn-primary">
            {{ $t('hero.ctaPrimary') }}
            <svg width="20" height="20" viewBox="0 0 20 20" fill="none"><path d="M4 10h12m0 0l-4-4m4 4l-4 4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
          </a>
          <a href="#demo" class="btn btn-outline">{{ $t('hero.ctaSecondary') }}</a>
        </div>

        <div class="stats">
          <div class="stat"><strong>{{ $t('hero.stat1Value') }}</strong><span>{{ $t('hero.stat1Label') }}</span></div>
          <div class="stat-line" />
          <div class="stat"><strong>{{ $t('hero.stat2Value') }}</strong><span>{{ $t('hero.stat2Label') }}</span></div>
          <div class="stat-line" />
          <div class="stat"><strong>{{ $t('hero.stat3Value') }}</strong><span>{{ $t('hero.stat3Label') }}</span></div>
        </div>
      </div>

      <div class="hero-visual">
        <div class="video-wrap">
          <img class="hero-image" :src="heroImg" :alt="$t('hero.imageAlt')" />
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
  width: 100%;
  max-width: 420px;
  border-radius: 24px;
  overflow: hidden;
  background: var(--off-white);
  box-shadow: var(--shadow-lg);
  display: flex;
}

.hero-image {
  width: 100%;
  height: auto;
  object-fit: contain;
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
    max-width: 320px;
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
