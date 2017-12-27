import Vue from "vue";
import VueResource from "vue-resource";
import VueAuth from "@websanova/vue-auth";
import App from "./components/app.vue";
import store from "./store";
import router from "./routes";
import {i18n} from "./lang";
import FontAwesomeIcon from "@fortawesome/vue-fontawesome";

import "popper.js";
import "bootstrap";
import moment from "moment";
Vue.use(VueResource);
Vue.http.options.root = "/api/v1";
Vue.router = router;

Vue.use(VueAuth, {
	auth: require("@websanova/vue-auth/drivers/auth/bearer.js"),
	http: require("@websanova/vue-auth/drivers/http/vue-resource.1.x.js"),
	router: require("@websanova/vue-auth/drivers/router/vue-router.2.x.js")
});

Vue.auth.options.rolesVar = "role";
Vue.auth.options.logoutData.makeRequest = true;
Vue.auth.options.logoutData.redirect = "/login";

Vue.component("FontAwesomeIcon", FontAwesomeIcon);

Vue.mixin({
	data () {
		return {
			moment: moment,
		};
	}
});

new Vue({
	el: "#app",
	store,
	router,
	i18n,
	render: (h) => {return h(App);}
});
