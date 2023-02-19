<template>
	<section class="container">
		<p class="heading">{{ heading }}</p>
		<div class="carousel">
			<div class="tape">
				<ul>
					<li v-for="image in images" class="slot">
				      <img :src="image.path" />
				    </li>
				</ul>
			</div>
			<div class="tape">
				<ul>
					<li v-for="image in images" class="slot">
				      <img :src="image.path" />
				    </li>
				</ul>
			</div>
		</div>
	</section>
</template>

<script>
export default {
	name: 'Carousel',
	data() {
		return {
			images: [
				{ id: 1, path: 'images/cna.png' },
				{ id: 2, path: 'images/dell.png' },
				{ id: 3, path: 'images/discovery.png' },
				{ id: 4, path: 'images/fujitsu.png' },
				{ id: 5, path: 'images/milliman.png' },
				{ id: 6, path: 'images/natgeo.png' },
				{ id: 7, path: 'images/phuel.png' },
				{ id: 8, path: 'images/primo.png' },
				{ id: 9, path: 'images/sap.png' }
			]
		}
	},
	mounted() {
		let moveTape = anime({
			targets: '.tape',
			translateX: -2000,
			direction: 'normal',
			easing: 'linear',
			duration: 8000,
			loop: true
		})
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
.tape {
/*	width: 1600px;
  	position: relative;
	overflow-x: hidden;*/
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