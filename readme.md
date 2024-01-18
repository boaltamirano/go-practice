## 1. Inicializar un módulo (require)
go mod init github.com/username/module


## Descargar una dependencia que necesites
go get github.com/donvito/hellomod


## Limpiar dependencias sin utilizar (optional)
go mod tidy


## Información de los módulos cacheados (optional)
go mod download -json