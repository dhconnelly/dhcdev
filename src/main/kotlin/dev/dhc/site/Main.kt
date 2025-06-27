package dev.dhc.site

import org.http4k.server.Jetty
import org.http4k.server.asServer

fun main() {
    val server = App().asServer(Jetty(8000)).start()
    println("serving at http://localhost:${server.port()}")
    server.block()
}
