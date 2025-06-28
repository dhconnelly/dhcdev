package dev.dhc.site

import org.http4k.config.Environment
import org.http4k.config.Environment.Companion.ENV
import org.http4k.config.Environment.Companion.JVM_PROPERTIES
import org.http4k.config.EnvironmentKey
import org.http4k.lens.boolean
import org.http4k.lens.int
import org.http4k.lens.nonBlankString

val port = EnvironmentKey.int().required("port")
val dir = EnvironmentKey.nonBlankString().required("dir")
val shouldBuild = EnvironmentKey.boolean().defaulted("build", true)
val shouldServe = EnvironmentKey.boolean().defaulted("serve", true)
val defaults = Environment.defaults(
    port of 9000,
    dir of "./pages",
)

fun main() {
    val env = (JVM_PROPERTIES overrides ENV overrides defaults)
    println(env)

    val buildDir = dir(env)
    val serveDir = if (shouldBuild(env)) render(buildDir) else buildDir
    if (shouldServe(env)) serve(port(env), serveDir).block()
}
