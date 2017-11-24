<template>
	<div class="schedule d-flex flex-column">
		<h1>Schedules</h1>

		<div class="flex-table">
			<div class="flex-table-row flex-table-header date-row">
				<div class="flex-table-row-item time-column"></div>
				<div v-for="(day, index) in days" class="flex-table-row-item date-header">
					{{day}}
				</div>
			</div>
			<div v-for="index in 48" class="flex-table-row time-row">
				<div class="flex-table-row-item time-column"></div>
				<div v-for="(day, index) in days" class="flex-table-row-item cell">

				</div>
			</div>
		</div>

	</div>
</template>

<script>
export default {
	data () {
		return {
			msg: "It works!",
			days: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"]
		};
	}
};
</script>

<style lang="scss" scoped>

	$md: 500px;

	.schedule {
		height:100%;
		width: 100%;
	}
	.flex-table {

		width: 100%;
		overflow-y: auto;

		display: flex;
		flex-flow: column nowrap;
		justify-content: space-between;
		background-color: white;

		&-header {
			display: none;
		}

		&-row {
			width: 100%;
			@media(min-width: $md) {
				display: flex;
				flex-flow: row nowrap;
				flex-grow: 1;
				flex-basis: 0;
			}
			&-item {
				display: flex;
				flex-flow: nowrap;
				flex-grow: 1;
				flex-basis: 0;
				padding: 0.5em;

			}
		}
	}

	@mixin fix-flex() {
		flex-grow: initial;
		flex-basis: initial;
	}


	@for $i from 1 through 10 {
		.flex-grow-#{$i} {
			flex-grow: $i;
		}
	}

	.date-header {
		background-color: #4F5D75;
		justify-content: center;
		align-items: center;
	}

	.cell {
		background-color: #f2f2f2;
		justify-content: center;
		align-items: center;
		height: 50px;

		border: 1px solid black;
		&:not(:last-child){
			margin-right: -1px;
		}
	}

	.time-row {
		&:not(:last-child){
			margin-bottom: -1px;
		}
	}

	.date-row {
		height: 50px;
		@include fix-flex();
	}

	.time-column {
		background-color: #00a6d6;
		width: 50px;
		@include fix-flex();
	}

</style>
