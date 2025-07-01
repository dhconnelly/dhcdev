package dev.dhc.site

import java.nio.file.Path
import kotlin.io.path.createFile
import kotlin.io.path.createParentDirectories
import kotlin.io.path.createTempDirectory
import kotlin.io.path.div
import kotlin.io.path.readText
import kotlin.io.path.writeText
import kotlin.test.Test
import kotlin.test.assertEquals

fun writeDirectory(tree: Map<String, String>): Path {
    val dir = createTempDirectory()
    for ((rel, content) in tree) {
        val path = dir / rel
        path.createParentDirectories()
        path.createFile().writeText(content)
    }
    return dir
}

class RenderTest {
    @Test
    fun testRender() {
        val content = mapOf(
            "index.md" to """
                === outline ===
                # title
                
                blah blah
                
                ## subtitle
                
                - a
                - b
                - c
            """.trimIndent(),

            "blog/post1.md" to """
                === list ===
                
                
                1. foo
                2. bar
                3. baz
            """.trimIndent(),

            "blog/img.png" to """
                1. should
                2. not
                3. render
            """.trimIndent(),

            "out/blargh.html" to """
                - not
                - this
                - either
            """.trimIndent(),
        )

        val expected = mapOf(
            "index.html" to PageMaker.render(content["index.md"]!!.reader().buffered()),

            "blog/post1.html" to PageMaker.render(content["blog/post1.md"]!!.reader().buffered()),

            "blog/img.png" to """
                1. should
                2. not
                3. render
            """.trimIndent(),

            "out/blargh.html" to """
                - not
                - this
                - either
            """.trimIndent(),
        )

        val srcDir = writeDirectory(content)
        val dstDir = createTempDirectory()
        render(srcDir).to(dstDir)

        for ((rel, want) in expected) {
            val path = dstDir / rel
            val got = path.readText()
            assertEquals(want, got)
        }
    }
}
