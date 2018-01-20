import Vue from "vue";
import Vuex from "vuex";
import actions from "./actions.js";
//import * as getters from './getters.js';
import mutations from "./mutations.js";
import schedule from "./modules/schedule";

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== "production";

export default new Vuex.Store({
	actions,
	//getters,
	mutations,
	modules: {schedule: schedule},
	state: {
		user: {},
		locale: "en",
		loadedLanguages: ["en"],
		locations: [],
		notification: {},
		refreshInterval: 300000 //= 5 minutes
	},
	strict: debug
});
