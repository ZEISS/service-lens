import type { NextRequest } from 'next/server.js'
import winston from 'winston'

const { combine, timestamp, json } = winston.format

const logger = winston.createLogger({
  level: 'info',
  format: combine(timestamp(), json()),
  transports: [
    new winston.transports.Console({
      format: combine(timestamp(), json())
    })
  ]
})

export const logRequest = (
  req: NextRequest,
  params: unknown,
  next: () => void
) => {
  logger.info(`${req.method} ${req.url}`)

  return next()
}
