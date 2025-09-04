import "@typespec/http";
import "@typespec/rest";
import "../models/common.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;

namespace ShitsumonBako;

/**
 * ファイル関連のエンドポイント
 */
@route("/api/files")
interface Files {
  /*
   * 画像ファイルアップロード
   */
  @post
  @route("/images")
  uploadImage(@header authorization?: string, @body image: bytes):
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
  deleteImage(@header authorization: string, @path imageId: string):
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
