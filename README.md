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