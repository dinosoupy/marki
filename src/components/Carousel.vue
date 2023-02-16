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
				{ id: 1, path: 'https://picsum.photos/200' },
				{ id: 2, path: 'https://picsum.photos/210' },
				{ id: 3, path: 'https://picsum.photos/220' },
				{ id: 4, path: 'https://picsum.photos/230' },
				{ id: 5, path: 'https://picsum.photos/240' },
				{ id: 6, path: 'https://picsum.photos/250' },
				{ id: 7, path: 'https://picsum.photos/260' },
				{ id: 8, path: 'https://picsum.photos/270' },
				{ id: 9, path: 'https://picsum.photos/280' },
				{ id: 10, path: 'https://picsum.photos/290' },
				{ id: 11, path: 'https://picsum.photos/300' },
				{ id: 12, path: 'https://picsum.photos/310' },
				{ id: 13, path: 'https://picsum.photos/320' },
				{ id: 14, path: 'https://picsum.photos/330' }
			]
		}
	},
	mounted() {
		let moveTape = anime({
			targets: '.tape',
			translateX: -1600,
			direction: 'normal',
			easing: 'linear',
			duration: 20000,
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
	width: 1600px;
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
	width: 100px;
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