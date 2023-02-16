<script>
import Navbar from './components/Navbar.vue'

export default {
  name: 'App',
  components: {
      Navbar
  }
}
</script>

<template>
  <Navbar></Navbar>
  <div class="routerView">
    <router-view/>
  </div>
</template>

<style>
body {
  margin: 0px;
}
.routerView {
  overflow: hidden;
  margin-top: 118px; /* Margin top is needed to offset the entire page content div below the navbar on page load. Navbar is a position: fixed type element so that content can go behind it when scrolling. This has the side effect of making normal content ignore navbar and hide behind it. For future reference, do not set a contant height property to navbar element - it will break the drop down animation. Always find out what height it happens to be and set margin-top here accordingly.*/
}
@media screen and (max-width:600px) {
  .routerView {
    margin-top: 80px; 
  }
}
</style>

