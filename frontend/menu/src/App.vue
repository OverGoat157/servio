<script setup>
import { watch } from 'vue'
import { useRoute } from 'vue-router'
import { loadRestaurant, restaurant } from './stores/restaurant'

const route = useRoute()

watch(
  () => route.params.slug,
  (slug) => {
    if (slug && !restaurant.id) {
      loadRestaurant(slug)
    }
  },
  { immediate: true }
)
</script>

<template>
  <router-view />
</template>
