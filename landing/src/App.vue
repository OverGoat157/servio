<script setup>
import NavBar from './components/NavBar.vue'
import HeroSection from './components/HeroSection.vue'
import FeaturesSection from './components/FeaturesSection.vue'
import HowItWorks from './components/HowItWorks.vue'
import DemoSection from './components/DemoSection.vue'
import PricingSection from './components/PricingSection.vue'
import CtaSection from './components/CtaSection.vue'
import FooterSection from './components/FooterSection.vue'
</script>

<template>
  <NavBar />
  <main>
    <HeroSection />
    <FeaturesSection />
    <HowItWorks />
    <DemoSection />
    <PricingSection />
    <CtaSection />
  </main>
  <FooterSection />
</template>
