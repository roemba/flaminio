import * as types from "./mutation-types";

export default {

	[types.UPDATE_USER](state, payload) {
		state.user = payload.user;
	}

};
