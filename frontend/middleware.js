import { NextResponse } from "next/server";
import { signOut } from "next-auth/react";


const secret = process.env.JWT_SECRET || "secret";

export default async function middleware(req, res, next) {
    const jwt = req.cookies.get("next-auth.session-token");
    const url = req.url;
    const { pathname } = req.nextUrl;
    if (pathname.startsWith("/")) {
        
        if (jwt !== undefined){
            const claims = atob(jwt.split('.')[1])
            const x = JSON.parse(claims)
            const exp = new Date(x.user['exp'] * 1000)
            const iat = new Date(x['iat'] * 1000)
            // console.log(exp, '|', iat)
            
            if (exp < new Date() || iat > new Date()) {
                console.log('หมดอายุ')
                // req.cookies.set("next-auth.session-token", { path: "/" });
                
            }
        }
        
    }

    return NextResponse.next();
}