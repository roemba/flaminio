import moment from "moment";
import Vue from "vue";
import * as notifications from "../notification-types";
import {i18n} from "@/lang";
import constants from "../../global-constants";
import * as actions from "../action-types";
import * as mutations from "../mutation-types";


export default {
	state: {
		reservations: [],
		selectedDate: moment()
	},
	mutations: {
		[mutations.UPDATE_RESERVATIONS](state, payload) {
			state.reservations = payload.reservations;
		},
		[mutations.CHANGE_LOCALE](state, payload) {
			state.selectedDate.locale(payload.locale);
		},
	},
	actions: {
		[actions.FETCH_RESERVATIONS]({commit}, payload) {
			Vue.http.get("reservations?date=" + payload.date.format(constants.ISO8601DATE)).then((response) => {
				response.json().then((data) => {
					commit(mutations.UPDATE_RESERVATIONS, {reservations: data});
				}).catch((error) => {
					throw(error);
				});
			}).catch(() => {
				commit(mutations.SHOW_NOTIFICATION,
					{
						type: notifications.CRITICAL,
						text: i18n.t("errors.loadReservationFailed")
					});
			});
		},
	}
};
