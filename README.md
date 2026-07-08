# Scripting & Automation Problems

## Hakkında
Bu repo, sistem yönetimi ve DevOps süreçlerinde karşılaşılan gerçek dünya senaryolarını çözmek amacıyla oluşturulmuştur. 

Temel amaç; yapay zeka araçlarına bağımlı kalmadan, çıplak programlama mantığını, veri yapılarını ve algoritmaları kullanarak sağlam, ölçeklenebilir ve güvenli altyapı araçları geliştirmektir. Repodaki her bir problem, spesifik bir sistem ihtiyacına veya otomasyon açığına odaklanır ayrıca temel programlama ve scripting yeteneklerini geliştirmeyi hedefler.

## Kullanılan Teknolojiler
Projeler, ihtiyaca ve performans gereksinimlerine göre aşağıdaki dillerle geliştirilmektedir:
* **Bash:** Linux ortamlarında sunucu otomasyonu, dosya/dizin manipülasyonu ve hızlı I/O yönlendirmeleri.
* **Python:** Büyük boyutlu JSON/metin log analizi, veri ayrıştırma (parsing) ve metrik raporlama.
* **Go (Golang):** Yüksek performanslı ağ (network) araçları, eşzamanlılık (concurrency), worker pool mimarileri ve port/sağlık taramaları.

## Proje Yapısı
Depo, zorluk seviyesi ve konsept bazında modüler olarak klasörlendirilmiştir. Her klasör kendi bağımsız çözümünü ve senaryosunu içerir.

Her bir problemin kendi klasöründe, sınırlarına ve beklenen çıktılara dair kendi `README.md` dosyası bulunmaktadır. Scriptler doğrudan terminal veya uygun bir sunucu ortamında çalıştırılmak üzere tasarlanmıştır.