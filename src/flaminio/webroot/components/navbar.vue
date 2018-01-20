<template>
	<nav class="navbar navbar-expand-md navbar-dark">
		<router-link class="navbar-brand" :to="{name: 'dashboard'}">Flamin.io</router-link>
		<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"/>
		</button>
		<div class="collapse navbar-collapse" id="navbarNav">
			<ul class="navbar-nav mr-auto">
				<li :class="['nav-item', {active: $route.name === 'dashboard'}]">
					<router-link class="nav-link" :to="{name: 'dashboard'}">{{ $t("navbar.dashboardLabel") }}</router-link>
				</li>
				<li :class="['nav-item', {active: $route.name === 'schedule'}]">
					<router-link class="nav-link" :to="{name: 'schedule'}">{{ $t("navbar.scheduleLabel") }}</router-link>
				</li>
				<li :class="['nav-item', {active: $route.name === 'users'}]">
					<router-link class="nav-link" :to="{name: 'users'}">{{ $t("navbar.usersLabel") }}</router-link>
				</li>
				<li class="nav-item">
					<div class="dropdown">
						<a class="nav-link dropdown-toggle" href="#" role="button" id="dropdownSettings" data-toggle="dropdown"
							aria-haspopup="true" aria-expanded="false">{{ $t("navbar.settingsLabel") }}</a>
						<div class="dropdown-menu" aria-labelledby="dropdownSettings">
							<a class="dropdown-item" href="#">{{ $t("navbar.settings.sequences") }}</a>
							<a class="dropdown-item" href="#">{{ $t("navbar.settings.locations") }}</a>
						</div>
					</div>
				</li>
			</ul>
			<div v-if="notificationObject" :class="[notificationObject.class, 'navbar-nav', 'notification', 'align-self-stretch', 'mr-2', 'p-2']">
				<p><font-awesome-icon class="mr-1" :icon="notificationObject.icon" />{{ notificationText }}</p>
				<a class="ml-1" @click="closeNotification()" role="button" href="#"><font-awesome-icon :icon="closeIcon" /></a>
			</div>
			<div class="navbar-nav dropdown mr-1 align-self-stretch">
				<button class="btn btn-secondary dropdown-toggle" type="button" id="dropdownLanguage" data-toggle="dropdown"
					aria-haspopup="true" aria-expanded="false"><font-awesome-icon :icon="languageIcon" /></button>
				<div class="dropdown-menu dropdown-menu-right" aria-labelledby="dropdownLanguage">
					<a @click="changeLanguage('en')" class="dropdown-item language-container" href="#">
						<span class="flag-icon flag-icon-gb mr-3"/>{{ $t("navbar.languages.english") }}
					</a>
					<a @click="changeLanguage('nl')" class="dropdown-item language-container" href="#">
						<span class="flag-icon flag-icon-nl mr-3"/>{{ $t("navbar.languages.dutch") }}
					</a>
				</div>
			</div>
			<div class="navbar-nav dropdown align-self-stretch">
				<button class="btn btn-secondary dropdown-toggle" type="button" id="dropdownUser" data-toggle="dropdown"
					aria-haspopup="true" aria-expanded="false">{{ dropdownLabel }}</button>
				<div class="dropdown-menu dropdown-menu-right" aria-labelledby="dropdownUser">
					<a @click="$auth.logout()" :class="['dropdown-item', {disabled: !$auth.check()}]" href="#">{{ $t("navbar.logout") }}</a>
				</div>
			</div>
		</div>
	</nav>
</template>

<script>
import {faGlobe, faExclamationTriangle, faExclamationCircle, faTimesCircle, faTimes} from "@fortawesome/fontawesome-free-solid";
import * as actions from "../store/action-types";
import * as mutations from "../store/mutation-types";
import * as notifications from "../store/notification-types";

export default {
	computed: {
		dropdownLabel: function() {
			if (this.$auth.check()) {
				return this.$auth.user().firstname + " " + this.$auth.user().lastname;
			}
			return this.$t("navbar.loginRequest");
		},
		languageIcon: () => {
			return faGlobe;
		},
		closeIcon: () => {
			return faTimes;
		},
		notificationText: function () {
			return this.$store.state.notification.text;
		},
		notificationObject: function () {
			switch (this.$store.state.notification.type) {
			case notifications.WARNING:
				return {icon: faExclamationTriangle, class: {
					"alert-warning": true
				}};
			case notifications.INFO:
				return {icon: faExclamationCircle, class: {
					"alert-info": true
				}};
			case notifications.CRITICAL:
				return {icon: faTimesCircle, class: {
					"alert-danger": true
				}};
			}
			return false;
		}
	},
	methods: {
		changeLanguage: function (lang) {
			this.$store.dispatch(actions.LOAD_LANGUAGE, lang);
		},
		closeNotification: function () {
			this.$store.commit(mutations.DELETE_NOTIFICATION);
		}
	}
};
</script>

<style lang="scss" scoped>
	nav {
		background-color: #0A2239;
		z-index: 3;
	}

	.navbar-nav.notification {
		border-radius: $border-radius;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.navbar-nav.notification p {
		margin: 0;
		padding: 0 5px;
	}

	.dropdown-menu {
		min-width: initial;
	}
</style>
