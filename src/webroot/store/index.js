import Vue from "vue";
import Vuex from "vuex";
import actions from "./actions.js";
//import * as getters from './getters.js';
import mutations from "./mutations.js";

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== "production";

export default new Vuex.Store({
	actions,
	//getters,
	mutations,
	state: {
		user: {}
	},
	strict: debug
});
