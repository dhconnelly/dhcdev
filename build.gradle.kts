plugins {
    kotlin("jvm") version "2.1.20"
    application
    id("com.github.johnrengelman.shadow") version "8.1.1"
}

group = "dev.dhc"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    implementation("org.jetbrains.kotlin:kotlin-stdlib")
    implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:1.10.2")
    implementation("com.squareup.moshi:moshi:1.15.2")
    implementation("org.http4k:http4k-core:6.15.0.1")
    implementation("org.http4k:http4k-server-jetty:6.15.0.1")
    implementation("org.http4k:http4k-format-moshi:6.15.0.1")
    implementation("org.http4k:http4k-config:6.15.0.1")
    implementation("org.commonmark:commonmark:0.25.0")
    implementation("org.commonmark:commonmark-ext-heading-anchor:0.25.0")

    testImplementation("org.jetbrains.kotlin:kotlin-test-junit5")
    testImplementation("org.junit.jupiter:junit-jupiter:5.10.0")
}

application {
    mainClass = "dev.dhc.site.MainKt"
}

tasks.test {
    useJUnitPlatform()
}

kotlin {
    jvmToolchain(21)
}

tasks.shadowJar {
    manifest.attributes["Main-Class"] = application.mainClass.get()
    archiveBaseName.set(project.name)
    archiveClassifier.set("")
    archiveVersion.set("")
    mergeServiceFiles()
}
