import { GoogleAuthProvider } from "firebase/auth";

export function getGoogleAuthProvider() {
  const provider = new GoogleAuthProvider();
  provider.addScope('profile');
  provider.addScope('email');

  return provider;
}