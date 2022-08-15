import { useSession, signIn, signOut } from "next-auth/react"
import { getToken } from "next-auth/jwt"


export default function Home() {
  const { data: session } = useSession()
  
  
  if (!session) {
    return (
      <div>
        <h1>Sign in</h1>
        <button onClick={signIn}>Sign in</button>
      </div>
    )
  } else {
    return (
      <div>
        <h1>Welcome {session.user.name}</h1>
        <button onClick={signOut}>Sign out</button>
      </div>
    )
  }

}
