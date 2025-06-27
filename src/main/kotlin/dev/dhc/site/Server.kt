package dev.dhc.site

import org.http4k.core.HttpHandler
import org.http4k.core.Method
import org.http4k.core.Request
import org.http4k.core.Response
import org.http4k.core.Status.Companion.OK
import org.http4k.core.then
import org.http4k.filter.ServerFilters.CatchLensFailure
import org.http4k.routing.bind
import org.http4k.routing.routes

class Context

class Server(private val context: Context): HttpHandler {
    val routes = routes(
        "/healthz" bind Method.GET to { _: Request -> Response(OK) }
    )

    override fun invoke(request: Request) = CatchLensFailure.then(routes)(request)
}

