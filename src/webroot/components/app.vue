<template>
	<div class="wrapper d-flex flex-column">
		<navbar class="up"></navbar>
		<div v-if="$auth.ready()" class="down">
			<router-view></router-view>
		</div>
		<p v-else>Loading...</p>
	</div>
</template>

<script>
import NavBar from "./navbar.vue";

export default {
	name: "app",
	components: {navbar: NavBar}
};
</script>

<style lang="scss">
	.wrapper {
		height: 100%;
	}

	.down {
		display: flex;
		flex-grow: 1;
	}
</style>

<style lang="scss" global>
	@import "~bootstrap/scss/bootstrap.scss";
	body {
		background-color: #132E32;
		color: white;
	}

	html, body {
		height:100%;
	}
</style>

