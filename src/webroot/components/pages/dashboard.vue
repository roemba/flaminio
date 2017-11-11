<template>
	<div class="container-fluid dashboard">
		<p>You are now viewing the dashboard</p>
	</div>
</template>

<script>
export default {
	data () {
		return {
			msg: "It works!"
		};
	},
	computed: {
		testMessage: () => {
			return "This works!";
		}
	}
};
</script>

<style lang="scss">
	.dashboard {
		height: 100%;
		background-color: #00A6D6;
	}
</style>
