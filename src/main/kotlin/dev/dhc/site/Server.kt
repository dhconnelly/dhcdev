package dev.dhc.site

import org.http4k.core.HttpHandler
import org.http4k.core.Method
import org.http4k.core.Method.GET
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
        "/healthz" bind GET to ::health,
        "/statusz" bind GET to ::status,
        "/varz" bind GET to ::vars,
        "/refresh" bind GET to ::refresh,
        "/*" bind GET to ::serve,
    )

    fun health(request: Request): Response = Response(OK)
    fun status(request: Request): Response = Response(OK)
    fun vars(request: Request): Response = Response(OK)
    fun refresh(request: Request): Response = Response(OK)
    fun serve(request: Request): Response = Response(OK)

    override fun invoke(request: Request) = CatchLensFailure.then(routes)(request)
}

