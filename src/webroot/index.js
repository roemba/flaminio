import Vue from "vue";
import VueResource from "vue-resource";
import VueAuth from "@websanova/vue-auth";
import App from "./components/app.vue";
import store from "./store";
import router from "./routes";

import "popper.js";
import "bootstrap";
Vue.use(VueResource);
Vue.http.options.root = "/api/v1";
Vue.router = router;

Vue.use(VueAuth, {
	auth: require("@websanova/vue-auth/drivers/auth/bearer.js"),
	http: require("@websanova/vue-auth/drivers/http/vue-resource.1.x.js"),
	router: require("@websanova/vue-auth/drivers/router/vue-router.2.x.js")
});

Vue.auth.options.rolesVar = "role";
Vue.auth.options.logoutData.redirect = "/login";


new Vue({
	el: "#app",
	store,
	router,
	render: (h) => {return h(App);}
});
