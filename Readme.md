# ToDo API 

## 概要
クリーンアーキテクチャーを用いた ToDoリスト APIです。

## ディレクトリ構成

### /application

ビジネスロジックを配置

- /application/auth　認証機能のビジネスロジック
- /application/services メイン機能のビジネスロジック
- /application/integration サービスを初期化するときに必要な認証機能とメイン機能を統合した構造体を定義

### /cmd

アプリケーションのエントリーポイントを配置

### /config

データベースの接続情報を配置

### /di

依存性関係を管理して、リポジトリとサービスの初期化を行う

### /domain

ドメイン層のコードを配置

- /domain/entity 各entityの構造体を定義（外部からはアクセスできない仕様）
- /domain/interface リポジトリのインターフェース

### /dto

クライアントとやり取りするために使うオブジェクトを配置


### /infrastructure

インフラストラクチャ層のコードを配置
各種gatewayはdomain層で定義されているinterfaceに依存させる形で実装する

- /infrastructure/auth 認証機能に関するデータベースとやり取りするために使うオブジェクトとその操作方法
- /infrastructure/database データベースとやり取りするために使うオブジェクト
- /infrastructure/gateway データベースとのやり取りの操作方法
- /infrastructure/mysql データベースの接続とmigration機能
- /infrastructure/router ルーティング設定


### /migrations

migrationファイルを配置

### /presentation

- /presentation/auth/controller_auth.go 認証機能に関するユーザからのリクエストを受け取りレスポンスを返す
- /presentation/auth/middleware.go リクエストを行う前に認証を行うためのミドルウェアを配置
- /presentation/handler ユーザからのリクエストを受け取りレスポンスを返す

## /transform

model(データベースとのやり取りを行うオブジェクト)とentityのデータ変換機能
dto(クライアントとのやり取りを行うオブジェクト)とentityのデータ変換機能

### 環境設定
ローカルサーバのポート番号を配置