import { AcceptOAuth2ConsentRequestSession } from "@ory/client";
import { ActionFunctionArgs, LoaderFunctionArgs, redirect } from "@remix-run/node";
import { Form, useLoaderData } from "@remix-run/react";
import { oauth2Api } from "~/utils/ory.server";

export async function loader({ request }: LoaderFunctionArgs) {
  const url = new URL(request.url);
  const consentChallenge = url.searchParams.get("consent_challenge");

  if (!consentChallenge) {
    throw new Error("Not Found");
  }

  const consentRequest = await oauth2Api.getOAuth2ConsentRequest({ consentChallenge});
  const { skip, client, requested_scope, requested_access_token_audience, subject } = consentRequest.data;

  if (skip || client?.skip_consent) {
    const acceptConsentResponse = await oauth2Api.acceptOAuth2ConsentRequest({ consentChallenge, acceptOAuth2ConsentRequest: {
      grant_scope: requested_scope,
      grant_access_token_audience: requested_access_token_audience,
      session: {
        access_token: {},
        id_token: {},
      },
    } });

    return redirect(acceptConsentResponse.data.redirect_to);
  }

  return {
    consentChallenge,
    requested_scope,
    user: subject,
    client,
  }

}

export async function action({ request }: ActionFunctionArgs) {
  const body = await request.formData();
  const consentChallenge = body.get("challenge");

  // To get the user from the form data, you can use the following code:
  // const user = body.get("user");

  if (!consentChallenge) {
    return redirect("/login");
  }

  if (body.get("submit") === "Deny access") {
    const rejectConsentRequest = await oauth2Api.rejectOAuth2ConsentRequest({ consentChallenge: consentChallenge.toString(), rejectOAuth2Request: {
      error: "access_denied",
      error_description: "The resource owner denied the request",
    } });

    return redirect(rejectConsentRequest.data.redirect_to);
  }

  try {
    const grantScope = body.getAll("grant_scope").map((scope) => scope.toString());

    const session = {
      // You can add more data to the session object if needed
      access_token: {},
      id_token: {
        tenant: 'test-tenant'
      },
    } satisfies AcceptOAuth2ConsentRequestSession;

    const consentRequest = await oauth2Api.getOAuth2ConsentRequest({ consentChallenge: consentChallenge.toString() });
    const acceptConsentRequest = await oauth2Api.acceptOAuth2ConsentRequest({ consentChallenge: consentChallenge.toString(), acceptOAuth2ConsentRequest: {
        grant_scope: grantScope,
        grant_access_token_audience: consentRequest.data.requested_access_token_audience,
        session,
      }
    });

    return redirect(acceptConsentRequest.data.redirect_to);
  } catch (error) {
    console.error('error:', error);

    return null;
  }
}

export default function ConsentPage() {
  const data = useLoaderData<typeof loader>();

  return (
    <div>
      <h1>Consent Page</h1>
      <Form method="POST">
        <input type="hidden" name="challenge" value={data.consentChallenge} />
        <input type="hidden" name="user" value={data.user} />
        {data.client?.logo_uri && <img src={data.client.logo_uri} alt={data.client.client_name} />}
        <p>Client: {data.client?.client_name}</p>
        <div>
          <h2>Requested Permissions</h2>
          {data.requested_scope?.map((scope) => (
            <div key={scope}>
              <input type="checkbox" name="grant_scope" value={scope} />
              <label htmlFor={scope}>{scope}</label>
            </div>
          ))}
        </div>
        <p>Do you want to be asked next time when this application wants to access your data? The application will
            not be able to ask for more permissions without your consent.</p>
        <ul>
          {data.client?.policy_uri && <li><a
            href={data.client.policy_uri}
            target="_blank"
            rel="noreferrer"
          >Read the privacy policy</a></li>}
          {data.client?.tos_uri && <li><a
            href={data.client.tos_uri}
            target="_blank"
            rel="noreferrer"
          >Read the terms of service</a></li>}
        </ul>
        <div>
          <input type="checkbox" name="remember" id="remember" value={1} />
          <label className="ml-4" htmlFor="remember">Do not ask me again</label>
        </div>
        <div className="flex gap-2">
          <button className="bg-blue-300 rounded px-4 py-2" id="accept" type="submit" name="submit" value="Allow access">Accept</button>
          <button className="bg-red-300 rounded px-4 py-2" id="reject" type="submit" name="submit" value="Deny access">Deny</button>
        </div>
      </Form>
    </div>
  )
}