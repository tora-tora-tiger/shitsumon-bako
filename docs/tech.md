# Webサービス「質問箱」技術構成

| 項目 | 技術/ツール | 役割 |
| :--- | :--- | :--- |
| **API仕様** | 📜 **TypeSpec** | APIの設計図。ここからフロントとバックエンド両方の型定義を生成する。 |
| **フロントエンド** | ⚛️ **Next.js (React)** | TypeSpecから生成された型を利用して、安全にAPI通信を実装する。 |
| **バックエンド** | 🐹 **Go** (+ Gin or Echo) | TypeSpecから生成された型を利用して、仕様に沿ったAPIロジックを実装する。 |
| **データベース** | 🐘 **PostgreSQL** | データの永続化。 |
| **ファイルストレージ** | 📦 **MinIO** | オンプレミス環境でS3互換のストレージを構築し、画像などの大容量ファイルを保管する。 |


## フロントエンド
**言語・フレームワーク**
- TypeScript + Next.js (App Router)
- React 18+ (最新安定版)

**スタイリング**
- Tailwind CSS: ユーティリティファーストのCSS
- Headless UI: アクセシブルなUI コンポーネント

**状態管理**
- React Query (TanStack Query): サーバー状態の管理
- Zustand: クライアント状態の管理

**フォーム管理**
- React Hook Form: フォームのバリデーションと状態管理
- Zod: スキーマバリデーション

**画像処理**
- Next.js Image: 最適化された画像表示
- react-dropzone: ファイルアップロード

**通知**
- react-hot-toast: トースト通知

## バックエンド
**言語・フレームワーク**
- Go 1.21+
- Echo: HTTP フレームワーク
- oapi-codegen: OpenAPI からのコード生成

**認証・セキュリティ**
- JWT: 認証トークン
- bcrypt: パスワードハッシュ化
- CORS設定: フロントエンドからのアクセス制御

**データベースアクセス**
- GORM: ORM
- golang-migrate: マイグレーション管理

**ファイル処理**
- MinIO Go SDK: S3互換ストレージへのアクセス
- imaging または resize: サーバーサイド画像リサイズ

**ロガー・監視**
- slog (標準): 構造化ログ
- OpenTelemetry (将来的): トレーシング・メトリクス

**バリデーション**
- go-playground/validator: リクエストバリデーション

## データベース
**PostgreSQL テーブル設計概要**
- users: ユーザー情報（ID、メール、パスワードハッシュ、表示名、アイコンURL等）
- profiles: プロフィール詳細（自己紹介文など運営が追加する項目）
- questions: 質問データ（質問者ID、回答者ID、内容、画像URL、ステータス等）
- answers: 回答データ（質問ID、内容、画像URL、公開設定、投稿日時等）
- notifications: 通知データ
- blocks: ブロック関係
- ng_words: NGワード設定

**インデックス設計**
- 質問一覧取得用（回答者ID + ステータス + 作成日時）
- 公開回答取得用（回答者ID + 公開フラグ + 作成日時）

## 開発環境
**コンテナ化**
- Docker Compose: 開発環境の構築
  - PostgreSQL コンテナ
  - MinIO コンテナ
  - Redis コンテナ（セッション・キャッシュ用）

**API 開発**
- TypeSpec: API 仕様記述
- Swagger UI: API ドキュメント確認

**開発ツール**
- Air: Go のホットリロード
- Next.js dev server: フロントエンドのホットリロード

## デプロイメント構成（検討中）
**オンプレミス想定**
- Nginx: リバースプロキシ・静的ファイル配信
- Let's Encrypt: SSL証明書
- systemd: サービス管理

**CI/CD**
- GitHub Actions: 自動テスト・ビルド・デプロイ

## 保留中の選択事項
4. **デプロイ方式**: Docker vs バイナリ直接実行（後で検討）