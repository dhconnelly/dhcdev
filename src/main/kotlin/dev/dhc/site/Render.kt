package dev.dhc.site

import org.commonmark.parser.Parser
import org.commonmark.renderer.html.HtmlRenderer
import java.nio.file.Path
import kotlin.io.path.ExperimentalPathApi
import kotlin.io.path.copyTo
import kotlin.io.path.createFile
import kotlin.io.path.createParentDirectories
import kotlin.io.path.createTempDirectory
import kotlin.io.path.div
import kotlin.io.path.extension
import kotlin.io.path.nameWithoutExtension
import kotlin.io.path.reader
import kotlin.io.path.relativeTo
import kotlin.io.path.walk
import kotlin.io.path.writeText

@OptIn(ExperimentalPathApi::class)
fun render(srcDir: Path): Path {
    val parser = Parser.builder().build()
    val renderer = HtmlRenderer.builder().build()

    val dstDir = createTempDirectory()
    println("building $srcDir to $dstDir")
    for (path in srcDir.walk()) {
        val rel = dstDir / path.relativeTo(srcDir)
        if (path.extension == "md") {
            val dst = rel.parent / (path.nameWithoutExtension + ".html")
            val node = path.reader().use { parser.parseReader(it) }
            val html = renderer.render(node)
            dst.createParentDirectories().createFile().writeText(html)
        } else {
            val dst = rel
            dst.createParentDirectories()
            path.copyTo(dst)
        }
    }

    println("built $srcDir to $dstDir")
    return dstDir
}
