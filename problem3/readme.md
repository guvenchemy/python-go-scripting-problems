# JSON API Log Analyzer

## Problem Tanımı

Mikroservis mimarilerinde ve dağıtık sistemlerde üretilen log verileri, makinelerin kolayca işleyebilmesi, indeksleyebilmesi ve merkezi sistemlerde (Elasticsearch, Loki vb.) sorgulanabilmesi amacıyla genellikle satır tabanlı JSON formatında (JSON Lines) tutulur.

Ancak bu log dosyalarının boyutları çok büyük seviyelere ulaştığında, sistemin genel sağlık durumunu, hata oranlarını ve performans darboğazlarını hızlıca tespit edecek hafif (lightweight) ve bağımsız bir analiz aracına ihtiyaç duyulmaktadır.

## Gereksinimler & Sınırlar

* *Bellek Dostu Akış Yönetimi (Stream I/O):* Analiz edilecek log dosyalarının boyutları gigabaytlarca olabileceğinden, dosyanın tamamı tek seferde belleğe (RAM) yüklenmemelidir. Bellek tüketimini minimumda tutmak adına dosya satır satır okunmalı ve işlenmelidir.

* *Hata Toleransı ve Kararlılık (Resilience):* Canlı sistemlerde log dosyalarının içine bozuk, yarım kalmış veya geçersiz JSON satırları sızabilir. Araç, geçersiz bir veri formatıyla karşılaştığında kesinlikle çökmemeli (panic/exception fırlatmamalı); ilgili satırı güvenle es geçerek çalışmaya devam etmeli ancak raporlama için bozuk satır sayısını kayıt altında tutmalıdır.

* *Metrik Dağılımı ve Sayımı:* Araç, başarıyla işlenen her satırdaki log seviyelerini (INFO, WARN, ERROR) ve HTTP durum kodlarını (status) kendi içlerinde gruplayarak frekans analizi (hangisinden kaç adet var) yapabilmelidir.

* *Hata İzleme (Error Tracking):* Seviyesi ERROR olarak işaretlenmiş tüm kayıtlar özel olarak incelenmeli; sistemdeki hataların hangi API uç noktalarında (path) yoğunlaştığını görebilmek adına hata alan benzersiz dizinlerin bir listesi çıkarılmalıdır.

* *Performans Analizi (SRE Metrikleri):* Sistem performansını ve kullanıcı deneyimini ölçümlemek adına, başarıyla işlenen tüm isteklerin yanıt süreleri (response_time_ms) kümülatif olarak hesaplanmalı ve dosya sonundaki raporda Ortalama Yanıt Süresi (Average Response Time) ms cinsinden sunulmalıdır. Veri kümesinin tamamen boş veya geçersiz olduğu ekstrem senaryolarda sıfıra bölme (ZeroDivisionError) gibi mantıksal çökmelerin önüne geçilmelidir.