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
    <FeatureGridv2
      heading="WE GET RESULTS"
      :listOfCards="[['images/tell-your-story.webp', 'TELL YOUR STORY', 'Whether you are an artist, a company, or an individual — you have a story to tell. Let us help you tell that story through video. From pre-production all the way through to the premiere, we will work with you to bring your story to life.'],
                     ['images/high-quality-experience.webp', 'HIGH QUALITY EXPERIENCE', 'We work across many video genres - corporate video, documentaries, websites, advertising, social media - you name it! Whatever your business objective, we can provide script-to-screen solutions with top-of-the-line professional light and sound equipment. Whether shooting on location or in a studio, we will create the feel that conveys your message and we promise to have fun along the way!'],
                     ['images/start-today.webp', 'START TODAY', 'Your first step is to get in touch with us about your project. We can point you in the direction of your next steps and even give you a quote. We do book up, so contact us soon so we can start the pre-production stages of your video and help you hit the deadlines you need.']]">
    </FeatureGridv2>
  </div>
</template>

<script>
// @ is an alias to /src
import TextHero from '@/components/TextHero.vue'
import Carousel from '@/components/Carousel.vue'
import FeatureGrid from '@/components/FeatureGrid.vue'
import FeatureGridv2 from '@/components/FeatureGridv2.vue'

export default {
  name: 'HomeView',
  components: {
    TextHero,
    Carousel,
    FeatureGrid,
    FeatureGridv2
  }
}
</script>
