<template>
  <div class="home">
    <!-- Hero and CTA -->
    <TextHero 
      heading="Lorem ipsum is placeholder text commonly used in the graphic." 
      subheading="Check out our new font generator and level up your social bios." 
      ctaText="See plans">
    </TextHero>
    <!-- Recent work carousel -->
    <Carousel 
      heading="Latest projects">
    </Carousel>
  </div>
</template>

<script>
// @ is an alias to /src
import TextHero from '@/components/TextHero.vue'
import Carousel from '@/components/Carousel.vue'

export default {
  name: 'HomeView',
  components: {
    TextHero,
    Carousel
  }
}
</script>
