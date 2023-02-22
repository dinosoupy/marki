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
			<div class="tape">
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
	mounted() {
    let moveTape = anime({
			targets: '.tape',
			translateX: -1723,
			duration: 12000,
			easing: 'linear',
			loop: true
		});
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
	width: 100%;
	max-width: 800px;
  position: relative;
	overflow-x: hidden;
}
.tape {
	padding: 1.5em 0 1.5em 0;
}
ul {
	display: flex;
	flex-direction: row;
	justify-content: space-evenly;
	list-style: none;
	margin: 0;
	padding-inline-start: 0;
}
.slot {
	margin-right: 2em;
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

