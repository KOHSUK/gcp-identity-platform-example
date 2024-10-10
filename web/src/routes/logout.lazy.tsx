import { createLazyFileRoute } from '@tanstack/react-router'
import { signOut } from 'firebase/auth';
import { auth } from '../services/firebase/auth';

export const Route = createLazyFileRoute('/logout')({
  component: Logout,
})

function Logout() {
  const handleLogout = async () => {
    try {
      await signOut(auth);
    } catch (error) {
      console.error(error);
    }
  }

  return (
    <div className="p-2">
      <h3>Logout</h3>
      <button onClick={handleLogout}>Logout</button>
    </div>
  )
}