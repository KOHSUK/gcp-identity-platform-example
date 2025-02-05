import { ActionFunctionArgs, LoaderFunctionArgs, redirect } from "@remix-run/node";
import { Form, useLoaderData } from "@remix-run/react";
import urlJoin from "url-join";
import { authenticator } from "~/utils/auth.server";
import { oauth2Api } from "~/utils/ory.server";
import { AxiosError } from 'axios'

export async function loader({
  request,
}: LoaderFunctionArgs) {
  const url = new URL(request.url);
  const codeChallenge = url.searchParams.get("login_challenge");
  const tenant = url.searchParams.get("tenant");

  if (!codeChallenge) {
    throw Response.json("Not Found", { status: 404 });
  }

  try {
    const loginRequest = await oauth2Api.getOAuth2LoginRequest({ loginChallenge: codeChallenge })
    const { skip, subject, oidc_context } = loginRequest.data

    // if hydra was already able to authenticate the user, skip will be true and we do not need to re-authenticate
    // the user.
    if (skip) {
      const acceptLoginResponse =  await oauth2Api.acceptOAuth2LoginRequest({ loginChallenge: codeChallenge, acceptOAuth2LoginRequest: { subject } })
      return redirect(acceptLoginResponse.data.redirect_to)
    }

    return {
      codeChallenge,
      action: urlJoin(process.env.BASE_URL || "", "/login"),
      hint: oidc_context?.login_hint || "",
      tenant: tenant || undefined,
    }
  } catch (error) {
    console.log('error:', error)
    throw Response.json("Something went wrong", { status: 500 });
  }
}

export async function action({
  request,
}: ActionFunctionArgs) {
  try {
    const user = await authenticator.authenticate("user-pass", request.clone());
    const body = await request.formData();
    const loginChallenge = body.get("code_challenge");

    if (!loginChallenge) {
      return redirect("/login");
    }

    const loginRequest = await oauth2Api.getOAuth2LoginRequest({
      loginChallenge: loginChallenge.toString(),
    });

    const acceptLoginResponse = await oauth2Api.acceptOAuth2LoginRequest({
      loginChallenge: loginChallenge.toString(),
      acceptOAuth2LoginRequest: {
        subject: user.uid,
        remember: body.get("remember") === "1",
        remember_for: 3600,
      },
    });

    return redirect(acceptLoginResponse.data.redirect_to);
  } catch (error) {
    console.error('error:', error);
    if (error instanceof AxiosError) {
      if (error.response?.status === 401) {
        return redirect("/login");
      }
    }

    return null;
  }
}

export default function LoginPage() {
  const data = useLoaderData<typeof loader>();

  return (
    <div className="p-4">
      <h1 className="mb-4 text-lg">Login</h1>
      <Form method="post">
        <div className="flex flex-col gap-2">
          <div>
            <label htmlFor="email">Email</label>
            <input className="border ml-4" type="email" name="email" required />
          </div>
          <div>
            <label htmlFor="password">Password</label>
            <input
              className="border ml-4"
              type="password"
              name="password"
              autoComplete="current-password"
              required
            />
          </div>
          <input type="hidden" name="code_challenge" value={data.codeChallenge} />
          <input type="hidden" name="tenant" value={data.tenant} />
          <div>
            <input type="checkbox" name="remember" id="remember" value={1} />
            <label className="ml-4" htmlFor="remember">Remember me</label>
          </div>
          <div>
            <button className="bg-blue-300 rounded px-4 py-2">Sign In</button>
          </div>
        </div>
      </Form>
    </div>
  )
}