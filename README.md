# Testing Golang & Nginx dan deploy dengan docker mengunakan tools Gitlab CI/CD
Pada kesempatan kali ini saya akan mencoba mendeploy Golang dengan Docker kontainer menggunakan tools Gitlab CI/CD. saya akan mencoba Menguji layanan Golang secara lokal.

# Prasyarat
- Dockerfile
- Docker Compose
- file .gitlab_ci.yml
 file .gitlab_ci.yml ini adalah file yang digunakan untuk menjalankan Jobs dan mengirimkan hasilnya kembali ke GitLab


## 1. Buat repository untuk Project Gitlab CI
```sh
Name project: "nama project"
Visibility Level : Public
```
## 2. Clone repository Gitlab ke local
```sh
git clone http://domain-gitlab-server/username-gitlab-kamu/nama project
```
## 3. clone repository Golang Test dari repository github pada link berikut 

```sh
git clone https://github.com/falyan/Golang-Test
```
setelah repository Github ini di clone pindahkan semua file kedalam direktori project yang telah di clone dari repository Gitlab kamu

## 4. buat file konfigurasi Gitlab Runner dengan nama file .gitlab_ci.yml
file konfigurasi Gitlab Runner dibuat didalam direktory project.
```sh
#.gitlab_ci.yml
stages:
- deploy

deploy-go:
  stage: deploy
  script:
    - docker-compose up -d 
  only:
    - master
  tags:
    - golang-test

```
file .gitlab_ci.yml ini adalah file yang digunakan untuk menjalankan Jobs dan mengirimkan hasilnya kembali ke GitLab. Ini digunakan bersama dengan GitLab CI (Continuous Integration).

pada file diatas terdiridari Jobs dengan stage deploy
- Jobs Deploy
Jobs ini berguna untuk menjalankan image yang sudah dibangun dari Dockerfile 

## 5. Push semua file ke repository Gitlab CI
```sh
git add .
git commit -m "Initial commit"
git push
```
setelah perintah git push dijalankan akan diminta username dan password gitlab 
```sh
username gitlab : "username-kamu"
password : "password-gitlab-kamu"
```

## 6. Daftar GitLab Runner sebagai Executor Shell
```sh
sudo gitlab-runner register -n \
  --url http://domain_gitlab_kamu/ \
  --registration-token [REGISTRATION_TOKEN] \
  --executor shell \
  --description "golang-test" \
  --tag-list "golang-test"
```

## 7. Set Permission dan verifikasi GitLab Runner
```sh
sudo usermod -aG docker gitlab-runner
sudo -u gitlab-runner -H docker info
sudo gitlab-runner start
sudo gitlab-runner list
sudo gitlab-runner status
sudo gitlab-runner verify
```
jika Gitlab Runner telah terverifikasi kamu dapat melihat statusnya pada dashboard Gitlab yaitu pada
```sh
# buka browser akses ke http://domain-gitlab-server/username-kamu/nama_project_kamu/settings/ci_cd 
# Expand the Runners, kemudian temukan Available specific runners.
```

## 7. Menguji Pipeline

```sh
# buka browser akses ke http:/domain-gitlab-server/username-kamu/nama_project_kamu/-/pipelines
# klik Run Pipeline > Run Pipeline
```

Jobs dari Pipeline sudah berjalan silahkan tunggu hasilnya dan jika Jobs gagal dengan status "ERROR: Job failed (system failure): prepare environment ..." jalankan perintah berikut

```sh
sudo rm -r /home/gitlab-runner/.bash_logout
sudo rm -r /home/username-kamu/.bash_logout
```

untuk mengakses berhasil atau tidaknya deploy bisa akses lewat http://ip_HOST_kamu:90/api/











**Credit by Falyan Zuril**

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [dill]: <https://github.com/joemccann/dillinger>
   [git-repo-url]: <https://github.com/joemccann/dillinger.git>
   [john gruber]: <http://daringfireball.net>
   [df1]: <http://daringfireball.net/projects/markdown/>
   [markdown-it]: <https://github.com/markdown-it/markdown-it>
   [Ace Editor]: <http://ace.ajax.org>
   [node.js]: <http://nodejs.org>
   [Twitter Bootstrap]: <http://twitter.github.com/bootstrap/>
   [jQuery]: <http://jquery.com>
   [@tjholowaychuk]: <http://twitter.com/tjholowaychuk>
   [express]: <http://expressjs.com>
   [AngularJS]: <http://angularjs.org>
   [Gulp]: <http://gulpjs.com>

   [PlDb]: <https://github.com/joemccann/dillinger/tree/master/plugins/dropbox/README.md>
   [PlGh]: <https://github.com/joemccann/dillinger/tree/master/plugins/github/README.md>
   [PlGd]: <https://github.com/joemccann/dillinger/tree/master/plugins/googledrive/README.md>
   [PlOd]: <https://github.com/joemccann/dillinger/tree/master/plugins/onedrive/README.md>
   [PlMe]: <https://github.com/joemccann/dillinger/tree/master/plugins/medium/README.md>
   [PlGa]: <https://github.com/RahulHP/dillinger/blob/master/plugins/googleanalytics/README.md>
