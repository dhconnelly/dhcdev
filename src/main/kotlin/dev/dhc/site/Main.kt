package dev.dhc.site

import org.http4k.config.Environment
import org.http4k.config.Environment.Companion.ENV
import org.http4k.config.Environment.Companion.JVM_PROPERTIES
import org.http4k.config.EnvironmentKey
import org.http4k.lens.boolean
import org.http4k.lens.int
import org.http4k.lens.nonBlankString
import org.http4k.server.Jetty
import org.http4k.server.asServer
import kotlin.io.path.ExperimentalPathApi
import kotlin.io.path.Path
import kotlin.io.path.copyToRecursively
import kotlin.io.path.createTempDirectory

val port = EnvironmentKey.int().required("port")
val dir = EnvironmentKey.nonBlankString().required("dir")
val build = EnvironmentKey.boolean().defaulted("build", true)
val serve = EnvironmentKey.boolean().defaulted("serve", true)
val defaults = Environment.defaults(
    port of 9000,
    dir of "./pages",
)

@OptIn(ExperimentalPathApi::class)
fun main() {
    val env = JVM_PROPERTIES overrides ENV overrides defaults
    val buildDir = Path(dir(env))

    val serveDir = if (build(env)) {
        val serveDir = createTempDirectory()
        println("building $buildDir to $serveDir")
        buildDir.copyToRecursively(serveDir, followLinks=false)
        println("built $buildDir to $serveDir")
        serveDir
    } else {
        buildDir
    }

    if (serve(env)) {
        val ctx = Context(dir = serveDir)
        val server = Server(ctx).asServer(Jetty(port(env))).start()
        println("serving ${ctx.dir} at http://localhost:${server.port()}")
        server.block()
    }
}
