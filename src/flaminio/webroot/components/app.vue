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
import * as actions from "../store/action-types";

export default {
	name: "app",
	components: {navbar: NavBar},
	created: function () {
		this.$auth.ready(function () {
			this.$store.dispatch(actions.LOAD_LANGUAGE, this.$auth.user().preferred_locale);
			this.refreshStoredData();
		});
	},
	methods: {
		refreshStoredData: function (){
			this.$store.dispatch(actions.GET_LOCATIONS);

			return new Promise((resolve) => {
				setTimeout(() => {resolve();}, this.$store.state.refreshInterval);
			}).then(() => {
				return this.refreshStoredData();
			});
		}
	}
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
	@import "../scss/flag-icon-base.scss";
	@import url('https://fonts.googleapis.com/css?family=Roboto:400,500,700');

	body {
		background-color: #132E32;
		color: white;
		font-family: 'Roboto', sans-serif;
	}

	html, body {
		height:100%;
	}

	.container-fluid {
		padding: 0 0 0 0;
	}

	button {
		cursor: pointer;
	}
</style>

