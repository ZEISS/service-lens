import { logger } from '@/lib/logger'
import { NextRequest, NextResponse } from 'next/server'

function withErrorHandler(
  fn: (request: NextRequest, ...args: any[]) => Promise<NextResponse>
) {
  return async function (request: NextRequest, ...args: any[]) {
    try {
      throw new Error('test')
    } catch (error) {
      // Log the error to a logging system
      logger.error({ error, requestBody: request, location: fn.name })
      // Respond with a generic 500 Internal Server Error
      return NextResponse.json(
        { message: 'Internal Server Error' },
        { status: 500 }
      )
    }
  }
}

export default withErrorHandler
