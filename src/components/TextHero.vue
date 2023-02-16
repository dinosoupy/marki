<template>
	<section class="container">
		<p class="heading">{{ heading }}</p>
		<p class="subheading">{{ subheading }}</p>
		<button class="callToAction">{{ ctaText }}</button>
	</section>
</template>

<script>
export default {
	name: 'TextHero',
	props: {
		heading: String,
		subheading: String,
		ctaText: String
	}
}
</script>

<style scoped>
.container {
	background: orange;
	display: flex;
	flex-direction: column;
	align-items: center;
	margin: 0;
	padding: 2rem;
}
.heading {
	font-family: Pompiere;
	font-size: 4em;
	text-align: center;
	max-width: 70%;
	padding: 1rem;
	margin: 0;
}
.subheading {
	font-family: Roboto;
	font-size: 1.2em;
	text-align: center;
	max-width: 70%;
	padding: 1rem;
	margin: 0;
}
.callToAction {
	background: black;
	color: white;
	font-size: 1em;
	font-weight: 400;
	border-radius: 50px;
	margin: 1rem;
	padding: 0.8rem 1.8rem;
}
@media screen and (max-width:600px) {
	.container {
		padding: 2rem 1.5rem;
	}
	.heading {
		max-width: 100%;
		font-size: 3em;
	}
}
</style>