import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi";
import "../models/common.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.OpenAPI;

namespace ShitsumonBako;

/**
 * ファイル関連のエンドポイント
 */
@tag("Files")
@route("/api/files")
@useAuth(BearerAuth)
interface Files {
  /**
   * 画像ファイルアップロード
   */
  @post
  @route("/images")
  @operationId("uploadImage")
  @summary("画像ファイルをアップロード")
  @useAuth([])
  uploadImage(@body image: bytes):
    | {
        @statusCode _: 201;
        @body response: ApiResponse<ImageFile>;
      }
    | {
        @statusCode _: 400;
        @body error: ApiResponse<ErrorDetails>;
      }
    | {
        @statusCode _: 413;
        @body error: ApiResponse<ErrorDetails>;
      }
    | {
        @statusCode _: 415;
        @body error: ApiResponse<ErrorDetails>;
      };

  /**
   * 画像ファイル取得
   */
  @get
  @route("/images/{imageId}")
  @operationId("getImage")
  @summary("画像ファイルを取得")
  @useAuth([])
  getImage(@path imageId: string): {
    @statusCode _: 200;
    @header `content-type`: string;
    @body content: bytes;
  } | {
    @statusCode _: 404;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * 画像ファイル削除
   */
  @delete
  @route("/images/{imageId}")
  @operationId("deleteImage")
  @summary("画像ファイルを削除")
  deleteImage(@path imageId: string):
    | {
        @statusCode _: 200;
        @body response: ApiResponse<{}>;
      }
    | {
        @statusCode _: 401;
        @body error: ApiResponse<ErrorDetails>;
      }
    | {
        @statusCode _: 404;
        @body error: ApiResponse<ErrorDetails>;
      }
    | {
        @statusCode _: 403;
        @body error: ApiResponse<ErrorDetails>;
      };
}
