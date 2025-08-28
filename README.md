# Mini Sekil Albomu

## Layihenin Meqsedleri
Bu layihə, Go proqramlaşdırma dilində sadə bir HTTP server yaratmaq və şəkil albomu API-si qurmaq məqsədi ilə hazırlanmışdır. Layihə real həyat nümunələrinə uyğun olaraq CRUD əməliyyatlarını öyrətmək və praktiki bacarıqları inkişaf etdirmək üçün nəzərdə tutulub.

## Istifade Edilen Texnologiyalar
- Go proqramlaşdırma dili
- net/http paketi
- encoding/json paketi
- Slice ve struct-lar
- JSON fayl ilə işləmə anlayisi
- Git və GitHub versiya idarəsi

## Endpoint-lar
1. GET /
   - Xoş gəlmisiniz mesajı göstərir.

2. GET /photos
   - Mövcud şəkillərin siyahısını JSON formatında qaytarır.

3. GET /photo/{id}
   - Verilən ID-yə uyğun şəkili JSON formatında qaytarır.

4. POST /photos
   - Yeni şəkil əlavə edir. JSON body ilə məlumat göndərilməlidir.

5. PUT /photos/{id}
   - Mövcud şəkili yeniləyir. JSON body ilə məlumat göndərilməlidir.

6. DELETE /photos/{id}
   - Mövcud şəkili silir.

## Layihenin Oyrendiklerim
- Go dilində HTTP server yaratmaq
- REST API dizayn prinsipləri
- JSON encode ve decode islemleri
- Slice ve struct-larla CRUD emeliyyatlari
- Error handling ve status kodlarinin istifadəsi
- Git ve GitHub ile versiya idarəsi

## Istifade Qaydalari
1. Layihenin fayllarini clone edin.
2. Go modulunu initialize edin.
3. `server.StartServer()` funksiyasini ishledin.
4. Curl veya Postman vasitesi ile endpoint-lari test edin.
