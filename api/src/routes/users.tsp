import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi";
import "../models/common.tsp";
import "../models/user.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.OpenAPI;

namespace ShitsumonBako;

/**
 * ユーザー関連のエンドポイント
 */
@tag("Users")
@route("/api/users")
@useAuth(BearerAuth)
interface Users {
  /**
   * ユーザー登録
   */
  @post
  @operationId("registerUser")
  @summary("新規ユーザー登録")
  @useAuth([])
  register(@body request: UserRegistrationRequest): {
    @statusCode statusCode: 201;
    @body response: ApiResponse<User>;
  } | {
    @statusCode statusCode: 400;
    @body error: ApiResponse<ErrorDetails>;
  } | {
    @statusCode statusCode: 409;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * ログイン
   */
  @post
  @route("/login")
  @operationId("loginUser")
  @summary("ユーザーログイン")
  @useAuth([])
  login(@body request: LoginRequest): {
    @statusCode statusCode: 200;
    @body response: ApiResponse<LoginResponse>;
  } | {
    @statusCode statusCode: 401;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * ログアウト
   */
  @post
  @route("/logout")
  @operationId("logoutUser")
  @summary("ユーザーログアウト")
  logout(): {
    @statusCode statusCode: 200;
    @body response: ApiResponse<{}>;
  } | {
    @statusCode statusCode: 401;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * 現在のユーザー情報取得
   */
  @get
  @route("/me")
  @operationId("getCurrentUser")
  @summary("ログイン中のユーザー情報取得")
  getCurrentUser(): {
    @statusCode statusCode: 200;
    @body response: ApiResponse<User>;
  } | {
    @statusCode statusCode: 401;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * プロフィール更新
   */
  @patch(#{implicitOptionality: true})
  @route("/me")
  @operationId("updateUserProfile")
  @summary("ユーザープロフィール更新")
  updateProfile(@body request: ProfileUpdateRequest): {
    @statusCode statusCode: 200;
    @body response: ApiResponse<User>;
  } | {
    @statusCode statusCode: 401;
    @body error: ApiResponse<ErrorDetails>;
  } | {
    @statusCode statusCode: 400;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * 公開ユーザー情報取得
   */
  @get
  @route("/{userId}")
  @operationId("getPublicUserInfo")
  @summary("公開ユーザー情報取得")
  @useAuth([])
  getPublicUserInfo(@path userId: string): {
    @statusCode statusCode: 200;
    @body response: ApiResponse<PublicUserInfo>;
  } | {
    @statusCode statusCode: 404;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * ユーザーブロック
   */
  @post
  @route("/block")
  @operationId("blockUser")
  @summary("ユーザーをブロック")
  blockUser(
    @body request: {
      userId: string;
    },
  ): {
    @statusCode statusCode: 201;
    @body response: ApiResponse<BlockedUser>;
  } | {
    @statusCode statusCode: 401;
    @body error: ApiResponse<ErrorDetails>;
  } | {
    @statusCode statusCode: 400;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * ブロックユーザー一覧取得
   */
  @get
  @route("/blocked")
  @operationId("getBlockedUsers")
  @summary("ブロックしたユーザー一覧取得")
  getBlockedUsers(...PaginationQuery): {
    @statusCode statusCode: 200;
    @body response: PaginatedResponse<BlockedUser>;
  } | {
    @statusCode statusCode: 401;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * ユーザーブロック解除
   */
  @delete
  @route("/block/{userId}")
  @operationId("unblockUser")
  @summary("ユーザーのブロックを解除")
  unblockUser(@path userId: string): {
    @statusCode statusCode: 200;
    @body response: ApiResponse<{}>;
  } | {
    @statusCode statusCode: 401;
    @body error: ApiResponse<ErrorDetails>;
  } | {
    @statusCode statusCode: 404;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * NGワード追加
   */
  @post
  @route("/ng-words")
  @operationId("addNgWord")
  @summary("NGワードを追加")
  addNgWord(@body request: NgWordCreateRequest): {
    @statusCode statusCode: 201;
    @body response: ApiResponse<NgWord>;
  } | {
    @statusCode statusCode: 401;
    @body error: ApiResponse<ErrorDetails>;
  } | {
    @statusCode statusCode: 400;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * NGワード一覧取得
   */
  @get
  @route("/ng-words")
  @operationId("getNgWords")
  @summary("NGワード一覧取得")
  getNgWords(...PaginationQuery): {
    @statusCode statusCode: 200;
    @body response: PaginatedResponse<NgWord>;
  } | {
    @statusCode statusCode: 401;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * NGワード削除
   */
  @delete
  @route("/ng-words/{ngWordId}")
  @operationId("deleteNgWord")
  @summary("NGワードを削除")
  deleteNgWord(@path ngWordId: string): {
    @statusCode statusCode: 200;
    @body response: ApiResponse<{}>;
  } | {
    @statusCode statusCode: 401;
    @body error: ApiResponse<ErrorDetails>;
  } | {
    @statusCode statusCode: 404;
    @body error: ApiResponse<ErrorDetails>;
  };
}
