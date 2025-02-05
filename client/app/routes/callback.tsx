import { useLoaderData } from "@remix-run/react";
import { Jwt } from "jsonwebtoken";
import { authorizationCodeGrant, verifyIdToken } from "~/utils/auth.client";

export async function clientLoader() {
  const tokens = await authorizationCodeGrant();

  let verifiedIdToken: Jwt | undefined;

  if (tokens?.id_token) {
    verifiedIdToken = await verifyIdToken(tokens.id_token);
  }

  return {
    tokens,
    verifiedIdToken,
  }
}

export default function Page() {
  const { tokens, verifiedIdToken } = useLoaderData<typeof clientLoader>(); 

  return (
    <div>
      <div>Tokens</div>
      {JSON.stringify(tokens)}
      <div>
        <div>ID Token:</div>
        <div>{verifiedIdToken && JSON.stringify(verifiedIdToken)}</div>
      </div>
    </div>
  );
}
