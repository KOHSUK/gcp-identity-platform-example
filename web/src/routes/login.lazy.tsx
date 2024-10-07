import { createLazyFileRoute } from '@tanstack/react-router'
import { GoogleAuthProvider, signInWithPopup } from 'firebase/auth';
import { auth } from '../services/firebase/auth';
import { getGoogleAuthProvider } from '../services/firebase/google';

export const Route = createLazyFileRoute('/login')({
  component: Login,
})

function Login() {
  const googleProvider = getGoogleAuthProvider();

  const handleLogin = async () => {
    try {
      const result = await signInWithPopup(auth, googleProvider);
      // The signed-in user info.
      const user = result.user;
      // This gives you a Facebook Access Token.
      const credential = GoogleAuthProvider.credentialFromResult(result);
      const token = credential?.accessToken;

      console.log("user", user);
      console.log("token", token);
    } catch (error) {
      console.error(error);
    }
  }

  return (
    <div className="p-2">
      <h3>Login</h3>
      <button onClick={handleLogin}>Login with Google</button>
    </div>
  )
}