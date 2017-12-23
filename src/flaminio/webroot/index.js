import Vue from "vue";
import VueResource from "vue-resource";
import VueAuth from "@websanova/vue-auth";
import App from "./components/app.vue";
import store from "./store";
import router from "./routes";

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
Vue.auth.options.logoutData.redirect = "/login";

function findScrollbarWidth() {
	let scrollDiv = document.createElement("div");
	scrollDiv.style.cssText = "width: 100px;" +
		"height: 100px;" +
		"overflow: scroll;" +
		"position: absolute;" +
		"top: -9999px;";
	document.body.appendChild(scrollDiv);

	// Get the scrollbar width
	const scrollbarWidth = scrollDiv.offsetWidth - scrollDiv.clientWidth;

	// Delete the DIV
	document.body.removeChild(scrollDiv);
	return scrollbarWidth;
}

Vue.mixin({
	data () {
		return {
			moment: moment,
			scrollBarWidth: findScrollbarWidth()
		};
	}
});

new Vue({
	el: "#app",
	store,
	router,
	render: (h) => {return h(App);}
});
