import appInsights from 'applicationinsights'

const connectionString = process.env.APPLICATIONINSIGHTS_CONNECTION_STRING

appInsights
  .setup(connectionString)
  .setAutoCollectConsole(true, true)
  .setAutoCollectHeartbeat(true)
  .setAutoCollectPerformance(true)
  .setAutoCollectExceptions(true)
  .setAutoCollectConsole(true)
  .setAutoCollectRequests(true)
  .setAutoCollectDependencies(true)
  .setDistributedTracingMode(appInsights.DistributedTracingModes.AI_AND_W3C)
  .setSendLiveMetrics(true)
  .setUseDiskRetryCaching(true)
  .start()
