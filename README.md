### Инструкция по запуску
- Первый варинат:
  - cd Computer_clubs_management
  - go build -o task.exe cmd/main.go
  - ./task.exe test_file.txt
- Второй вариант:
  - cd Computer_clubs_management
  - docker build -t clubs .
  - docker run -it clubs

