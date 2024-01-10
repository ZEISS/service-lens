import { Session } from 'next-auth'
import { NextResponse } from 'next/server'
import { headers } from 'next/headers'
import type { NextRequest } from 'next/server'

export const middleware = async (request: NextRequest) => {
  const { origin, protocol, host } = request.nextUrl
  const baseUrl =
    request.headers.get('x-original-proto') === 'http' && protocol === 'https:'
      ? `http://${host}`
      : origin

  const requestHeaders = new Headers(request.headers)
  requestHeaders.set('x-pathname', request.nextUrl.pathname)

  const res = await fetch(`${baseUrl}/api/auth/session`, {
    headers: {
      cookie: headers().get('cookie') ?? ''
    },
    cache: 'no-store'
  })

  const session: Session = await res.json()
  const isLoggedIn = session !== null
  const pathname = request.nextUrl.pathname

  if (pathname != '/login' && !isLoggedIn) {
    return NextResponse.redirect(new URL('/login', origin))
  }

  return NextResponse.next({
    request: {
      headers: requestHeaders
    }
  })
}

export const config = {
  matcher: [
    /*
     * Match all request paths except for the ones starting with:
     * - api (API routes)
     * - _next/static (static files)
     * - _next/image (image optimization files)
     * - favicon.ico (favicon file)
     */
    '/((?!api|_next/static|_next/image|favicon.ico).*)',
    '/dashboard/:path*',
    '/account/:path*'
  ]
}
