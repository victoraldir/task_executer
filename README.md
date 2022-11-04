[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="#">
    <img src="assets/workflow_icon.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">Task Parser / Executer</h3>

  <p align="center">
    Command line application that can read YAML files and perform tasks. written in Go
    <br />
    <a href="https://about.me/victoraldir"><strong>By Victor Hugo ❤️ »</strong></a>
    <br />
    <br />
  </p>
</p>

## Summary

This is a simple command line application that can load tasks from YAML files via STDIN and to perform actions

## Technology

- __Language__: Go 1.19
- __Mock library__: golang/mock

## Run tests

`./run-tests.sh`

## Execute sample tasks 

`go run cmd/main.go < sample/task_flow.yaml`


[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/victoraldir
[dot]: assets/dot-on.png
