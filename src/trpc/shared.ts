declare global {
  interface BigInt {
    toJSON(): string
  }
}

function getBaseUrl() {
  if (typeof window !== 'undefined') return ''
  if (process.env.VERCEL_URL) return `https://${process.env.VERCEL_URL}`
  return 'http://localhost:3000'
}

export function getUrl() {
  return getBaseUrl() + '/api/trpc'
}

BigInt.prototype.toJSON = function () {
  return String(this)
}
