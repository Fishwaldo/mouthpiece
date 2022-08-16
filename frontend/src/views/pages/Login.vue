<template>
  <div class="bg-light min-vh-100 d-flex flex-row align-items-center">
    <CContainer>
      <CRow class="justify-content-center">
        <CCol :md="8">
          <CCardGroup>
            <CCard class="p-4">
              <CCardBody>
                <CForm novalidate @submit="handleLogin">
                  <h1>Login</h1>
                  <p class="text-medium-emphasis">Sign In to your account</p>
                  <CInputGroup class="mb-3">
                    <CInputGroupText>
                      <CIcon icon="cil-user" />
                    </CInputGroupText>
                    <CFormInput
                      type="email"
                      placeholder="Username"
                      autocomplete="username"
                      v-bind:valid="this.v$.user.username.$dirty ? !this.v$.user.username.$error : null"
                      v-bind:invalid="this.v$.user.username.$dirty ? this.v$.user.username.$error : null"
                      v-model="this.v$.user.username.$model"
                    />
                    <CFormFeedback v-for="msg in this.v$.user.username.$errors" :key="msg.id" invalid>
                      {{ msg.$message }}
                    </CFormFeedback>
                  </CInputGroup>
                  <CInputGroup class="mb-4">
                    <CInputGroupText>
                      <CIcon icon="cil-lock-locked" />
                    </CInputGroupText>
                    <CFormInput
                      type="password"
                      placeholder="Password"
                      autocomplete="current-password"
                      v-bind:valid="this.v$.user.password.$dirty ? !this.v$.user.password.$error : null"
                      v-bind:invalid="this.v$.user.password.$error"
                      v-model="this.v$.user.password.$model"
                    />
                    <CFormFeedback v-for="msg in this.v$.user.password.$errors" :key="msg.id" invalid>
                      {{ msg.$message }}
                    </CFormFeedback>
                  </CInputGroup>

                  <CRow>
                    <CAlert v-if="formresult" color="danger">{{ formresult }}</CAlert>
                    <CCol :xs="6">
                      <CButton color="primary" class="px-4" type="submit" :disabled="this.v$.$invalid">
                        Login
                      </CButton>
                    </CCol>
                    <CCol :xs="6" class="text-right">
                      <CButton color="link" class="px-0">
                        Forgot password?
                      </CButton>
                    </CCol>
                  </CRow>
                </CForm>
              </CCardBody>
            </CCard>
            <CCard class="text-white bg-primary py-5" style="width: 44%">
              <CCardBody class="text-center">
                <div>
                  <h2>SSO Login</h2>
                  <p>
                    Login with your SSO Accounts to access the full
                    functionality
                  </p>
                  <CContainer>
                    <CRow class="align-items-center">
                      <CCol :xs="3" class="align-self-center" v-if="show.github">
                        <div>
                          <CButton component="a" color="light" variant="outline" class="mt-3" size="sm" href="/auth/github/login?from=/static/">
                            <CIcon icon="cibGithub"/>Github
                          </CButton>
                        </div>
                      </CCol>
                      <CCol :xs="3" class="align-self-center" v-if="show.microsoft">
                        <div>
                          <CButton component="a" color="light" variant="outline" class="mt-3" size="sm" href="/auth/microsoft/login?from=/static/">
                            <CIcon icon="cilSettings"/>Microsoft
                          </CButton>
                        </div>
                      </CCol>
                      <CCol :xs="3" class="align-self-center" v-if="show.google">
                        <div>
                          <CButton component="a" color="light" variant="outline" class="mt-3" size="sm" href="/auth/google/login?from=/static/">
                            <CIcon icon="cibGoogle"/>Google
                          </CButton>
                        </div>
                      </CCol>
                      <CCol :xs="3" class="align-self-center" v-if="show.dev">
                        <div>
                          <CButton component="a" color="light" variant="outline" class="mt-3" size="sm" href="/auth/dev/login?from=/static/">
                            <CIcon icon="cilSettings"/>Developer
                          </CButton>
                        </div>
                      </CCol>
                    </CRow>
                  </CContainer>
                </div>
              </CCardBody>
            </CCard>
          </CCardGroup>
        </CCol>
      </CRow>
    </CContainer>
  </div>
</template>

<script>
//import { Form, Field, ErrorMessage } from 'vee-validate'
import User from '@/models/user'
import SSOProviders from '@/models/ssoproviders'
import useVuelidate from '@vuelidate/core'
import { required, email } from '@vuelidate/validators'

export default {
  name: 'Login',
  setup: () => ({ v$: useVuelidate() }),
  data() {
    return {
      user: new User('', ''),
      show: new SSOProviders(Array(0)),
      formresult: ""
    }
  },
  validations () {
    return {
      user: {
        username: {
          required,
          email
        },
        password: {
          required
        }
      }
    }
  },
  mounted() {
    console.log('mounting login') 
    this.$store.dispatch('auth/providers').then(() => {
    //console.log("providers" + this.$store.state.auth.providers)
    //console.log(this.$store.state.auth.providers.github)
    this.show = this.$store.state.auth.providers
    //console.log(this.show.github)
    })
    this.$store.dispatch('auth/feconfig').then(() => {
    this.config = this.$store.state.auth.config
    console.log(this.config)
    })
  },
  computed: {
    loggedIn() {
      return this.$store.state.auth.status.loggedIn
    },
  },
  created() {
    if (this.loggedIn) {
      console.log('Logged In')
    }
  },
  methods: {
    async  handleLogin() {
      const isValid = await this.v$.$validate()
      if (!isValid) { 
        return
      }
      this.loading = true
      this.$store.dispatch('auth/login', this.user).then(
        () => {
          this.$router.push('/dashboard')
        },
        (error) => {
          this.formresult = error.response.data.error
        },
      )
    },
    handleProviderAuth(provider) {
        this.$auth.authenticate(provider)
    },
  },
}
</script>
