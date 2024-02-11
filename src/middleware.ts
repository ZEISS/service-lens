import { Session } from 'next-auth'
import { NextResponse } from 'next/server'
import { headers } from 'next/headers'
import type { NextRequest } from 'next/server'
import { cookies } from 'next/headers'
import { URLPattern } from 'urlpattern-polyfill'

const PATTERNS = [
  {
    pattern: new URLPattern({ pathname: '/teams/:team' }),
    handler: ({ pathname }: { pathname: URLPatternComponentResult }) =>
      pathname.groups
  }
]

const params = (url: string) => {
  const input = url.split('?')[0]

  for (const { pattern, handler } of PATTERNS) {
    const patternResult = pattern.exec(input)
    if (patternResult !== null && 'pathname' in patternResult) {
      return handler(patternResult)
    }
  }

  return {}
}

export const middleware = async (request: NextRequest) => {
  const { origin, protocol, host } = request.nextUrl
  const isHttps =
    request.headers.get('x-original-proto') === 'http' && protocol === 'https:'
  const baseUrl = process.env.BASE_URL ?? isHttps ? `http://${host}` : origin

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

  if (!pathname.startsWith('/login') && !isLoggedIn) {
    return NextResponse.redirect(new URL('/login', origin))
  }

  const cookiesList = cookies()
  const hasScope = cookiesList.has('scope')

  if (!hasScope && isLoggedIn && !pathname.startsWith('/home')) {
    return NextResponse.redirect(new URL(`/home`, origin), {
      status: 302
    })
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
    '/teams/:path*',
    '/settings/:path*',
    '/home/:path*',
    '/account/:path*'
  ]
}
