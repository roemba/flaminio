<template>
	<div class="container-fluid schedule d-flex flex-column">
		<h1>Schedules</h1>

		<div class="flex-table-row flex-table-header date-row">
			<div class="time-column"></div>
			<div v-for="(day, index) in days" class="flex-table-row-item date-header">
				{{day}}
			</div>
		</div>
		<div class="flex-table">
			<div class="flex-table-row time-row">
				<div class="time-column"></div>
				<div v-for="(day, index) in days" class="flex-table-row-item cell">

				</div>
			</div>
			<div v-for="(time,index) in times" class="flex-table-row time-row">
				<div class="time-column"><label>{{time}}</label></div>
				<div v-for="(day, index) in days" class="flex-table-row-item cell">

				</div>
			</div>
		</div>

	</div>
</template>

<script>
import * as mutations from "../../store/mutation-types";

export default {
	data () {
		return {
			days: [],
			times: []
		};
	},
	created: function () {
		this.$store.commit(mutations.CHANGE_LOCALE,{locale: "nl"});
		this.days = this.moment.weekdays();

		const division = 30;
		let currentTime = this.moment("00:00", "HH:mm");

		for (currentTime; currentTime.isBefore(this.moment("23:31", "HH:mm")); currentTime.add(division, "minutes")) {
			this.times.push(currentTime.format("HH:mm"));
		}
		this.times.shift();
		console.log(this.times);
	}
};
</script>

<style lang="scss" scoped>
	.schedule {
		height:100%;
		background-color: $f-blue-1;
	}

	.flex-table {

		width: 100%;
		overflow-y: auto;

		display: flex;
		flex-flow: column nowrap;
		justify-content: space-between;

		&-header {
			display: none;
		}

		&-row {
			width: 100%;
			@include breakpoint(md) {
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
		background-color: $f-grey-1;
		justify-content: center;
		align-items: center;
	}

	.cell {
		background-color: #f2f2f2;
		justify-content: center;
		align-items: center;
		min-height: 50px;

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
		position: relative;
		width: 50px;

		label {
			position: absolute;
			top: 0;
			left: 0;
			right: 0;
			margin: 0;
			transform: translateY(calc(-50% - 2px));
			text-align: right;
		}
	}

</style>
