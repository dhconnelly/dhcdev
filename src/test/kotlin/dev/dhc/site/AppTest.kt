package dev.dhc.site

import com.sun.org.apache.bcel.internal.generic.GETFIELD
import org.http4k.core.Method
import org.http4k.core.Method.GET
import org.http4k.core.Request
import org.http4k.core.Status.Companion.OK
import kotlin.test.Test
import kotlin.test.assertContains
import kotlin.test.assertEquals

class AppTest {
    @Test
    fun testApp() {
        val app = App()
        val response = app(Request(GET, "/foo"))
        assertEquals(OK, response.status)
        assert("Hello" in response.bodyString())
    }
}
