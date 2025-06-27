package dev.dhc.site

import org.http4k.core.HttpHandler
import org.http4k.core.Request
import org.http4k.core.Response
import org.http4k.core.Status.Companion.OK

class App: HttpHandler {
    override fun invoke(request: Request): Response =
        Response(OK).body("Hello, ${request.header("user-agent") ?: "unknown"}")
}
