import {i18n} from "@/lang";
import moment from "moment";
import * as types from "./mutation-types";

export default {

	[types.UPDATE_USER](state, payload) {
		state.user = payload.user;
	},

	[types.CHANGE_LOCALE](state, payload) {
		//TODO Check moment locales for compatibility with i18n locales, and check for lazy loading of those too
		moment.locale(payload.locale);
		state.locale = payload.locale;
		i18n.locale = payload.locale;
		document.querySelector("html").setAttribute("lang", payload.locale);
	},

	[types.ADD_LOADED_LANGUAGE](state, payload) {
		state.loadedLanguages.push(payload.locale);
	},

	[types.UPDATE_LOCATIONS](state, payload) {
		state.locations = payload.locations;
	},

	[types.SHOW_NOTIFICATION](state, payload) {
		state.notification = payload;
	},

	[types.DELETE_NOTIFICATION](state) {
		state.notification = {};
	}

};
