import * as fs from "node:fs"
import * as path from "node:path"

// Define paths
const packageJsonPath = path.join(__dirname, "package.json")
const chartYamlPath = path.join(__dirname, "helm/charts/service-lens/", "Chart.yaml")

interface PackageJson {
  name: string
  version: string
  description?: string
  main?: string
  scripts?: { [key: string]: string }
  keywords?: string[]
  author?: string
  license?: string
  dependencies?: { [key: string]: string }
  devDependencies?: { [key: string]: string }
  [key: string]: any // Allow for additional properties
}

// Function to update Chart.yaml with npm version
const updateChartVersion = async () => {
  try {
    // Read package.json
    const packageJsonData = await fs.promises.readFile(packageJsonPath, "utf8")
    const packageJson = JSON.parse(packageJsonData) as PackageJson
    const npmVersion = packageJson.version

    // Read Chart.yaml
    const chartYamlData = await fs.promises.readFile(chartYamlPath, "utf8")

    // Update the version in Chart.yaml
    const updatedYaml = chartYamlData.replace(/version: .*/, `version: ${npmVersion}`)

    // Update the appVersion in Chart.yaml
    const finalYaml = updatedYaml.replace(/appVersion: .*/, `appVersion: ${npmVersion}`)

    // Write back updated Chart.yaml
    await fs.promises.writeFile(chartYamlPath, finalYaml, "utf8")
    console.log(`Updated Chart.yaml to version ${npmVersion}`)
  } catch (error) {
    console.error("Error:", error)
  }
}

// Execute the function
updateChartVersion()
