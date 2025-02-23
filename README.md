# home-assistant-util-api

## 概要

Home Assistantのちょっと手が届かないところを補完するためのAPIです。

## 機能

### 祝日判定系API

祝日は日本の祝日を想定。
GoogleCalenderAPIが祝日だと言ったら祝日として扱います。

- 当日が祝日かどうかを取得するAPI
- 当日が祝日扱いかどうかを取得するAPI
  - 上記APIに中部電力のスマートライフプランにて休日扱いの日も追加したもの

## インストール

```bash
cp .env.sample .env
docker-compose up -d
```

```bash
curl http://localhost:8000/isholiday
curl http://localhost:8000/isholidayrate
```
