package basic

import io.github.oshai.kotlinlogging.DelegatingKLogger
import io.github.oshai.kotlinlogging.KotlinLogging
import kotlinx.coroutines.runBlocking
import kotlinx.coroutines.launch
import io.github.oshai.kotlinlogging.withLoggingContext
import io.github.oshai.kotlinlogging.coroutines.withLoggingContextAsync
import kotlinx.coroutines.delay
import kotlinx.coroutines.slf4j.MDCContext
import kotlin.time.Duration.Companion.seconds

private val logger = KotlinLogging.logger {}

fun main() {
    withLoggingContext("context" to "main()") {
        println(logger.isTraceEnabled())
        withLoggingContext("context" to "runBlocking") {
            runBlocking(MDCContext()) {
                logger.trace{"runBlocking.."}
                withLoggingContextAsync("context" to "launch") {
                    launch(MDCContext()) {
                        logger.info{"Launch.."}
                        delay(3.seconds)
                        logger.info{"Exit launch.."}
                    }
                }
                logger.info{"Exit runBlocking.."}
            }
            logger.info{"Exit main.."}
        }
    }
}