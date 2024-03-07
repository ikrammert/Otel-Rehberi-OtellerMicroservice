# Otel Rehberi

Bu proje, otel rehberi uygulamasının oteller mikroservisini içermektedir. Bu mikroservis aracılığıyla, REST protokolü kullanılarak aşağıdaki işlemler gerçekleştirilir:

## İşlevler

### Otel oluşturma
### Otel Kaldırma
### Otel iletişim bilgisi ekleme
### Otel iletişim bilgisi kaldırma
### Otel yetkililerinin listelenmesi
### Otel ile ilgili iletişim bilgilerinin de yer aldığı detay bilgilerin getirilmesi
### Otellerin bulundukları konuma göre istatistiklerini çıkartan bir rapor talebi
### Sistemin oluşturduğu raporların listelenmesi
### Sistemin oluşturduğu bir raporun detay bilgilerinin getirilmesi

## Kullanılan Teknolojiler
**Veritabanı:** MongoDB
Esnek veri modelleme ve ölçeklenebilirlik avantajları nedeniyle tercih edildi.

**Framework:** gin-gonic/gin
Hızlı geliştirme ve yüksek performans için ideal bir Go web framework'ü.

## Kullanılan Servisler
**Message Queue Sistemi** RabbitMQ
RabbitMQ, güvenilir mesaj sıralama ve dağıtma işlevleri sunan bir açık kaynaklı mesajlaşma yazılımıdır. Proje için RabbitMQ tercih edilmiştir çünkü:


- Güvenilirlik ve Dayanıklılık: RabbitMQ, mesajları güvenilir bir şekilde iletebilir ve depolayabilir, böylece veri kaybı riskini azaltır.
- Esneklik: RabbitMQ, çeşitli protokollerle uyumlu olduğu için farklı dillerde ve platformlarda kullanılabilir.
- Yüksek Performans: RabbitMQ, yüksek verimlilik ve düşük gecikme süreleri sağlar, böylece uygulamanın hızlı ve verimli bir şekilde çalışmasını sağlar.

**Logging-ELK**
ELK: Loglama, analiz ve görselleştirme için kullanılan bir yazılım yığınıdır.

**Diğer Message Queue Sistemleriyle Karşılaştırma:**

**Apache Kafka**: Apache Kafka, yüksek hacimli ve düşük gecikmeli veri akışları için uygundur, ancak genellikle RabbitMQ'dan daha karmaşıktır ve daha fazla yapılandırma gerektirebilir.
**ActiveMQ**: ActiveMQ, genel amaçlı bir mesajlaşma aracıdır ancak bazı durumlarda RabbitMQ kadar yüksek performans sağlamayabilir ve daha az ölçeklenebilir olabilir.
**Redis Pub/Sub**: Redis, basit mesajlaşma senaryoları için kullanılabilir ancak RabbitMQ kadar geniş özellik yelpazesine sahip değildir ve daha karmaşık mesajlaşma senaryolarını desteklemeyebilir.

## İşlevleri Gerçekleştirme

- Otel oluşturma: POST /otels
- Otel kaldırma: DELETE /otels/{otelID}
- Otel iletişim bilgisi ekleme: POST /otels/{otelID}/contact
- Otel iletişim bilgisi kaldırma: DELETE /otels/{otelID}/contact/{contactID}
- Otel yetkililerinin listelenmesi: GET /otels/{otelID}/managers
- Otel detay bilgilerinin getirilmesi: GET /otels/{otelID}
- Konuma göre istatistik raporu talebi: POST /reports
- Oluşturulan raporların listelenmesi: GET /reports
- Bir raporun detay bilgilerinin getirilmesi: GET /reports/{reportID}

### Log için

```
cd elk-stack
docker-compose up -d
```

- Kibana'ya git localhost:5601 (Şifre ./elk-stack/.env bulunuyor) 