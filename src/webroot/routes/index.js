import Vue from "vue";
import VueRouter from "vue-router";
import Login from "../components/pages/login.vue";
import Dashboard from "../components/pages/dashboard.vue";
import Forbidden from "../components/pages/403.vue";
import Schedule from "../components/pages/schedule.vue";

Vue.use(VueRouter);

export default new VueRouter({
	routes: [
		//{path: '*', redirect: "/"},
		{
			path: "/login",
			name: "login",
			meta: {auth: false},
			component: Login
		},
		{
			path: "/",
			name: "dashboard",
			meta: {auth: true},
			component: Dashboard
		},
		{
			path: "/schedule",
			name: "schedule",
			meta: {auth: true},
			component: Schedule
		},
		{
			path: "/403",
			name: "error-403",
			component: Forbidden
		}

	]
});
