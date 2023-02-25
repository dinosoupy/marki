<template>
	<section class="container">
		<p class="heading">{{ heading }}</p>
		<br>
		<div class="grid">
			<FeatureCard
	      heading="300+"
	      text="videos created">
	    </FeatureCard>
	    <div class="verticalLine"></div>
	    <FeatureCard
	      heading="10M+"
	      text="views across our site">
	    </FeatureCard>
	    <div class="verticalLine"></div>
	    <FeatureCard
	      heading="15+"
	      text="networks and platforms">
	    </FeatureCard>
		</div>
	</section>
</template>

<script>
	import FeatureCard from '@/components/FeatureCard.vue'

	export default {
		name: 'FeatureGrid',
		props: {
			heading: String
		},
		components: {
			FeatureCard
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
	overflow-x: hidden;
}
.grid {
	display: flex;
	flex-direction: row;
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