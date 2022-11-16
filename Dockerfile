FROM golang:latest
# Keyword FROM ini digunakan untuk inisialisasi build stage dan juga menentukan basis Image yang digunakan.
# Informasi golang:alpine di sini adalah basis image yang dimaksud, yaitu image bernama golang.
# Tag bernama alpine yang tersedia di laman officila Docker Hub Golang https://hub.docker.com/_/golang.
# Dalam Image golang:alpine sudah tersedia beberapa utilitas untuk keperluan build aplikasi Golang.
# Image golang:alpine basisnya adalah Alpine OS.

WORKDIR /app
# Digunakan untuk menentukan working directory yang pada konteks ini adalah /app.
# Statement ini menjadikan semua statement RUN di bawahnya akan dieksekusi pada working directory.

COPY . .
# Digunakan untuk meng-copy file pada argument pertama
# . pertama merepresentasikan direktori yang aktif pada host atau komputer kita.
# . kedua merepresentasikan direktory di image golang:alpine yaitu /app.
# Dengan ini isi /app adalah sama persis seperti isi folder project go-sys.

RUN go mod tidy
# Digunakan untuk mem-validasi dependency pada golang.
# Disinilah kita memerlukan GIT yang di install di atas.

RUN go build -o binary
# Command go build digunakan untuk build binary atau executable dari kode program Go.
# Dengan ini source code dalam working directory akan di-build ke executable dengan nama binary.

ENTRYPOINT ["/app/binary"]
# Statement ini digunakan untuk menentukan entrypoint container sewaktu dijalankan.
# Jadi khusus statement ENTRYPOINT ini pada contoh di atas adalah yang efeknya baru kelihatan ketika Image di-run ke container.
# Sewaktu proses build aplikasi ke Image maka efeknya belum terlihat.
# Dengan statement tersebut nantinya sewaktu container jalan, maka executable binary dari aplikasi akan dijalankan di container sebagai entrypoint.
