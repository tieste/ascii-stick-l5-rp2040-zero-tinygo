# ascii-stick-l5-rp2040-zero-tinygo

アスキースティックL5をRP2040-ZeroとTinyGoでUSB化する

# build

tinygo flash --target waveshare-rp2040-zero --size short ./main.go

# 概要

アスキースティックL5に PS Vita 用のアナログスティックを追加する改造を施したものを、RP2040-Zeroで動かすためのTinyGoプログラムです。

## 諸元

| 種類               | 個数                |
| ------------------ | ------------------- |
| ボタン             | A, B, START, SELECT |
| ハットスイッチ     | 1式                 |
| アナログスティック | X軸, Y軸            |

