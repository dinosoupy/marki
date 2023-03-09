<template>
	<section class="container">
		<p class="heading">{{ heading }}</p>
		<br>
		<ul class="grid">
			<li v-for="(card, index) in listOfCards" :key="index" class="card">
				<FeatureCardv2
					:imagePath="card[0]"
					:cardHeading="card[1]"
					:cardText="card[2]">
				</FeatureCardv2>
			</li>
		</ul>
	</section>
</template>

<script>
	import FeatureCardv2 from '@/components/FeatureCardv2.vue'

	export default {
		name: 'FeatureGridv2',
		props: {
			heading: String,
			listOfCards: {
				type: Array,
				required: true,
			},
		},
		components: {
			FeatureCardv2
		}
	}
</script>

<style scoped>
.container {
	display: flex;
	flex-direction: column;
	align-items: center;
	margin: 0;
	padding: 2rem;
}
.grid {
	display: flex;
  flex-direction: row;
  list-style: none;
	margin: 0;
	padding: 2rem;
	justify-content: space-evenly;
	width: 100%;
}
.heading {
	font-family: Pompiere;
	font-size: 3em;
	color: orange;
	text-align: center;
	max-width: 70%;
	padding: 1rem;
	margin: 0;
}
.verticalLine {
	border-left: 1px solid;
	color: black;
	opacity: 10%;
  height: 300px;
}
</style>