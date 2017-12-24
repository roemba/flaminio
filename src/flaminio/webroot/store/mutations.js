import moment from "moment";
import * as types from "./mutation-types";

export default {

	[types.UPDATE_USER](state, payload) {
		state.user = payload.user;
	},

	[types.CHANGE_LOCALE](state, payload) {
		moment.locale(payload.locale);
		state.locale = payload.locale;
	},

	[types.UPDATE_LOCATIONS](state, payload) {
		state.locations = payload.locations;
	}

};
