<template>
	<div class="container-fluid schedule d-flex flex-column">
		<div class="toolbar p-2">
			<label for="selectedDate">Select date</label>
			<input class="form-control ml-1" type="text" v-model.lazy="selectedDateFunction" id="selectedDate">
		</div>

		<div class="schedule-container">
			<div class="schedule-top-container">
				<div class="schedule-top-container-time-container-offset"/>
				<div class="schedule-top-container-top-container-inner">
					<div class="schedule-location-container-outer">
						<div class="schedule-location-container-inner">
							<div v-for="location in locations" class="schedule-location">
								<h5>{{ location.name }}</h5>
							</div>
						</div>
					</div>
				</div>
				<div class="schedule-top-container-scrollbar-offset"/>
			</div>
			<div class="schedule-bottom-container">
				<div class="schedule-time-container" ref="timeContainer">
					<div class="schedule-time-container-inner">
						<div v-for="time in times" class="schedule-time-container-inner-entry">
							<span class="schedule-time-container-inner-entry-text">
								{{ time }}
							</span>
						</div>
					</div>
				</div>
				<div class="schedule-entry-container" @scroll="synchronizeScroll">
					<div class="flex-table">
						<div>
							<div v-for="time in times" class="horizontal-divider"/> <!-- eslint-disable-line vue/no-unused-vars -->
						</div>
						<div v-for="location in locations" class="flex-table-column">
							<div class="flex-table-column-holder">
								<div :style="reservationStyle(reservation)"
									v-for="reservation in getReservationsByLocation(location.uuid)"
									class="entry-container rounded">
									<reservationComp :reservation="reservation"/>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>

	</div>
</template>

<script>
import * as actions from "@/store/action-types";
import * as mutations from "@/store/mutation-types";
import reservationComp from "./reservation.vue";
import { mapState } from "vuex";

export default {
	components: {reservationComp},
	data: function () {
		return {
			times: [],
		};
	},
	computed: {
		...mapState({
			locations: (state) => {return state.locations;},
			reservations: (state) => {return state.schedule.reservations;},
			selectedDate: (state) => {return state.schedule.selectedDate;}
		}),
		selectedDateFunction: {
			get: function () {
				return this.selectedDate.format("L");
			},
			set: function (newValue) {
				this.$store.commit(mutations.UPDATE_SELECTED_DATE, {newValue});
				this.$store.dispatch(actions.FETCH_RESERVATIONS, {date: this.selectedDate});
			}
		}
	},
	created: function () {
		this.$store.dispatch(actions.FETCH_RESERVATIONS, {date: this.selectedDate});

		const division = 30;
		let currentTime = this.$moment("00:00", "HH:mm");

		for (currentTime; currentTime.isBefore(this.$moment("23:31", "HH:mm")); currentTime.add(division, "minutes")) {
			this.times.push(currentTime.format("HH:mm"));
		}
	},
	methods: {
		synchronizeScroll: function (event) {
			this.$refs.timeContainer.scrollTop = event.target.scrollTop;
		},
		getReservationsByLocation: function (locationUUID) {
			let reservationArray = [];
			if (this.reservations !== null) {
				for (let reservation of this.reservations) {
					if (reservation.location_id === locationUUID){
						reservationArray.push(reservation);
					}
				}
			}
			return reservationArray;
		},
		reservationStyle: function (reservation) {
			let styleObject = {
				left: "0",
				width: "100%",
				"z-index": "3"
			};
			const startTime = this.$moment(reservation.duration.start, this.ISO8601DATE_TIME);
			const endTime = this.$moment(reservation.duration.end, this.ISO8601DATE_TIME);
			const selectedTime = this.selectedDate;
			selectedTime.set({"hour": 0, "minute": 0, "second": 0});

			const rowHeightPx = 48.;
			const totalHoursDisplayed = 24.;
			const nOfMinutesInRow = 30.;
			const nOfMinutesInHour = 60.;

			const nofRows = (totalHoursDisplayed*nOfMinutesInHour)/nOfMinutesInRow;
			const minuteHeight = rowHeightPx/nOfMinutesInRow;
			const timeFraction = startTime.hour() + (startTime.minute()/nOfMinutesInHour);
			const totalHeight = rowHeightPx*nofRows;

			if (startTime.isSame(selectedTime, "day")) {
				const topOffset = Math.round(2.*rowHeightPx*timeFraction);
				styleObject.top = topOffset + "px";

				const minuteDifference = endTime.diff(startTime, "minutes");
				let height = minuteDifference*minuteHeight;
				if ((height + topOffset) > totalHeight) {
					height -= (height + topOffset - totalHeight);
				}
				styleObject.height = Math.round(height) + "px";
			} else if (endTime.isSame(selectedTime, "day")) {
				styleObject.top = "0";

				const minuteDifference = endTime.diff(selectedTime, "minutes");
				styleObject.height = Math.round(minuteDifference*minuteHeight) + "px";
			} else {
				styleObject.top = "0";
				styleObject.height = Math.round(totalHeight) + "px";
			}

			return styleObject;
		}
	}
};
</script>

<style lang="scss" scoped>
	$row-height: 48px;
	$time-column-width: 50px;

	h5 {
		margin-top: 5px;
		text-overflow: ellipsis;
		white-space: nowrap;
		overflow: hidden;
	}

	.toolbar {
		display: flex;
		align-items: center;
		width: 100%;
		background-color: $f-grey-1;
		label {
			margin: 0;
		}
		input {
			padding: 0;
			text-align: center;
			width: initial;
		}
	}

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
			justify-content: center;
			text-align: center;

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
		overflow: hidden;
	}

</style>
