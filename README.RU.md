[English](README.md) · Русский

<div align="center">

<img src="assets/logo.svg" width="360" alt="логотип txtpix">

# txtpix

Превращает изображения в текстовую графику, а файлы в base64 — прямо из терминала.

[![CI](https://github.com/exec-cmd/txtpix/actions/workflows/ci.yml/badge.svg)](https://github.com/exec-cmd/txtpix/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

</div>

## Возможности

- **ASCII-арт.** Переводит яркость пикселей в символы: `@%#*+=-:. `.
- **Unicode-арт.** Упаковывает две строки пикселей в одну строку символов с полублоками (`▀ ▄ █`).
- **base64.** Кодирует любой файл в текст base64 и декодирует его обратно.
- **Вывод.** Печатает в stdout или записывает тот же текст в файл через `-o`.

## Установка

Нужен Go 1.27.1 или новее.

### Сборка из исходников

```sh
make build
./build/bin/txtpix --help
```

Без `make`:

```sh
go build -o txtpix ./cmd/txtpix
```

### go install

```sh
go install github.com/exec-cmd/txtpix/cmd/txtpix@latest
```

### Готовые бинарники

Каждый тег `v*` публикует бинарники на странице GitHub Releases для linux/amd64, linux/arm64, windows/amd64 и darwin/arm64.

## Использование

```sh
txtpix convert <file> [flags]
```

| Флаг | По умолчанию | Описание |
| --- | --- | --- |
| `-a`, `--alg` | `base64` | Алгоритм: `base64`, `ascii`, `unicode`, `base64_decode` |
| `-o`, `--out` | нет | Дополнительно записать результат в этот файл |

### ASCII-арт

```sh
txtpix convert testdata/logo.png -a ascii
```

```text
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@%%%#%**@==@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@##***==@**@@@...@@@....@==@@=@@=@=@@@@
@@@**===**@##@@@@.@@..@@.@@=@=@@@@@=@@@@@
@@@==#*###@%%@@@@@@@@@@@@@=@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
```

### Unicode-арт

```sh
txtpix convert testdata/logo.png -a unicode
```

```text
█▀                                      ▀
           ▄▄                            
        ▄▄ ▀▀                            
        ▀▀      ▀█▀ ▄▄▀▀█▀ █▀▄ ▀  ▀▄▀    
     ███         █▄ ▀▀▄ █▄ █▄▀ ▄▄ ▄▀▄    
   ██                      ▀             
▄                                        
```

### base64 и декодирование

```sh
printf 'hello' > hello.txt
txtpix convert hello.txt -o hello.b64   # печатает aGVsbG8= и пишет его в hello.b64
txtpix convert -a base64_decode hello.b64
```

### Сохранение результата

```sh
txtpix convert testdata/logo.png -a ascii -o logo.txt
```

`-o` записывает текст в файл. Картинку он пока не создаёт.

## Примечания

- **Ширина.** Алгоритмы изображений используют ширину терминала, делённую на 1.5. Если ширина неизвестна, например при выводе в канал, берётся 41 колонка.
- **Форматы изображений.** `ascii` и `unicode` читают PNG, JPEG и GIF.
- **Пустые файлы** отклоняются с ошибкой.
- **`base64_decode`** печатает декодированные байты как текст, поэтому бинарные данные выйдут искажёнными.

## Разработка

```sh
make test     # форматирование, vet и тесты
make build    # сборка build/bin/txtpix
```

## Лицензия

MIT. См. [LICENSE](LICENSE).
