import { i18n } from './i18n'

// Pick a localized field value for the current locale, with fallback to Russian.
// `obj` is the raw API object (e.g. a menu item), `base` is the RU field name ('name', 'description', 'ingredients').
export function localized(obj, base) {
  if (!obj) return ''
  const locale = i18n.global.locale.value
  if (locale === 'en') {
    const en = obj[base + '_en']
    if (en && String(en).trim()) return en
  }
  return obj[base] || ''
}
