package dev.dhc.site

import org.http4k.client.JavaHttpClient
import org.http4k.core.Method.GET
import org.http4k.core.Request
import org.http4k.core.Status.Companion.NOT_FOUND
import org.http4k.core.Status.Companion.OK
import kotlin.io.path.Path
import kotlin.io.path.createTempDirectory
import kotlin.io.path.createTempFile
import kotlin.io.path.relativeTo
import kotlin.io.path.writeText
import kotlin.test.Test
import kotlin.test.assertEquals

private fun serve() = serve(Path("."))

class ServeTest {
    val client = JavaHttpClient()

    @Test
    fun testHealthz() {
        val srv = serve().at(0)
        val response = client(Request(GET, "http://localhost:${srv.port()}/healthz"))
        assertEquals(OK, response.status)
        srv.stop()
    }

    @Test
    fun testNotFound() {
        val srv = serve().at(0)
        val response = client(Request(GET, "http://localhost:${srv.port()}/foo"))
        assertEquals(NOT_FOUND, response.status)
        srv.stop()
    }

    @Test
    fun testStatic() {
        val dir = createTempDirectory()
        val content = "foo bar baz"
        val path = createTempFile(dir, suffix=".md").apply { writeText(content) }
        val rel = path.relativeTo(dir)

        val srv = serve(dir).at(0)
        val response = client(Request(GET, "http://localhost:${srv.port()}/$rel"))
        assertEquals(OK, response.status)
        assertEquals(content, response.bodyString())
        srv.stop()
    }
}
