<template>
  <div class="home">
    <!-- Hero and CTA -->
    <TextHero 
      heading="WE CREATE VIDEOS THAT CREATE BUSINESS" 
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
    <FeatureGrid
      heading="WE GET RESULTS">
    </FeatureGrid>
  </div>
</template>

<script>
// @ is an alias to /src
import TextHero from '@/components/TextHero.vue'
import Carousel from '@/components/Carousel.vue'
import FeatureGrid from '@/components/FeatureGrid.vue'

export default {
  name: 'HomeView',
  components: {
    TextHero,
    Carousel,
    FeatureGrid
  }
}
</script>
