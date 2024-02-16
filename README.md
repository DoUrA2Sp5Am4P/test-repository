# Распределенный вычислитель арифметических выражений
Представим, что такие простые операции, как сложение, вычитание, умножение, деление занимают много времени.
Поэтому я создал приложение на Go, которое производит вычисление выражений параллельно.
# Как установить?
```
mkdir Distributed_arithmetic_expression_evaluator
cd Distributed_arithmetic_expression_evaluator
git pull https://github.com/DoUrA2Sp5Am4P/test-repository.git
go build
Запустите Distributed_arithmetic_expression_evaluator
```
# Как использовать?
```
0. По желанию отредактируйте файл CONCURRENCY.go, чтобы изменить количество воркеров (По умолчанию 5)
1. Перейдите на 127.0.0.1:8080
Убедитесь, что на 127.0.0.1:8080 ничего не запущено!
2. Наслаждайтесь приложением. Документация по использованию каждой страницы размещена на wiki
```
# [Wiki](https://github.com/DoUrA2Sp5Am4P/test-repository/wiki)
На wiki размещены примеры запросов и описания web-страниц приложения
