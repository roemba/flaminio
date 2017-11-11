import Vue from "vue";
import * as types from "./action-types";
import * as mutations from "./mutation-types";

export default {

	[types.AUTHENTICATE]({commit}, payload) {
		Vue.auth.login({
			body: {email: payload.email, password: payload.password},
			rememberMe: payload.rememberMe,
			success: () => {
				commit(mutations.UPDATE_USER, {user: Vue.auth.user()});
			}
		});
	}

};
