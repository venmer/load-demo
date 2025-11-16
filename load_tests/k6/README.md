## Нагрузочное тестирование с k6

### Что такое?

- https://k6.io
- конфигурируется через JS
- много готовых настроек (например для GRPC)
- экспорт метрик в influxdb

### Ключевые преимущества

- генерация данных
- поддерживаемые протоколы
- большое количество метрик из коробки
- интеграция в CI
- расширяемость. Возможно писать собственные модули на Go
Можно написать на Go генератор данных и API использовать в своих сценариях
- можно прикрепить метрики к конкретном сценарию (?)


### Архитектура
TBD:
генератор -> k6 node -> балансировщик -> SUT


## Воркшоп

1. Установка k6 
- Debian, Ubuntu
```bash
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys C5AD17C747E3415A3642D57D77C6C491D6AC1D69
```
```bash
echo "deb https://dl.k6.io/deb stable main" | sudo tee /etc/apt/sources.list.d/k6.list
```
```bash
sudo apt-get update
```
```bash
sudo apt-get install k6
```
- macOS
```
brew install k6
```

1. Запуск первого теста
```bash
k6 run load_tests/k6/getheatlh.js 
```


Полезные ссылки:
- Testing guides https://grafana.com/docs/k6/latest/testing-guides/

