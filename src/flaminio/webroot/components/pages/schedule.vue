<template>
	<div class="container-fluid schedule d-flex flex-column">
		<h1>Schedules</h1>

		<div class="schedule-container">
			<div class="schedule-top-container">
				<div class="schedule-top-container-time-container-offset"/>
				<div class="schedule-top-container-top-container-inner">
					<div class="schedule-location-container-outer">
						<div class="schedule-location-container-inner">
							<div v-for="location in locations" class="schedule-location">
								{{ location.name }}
							</div>
						</div>
					</div>
				</div>
				<div class="schedule-top-container-scrollbar-offset"/>
			</div>
			<div class="schedule-bottom-container">
				<div class="schedule-time-container" ref="timeContainer">
					<div class="schedule-time-container-inner">
						<div v-for="(time, index) in times" class="schedule-time-container-inner-entry">
							<span class="schedule-time-container-inner-entry-text">
								{{ time }}
							</span>
						</div>
					</div>
				</div>
				<div class="schedule-entry-container" @scroll="synchronizeScroll">
					<div class="flex-table">
						<div>
							<div v-for="(time, index) in times" class="horizontal-divider"/>
						</div>
						<div v-for="location in locations" class="flex-table-column">
							<div class="flex-table-column-holder">
								<div class="entry-container rounded">Hallo!</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>

	</div>
</template>

<script>
import * as actions from "../../store/action-types";

export default {
	data () {
		return {
			reservations: [],
			times: []
		};
	},
	computed: {
		locations: function () {
			return this.$store.state.locations;
		}
	},
	created: function () {
		this.$store.dispatch(actions.GET_LOCATIONS);

		const division = 30;
		let currentTime = this.moment("00:00", "HH:mm");

		for (currentTime; currentTime.isBefore(this.moment("23:31", "HH:mm")); currentTime.add(division, "minutes")) {
			this.times.push(currentTime.format("HH:mm"));
		}
	},
	methods: {
		synchronizeScroll: function (event) {
			this.$refs.timeContainer.scrollTop = event.target.scrollTop;
		}
	}
};
</script>

<style lang="scss" scoped>
	$row-height: 48px;
	$time-column-width: 50px;

	.schedule {
		height: 100%;
		background-color: $f-blue-1;

		&-container {
			background-color: whitesmoke;
			color: black;
			height: 100%;
			display: flex;
			flex-direction: column;
			overflow: hidden;
		}

		&-top-container {
			display: flex;
			flex: none;
			border-bottom: black 1px solid;

			&-time-container-offset {
				display: flex;
				flex: none;
				flex-direction: column;
				min-width: $time-column-width;
			}

			&-top-container-inner {
				border-left: black 1px solid;
				margin-left: -1px;
				flex: 1 1 auto;
				display: flex;
				flex-direction: column;
				overflow: hidden;
			}

			&-scrollbar-offset {
				overflow: scroll;
				visibility: hidden;
				flex: none;
			}
		}
		&-location {
			border-right: black 1px solid;
			overflow: hidden;
			flex: 1 1 0;
			display: flex;

			&-container {
				&-outer {
					height: 80px;
					display: flex;
					flex: none;
					overflow: hidden;
				}
				&-inner {
					flex: 1 1 auto;
					display: flex;
					overflow: hidden;
				}
			}
		}

		&-bottom-container {
			overflow: hidden;
			display: flex;
			flex: 1 1 60%;
		}

		&-time-container {
			overflow-y: hidden;
			flex: 0 0 auto;

			&-inner {
				position: relative;
				border-right: black 1px solid;
				display: inline-block;
				min-width: $time-column-width;
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
