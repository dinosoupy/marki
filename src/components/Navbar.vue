<template>
  <nav>
    <img src="../assets/logo.png" />
    <div class="hamburger-container">
      <input type="checkbox" class="hamburger" id="hamburger" @click="toggleMenu" />
      <label for="hamburger">
        <span id="bread-top" class="bread-top"></span>
        <span id="bread-base" class="bread-base"></span>
      </label>
    </div>
    <ul id="links-container" class="links-container-desktop">
      <li class="link"><router-link to="/">HOME</router-link></li>
      <li class="link"><router-link to="/about">WHO WE ARE</router-link></li>
      <li class="link"><router-link to="/work">OUR WORK</router-link></li>
      <li class="link"><router-link to="/contact">BOOK A CALL</router-link></li>
    </ul>
    <ul id="links-container" class="links-container-mobile">
      <li class="link"><router-link to="/" @click="tapReact">HOME</router-link></li>
      <li class="link"><router-link to="/about" @click="tapReact">WHO WE ARE</router-link></li>
      <li class="link"><router-link to="/work" @click="tapReact">OUR WORK</router-link></li>
      <li class="link"><router-link to="/contact" @click="tapReact">BOOK A CALL</router-link></li>
    </ul>
  </nav>
</template>

<script>
export default {
	name: 'Navbar',
  methods: {
    toggleMenu() {
      var breadTop = document.getElementById("bread-top");
      var breadBase = document.getElementById("bread-base");

      var checkbox = document.getElementById("hamburger");

      if (checkbox.checked) {
        breadTop.style.top = "0px";
        breadBase.style.top = "0px";
        breadTop.style.transform = "rotate(-45deg) translatey(500%) translatex(-2%)";
        breadBase.style.transform = "rotate(45deg) translatey(-500%) translatex(2%)";
        
        var openAnimation = anime({
          targets: '.links-container-mobile',
          height: 150,
          easing: 'easeInOutExpo',
          duration: 750
        })
      } 
      else {    
        breadTop.style.transform = "none";
        breadBase.style.transform = "none";
        setTimeout(function(){
          breadTop.style.top = "5px";  
          breadBase.style.top = "-5px"; 
        }, 100);
 

        var closeAnimation = anime({
          targets: '.links-container-mobile',
          height: 0,
          easing: 'easeInOutExpo',
          duration: 750
        })
      }
    },
    tapReact() {
      document.getElementById('hamburger').click();
    }
  }
} 
</script>

<style scoped>
nav {
  background-color: #333;
  position: fixed;
  width: 100%;
  top: 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0;
  z-index: 1;
}

ul {
  display: flex;
  justify-content: right;
  list-style-type: none;
  margin: 0;
  padding: 0em 1.5em;
  overflow-y: hidden
}

li {
  padding: 50px 20px;
  text-align: center;
  font-family: Roboto, sans-serif;
  font-weight: 300;
  overflow-y: hidden
}

.links-container-mobile {
  display: none;
}

a {
  color: #9e9e9e;
  text-decoration: none;
}

.router-link-exact-active {
  color: #fff
}

img {
  height: 80px;
  width: 138px;
  padding: 0.5em 1.5em 0.5em 1.5em
}

.hamburger-container {
  display: none
}

@media screen and (max-width:600px) {
  nav {
    display: block; /* Change from display: flex as row flex is no longer needed */
    padding: 0.5em;
    overflow-x: hidden;
    width: 100%;
  }

  img {
    height: 50px;
    width: 50px;
    padding: 1em
  }

  .links-container-desktop {
    display: none;
  }

  /* Set height to 0 on page load (menu should be collapsed at first) */
  .links-container-mobile {
    display: block;
    height: 0;
  }

  li {
    padding: 10px;
    font-size: 12px
  }

  ul {
    background: inherit;
    flex-direction: column
  }

  .hamburger-container {
    display: block;
    postiion: relative;
    float: right;
    padding: 2.5em 2em
  }

  .hamburger {
    display: none
  }

  .bread-top,
  .bread-base {
    display: block;
    width: 25px;
    height: 2px;
    border-radius: 10px;
    position: relative;
    background: grey;
    border-radius: 3px;
    z-index: 1;
    transform-origin: 4px 0px;
    transition: transform 0.5s cubic-bezier(0.77,0.2,0.05,1.0),
                top 0.5s cubic-bezier(0.77,0.2,0.05,1.0)
  }

  .bread-top {
    top: 5px
  }

  .bread-base {
    top: -5px
  }

}
</style>