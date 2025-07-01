package dev.dhc.site

import org.commonmark.parser.Parser
import org.commonmark.renderer.html.HtmlRenderer
import java.io.BufferedReader
import java.nio.file.Path
import kotlin.io.path.copyTo
import kotlin.io.path.createFile
import kotlin.io.path.createParentDirectories
import kotlin.io.path.deleteExisting
import kotlin.io.path.deleteIfExists
import kotlin.io.path.div
import kotlin.io.path.extension
import kotlin.io.path.nameWithoutExtension
import kotlin.io.path.reader
import kotlin.io.path.relativeTo
import kotlin.io.path.walk
import kotlin.io.path.writeText

class PageException(val why: String): Exception("invalid page: $why")

fun tmpl(title: String, content: String) = """
    <!DOCTYPE html>
    <html>
        <head>
            <title>$title</title>
            <meta charset="utf-8" />
            <meta http-equiv="X-UA-Compatible" content="IE=edge" />
            <meta name="viewport" content="width=device-width, initial-scale=1" />
            <meta property="og:title" content="$title" />
            <meta name="author" content="Daniel Connelly" />
            <meta property="og:locale" content="en_US" />
            <meta property="og:type" content="website" />
            <meta name="twitter:card" content="summary" />
            <meta property="twitter:title" content="$title" />
            <link rel="stylesheet" type="text/css" href="/css/main.css">
        </head>
        <body>
            <main><div id="content">$content</div></main>
        </body>
    </html>
    """.trimIndent()

object PageMaker {
    val parser: Parser = Parser.builder().build()
    val renderer: HtmlRenderer = HtmlRenderer.builder().build()
    val titlePat: Regex = """^=== ([^=]+) ===$""".toRegex()

    fun render(r: BufferedReader): String {
        val title = titlePat.matchEntire(r.readLine())?.groups[1]?.value
            ?: throw PageException("failed to parse title")
        // TODO: code highlighting
        // TODO: anchors
        val node = parser.parseReader(r)
        val content = renderer.render(node)
        return tmpl(title, content)
    }

    fun render(from: Path, to: Path) {
        from.reader().buffered().use { r ->
            to.writeText(render(r))
        }
    }
}

class Renderable(val srcDir: Path) {
    fun to(dstDir: Path) {
        println("building $srcDir to $dstDir")
        for (path in srcDir.walk()) {
            val rel = dstDir / path.relativeTo(srcDir)
            if (path.extension == "md") {
                val dst = rel.parent / (path.nameWithoutExtension + ".html")
                dst.createParentDirectories()
                dst.deleteIfExists()
                dst.createFile()
                PageMaker.render(path, dst)
            } else {
                rel.createParentDirectories()
                rel.deleteIfExists()
                path.copyTo(rel)
            }
        }
        println("built $srcDir to $dstDir")
    }
}

fun render(srcDir: Path) = Renderable(srcDir)
