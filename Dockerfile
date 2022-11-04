FROM golang:alpine
# Keyword FROM ini digunakan untuk inisialisasi build stage dan juga menentukan basis Image yang digunakan.
# Informasi golang:alpine di sini adalah basis image yang dimaksud, yaitu image bernama golang.
# Tag bernama alpine yang tersedia di laman officila Docker Hub Golang https://hub.docker.com/_/golang.
# Dalam Image golang:alpine sudah tersedia beberapa utilitas untuk keperluan build aplikasi Golang.
# Image golang:alpine basisnya adalah Alpine OS.

RUN apk update && apk add --no-cache git
# Keyword RUN digunakan untuk menjalankan shell comamnd.
# Argument setelahnya, yaitu apk update && apk add --no-cache git akan dijalankan di Image golang:alpine yang sudah di-set sebelumnya.
# Command tersebut merupakan command Alpine OS yang kurang lebih gunanya adalah berikut:
# 1. Command apk update digunakan untuk meng-update index packages pada OS.
# 2. Command apk add --no-cache git digunakan untuk meng-install Git. Karena iamge golang:alpine by default, GIT berlum tersedia.adalah tidak tersedia.

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

# Build the image
# docker build -t go-sys-image . # flag -t digunakan untuk menentukan nama imaga

# Buat container baru berbasis image yang sudah kita buat.
# docker container create --name go-sys-container -p 8080:9000 go-sys-image

# Check container dengan
# docker container ls -a

# Start the container
# docker container start my-container-hello-world
# docker container ls

# Stop Container
# Untuk stop container bisa dengan command docker container stop <nama-container-atau-container-id>.

#docker container stop my-container-hello-world
# docker container ls

# Hapus Container
# Untuk hapus container bisa dengan command docker container rm <nama-container-atau-container-id>.
# docker container rm my-container-hello-world
# docker container ls

# Hapus Image
# Untuk hapus image bisa dengan command docker image rm <nama-image-atau-image-id>. O iya, untuk penghapusan image, harus dipastikan terlebih dahulu tidak ada container yang running menggunakan basis image yang ingin dihapus.
# docker image rm my-image-hello-world
# docker images