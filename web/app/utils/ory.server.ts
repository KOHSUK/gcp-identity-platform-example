import { Configuration, FrontendApi, IdentityApi, OAuth2Api } from '@ory/client'

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const baseOptions: any = {}

if (process.env.MOCK_TLS_TERMINATION) {
  baseOptions.headers = { "X-Forwarded-Proto": "https" }
}

const config = new Configuration({
  basePath: process.env.HYDRA_ADMIN_URL,
  // accessToken: process.env.ORY_API_KEY || process.env.ORY_PAT,
})

const identityApi = new IdentityApi(config)
const oauth2Api = new OAuth2Api(config)
const frontendApi = new FrontendApi(config)

export { identityApi, oauth2Api, frontendApi }