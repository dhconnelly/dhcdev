package dev.dhc.site

import org.http4k.config.Environment
import org.http4k.config.Environment.Companion.ENV
import org.http4k.config.Environment.Companion.JVM_PROPERTIES
import org.http4k.config.EnvironmentKey
import org.http4k.lens.boolean
import org.http4k.lens.int
import org.http4k.lens.nonBlankString
import kotlin.io.path.Path

val port = EnvironmentKey.int().required("port")
val sourceDir = EnvironmentKey.nonBlankString().required("source_dir")
val serveDir = EnvironmentKey.nonBlankString().required("serve_dir")
val shouldBuild = EnvironmentKey.boolean().defaulted("build", true)
val shouldServe = EnvironmentKey.boolean().defaulted("serve", true)
val defaults = Environment.defaults(
    port of 9000,
    sourceDir of "./pages",
    serveDir of "./out",
)

fun main() {
    val env = (JVM_PROPERTIES overrides ENV overrides defaults)
    if (shouldBuild(env)) render(Path(sourceDir(env))).to(Path(serveDir(env)))
    if (shouldServe(env)) serve(Path(serveDir(env))).at(port(env)).block()
}
