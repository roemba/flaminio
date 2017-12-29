import Vue from "vue";
import * as types from "./action-types";
import * as mutations from "./mutation-types";
import * as notifications from "../components/notification-types";
import {i18n} from "@/lang";

export default {

	[types.AUTHENTICATE]({dispatch, commit}, payload) {
		Vue.auth.login({
			body: {email: payload.email, password: payload.password},
			rememberMe: payload.rememberMe,
			success: () => {
				commit(mutations.UPDATE_USER, {user: Vue.auth.user()});
				dispatch(types.LOAD_LANGUAGE, Vue.auth.user().preferred_locale);
			},
			error: () => {
				Vue.router.push({name: "login", query: {invalid: true}});
			}
		});
	},

	[types.GET_LOCATIONS]({commit}, payload) {
		let uri = "locations";
		if (payload !== undefined && payload.hasOwnProperty("uuid")) {
			uri += "/" + payload.uuid;
		}

		Vue.http.get(uri).then((response) => {
			response.json().then((data) => {
				commit(mutations.UPDATE_LOCATIONS, {locations: data});
			}).catch(() => {
				commit(mutations.SHOW_NOTIFICATION, {type: notifications.CRITICAL, text: i18n.t("errors.loadLocationFailed")});
			});
		}).catch(() => {
			commit(mutations.SHOW_NOTIFICATION, {type: notifications.CRITICAL, text: i18n.t("errors.loadLocationFailed")});
		});
	},

	[types.LOAD_LANGUAGE]({state, commit}, locale) {
		if (i18n.locale !== locale) {
			if (state.loadedLanguages.indexOf(locale) === -1) {
				import(/* webpackChunkName: "lang-[request]"*/ `@/lang/${locale}`).then((msgs) => {
					i18n.setLocaleMessage(locale, msgs.default);
					commit(mutations.ADD_LOADED_LANGUAGE, {locale: locale});
					commit(mutations.CHANGE_LOCALE, {locale: locale});
				});
			} else {
				commit(mutations.CHANGE_LOCALE, {locale: locale});
			}
		}
	}
};
