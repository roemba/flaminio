<template>
	<div class="container-fluid dashboard">
		<p>You are now viewing the dashboard {{moment.weekdays()}}</p>
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

<style lang="scss" scoped>
	.dashboard {
		height: 100%;
		background-color: $f-blue-1;
	}
</style>
