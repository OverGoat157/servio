<template>
  <footer class="footer">
    <div class="container">
      <div class="footer-top">
        <div class="footer-brand">
          <a href="#" class="logo">
            <span class="logo-icon">M</span>
            <span>MenuLink</span>
          </a>
          <p>Шаблонные мини-приложения с&nbsp;меню для ресторанов. Заказы прямо в мессенджер.</p>
        </div>

        <div class="footer-links">
          <div class="col">
            <h4>Продукт</h4>
            <a href="#features">Возможности</a>
            <a href="#how">Как это работает</a>
            <a href="#pricing">Тарифы</a>
            <a href="#demo">Демо</a>
          </div>
          <div class="col">
            <h4>Компания</h4>
            <a href="#">О нас</a>
            <a href="#">Блог</a>
            <a href="#">Контакты</a>
          </div>
          <div class="col">
            <h4>Связаться</h4>
            <a href="#">hello@menulink.ru</a>
            <a href="#">+7 (800) 123-45-67</a>
            <div class="socials">
              <a href="#" class="social" aria-label="Telegram">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor"><path d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12 12-5.373 12-12S18.627 0 12 0zm5.562 8.161c-.18 1.897-.962 6.502-1.359 8.627-.168.9-.5 1.201-.82 1.23-.697.064-1.226-.46-1.9-.903-1.056-.692-1.653-1.123-2.678-1.799-1.185-.781-.417-1.21.258-1.911.177-.184 3.247-2.977 3.307-3.23.007-.032.014-.15-.056-.212s-.174-.041-.249-.024c-.106.024-1.793 1.14-5.061 3.345-.479.33-.913.492-1.302.484-.428-.008-1.252-.241-1.865-.44-.752-.245-1.349-.374-1.297-.789.027-.216.325-.437.893-.663 3.498-1.524 5.831-2.529 6.998-3.014 3.332-1.386 4.025-1.627 4.476-1.635.099-.002.321.023.465.141a.506.506 0 01.171.325c.016.093.036.306.02.472z"/></svg>
              </a>
              <a href="#" class="social" aria-label="WhatsApp">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor"><path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347z"/><path d="M12 0C5.373 0 0 5.373 0 12c0 2.127.556 4.12 1.525 5.855L0 24l6.335-1.652A11.94 11.94 0 0012 24c6.627 0 12-5.373 12-12S18.627 0 12 0z"/></svg>
              </a>
            </div>
          </div>
        </div>
      </div>

      <div class="footer-bottom">
        <span>&copy; 2025 MenuLink. Все права защищены.</span>
        <div class="legal">
          <a href="#">Политика конфиденциальности</a>
          <a href="#">Оферта</a>
        </div>
      </div>
    </div>
  </footer>
</template>

<style scoped>
.footer {
  background: linear-gradient(180deg, #1F2937, #111827);
  color: rgba(255, 255, 255, 0.65);
  padding: 64px 0 32px;
}

.footer-top {
  display: grid;
  grid-template-columns: 1fr 2fr;
  gap: 64px;
  padding-bottom: 40px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.logo {
  display: flex;
  align-items: center;
  gap: 10px;
  font-weight: 700;
  font-size: 20px;
  color: #fff;
  margin-bottom: 16px;
}

.logo-icon {
  width: 36px;
  height: 36px;
  background: var(--primary);
  color: #fff;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  font-weight: 700;
}

.footer-brand p {
  font-size: 14px;
  line-height: 1.6;
  max-width: 280px;
}

.footer-links {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 32px;
}

.col {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.col h4 {
  font-size: 14px;
  font-weight: 600;
  color: #fff;
  margin-bottom: 4px;
  text-transform: uppercase;
  letter-spacing: 1px;
}

.col a {
  font-size: 14px;
  color: rgba(255, 255, 255, 0.55);
  transition: color var(--ease);
}

.col a:hover {
  color: #fff;
}

.socials {
  display: flex;
  gap: 12px;
  margin-top: 4px;
}

.social {
  width: 36px;
  height: 36px;
  border-radius: 10px;
  background: rgba(255, 255, 255, 0.1);
  display: flex;
  align-items: center;
  justify-content: center;
  color: rgba(255, 255, 255, 0.55);
  transition: all var(--ease);
}

.social:hover {
  background: var(--primary);
  color: #fff;
}

.footer-bottom {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 24px;
  font-size: 13px;
}

.legal {
  display: flex;
  gap: 24px;
}

.legal a {
  color: rgba(255, 255, 255, 0.4);
  transition: color var(--ease);
}

.legal a:hover {
  color: #fff;
}

@media (max-width: 768px) {
  .footer {
    padding: 48px 0 24px;
  }

  .footer-top {
    grid-template-columns: 1fr;
    gap: 40px;
  }

  .footer-links {
    grid-template-columns: 1fr 1fr;
    gap: 24px;
  }

  .footer-bottom {
    flex-direction: column;
    gap: 12px;
    text-align: center;
  }
}
</style>
