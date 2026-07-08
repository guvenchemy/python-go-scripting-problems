## Problem Tanımı

Sunucu güvenliğinde en büyük risklerden biri, sisteme sızan saldırganların standart root kullanıcısı dışında, kendi oluşturdukları gizli hesaplara UID 0 (root yetkisi) atayarak arka kapı (backdoor) bırakmalarıdır. Ayrıca, sunucuda gereksiz yere terminal erişimine (login shell) sahip olan servis hesapları büyük bir zafiyet yüzeyi oluşturur.

Bu proje; sistemdeki kullanıcıları düzenli olarak denetleyen, yetkisiz root erişimlerini anında tespit eden ve terminal erişimi olan hesapların haritasını çıkaran, harici hiçbir bağımlılığı olmayan (sıfır Python/Go bağımlılığı) natif bir Bash güvenlik aracıdır.

## Mimari ve Özellikler

*   *Tek Geçişli (Single-Pass) Analiz:* Sistem kaynaklarını yormamak adına /etc/passwd dosyası tek bir döngüde okunur ve tüm ayrıştırma (parsing) işlemleri eşzamanlı yapılır.

*   *Native Parsing:* Metin işleme için awk veya sed gibi harici programlar çağırmak yerine, Bash'in dahili IFS (Internal Field Separator) mekanizması kullanılarak milisaniyeler içinde sonuç üretilir.

*   *Sıfır Yanlış Alarm (Zero False-Positive):* SRE "alert fatigue" (alarm yorgunluğu) prensibi göz önünde bulundurularak, script sadece gerçekten yetkisiz bir UID 0 tespit ettiğinde [CRITICAL ALERT] üretir. Tehlike yoksa sessizce çalışır.

*    *Legacy Sistem Uyumluluğu:* Harici bir kütüphane gerektirmediği için CentOS 7 gibi eski nesil sunucu ortamlarında (legacy) veya container'ların içinde sorunsuz çalışır.