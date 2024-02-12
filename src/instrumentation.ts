export async function register() {
  if (
    process.env.NEXT_RUNTIME === 'nodejs' &&
    process.env.APPLICATIONINSIGHTS_CONNECTION_STRING
  ) {
    await import('./lib/insights')
  }
}
