import org.apache.kafka.clients.producer.KafkaProducer
import org.apache.kafka.clients.producer.ProducerRecord
import org.apache.kafka.common.serialization.StringSerializer
import org.json.JSONObject
import java.math.BigDecimal
import java.math.RoundingMode
import java.util.Properties
import java.util.UUID
import kotlin.random.Random

fun main() {
    // --- Variables ---
    // Create producer at start
    val producer = createKafkaProducer()
    val topic = "transactions"

    // --- Generate and send transactions ---
    while (true) {
        // Generate Transaction object
        val transaction = JSONObject().apply {
            put("transaction_id", UUID.randomUUID().toString())
            put("user_id", Random.nextInt(1, 1000))
            put("amount", BigDecimal(Random.nextDouble(1.0, 10000.0)).setScale(2, RoundingMode.HALF_UP))
            put("timestamp", System.currentTimeMillis())
        }

        // Simple Transaction to Kafka
        producer.send(ProducerRecord(topic, transaction.toString()))
        println("Sent transaction: $transaction")

        // Pause every second!
        Thread.sleep(1000)
    }
}

fun createKafkaProducer(): KafkaProducer<String, String> {
    val props = Properties().apply {
        put("bootstrap.servers", "kafka:9092")
        put("key.serializer", StringSerializer::class.java.name)
        put("value.serializer", StringSerializer::class.java.name)
    }
    return KafkaProducer(props)
}
