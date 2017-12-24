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
			},
			error: () => {
				Vue.router.push({name: "login", query: {invalid: true}});
			}
		});
	},

	[types.GET_LOCATIONS]({commit}, payload) {
		let uri = "locations";
		if (payload.hasOwnProperty("uuid")) {
			uri += "/" + payload.uuid;
		}

		Vue.http.get(uri).then((response) => {
			response.json().then((data) => {
				commit(mutations.UPDATE_LOCATIONS, {locations: data});
			});
		}).catch((response) => {
			console.log(response.status);
			console.log("catched!");
		});
	}

};
