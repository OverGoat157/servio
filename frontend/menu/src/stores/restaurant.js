import { reactive, ref } from 'vue'
import { fetchMenu } from '../api/client'

export const restaurant = reactive({
  id: null,
  name: '',
  slug: '',
  description: null,
  logo: null,
  phone: null,
  address: null,
  working_hours: null,
  theme: 'default',
  categories: [],
  messengers: [],
})

export const loading = ref(true)
export const error = ref(null)

export async function loadRestaurant(slug) {
  loading.value = true
  error.value = null
  try {
    const data = await fetchMenu(slug)
    Object.assign(restaurant, data)
  } catch (e) {
    error.value = e.message
  } finally {
    loading.value = false
  }
}
