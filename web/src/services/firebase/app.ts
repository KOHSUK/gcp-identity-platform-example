import { initializeApp } from 'firebase/app';
import { getAuth } from 'firebase/auth';

console.log('GOOGLE_API_KEY', import.meta.env.VITE_GOOGLE_API_KEY);
console.log('GOOGLE_AUTH_DOMAIN', import.meta.env.VITE_GOOGLE_AUTH_DOMAIN);

const firebaseConfig = {
  apiKey: import.meta.env.GOOGLE_API_KEY,
  authDomain: import.meta.env.GOOGLE_AUTH_DOMAIN
};

const app = initializeApp(firebaseConfig);

export const auth = getAuth(app);