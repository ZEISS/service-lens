import winston from 'winston'

const { combine, timestamp, json } = winston.format

const logger = winston.createLogger({
  level: 'info',
  format: combine(timestamp(), json())
})

export { logger }
