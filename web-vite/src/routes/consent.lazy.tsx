import { createLazyFileRoute } from '@tanstack/react-router'

export const Route = createLazyFileRoute('/consent')({
  component: Consent,
})

function Consent() {
  const handleLogin = async () => {
    console.log('consent');
  }

  return (
    <div className="p-2">
      <h3>Consent</h3>
      <button onClick={handleLogin}>Consent</button>
    </div>
  )
}