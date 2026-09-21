import { revalidatePath } from "next/cache"
import { headers } from "next/headers"
import type { NextRequest } from "next/server"

import { auth } from "@/lib/auth"

export async function GET(request: NextRequest) {
  const path = request.nextUrl.searchParams.get("path")

  const session = await auth.api.getSession({
    headers: await headers(),
  })

  if (!session) {
    return new Response(JSON.stringify({ error: "User is not authenticated" }), {
      status: 401,
      headers: { "Content-Type": "application/json" },
    })
  }

  if (path) {
    revalidatePath(path)
    return Response.json({ revalidated: true, now: Date.now() })
  }

  return Response.json({
    revalidated: false,
    now: Date.now(),
    message: "Missing path to revalidate",
  })
}
