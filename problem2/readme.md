# Smart Log Rotator & Cleanup Script
## Problem Tanımı

Sunucularda veya Docker container'larında çalışan canlı uygulamalar, sistem diskine sürekli olarak log verisi yazmaktadır. Bu log dosyalarının boyutları kontrol edilmediğinde zamanla diski tamamen doldurarak sistemin kilitlenmesine ve servislerin çökmesine neden olmaktadır.

Bu durumun önüne geçmek amacıyla, disk alanını dinamik olarak yöneten, eski verileri arşivleyen ve sistem kaynaklarını optimize eden akıllı bir log döndürme (rotation) ve temizlik aracına ihtiyaç duyulmaktadır.
## Gereksinimler & Sınırlar

   * *Dinamik Dizin Taraması:* Araç, belirlenen bir log dizini içerisindeki tüm .log uzantılı aktif dosyaları otomatik olarak tespit edebilmeli ve birden fazla dosyanın varlığı durumunda her biri için bağımsız işlem yürütebilmelidir.
   * *Boyut Kontrolü ve Arşivleme:* Belirlenen kritik boyut sınırını (Örn: 10MB) aşan log dosyaları, veri kaybını önlemek adına geriye dönük tarih damgası verilerek sıkıştırılmış arşiv formatında (.tar.gz veya .zip) yedeklenmelidir. Sıkıştırma işlemi yapılırken gereksiz klasör ağaçları arşiv içine dahil edilmemelidir.
   * *Canlı Sistem Kararlılığı (Truncate):* Uygulamalar log dosyalarına anlık olarak yazmaya devam ettiği için, boyutu dolan orijinal log dosyası sistemden doğrudan silinmemelidir. Canlı süreçlerin (process/PID) dosya tanımlayıcılarını (file descriptor) bozmamak ve uygulamanın çökmesini engellemek adına, dosya silinmeden içeriği güvenli bir şekilde sıfırlanmalıdır (0 byte yapılmalıdır).
   * *Otomatik Disk Temizliği:* Sistemin diski kalıcı olarak şişirmesini önlemek adına, oluşturulan sıkıştırılmış arşiv yedekleri düzenli olarak denetlenmelidir. Belirlenen saklama süresinden (Örn: 30 gün) daha eski olan tüm arşiv dosyaları sistemden kalıcı olarak temizlenmelidir. Bu temizlik işlemi, aktif log dosyalarının boyutundan bağımsız olarak her çalışma periyodunda mutlak suretle tetiklenmelidir.
   * *İzlenebilirlik ve Loglama (Observability):* Aracın arka planda (Örn: cronjob) sorunsuz çalıştığından emin olmak amacıyla; yapılan her arşivleme, temizleme, hata veya normal durum çıktısı, başına standart zaman damgası eklenerek yapılandırılmış bir şekilde raporlanmalıdır.
