package dev.dhc.site

import kotlin.io.path.ExperimentalPathApi
import kotlin.io.path.Path
import kotlin.io.path.copyToRecursively
import kotlin.io.path.createTempDirectory

@OptIn(ExperimentalPathApi::class)
fun render(srcDir: String): String {
    val dstDir = createTempDirectory()
    println("building $srcDir to $dstDir")
    Path(srcDir).copyToRecursively(dstDir, followLinks=false)
    println("built $srcDir to $dstDir")
    return dstDir.toString()
}
