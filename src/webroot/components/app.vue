<template>
	<div class="wrapper d-flex flex-column">
		<navbar class="up"></navbar>
		<div v-if="$auth.ready()" class="container-fluid down">
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

<style lang="scss" scoped>
	.wrapper {
		height: 100%;
	}

	.down {
		display: flex;
		flex: 1;
		min-height: 0;
		background-color: #00A6D6;
		padding: 15px;
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

