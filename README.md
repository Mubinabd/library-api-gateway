# Library

**Project Description:**
Kitoblar, mualliflar, janrlar va yig'uvchilarni boshqarish xizmatlari.
Kutubxonadan kitob olib o'qish uchun qulaylashtirilgan keng imkoniyatlarga ega kichkina API service.
Ushbu tizim orqali foydalanuvchilar kutubxonadan ijaraga kitob olib o'qishlari mumkin bo'ladi.
Qaysi foydalanuvchi qaysi turdagi kitobni qanday janrdagi kitobni qachon olgan kutubxondan va kutubxonaga qachon qaytarishi
haqidagi ma'lumotlar keltirilgan.Kitobning muallifi tavsifi va kitob haqidagi malumotlarni ko'rib qidiruv tizimini
ham amalga oshirish mumkin.Foydalanuvchilar o'zlarining qanday kitob olganliklarini va qachon olib qachon 
qaytarganliklarini ham bilib bilishlari mumkin.

## Installation

1. Initialize a git repository and clone the project:
    ```sh
    git init
    git clone git@github.com/Mubinabd/library-exam.git
    ```
2. Create a database named `library` on port `5432`.
3. Update the `.env` file with the appropriate configuration.
   ```.env
   DB_HOST=localhost
   DB_USER=postgres
   DB_NAME=library
   DB_PASSWORD=pass
   DB_PORT=5432
   LOGPATH=logs/info.log
   ```

4. Use the following Makefile commands to manage the database migrations and set up the project:
    ```makefile
    # Set the database URL
    exp:
        export DBURL="postgres://postgres:1234@localhost:5432/library?sslmode=disable'"

    # Run migrations
    mig-up:
        migrate -path migrations -database ${DBURL} -verbose up

    # Rollback migrations
    mig-down:
        migrate -path migrations -database ${DBURL} -verbose down

    # Create a new migration
    mig-create:
        migrate create -ext sql -dir migrations -seq create_table

    # Create an insert migration
    mig-insert:
        migrate create -ext sql -dir migrations -seq insert_table

    # Generate Swagger documentation
    swag:
        swag init -g api/api.go -o api/docs

    # Clean up migrations (commented out by default)
    # mig-delete:
    #   rm -r db/migrations
    ```
5. Set the environment variable and run the project:
    ```sh
    make exp
    make mig-up
    go run main.go
    ```
6. Open the following URL to access the Swagger documentation:
    ```
    http://localhost:8090/api/swagger/index.html#/
    ```

## Features and Usages
1. Auth serviceda User uchun Register va Login bo'limi bor va Get bo'limida user faqat o'zining profile ni ko'ra oladi.
2. Author bo'limida author create, put, delete metodlarida faqat username adminga teng bo'lsagina admin qila yaratib 
o'zgartirib,o'chira oladi va idsi bo'yicha shu authorni barcha malumotlarini olsa bo'ladi.GetAll metodida name bo'yicha
filter qilinadi.Author bo'limiga qo'shimcha API qo'shilgan bo'lib bu metod GetAuthorBooks 1ta authorga tegishli barcha 
kitoblarni olib beradi.
3. Book bo'limida book create, put, delete metodida username adminga teng bo'lsa admin kitob yarata,o'chira va qayta 
o'zgaratira oladi va title bo'yicha shu bookni barcha ma'lumotlarini olsa bo'ladi. GetAll metodida title bo'yicha 
filter qilinadi.Book bo'limiga qo'shimcha API qo'shilgan bo'lib bu metod SearchTitleAndAuthor author yoki title 
bo'yicha qidiruv tizimini amalga oshiradi.
4. Borrower bo'limida borrower create,put,delete qilinadi bularni barchasini username adminga teng bo'lsagina admin qila
oladi va id si bo'yicha shu borrowerni barcha ma'lumotlarini olsa bo'ladi.GetAll metodida hech qanday parametri yo'q bo'lib 
hech qanday field bervormasdan va filter qo'ymasdan barcha borrowerlarni olsa bo'ladi. Borrower bo'limiga 3ta qo'shimcha 
API qo'shilgan bo'lib bularning birinchisi BorrowerBooks malum bir foydalanuvchi tomonidan olingan barcha kitoblarni 
olib beradi.Ikkinchisi GetOverdueBooks return date bilan hozirgi vaqtni solishtiradi va return date o'tib ketgan bo'lsa shu
kitoblarni chiqaradi.Uchinchisi HistoryUser foydalanuvchining kutubxonadan olgan kitoblarini tarixini ko'rsatadi.
5. Genre bo'limida genre create,put,delete metodida username adminga teng bo'lsa admin yangi janr yarata oladi va get metodida
name bo'yicha shu namega tegishli genreni qaytaradi. GetAll metodida name bo'yicha filter qilindai.Genre bo'limiga qo'shimcha 
API qo'shilgan bo'lib GetBooksByGenre bu metod orqali biror bir janrga tegishli barcha kitoblarni olishingiz mumkin bo'ladi.

## Dependencies

- **Scheduling**: [github.com/go-co-op/gocron](https://github.com/go-co-op/gocron)
- **Swagger**: [github.com/swaggo/swag](https://github.com/swaggo/swag)
- **Database**:
    - [database/sql](https://golang.org/pkg/database/sql/)
    - [github.com/lib/pq](https://github.com/lib/pq)
- **Environment Variables**: [github.com/joho/godotenv](https://github.com/joho/godotenv)
- **API Framework**: [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- **Migrations**: [github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate)
****
## Acknowledgments

- Mubina Bahodirova

## Special thanks to HUSAN MUSA