export default class SSOProviders {
  constructor(providers) {
    this.github = providers.includes("github")
    this.dev = providers.includes("dev")
    this.google = providers.includes("google")
    this.microsoft = providers.includes("microsoft")
  }
}
