package dev.dhc.site

import org.http4k.core.HttpHandler
import org.http4k.core.Method.GET
import org.http4k.core.Request
import org.http4k.core.Response
import org.http4k.core.Status.Companion.OK
import org.http4k.core.then
import org.http4k.filter.ResponseFilters
import org.http4k.filter.ServerFilters.CatchLensFailure
import org.http4k.routing.ResourceLoader.Companion.Directory
import org.http4k.routing.bind
import org.http4k.routing.routes
import org.http4k.routing.static
import org.http4k.server.Http4kServer
import org.http4k.server.Jetty
import org.http4k.server.asServer
import java.nio.file.Path

data class Context(val dir: Path)

fun server(ctx: Context): HttpHandler {
    val routes = routes(
        "/healthz" bind GET to { _: Request -> Response(OK) },
        static(Directory(ctx.dir.toString())),
    )
    val logger = ResponseFilters.ReportHttpTransaction { tx ->
        val caller = tx.request.source
        val method = tx.request.method
        val path = tx.request.uri
        val version = tx.request.version
        val status = tx.response.status
        println("${caller?.address}:${caller?.port} $method $path $version $status")
    }
    return logger.then(routes)
}

class Servable(val serveDir: Path) {
    fun at(port: Int): Http4kServer {
        val srv = server(Context(serveDir)).asServer(Jetty(port)).start()
        println("serving $serveDir at http://localhost:${srv.port()}")
        return srv
    }
}

fun serve(serveDir: Path) = Servable(serveDir)
