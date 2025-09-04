import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

using TypeSpec.Http;
using TypeSpec.Rest;

namespace ShitsumonBako;

/**
 * 共通レスポンス構造
 */
model ApiResponse<T> {
  message?: string;
  data?: T;
}

/**
 * エラー詳細情報
 */
model ErrorDetails {
  message: string;
  code?: string;
  details?: unknown;
}

/**
 * ページネーション用の共通パラメータ
 */
model PaginationQuery {
  @query page?: int32 = 1;
  // 1ページあたりの件数
  @query limit?: int32 = 20;
}

/**
 * ページネーション付きレスポンス
 */
model PaginatedResponse<T> {
  data: T[];
  pagination: {
    page: int32;
    limit: int32;
    total: int32;
    totalPages: int32;
  };
}

/**
 * 画像ファイル情報
 */
model ImageFile {
  id: string;
  originalName: string;
  mimeType: string;
  size: int64;
  url: string;
  createdAt: utcDateTime;
}

/**
 * 共通のタイムスタンプ
 */
model Timestamps {
  createdAt: utcDateTime;
  updatedAt: utcDateTime;
}

/**
 * ソート順序
 */
enum SortOrder {
  asc: "asc",
  desc: "desc",
}
