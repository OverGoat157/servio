<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

function goBack() {
  if (window.history.length > 1) {
    router.back()
  } else {
    router.push({ name: 'dashboard' })
  }
}
</script>

<template>
  <div class="page">
    <div class="page-header">
      <button class="back-link" @click="goBack">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M19 12H5M12 19l-7-7 7-7"/></svg>
        {{ $t('common.back') }}
      </button>
      <h1>{{ $t('helpTelegram.title') }}</h1>
      <p class="lead">
        {{ $t('helpTelegram.lead') }}
      </p>
    </div>

    <div class="card step">
      <div class="step-num">1</div>
      <div class="step-body">
        <h2>{{ $t('helpTelegram.step1Title') }}</h2>
        <ol>
          <li><a href="https://t.me/BotFather" target="_blank" rel="noopener">{{ $t('helpTelegram.step1Item1') }}</a></li>
          <li>{{ $t('helpTelegram.step1Item2') }}</li>
          <li>{{ $t('helpTelegram.step1Item3Pre') }} <b>{{ $t('helpTelegram.step1Item3Bold') }}</b> {{ $t('helpTelegram.step1Item3Post') }}</li>
          <li>{{ $t('helpTelegram.step1Item4Pre') }} <b>{{ $t('helpTelegram.step1Item4Bold') }}</b> {{ $t('helpTelegram.step1Item4Post') }}</li>
          <li>{{ $t('helpTelegram.step1Item5') }}
            <div class="code-block">123456789:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw</div>
          </li>
        </ol>
        <div class="callout warn">
          {{ $t('helpTelegram.step1Warn') }}
        </div>
      </div>
    </div>

    <div class="card step">
      <div class="step-num">2</div>
      <div class="step-body">
        <h2>{{ $t('helpTelegram.step2Title') }}</h2>
        <p>{{ $t('helpTelegram.step2Lead') }}</p>

        <h3>{{ $t('helpTelegram.step2SubA') }}</h3>
        <ol>
          <li><a href="https://t.me/userinfobot" target="_blank" rel="noopener">{{ $t('helpTelegram.step2A1') }}</a></li>
          <li>{{ $t('helpTelegram.step2A2') }}</li>
          <li>{{ $t('helpTelegram.step2A3') }}</li>
        </ol>

        <h3>{{ $t('helpTelegram.step2SubB') }}</h3>
        <ol>
          <li>{{ $t('helpTelegram.step2B1') }}</li>
          <li>{{ $t('helpTelegram.step2B2') }}</li>
          <li><a href="https://t.me/getmyid_bot" target="_blank" rel="noopener">{{ $t('helpTelegram.step2B3') }}</a></li>
          <li>{{ $t('helpTelegram.step2B4') }}</li>
          <li>{{ $t('helpTelegram.step2B5') }}</li>
        </ol>

        <div class="callout">
          {{ $t('helpTelegram.step2Note') }}
        </div>
      </div>
    </div>

    <div class="card step">
      <div class="step-num">3</div>
      <div class="step-body">
        <h2>{{ $t('helpTelegram.step3Title') }}</h2>
        <ol>
          <li>{{ $t('helpTelegram.step3Item1') }}</li>
          <li>{{ $t('helpTelegram.step3Item2Pre') }}
            <ul>
              <li>{{ $t('helpTelegram.step3Item2Bullet1') }}</li>
              <li>{{ $t('helpTelegram.step3Item2Bullet2') }}</li>
            </ul>
          </li>
          <li>{{ $t('helpTelegram.step3Item3') }}</li>
          <li>{{ $t('helpTelegram.step3Item4') }}</li>
        </ol>
      </div>
    </div>

    <div class="card step">
      <div class="step-num">4</div>
      <div class="step-body">
        <h2>{{ $t('helpTelegram.step4Title') }}</h2>
        <ol>
          <li>{{ $t('helpTelegram.step4Item1') }}</li>
          <li>{{ $t('helpTelegram.step4Item2') }}</li>
          <li>{{ $t('helpTelegram.step4Item3') }}</li>
          <li>{{ $t('helpTelegram.step4Item4') }}</li>
          <li>{{ $t('helpTelegram.step4Item5') }}</li>
        </ol>
      </div>
    </div>

    <div class="card faq">
      <h2>{{ $t('helpTelegram.faqTitle') }}</h2>
      <dl>
        <dt>{{ $t('helpTelegram.faqQ1') }}</dt>
        <dd>{{ $t('helpTelegram.faqA1') }}</dd>

        <dt>{{ $t('helpTelegram.faqQ2') }}</dt>
        <dd>{{ $t('helpTelegram.faqA2') }}</dd>

        <dt>{{ $t('helpTelegram.faqQ3') }}</dt>
        <dd>{{ $t('helpTelegram.faqA3') }}</dd>

        <dt>{{ $t('helpTelegram.faqQ4') }}</dt>
        <dd>{{ $t('helpTelegram.faqA4') }}</dd>
      </dl>
    </div>

    <router-link :to="{ name: 'dashboard' }" class="btn btn-primary bottom-link">
      {{ $t('helpTelegram.goToDashboard') }}
    </router-link>
  </div>
</template>

<style scoped>
.page {
  max-width: 780px;
  margin: 0 auto;
  padding: 32px 24px 80px;
}

.page-header {
  margin-bottom: 28px;
}

.back-link {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 14px;
  color: var(--text-secondary);
  margin-bottom: 12px;
}

.page-header h1 {
  font-size: 26px;
  font-weight: 700;
  margin-bottom: 6px;
}

.lead {
  font-size: 14px;
  color: var(--text-secondary);
  line-height: 1.5;
}

.step {
  display: flex;
  gap: 18px;
  padding: 22px 24px;
  margin-bottom: 14px;
}

.step-num {
  flex-shrink: 0;
  width: 34px;
  height: 34px;
  border-radius: 50%;
  background: var(--primary);
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 15px;
  font-weight: 700;
}

.step-body {
  flex: 1;
  min-width: 0;
}

.step h2 {
  font-size: 17px;
  font-weight: 700;
  margin-bottom: 12px;
}

.step h3 {
  font-size: 14px;
  font-weight: 700;
  margin-top: 16px;
  margin-bottom: 8px;
  color: var(--text);
}

.step p {
  font-size: 14px;
  color: var(--text-secondary);
  line-height: 1.6;
  margin-bottom: 10px;
}

.step ol,
.step ul {
  font-size: 14px;
  color: var(--text);
  line-height: 1.65;
  padding-left: 22px;
  margin-bottom: 10px;
}

.step ol li,
.step ul li {
  margin-bottom: 6px;
}

.step ul {
  padding-left: 18px;
  list-style: disc;
  margin-top: 4px;
}

code {
  background: var(--bg-secondary);
  padding: 2px 6px;
  border-radius: 4px;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 12.5px;
  color: var(--text);
}

.code-block {
  margin-top: 8px;
  padding: 10px 14px;
  background: var(--bg-secondary);
  border-radius: var(--radius);
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 12.5px;
  word-break: break-all;
  color: var(--text);
}

a {
  color: var(--primary);
  text-decoration: underline;
}

.callout {
  margin-top: 12px;
  padding: 12px 14px;
  background: var(--primary-light);
  border-left: 3px solid var(--primary);
  border-radius: 6px;
  font-size: 13px;
  line-height: 1.5;
  color: var(--text);
}

.callout.warn {
  background: #FEF3C7;
  border-left-color: #D97706;
}

.faq {
  padding: 22px 24px;
  margin-bottom: 14px;
}

.faq h2 {
  font-size: 17px;
  font-weight: 700;
  margin-bottom: 14px;
}

.faq dt {
  font-size: 14px;
  font-weight: 600;
  color: var(--text);
  margin-top: 14px;
}

.faq dt:first-of-type {
  margin-top: 0;
}

.faq dd {
  font-size: 13px;
  color: var(--text-secondary);
  line-height: 1.6;
  margin-top: 4px;
  margin-left: 0;
}

.bottom-link {
  display: inline-flex;
  margin-top: 8px;
}
</style>
