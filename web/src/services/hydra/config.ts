import { Configuration, JwkApi } from "@ory/client"

const baseOptions: Record<string, unknown> = {}

if (process.env.MOCK_TLS_TERMINATION) {
  baseOptions.headers = { "X-Forwarded-Proto": "https" }
}

const configuration = new Configuration({
  basePath: import.meta.env.VITE_HYDRA_ADMIN_URL,
  accessToken: import.meta.env.VITE_ORY_API_KEY || import.meta.env.VITE_ORY_PAT,
  baseOptions,
})

const hydraAdmin = new JwkApi(configuration)

export { hydraAdmin }

