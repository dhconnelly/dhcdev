package dev.dhc.site

import org.http4k.config.Environment
import org.http4k.config.EnvironmentKey
import org.http4k.lens.int
import org.http4k.server.Jetty
import org.http4k.server.asServer

val port = EnvironmentKey.int().required("port")

val defaultConfig = Environment.defaults(
    port of 9000,
)

fun main() {
    val env = Environment.JVM_PROPERTIES overrides
        Environment.ENV overrides
        defaultConfig
    val server = Server(Context()).asServer(Jetty(port(env))).start()
    println("serving at http://localhost:${server.port()}")
    server.block()
}
