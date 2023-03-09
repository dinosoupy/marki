<template>
	<section class="container">
		<img :src="imagePath" /><br>
		<p class="heading">{{ cardHeading }}</p>
		<p class="text">{{ cardText }}</p>
	</section>
</template>

<script>
export default {
	name: 'FeatureCardv2',
	props: {
		imagePath: String,
		cardHeading: String,
		cardText: String
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
	width: 320px;
}
img {
	height: 320px;
	width: 320px;
	border-radius: 50%;
}
.heading {
	font-family: Roboto;
	font-size: 1.4em;
	font-weight: 300;
	text-align: center;
	max-width: 100%;
	padding: 1rem;
	margin: 0;
	opacity: 0.6;
}
.text {
	font-family: Roboto;
	font-size: 1em;
	opacity: 0.6;
	text-align: center;
	line-height: 1.5em;
}
</style>