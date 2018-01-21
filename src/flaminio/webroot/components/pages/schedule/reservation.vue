<template>
	<div class="entry w-100 h-100 p-2" :style="{'background-color': '#'+reservation.color, color: textColor}">
		<h6>{{ reservation.name }}</h6>
		<p>{{ timeStamp }}</p>
		<p>{{ reservation.description }}</p>
	</div>
</template>

<script>
import {blackOrWhite} from "@/utility";

export default {
	props: {
		reservation: {
			type: Object,
			required: true
		}
	},
	computed: {
		timeStamp: function() {
			const start = this.$moment(this.reservation.duration.start, this.ISO8601DATE_TIME).locale(this.$store.state.locale);
			const end = this.$moment(this.reservation.duration.end, this.ISO8601DATE_TIME).locale(this.$store.state.locale);
			if (start.isSame(end, "day")){
				return start.format("HH:mm") + " - " + end.format("HH:mm");
			}
			return start.format("L HH:mm") + " - " + end.format("L HH:mm");
		},
		textColor: function() {
			return blackOrWhite("#" + this.reservation.color);
		}
	}
};
</script>

<style lang="scss" scoped>
	.entry {
		cursor: pointer;
	}

	h6 {
		font-weight: 500;
		margin-bottom: 0;
	}
</style>
