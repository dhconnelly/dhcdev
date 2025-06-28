package dev.dhc.site

import org.http4k.config.Environment
import org.http4k.config.EnvironmentKey
import org.http4k.lens.int
import org.http4k.lens.nonBlankString
import org.http4k.server.Jetty
import org.http4k.server.asServer

val port = EnvironmentKey.int().required("port")
val dir = EnvironmentKey.nonBlankString().required("dir")

val defaultConfig = Environment.defaults(
    port of 9000,
    dir of "./pages",
)

fun main() {
    val env = Environment.JVM_PROPERTIES overrides
        Environment.ENV overrides
        defaultConfig
    val context = Context(dir(env))
    val server = Server(context).asServer(Jetty(port(env))).start()
    println("serving ${context.directory} at http://localhost:${server.port()}")
    server.block()
}
