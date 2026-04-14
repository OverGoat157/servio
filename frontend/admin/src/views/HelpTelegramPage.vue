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
        Назад
      </button>
      <h1>Подключение Telegram</h1>
      <p class="lead">
        Пошаговая инструкция: создайте бота, добавьте его в чат и пропишите данные в настройках мессенджеров.
      </p>
    </div>

    <div class="card step">
      <div class="step-num">1</div>
      <div class="step-body">
        <h2>Создать бота в @BotFather</h2>
        <ol>
          <li>Откройте в Telegram <a href="https://t.me/BotFather" target="_blank" rel="noopener">@BotFather</a>.</li>
          <li>Отправьте команду <code>/newbot</code>.</li>
          <li>Введите <b>имя бота</b> — оно будет отображаться в чате (например: <i>Суши Море Заказы</i>).</li>
          <li>Введите <b>username бота</b> — должен быть уникальным и заканчиваться на <code>bot</code> (например: <code>sushi_more_orders_bot</code>).</li>
          <li>BotFather пришлёт сообщение с токеном — выглядит так:
            <div class="code-block">123456789:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw</div>
          </li>
        </ol>
        <div class="callout warn">
          Скопируйте токен и сохраните его в надёжном месте. Никому его не показывайте — это пароль к вашему боту.
        </div>
      </div>
    </div>

    <div class="card step">
      <div class="step-num">2</div>
      <div class="step-body">
        <h2>Получить Chat ID — куда бот будет слать заказы</h2>
        <p>Нужно решить, куда будут приходить заказы: лично вам или в групповой чат, где их видят все сотрудники.</p>

        <h3>Вариант А — в личный чат</h3>
        <ol>
          <li>Откройте <a href="https://t.me/userinfobot" target="_blank" rel="noopener">@userinfobot</a> и нажмите <b>Start</b>.</li>
          <li>Бот пришлёт ваш числовой ID — например <code>123456789</code>. Это и есть <b>Chat ID</b>.</li>
          <li>Обязательно откройте диалог со <b>своим</b> ботом (которого создавали на шаге 1) и нажмите <b>Start</b> — иначе бот не сможет вам писать.</li>
        </ol>

        <h3>Вариант Б — в групповой чат (рекомендуется для команды)</h3>
        <ol>
          <li>Создайте группу в Telegram (например «Заказы Суши Море») и добавьте туда коллег.</li>
          <li>Добавьте в ту же группу <b>своего бота</b> по username (который придумали на шаге 1).</li>
          <li>Добавьте туда же <a href="https://t.me/getmyid_bot" target="_blank" rel="noopener">@getmyid_bot</a> — он пришлёт <b>Group ID</b>, начинающийся с минуса, например <code>-1001234567890</code>.</li>
          <li>После этого @getmyid_bot можно удалить из группы.</li>
          <li><b>Своего бота</b> из группы удалять нельзя — иначе заказы не будут доставляться.</li>
        </ol>

        <div class="callout">
          У групп Chat ID всегда отрицательный и начинается с минуса. Копируйте его целиком, вместе с минусом.
        </div>
      </div>
    </div>

    <div class="card step">
      <div class="step-num">3</div>
      <div class="step-body">
        <h2>Прописать данные в админке</h2>
        <ol>
          <li>Откройте карточку ресторана → <b>Мессенджеры</b>.</li>
          <li>В блоке <b>Telegram</b> заполните:
            <ul>
              <li><b>Bot Token</b> — токен от @BotFather из шага 1</li>
              <li><b>Chat ID</b> — число из шага 2</li>
            </ul>
          </li>
          <li>Нажмите <b>Сохранить</b>.</li>
          <li>Включите тумблер <b>Включено</b>.</li>
        </ol>
      </div>
    </div>

    <div class="card step">
      <div class="step-num">4</div>
      <div class="step-body">
        <h2>Проверить работу</h2>
        <ol>
          <li>Откройте меню ресторана с телефона.</li>
          <li>Добавьте блюдо в корзину, заполните имя и телефон.</li>
          <li>Выберите способ получения и мессенджер <b>Telegram</b>.</li>
          <li>Нажмите <b>Отправить заказ</b>.</li>
          <li>В вашем чате должно появиться сообщение от бота с составом заказа.</li>
        </ol>
      </div>
    </div>

    <div class="card faq">
      <h2>Частые проблемы</h2>
      <dl>
        <dt>Нажимаю «Отправить заказ», но в Telegram ничего не приходит</dt>
        <dd>Проверьте, что бот добавлен в группу и не удалён. Если это личный чат — откройте диалог с ботом и нажмите <b>Start</b>.</dd>

        <dt>Ошибка «Telegram не настроен»</dt>
        <dd>Bot Token или Chat ID пустые, либо тумблер <b>Включено</b> выключен.</dd>

        <dt>Ошибка «Не удалось отправить в Telegram»</dt>
        <dd>Неверный токен или Chat ID. Проверьте, что скопировали их без лишних пробелов. Для группы Chat ID должен начинаться с минуса.</dd>

        <dt>Заказы перестали приходить после того, как добавил нового сотрудника в группу</dt>
        <dd>Скорее всего кто-то случайно удалил бота из группы. Добавьте его обратно по username.</dd>
      </dl>
    </div>

    <router-link :to="{ name: 'dashboard' }" class="btn btn-primary bottom-link">
      Перейти в панель управления
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
