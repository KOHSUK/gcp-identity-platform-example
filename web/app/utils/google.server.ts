import { FirebaseOptions, initializeApp } from 'firebase/app';
import { getAuth as getFirebaseAuth, GoogleAuthProvider } from 'firebase/auth';

const firebaseConfig = {
  apiKey: process.env.GOOGLE_API_KEY,
  authDomain: process.env.GOOGLE_AUTH_DOMAIN,
} satisfies FirebaseOptions;


export function getAuth(tenantId?: string) {
  const app = initializeApp(firebaseConfig);
  const auth = getFirebaseAuth(app)

  if (tenantId) {
    auth.tenantId = tenantId;
  }

  return auth;
}

export function getGoogleAuthProvider() {
  const provider = new GoogleAuthProvider();
  provider.addScope('profile');
  provider.addScope('email');

  return provider;
}

