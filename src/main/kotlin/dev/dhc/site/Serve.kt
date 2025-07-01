package dev.dhc.site

import com.squareup.moshi.JsonAdapter
import com.squareup.moshi.Moshi
import com.squareup.moshi.adapter
import com.squareup.moshi.kotlin.reflect.KotlinJsonAdapterFactory
import io.micrometer.core.instrument.Measurement
import io.micrometer.core.instrument.Meter
import io.micrometer.core.instrument.MeterRegistry
import io.micrometer.core.instrument.simple.SimpleMeterRegistry
import org.http4k.core.Body
import org.http4k.core.HttpHandler
import org.http4k.core.Method.GET
import org.http4k.core.Request
import org.http4k.core.Response
import org.http4k.core.Status.Companion.OK
import org.http4k.core.then
import org.http4k.core.with
import org.http4k.filter.MicrometerMetrics
import org.http4k.filter.ResponseFilters
import org.http4k.filter.ServerFilters
import org.http4k.format.ConfigurableMoshi
import org.http4k.format.json
import org.http4k.lens.BiDiBodyLens
import org.http4k.routing.ResourceLoader.Companion.Directory
import org.http4k.routing.bind
import org.http4k.routing.routes
import org.http4k.routing.static
import org.http4k.server.Http4kServer
import org.http4k.server.Jetty
import org.http4k.server.asServer
import java.nio.file.Path


@OptIn(ExperimentalStdlibApi::class)
data class Context(
    val dir: Path,
    val registry: MeterRegistry = SimpleMeterRegistry(),
    val moshi: ConfigurableMoshi = ConfigurableMoshi(
        Moshi.Builder().addLast(KotlinJsonAdapterFactory())
    ),
    val metricsLens: BiDiBodyLens<Metrics> = moshi.autoBody<Metrics>().toLens(),
)

data class Metrics(val measurements: Map<String, List<String>>)

fun varz(ctx: Context) = { _: Request ->
    val metrics = Metrics(ctx.registry.meters.associate { meter ->
        meter.id.toString() to meter.measure().take(5).map { it.toString() }
    })
    Response(OK).with(ctx.metricsLens of metrics)
}

fun server(ctx: Context) =
    ServerFilters.CatchLensFailure
        .then(ServerFilters.MicrometerMetrics.RequestCounter(ctx.registry))
        .then(ServerFilters.MicrometerMetrics.RequestTimer(ctx.registry))
        .then(ResponseFilters.ReportHttpTransaction { tx ->
            val caller = tx.request.source
            val method = tx.request.method
            val path = tx.request.uri
            val version = tx.request.version
            val status = tx.response.status
            println("${caller?.address}:${caller?.port} $method $path $version $status")
        })
        .then(routes(
            "/varz" bind GET to varz(ctx),
            static(Directory(ctx.dir.toString())),
        ))

class Servable(val serveDir: Path) {
    fun at(port: Int): Http4kServer {
        val ctx = Context(serveDir)
        val srv = server(ctx).asServer(Jetty(port)).start()
        println("serving $serveDir at http://localhost:${srv.port()}")
        return srv
    }
}

fun serve(serveDir: Path) = Servable(serveDir)
