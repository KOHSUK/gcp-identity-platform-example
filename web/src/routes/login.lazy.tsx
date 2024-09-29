import { createLazyFileRoute } from '@tanstack/react-router'
import { GoogleAuthProvider, signInWithPopup } from 'firebase/auth';
import { auth } from '../services/firebase/app';

export const Route = createLazyFileRoute('/login')({
  component: Login,
})

function Login() {
  const googleProvider = new GoogleAuthProvider();

  return (
    <div className="p-2">
      <h3>Login</h3>
      <button onClick={() => signInWithPopup(auth, googleProvider)}>Login with Google</button>
    </div>
  )
}