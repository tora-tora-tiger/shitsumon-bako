import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

using TypeSpec.Http;
using TypeSpec.Rest;

namespace ShitsumonBako;

/**
 * レスポンス構造
 */
model ApiResponse<T> {
  statusCode: int32;
  message?: string;
  data?: T;
}

/**
 * エラーレスポンス
 */
model ErrorResponse {
  message: string;
  code?: string;
  details?: unknown;
}

/**
 * ページネーション用の共通パラメータ
 */
model PaginationQuery {
  @query page?: int32 = 1;
  @query limit?: int32 = 20;
}

/**
 * ページネーション付きレスポンス
 */
model PaginatedResponse<T> {
  success: true;
  data: T[];
  pagination: {
    page: int32;
    limit: int32;
    total: int32;
    totalPages: int32;
    hasNext: boolean;
    hasPrev: boolean;
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
