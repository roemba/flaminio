<template>
	<div class="container-fluid schedule">
		<h1>Schedules</h1>
		<table class="table table-inverse">
			<thead>
				<tr>
					<th>Status</th>
					<th>Name</th>
					<th>Start enrol date</th>
					<th>End enrol date</th>
					<th>Actions</th>
				</tr>
			</thead>
		</table>
	</div>
</template>

<script>
export default {
	data () {
		return {
			msg: "It works!"
		};
	}
};
</script>

<style lang="scss">
	.schedule {
		background-color: #00A6D6;
	}
</style>
