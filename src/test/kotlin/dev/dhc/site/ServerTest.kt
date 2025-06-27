package dev.dhc.site

import org.http4k.core.Method.GET
import org.http4k.core.Request
import org.http4k.core.Status.Companion.OK
import kotlin.test.Test
import kotlin.test.assertEquals

class ServerTest {
    @Test
    fun testHealthz() {
        val app = Server(Context())
        val response = app(Request(GET, "/healthz"))
        assertEquals(OK, response.status)
    }
}
