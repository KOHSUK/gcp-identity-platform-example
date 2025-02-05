export const Env = {
  AUTH_ISSUER: import.meta.env.VITE_AUTH_ISSUER ? import.meta.env.VITE_AUTH_ISSUER : "",
  AUTH_CLIENT_ID: import.meta.env.VITE_AUTH_CLIENT_ID ? import.meta.env.VITE_AUTH_CLIENT_ID : "",
  AUTH_REDIRECT_URI: import.meta.env.VITE_AUTH_REDIRECT_URI ? import.meta.env.VITE_AUTH_REDIRECT_URI : "",
  AUTH_SCOPE: import.meta.env.VITE_AUTH_SCOPE ? import.meta.env.VITE_AUTH_SCOPE : "",
}