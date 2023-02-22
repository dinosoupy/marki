<template>
	<section class="container">
		<p class="heading">{{ heading }}</p>
		<p class="text">{{ text }}</p>
	</section>
</template>

<script>
export default {
	name: 'FeatureCard',
	props: {
		heading: String,
		text: String
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
.heading {
	font-family: Roboto;
	font-size: 1.2em;
	font-weight: 400;
	text-align: center;
	max-width: 70%;
	padding: 1rem;
	margin: 0;
}
.text {
	font-family: Roboto;
}
</style>