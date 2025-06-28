package dev.dhc.site

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

class ServerTest {
    @Test
    fun testHealthz() {
        val app = Server(Context(dir = Path("")))
        val response = app(Request(GET, "/healthz"))
        assertEquals(OK, response.status)
    }

    @Test
    fun testNotFound() {
        val app = Server(Context(dir = Path("")))
        val response = app(Request(GET, "foo"))
        assertEquals(NOT_FOUND, response.status)
    }

    @Test
    fun testStatic() {
        val dir = createTempDirectory()
        val path = createTempFile(dir, suffix=".md")
        val content = "foo bar baz"
        path.writeText(content)
        val rel = path.relativeTo(dir)

        val app = Server(Context(dir = dir))
        val response = app(Request(GET, "/$rel"))
        assertEquals(OK, response.status)
        assertEquals(content, response.bodyString())
    }
}
