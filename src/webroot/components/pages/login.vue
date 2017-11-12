<template>
	<div class="outer-container">
		<div :class="['container', 'rounded', [$route.query.invalid ? 'incorrect-password' : '']]">
			<p v-if="!$route.query.invalid">Please login using your provided user credentials</p>
			<p v-else>Invalid username or password. Please try again.</p>
			<form @submit.prevent="login()" autocomplete="on" novalidate>
				<div class="form-group">
					<label for="email">Email address</label>
					<input v-model="email" type="email" :class="['form-control', [emailIsValid ? 'is-valid' : 'is-invalid']]" id="email" placeholder="Enter email" required>
					<div v-if="!emailIsValid" class="invalid-feedback">
						Please provide a valid email
					</div>
				</div>
				<div class="form-group">
					<label for="password">Password</label>
					<input v-model="password" type="password" :class="['form-control', [passwordIsValid ? 'is-valid' : 'is-invalid']]" id="password" placeholder="Password" required>
					<div v-if="!passwordIsValid" class="invalid-feedback">
						Please provide a valid password
					</div>
				</div>
				<div class="d-flex justify-content-between">
					<button type="button" class="btn btn-secondary">Reset password</button>
					<div>
						<div class="form-check form-check-inline pr-4">
							<label class="form-check-label">
								<input v-model="rememberMe" class="form-check-input" type="checkbox">
								Remember me
							</label>
						</div>
						<button type="submit" class="btn btn-primary">Login</button>
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

<style lang="scss">
	.outer-container {
		position: fixed;
		top: 0;
		left: 0;
		display: flex;
		align-items: center;
		width: 100%;
		height: 100%;
	}

	.container {
		max-width: 540px;
		padding: 15px;
		background-color: #00A6D6;
	}

	.invalid-feedback {
		color: bisque;
	}

	.incorrect-password {
		background-color: #ef2b2b;
	}
</style>
