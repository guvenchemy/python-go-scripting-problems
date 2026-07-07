# Concurrent TCP Port Scanner

## Problem Tanımı
Altyapı yönetiminde sunucuların dış dünyaya açık portlarının (servislerinin) ve firewall (güvenlik duvarı) kurallarının düzenli olarak denetlenmesi SRE ve güvenlik süreçleri için kritiktir. Mevcut tarama araçlarının hantal kalabildiği veya otomasyon süreçlerine doğrudan entegre edilemediği senaryolarda; hafif, bağımsız ve sistem kaynaklarını tüketmeden çok hızlı sonuç üretebilen bir port tarama aracına ihtiyaç duyulmaktadır.

## Gereksinimler & Sınırlar

* **Girdi Doğrulama:** Araç, kullanıcıdan taranacak hedef IP adresini ve port aralığını (Örn: 20-1000) dinamik olarak almalı, regex veya ağ kütüphaneleriyle girdilerin standartlara uygunluğunu (Geçerli IP ve 1-65535 arası port sınırları) doğrulamalıdır.
* **Zaman Aşımı ve Hata Yönetimi:** Kapalı veya filtre uygulanmış portlara atılan isteklerin programı askıda bırakmasını önlemek adına, her TCP bağlantı denemesine (`net.DialTimeout`) katı bir zaman aşımı (Örn: 500ms) sınırı getirilmelidir.
* **Eşzamanlılık (Concurrency):** Binlerce portun tek tek (senkron) taranmasının yaratacağı zaman kaybını önlemek amacıyla, her port tarama işlemi eşzamanlı olarak yürütülmelidir.
* **Sistem Kaynak Yönetimi (Worker Pool):** Kontrolsüz şekilde binlerce eşzamanlı bağlantı (goroutine) açılması, işletim sisteminde `too many open files` (dosya tanımlayıcı limiti) hatasına yol açar. Bu riski engellemek adına, sisteme aşırı yük bindirmeyen sabit sayıda (Örn: 100) işçiden oluşan bir **Worker Pool Mimarisi** kurulmalı; iş dağıtımı asenkron bir `buffered channel` kuyruğu üzerinden yönetilmelidir.
* **Senkronizasyon:** Tüm işçilerin görevlerini tamamen bitirdiğinden ve açık TCP bağlantılarının (resource leak önlemek adına) güvenle kapatıldığından emin olmak için `sync.WaitGroup` senkronizasyon mekanizması entegre edilmelidir.