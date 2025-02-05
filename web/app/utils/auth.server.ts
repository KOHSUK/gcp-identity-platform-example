import { FormStrategy } from "remix-auth-form";
import { Authenticator } from "remix-auth";
import { sessionStorage } from "./session.server";
import { getAuth } from "./google.server";
import { signInWithEmailAndPassword, User } from "firebase/auth";

// Create an instance of the authenticator, pass a generic with what
// strategies will return and will store in the session
export const authenticator = new Authenticator<User>(sessionStorage);

// Tell the Authenticator to use the form strategy
authenticator.use(
  new FormStrategy(async ({ form }) => {
    const email = form.get("email");
    const password = form.get("password");

    if (!email || !password) {
      throw new Error("The username / password combination is not correct")
    }

    const tenant = form.get("tenant");

    const auth = getAuth(tenant?.toString());

    const result = await signInWithEmailAndPassword(auth, email.toString(), password.toString());

    const user = result.user;

    // the type of this user must match the type you pass to the Authenticator
    // the strategy will automatically inherit the type if you instantiate
    // directly inside the `use` method
    return user;
  }),
  // each strategy has a name and can be changed to use another one
  // same strategy multiple times, especially useful for the OAuth2 strategy.
  "user-pass"
);