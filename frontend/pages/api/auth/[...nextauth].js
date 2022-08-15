import NextAuth from "next-auth"
import CredentialsProvider from "next-auth/providers/credentials"
import jwt from "jsonwebtoken"

const authOptions = {
    // Configure one or more authentication providers
    providers: [
        CredentialsProvider({
            id: 'username-login',
            name: "Login",
            credentials: {
                username: { label: "Username", type: "text" },
                password: { label: "Password", type: "password" }
            },
            async authorize(credentials, req) {
                const res = await fetch("http://localhost:8080/login", {
                    method: 'POST',
                    body: JSON.stringify(credentials),
                    headers: { "Content-Type": "application/json" }
                })

                const user = await res.json()

                if (!res.ok) {
                    throw new Error(user.message)
                }
                if (res.ok && user) {
                    return user
                }
                return null
            }
        })
    ],
    // Configure custom routes
    jwt: {
        maxAge: 30 * 24 * 60 * 60,
        async encode({ secret, token}) {
            return jwt.sign(token, secret)
        },
        async decode({ secret, token }) {
            return jwt.verify(token, secret)
        }
    },
    session: {
        jwt: true,
        maxAge: 30 * 24 * 60 * 60,        
    },
    // Configure custom session management
    callbacks: {
        async signIn({ user, account, profile, email, credentials }) {
            return user
        },
        async jwt({ token, user, account, profile, isNewUser }) {
            console.log(token)
            if (token?.length > 0) {
                console.log(token)
            }

            if (account && user) {
                const tk = user.data['token']
                const claims = atob(tk.split('.')[1])
                const x = JSON.parse(claims)
                const isSignIn = (user) ? true : false
                if (isSignIn) { token.auth_time = Math.floor(Date.now() / 1000) }
                return {
                    ...token,
                    exp: x['exp'],
                    accessToken: user.data['token'],
                    user: {
                        userId: x['userid'],
                        userName: x['username'],
                        exp: x['exp'],
                        iat: x['iat']
                    },
                }
            }

            return Promise.resolve(token)
        },
        async session({ session, token }) {
            session.accessToken = token.accessToken
            session.user = token.user
            return session
        }
    },
    pages: {
        signIn: "/auth/signin",

    },
    secret: process.env.JWT_SECRET,
    debug: process.env.NODE_ENV === "production",
}

export default NextAuth(authOptions)


