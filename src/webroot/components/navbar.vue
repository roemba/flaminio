<template>
	<nav class="navbar navbar-expand-md navbar-dark">
		<router-link class="navbar-brand" :to="{name: 'dashboard'}">Flamin.io</router-link>
		<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
		<div class="collapse navbar-collapse" id="navbarNav">
			<ul class="navbar-nav mr-auto">
				<li :class="['nav-item', {active: $route.name === 'dashboard'}]">
					<router-link class="nav-link" :to="{name: 'dashboard'}">Dashboard</router-link>
				</li>
				<li :class="['nav-item', {active: $route.name === 'schedule'}]">
					<router-link class="nav-link" :to="{name: 'schedule'}">Schedule</router-link>
				</li>
				<li class="nav-item">
					<div class="dropdown">
						<a class="nav-link dropdown-toggle" href="#" role="button" id="dropdownSettings" data-toggle="dropdown"
							aria-haspopup="true" aria-expanded="false">Settings</a>
						<div class="dropdown-menu" aria-labelledby="dropdownSettings">
							<a class="dropdown-item" href="#">Courses</a>
							<a class="dropdown-item" href="#">Rooms</a>
							<a class="dropdown-item" href="#">Users</a>
						</div>
					</div>
				</li>
			</ul>
			<div class="navbar-nav dropdown">
				<button class="btn btn-secondary dropdown-toggle" type="button" id="dropdownUser" data-toggle="dropdown"
					aria-haspopup="true" aria-expanded="false">{{dropdownLabel}}</button>
				<div class="dropdown-menu dropdown-menu-right" aria-labelledby="dropdownUser">
					<a @click="$auth.logout()" :class="['dropdown-item', {disabled: !$auth.check()}]" href="#">Logout</a>
				</div>
			</div>
		</div>
	</nav>
</template>

<script>
export default {
	computed: {
		dropdownLabel: function() {
			if (this.$auth.check()) {
				return this.$auth.user().firstname + " " + this.$auth.user().lastname;
			}
			return "Please login";
		}
	}
};
</script>

<style lang="scss">
	nav {
		background-color: #0A2239;
	}
</style>
