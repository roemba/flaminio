<template>
	<div class="container-fluid schedule d-flex flex-column">
		<h1>Schedules</h1>

		<div class="flex-table-row flex-table-header date-row">
			<div class="time-column"></div>
			<div v-for="(day, index) in days" class="flex-table-row-item date-header">
				{{day}}
			</div>
		</div>
		<div class="schedule-bottom-container">
			<div class="schedule-time-container" ref="timeContainer">
				<div class="schedule-time-container-inner">
					<div v-for="(time, index) in times" class="schedule-time-container-inner-entry">
						<span class="schedule-time-container-inner-entry-text">
							{{time}}
						</span>
					</div>
				</div>
			</div>
			<div class="schedule-entry-container" @scroll="synchronizeScroll">
				<div class="flex-table">
					<div>
						<div v-for="(time, index) in times" class="horizontal-divider"></div>
					</div>
					<div v-for="day in days" class="flex-table-column">
						<div class="flex-table-column-holder">
							<div class="entry-container rounded">Hallo!</div>
						</div>
					</div>
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
	methods: {
		synchronizeScroll: function (event) {
			this.$refs.timeContainer.scrollTop = event.target.scrollTop;
		}
	},
	created: function () {
		this.$store.commit(mutations.CHANGE_LOCALE,{locale: "nl"});
		this.days = this.moment.weekdays();

		const division = 30;
		let currentTime = this.moment("00:00", "HH:mm");

		for (currentTime; currentTime.isBefore(this.moment("23:31", "HH:mm")); currentTime.add(division, "minutes")) {
			this.times.push(currentTime.format("HH:mm"));
		}
	}
};
</script>

<style lang="scss" scoped>
	$row-height: 48px;

	.schedule {
		height: 100%;
		background-color: $f-blue-1;

		&-bottom-container {
			overflow: hidden;
			display: flex;
			flex: 1 1 auto;
			background-color: whitesmoke;
			color: black;
		}

		&-time-container {
			overflow-y: hidden;
			flex: 0 0 auto;

			&-inner {
				position: relative;
				border-right: black 1px solid;
				display: inline-block;
				min-width: 40px;
				padding: 0 4px;
				white-space: nowrap;

				&-entry {
					position: relative;
					height: $row-height;
					text-align: center;

					&-text {
						display: block;
						position: relative;
						top: -10px;
						line-height: 1;
					}

					&:first-child > .schedule-time-container-inner-entry-text {
						display: none;
					}
				}
			}
		}

		&-entry-container {
			width: 100%;
			overflow-y: scroll;
			overflow-x: auto;
			flex: 1 1 auto;
		}
	}

	.flex-table {
		display: flex;
		position: relative;

		&-column {
			border-right: black 1px solid;
			position: relative; //?
			padding-right: 12px;
			flex: 1 1 auto;

			&-holder {
				position: relative;
				height: 100%;
				width: 100%;
			}
		}
	}

	.horizontal-divider {
		height: $row-height;

		&::after {
			content: '';
			border-bottom: black 1px solid;
			position: absolute;
			width: 100%;
			margin-top: -1px;
			z-index: 2;
		}
	}

	.entry-container {
		position: absolute;
		outline: none;
	}

</style>
