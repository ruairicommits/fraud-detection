import org.json.JSONObject
import java.math.BigDecimal
import java.math.RoundingMode
import java.util.UUID
import kotlin.random.Random

fun main() {
    while (true) {
        // --- Variables ---
        val transaction_id = UUID.randomUUID().toString()
        var user_id = Random.nextInt(1, 1000)
        // NB: We are using currency, so we want 2 decimal place and round up
        var amount = BigDecimal(Random.nextDouble(1.0, 10000.0)).setScale(2, RoundingMode.HALF_UP)
        var timestamp = System.currentTimeMillis()

        // --- Transaction object ---
        val transaction = JSONObject().apply {
            put("transaction_id", transaction_id)
            put("user_id", user_id)
            put("amount", amount)
            put("timestamp", timestamp)
        }

        println(transaction.toString()) // Print
        // todo - send to Kafka/HTTP endpoint
        Thread.sleep(1000) // Simulate 1-second intervals for testing
    }
}
