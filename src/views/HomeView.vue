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
      heading="TRUSTED BY"
      :images="['images/cna.png', 
                'images/dell.png',
                'images/discovery.png',
                'images/fujitsu.png',
                'images/milliman.png',
                'images/natgeo.png',
                'images/phuel.png',
                'images/primo.png',
                'images/sap.png']">
    </Carousel>
    <FeatureCard
      heading="300+"
      text="videos created">
    </FeatureCard>
  </div>
</template>

<script>
// @ is an alias to /src
import TextHero from '@/components/TextHero.vue'
import Carousel from '@/components/Carousel.vue'
import FeatureCard from '@/components/FeatureCard.vue'

export default {
  name: 'HomeView',
  components: {
    TextHero,
    Carousel,
    FeatureCard
  }
}
</script>
