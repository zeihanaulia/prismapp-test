# Test for backend engineer Prismapp

## Challenge

Please create a chat backend application that exposes 3 public API:

1. API to send message:
   * Accept one parameter: the string message to sent.
   * This API should be a pull API (e.g. REST/GraphQL, etc).

2. API to retrieve messages:
   * To retrieve all previously-sent messages.
   * This API should be also an pull API.

3. API to retrieve message at realtime:
   * This will create a long lived connection that will retrieve message at realtime.
   * This API should be a push API (Websocket/MQTT, etc).

Some notes:

1. Feel free to use whatever tools/library necessary.
2. No need to think about authentication, etc.
3. The goal here is not only to measure your technical capabilities, but also to understand how you make design decision and its tradeoffs. So, creating a simple design-decision document would be huge plus for us. Short and concise document is preferred.

## Identify Issue

Tugasnya adalah membuat backend untuk chat application. Tugas dibagi menjadi 3 API:

1. API untuk mengirimkan pesan, API ini bertugas melakukan pengiriman pesan dan pesan akan dicatat kedalam storage
2. API untuk mengambil pesan yang telah dikirimkan, disini semua chat akan ditampilkan.
3. API untuk mendapatkan pesan secara realtime, jadi ketika user mengirimkan dari API 1 maka API 3 akan secara langsung mendapat pesannya

## Solution

Kita membutuhkan suatu proses yang memecah antara pesan yang tersimpan distorage dan pesan yang langsung dikirimkan.
~~Ketika API no 1 menerima pesan maka akan ada "worker" yang mem-"publish" ke dua service yang berberda yaitu, penyimpanan dan penerimaan pesan ke client.~~
Ketika API no 1 menerima pesan maka ada 2 step yang dilakukan , pertama kirim pesan ke websocket lalu menginput pesan kedalam database

## Use Case

Dengan ini kita bisa membataskan konteks dalam 3 use case yaitu:

    - Send New Message in a room of chat.
    - Retrieve Previous Message in a room of chat.
    - Retrieve Message in realtime.

## Design API

1. [POST] /chats    - Membuat chat baru - satu arah
2. [GET]  /chats    - Membaca semua chat
3. [GET]  ws:/realtime - Membaca chat terbaru secara realtime / dua arah


## Requirement 

1. Programming Language:
    * Go Langunage 19.3

2. Persistance Storage
    * Mariadb 10.1

3. Realtime
    * Websocket

## Folder structure

```
    - cmd                   = folder execute
        - chatd             = berisi file index untuk chat yang satu arah - REST
        - realtimed         = berisi file index untuk realtime - websocket
    - doc                   = dokumentasi Architecture Decision Record (ADR)
        - decision          = seluruh keputusan dan alasan mengenai architecture dicatat
    - docker                = folder untuk docker file dan settingnya
        - mariadb           = docker filde dan setting untuk mariadb
    - src                   = folder applikasi
        - chat              = applikasi chat satu arah - REST
            - entities      = data dari domain bisnis
            - gateway       = gateway ke utility semisal gateway untuk ke database
            - handlers      = handler applikasi semacam controller atau tempat menangani request
            - server        = run server applikasi
            - usecase       = folder seluruh usecase
                - retrieve  = usecase retrieve all previous message
                - send      = usecase send new message
            - utility       = folder detail / alat pendukung apllikasi seperti database
                - mysql     = berisi kode untuk integrasi ke mysql
                - websocket = berisi kode untuk integrasi ke websocket
        - websocket         = applikasi dua arah
            - ...
    - views                 = folder yang berisi html
```

## How to running

1. Install go language
2. Install mariadb 10.1
3. Create a new database with name : `prismapp-test`
4. Create new table
    ```SQL
        CREATE TABLE `room` (
            `id` varchar(36) NOT NULL,
            `message` longtext NOT NULL,
            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
            PRIMARY KEY (`id`)
        ) ENGINE=InnoDB DEFAULT CHARSET=latin1;

        // file dump ada di folder docker/mariadb/docker-entrypoint-initdb.d
    ```
5. clone this project from 
6. open 2 terminal 
7. in first terminal goto project path:
    ```bash
    cd /cmd/realtimed
    ```
8. then run the project
    ```bash
      go run main.go
      or
      go build -o realtimed && ./realtimed
    ```

9. in second terminal go to project path
    ```bash
       cd /cmd/chatd
    ```

10. make sure the `realtimed` is runnig in port :7070 then run the chatd project
    ```bash
      go run main.go
      or
      go build -o chatd && ./chatd
    ```

11. go to browser open localhost:8080 
12. or goto postman  to send message or retrieve all previous message
13. or with curl

```bash
    // send message
    curl -X POST http://localhost:8080/chats -F 'message=test test'

    // retrieve all previous message
    curl -X GET  http://localhost:8080/chats

```

Example in Gif

https://i.imgur.com/OyU763Z.gif
