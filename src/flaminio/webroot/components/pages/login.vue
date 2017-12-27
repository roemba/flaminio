<template>
	<div class="outer-container">
		<div :class="['container', 'rounded', [$route.query.invalid ? 'incorrect-password' : '']]">
			<p v-if="!$route.query.invalid">{{ $t("login.header") }}</p>
			<p v-else>{{ $t("login.headerInvalid") }}</p>
			<form @submit.prevent="login()" autocomplete="on" novalidate>
				<div class="form-group">
					<label for="email">{{ $t("login.email") }}</label>
					<input v-model="email" type="email" :class="['form-control', [emailIsValid ? 'is-valid' : 'is-invalid']]" id="email" :placeholder="$t('login.emailPlaceholder')" required>
					<div v-if="!emailIsValid" class="invalid-feedback">
						{{ $t("login.invalidEmail") }}
					</div>
				</div>
				<div class="form-group">
					<label for="password">{{ $t("login.password") }}</label>
					<input v-model="password" type="password"
						:class="['form-control', [passwordIsValid ? 'is-valid' : 'is-invalid']]" id="password"
						:placeholder="$t('login.passwordPlaceholder')" required>
					<div v-if="!passwordIsValid" class="invalid-feedback">
						{{ $t("login.invalidPassword") }}
					</div>
				</div>
				<div class="d-flex justify-content-between">
					<button type="button" class="btn btn-secondary">{{ $t("login.resetPassword") }}</button>
					<div>
						<div class="form-check form-check-inline pr-4">
							<label class="form-check-label">
								<input v-model="rememberMe" class="form-check-input" type="checkbox">
								{{ $t("login.rememberMe") }}
							</label>
						</div>
						<button type="submit" class="btn btn-primary">{{ $t("login.login") }}</button>
					</div>
				</div>
			</form>
		</div>
	</div>
</template>

<script>
import * as actions from "../../store/action-types";

export default {
	data(){
		return {
			storedEmail: "",
			storedPassword: "",
			emailIsValid: false,
			passwordIsValid: false,
			rememberMe: false
		};
	},
	computed: {
		email: {
			get () {
				return this.storedEmail;
			},
			set (value) {
				const emailRe = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/; // eslint-disable-line no-useless-escape
				this.emailIsValid = !!emailRe.exec(value);
				this.storedEmail = value;
			}
		},
		password: {
			get () {
				return this.storedPassword;
			},
			set (value) {
				this.passwordIsValid = !!value;
				this.storedPassword = value;
			}
		}
	},
	methods: {
		login: function () {
			if (this.passwordIsValid && this.emailIsValid) {
				this.$store.dispatch(actions.AUTHENTICATE, {email: this.storedEmail, password: this.storedPassword,
					rememberMe: this.rememberMe});
			}
		}
	}
};
</script>

<style lang="scss" scoped>
	.outer-container {
		position: fixed;
		top: 0;
		left: 0;
		display: flex;
		align-items: center;
		width: 100%;
		height: 100%;
		background-color: $f-green-1;
	}

	.container {
		max-width: 540px;
		padding: 15px;
		background-color: $f-blue-1;
	}

	.invalid-feedback {
		color: bisque;
	}

	.incorrect-password {
		background-color: $f-red-1;
	}
</style>
