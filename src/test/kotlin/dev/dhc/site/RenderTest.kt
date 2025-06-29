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

fun writeDirectory(tree: Map<Path, String>): Path {
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
                # title
                
                blah blah
                
                ## subtitle
                
                - a
                - b
                - c
            """.trimIndent(),

            "blog/post1.md" to """
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
        ).mapKeys { Path.of(it.key) }

        val expected = mapOf(
            "index.html" to """
                <h1>title</h1>
                <p>blah blah</p>
                <h2>subtitle</h2>
                <ul>
                <li>a</li>
                <li>b</li>
                <li>c</li>
                </ul>

            """.trimIndent(),

            "blog/post1.html" to """
                <ol>
                <li>foo</li>
                <li>bar</li>
                <li>baz</li>
                </ol>

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
        ).mapKeys { Path.of(it.key) }

        val srcDir = writeDirectory(content)
        val dstDir = render(srcDir)

        for ((rel, want) in expected) {
            val path = dstDir / rel
            val got = path.readText()
            assertEquals(want, got)
        }
    }
}
