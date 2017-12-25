<template>
	<div class="wrapper d-flex flex-column">
		<navbar class="up"/>
		<div v-if="$auth.ready()" :style="{'background-color': $store.state.viewportBackgroundColor}" class="container-fluid down">
			<router-view/>
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
		padding: 0;
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

	.container-fluid {
		padding: 15px;
	}
</style>

