<template>
	<div class="container-fluid dashboard">
		<h1 v-t="{path: 'dashboard.welcome', args: {name:$auth.user().firstname}}"/>
	</div>
</template>

<script>
export default {

};
</script>

<style lang="scss" scoped>
	.dashboard {
		height: 100%;
		background-color: $f-blue-1;
	}
</style>
