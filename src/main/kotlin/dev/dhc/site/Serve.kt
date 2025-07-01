package dev.dhc.site

import org.http4k.core.HttpHandler
import org.http4k.core.Method.GET
import org.http4k.core.Request
import org.http4k.core.Response
import org.http4k.core.Status.Companion.OK
import org.http4k.core.then
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

class Server(ctx: Context): HttpHandler {
    val routes = routes(
        "/healthz" bind GET to { _: Request -> Response(OK) },
        static(Directory(ctx.dir.toString())),
    )

    override fun invoke(request: Request) = CatchLensFailure.then(routes)(request)
}

class Servable(val serveDir: Path) {
    fun at(port: Int): Http4kServer {
        val srv = Server(Context(serveDir)).asServer(Jetty(port)).start()
        println("serving $serveDir at http://localhost:${srv.port()}")
        return srv
    }
}

fun serve(serveDir: Path) = Servable(serveDir)
