<template>
	<section class="container">
		<p class="heading">{{ heading }}</p>
		<div class="carousel">
			<div class="tape" ref="tape">
				<ul>
					<li v-for="(image, index) in images" :key="index" class="slot">
				    <img :src="image" />
				  </li>
				</ul>
			</div>
			<div class="tape" ref="tape">
				<ul>
					<li v-for="(image, index) in images" :key="index" class="slot">
				    <img :src="image" />
				  </li>
				</ul>
			</div>
		</div>
	</section>
</template>

<script>
export default {
	name: 'Carousel',
	props: {
		heading: String,
		images: {
			type: Array,
			required: true,
		},
	},
	data() {
		return {
			tapeWidth: 0
		}
	},
	mounted() {
		this.tapeWidth = this.$refs.tape.offsetWidth;
		this.startCarousel();
	},
	methods: {
		startCarousel() {
			let moveTape = anime({
				targets: '.tape',
				translateX: -this.tapeWidth,
				direction: 'normal',
				easing: 'linear',
				duration: 8000,
				loop: true
			});
		}
	}
}
</script>

<style>
.container {
	display: flex;
	flex-direction: column;
	align-items: center;
	margin: 0;
	padding: 2rem;
	width: 100%;
	overflow-x: hidden;
}
.carousel {
	display: grid;
  grid-template-columns: auto auto;
	background-color: white;
	width: 100%;
	max-width: 800px;
  position: relative;
	overflow-x: hidden;
}
ul {
	display: flex;
	flex-direction: row;
	justify-content: space-evenly;
	list-style: none;
}
.slot {
	margin: 1em;
}
.slot img {
	height: 100px;
	width: auto;
}
.heading {
	font-family: Pompiere;
	font-size: 4em;
	text-align: center;
	max-width: 70%;
	padding: 1rem;
	margin: 0;
}
</style>
