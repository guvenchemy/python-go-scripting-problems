# Concurrent HTTP Health Checker 
## Problem Tanımı

Sistem mimarisinde yer alan çeşitli mikroservislerin, API uç noktalarının ve web sitelerinin erişilebilirlik durumlarını (sağlık kontrollerini) düzenli olarak izleyecek dinamik bir komut satırı (CLI) aracına ihtiyaç duyulmaktadır.

İzlenecek hedef adresler, sisteme bir metin dosyası girdi olarak verilerek dinamik şekilde beslenmelidir.
## Gereksinimler & Sınırlar

   * *Girdi Yönetimi:* Araç, hedef URL'lerin alt alta yazılı olduğu bir metin dosyasını (urls.txt) okuyabilmelidir. Dosya okunurken satır başı/sonu boşlukları veya geçersiz boş satırlar sisteme yük bindirmeden ayıklanmalıdır.

   * *Durum Raporlama:* Dosyadaki her bir URL'ye HTTP GET isteği atılmalı ve dönen HTTP durum kodu (Status Code) ile hedef adres, terminale anlamlı bir log formatında basılmalıdır.

   * *Hata ve İstisna Yönetimi (Error Handling):* Geçici ağ kesintileri, DNS çözümlenememe durumları (NXDOMAIN) veya yanıt vermeyen sunucular nedeniyle bir URL'ye ulaşılamadığında program kesinlikle çökmemeli (panic fırlatmamalı), ilgili hatayı yakalayıp loglamalı ve listedeki bir sonraki URL ile kararlı bir şekilde devam etmelidir.

   * *Zaman Aşımı (Timeout):* Yanıt vermeyen veya istekleri askıda bırakan sunucuların tüm izleme sistemini kilitlemesini önlemek adına, atılan her isteğe katı bir zaman aşımı (timeout) sınırı getirilmelidir.

   * *Performans ve Darboğaz Yönetimi (Concurrency):* Kontrol edilecek URL sayısı arttığında sistemin doğrusal olarak yavaşlamasını engellemek adına, istekler birbirini beklememeli (senkron olmamalı); eşzamanlı mimari kullanılarak tüm adreslere aynı anda istek çıkılabilmelidir.

   * *Sistem Kaynak Yönetimi:* Programın arka planda sürekli çalışabileceği göz önünde bulundurularak; açık kalan bağlantıların, bellek sızıntılarının (memory leak) ve işletim sistemi seviyesindeki dosya tanımlayıcı (file descriptor) tükenmelerinin önüne geçecek kaynak yönetim standartları uygulanmalıdır.
