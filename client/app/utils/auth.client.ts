import * as client from "openid-client";
import { Env } from "~/constant/env";
import IdTokenVerifier from "idtoken-verifier";
import { Jwt } from "jsonwebtoken";

export async function discovery() {
  try {
    const config = await client.discovery(
      // Authorization Server's Issuer Identifier
      new URL(Env.AUTH_ISSUER),
      // Client identifier at the Authorization Server
      Env.AUTH_CLIENT_ID,
      // Client Secret
      undefined,
      undefined,
      {
        execute: [client.allowInsecureRequests],
      }
    );

    return config;
  } catch (error) {
    console.error("Error in discovery", error);
    return null;
  }
}

export async function getAuthorizationEndpoint(tenant?: string) {
  const config = await discovery();
  if (!config) {
    return;
  }
  const code_verifier = client.randomPKCECodeVerifier();
  // save code_verifier to localstorage for later use
  localStorage.setItem("code_verifier", code_verifier);

  const code_challenge = await client.calculatePKCECodeChallenge(code_verifier);
  const state = client.randomState();
  // save state to localstorage for later use
  localStorage.setItem("state", state);

  let nonce!: string;

  // redirect user to as.authorization_endpoint
  const parameters: Record<string, string> = {
    redirect_uri: Env.AUTH_REDIRECT_URI,
    scope: Env.AUTH_SCOPE,
    code_challenge,
    code_challenge_method: "S256",
    state,
  };

  /**
   * We cannot be sure the AS supports PKCE so we're going to use nonce too. Use
   * of PKCE is backwards compatible even if the AS doesn't support it which is
   * why we're using it regardless.
   */
  if (!config.serverMetadata().supportsPKCE()) {
    console.log("AS does not support PKCE, using nonce");
    nonce = client.randomNonce();
    localStorage.setItem("nonce", nonce);
    parameters.nonce = nonce;
  }

  if (tenant) {
    parameters.tenant = tenant;
  }

  const redirectTo = client.buildAuthorizationUrl(config, parameters);

  return redirectTo;
}

export async function authorizationCodeGrant() {
  const config = await discovery();
  if (!config) {
    return;
  }

  const currentUrl = new URL(window.location.href);
  const pkceCodeVerifier = localStorage.getItem("code_verifier") || undefined;
  const expectedState = localStorage.getItem("state") || undefined;

  const tokens = await client.authorizationCodeGrant(config, currentUrl, {
    pkceCodeVerifier,
    expectedState,
  });

  return tokens;
}

export async function verifyIdToken(idToken: string): Promise<Jwt | undefined> {
  const config = await discovery();
  if (!config) {
    return undefined;
  }

  const verifier = new IdTokenVerifier({
    issuer: Env.AUTH_ISSUER,
    audience: Env.AUTH_CLIENT_ID,
  });

  try {
    // コールバックベースのverify関数をPromiseでラップ
    const payload = await new Promise<Jwt>((resolve, reject) => {
      verifier.verify(idToken, (err, payload) => {
        if (err) {
          reject(err);
          return;
        }
        resolve(payload as Jwt);
      });
    });

    return payload;
  } catch (error) {
    console.error("Error in verifyIdToken:", error);
    return undefined;
  }
}
