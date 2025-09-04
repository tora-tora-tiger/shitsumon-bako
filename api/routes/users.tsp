import "@typespec/http";
import "@typespec/rest";
import "../models/common.tsp";
import "../models/user.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;

namespace ShitsumonBako;

/**
 * ユーザー関連のエンドポイント
 */
@route("/api/users")
interface Users {
  /**
   * ユーザー登録
   */
  @post
  register(@body request: UserRegistrationRequest): {
    @statusCode _: 201;
    @body response: ApiResponse<User>;
  } | {
    @statusCode _: 400;
    @body error: ApiResponse<ErrorDetails>;
  } | {
    @statusCode _: 409;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * ログイン
   */
  @post
  @route("/login")
  login(@body request: LoginRequest): {
    @statusCode _: 200;
    @body response: ApiResponse<LoginResponse>;
  } | {
    @statusCode _: 401;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * ログアウト
   */
  @post
  @route("/logout")
  logout(@header authorization: string): {
    @statusCode _: 200;
    @body response: ApiResponse<{}>;
  } | {
    @statusCode _: 401;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * 現在のユーザー情報取得
   */
  @get
  @route("/me")
  getCurrentUser(@header authorization: string): {
    @statusCode _: 200;
    @body response: ApiResponse<User>;
  } | {
    @statusCode _: 401;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * プロフィール更新
   */
  @patch
  @route("/me")
  updateProfile(
    @header authorization: string,
    @body request: ProfileUpdateRequest,
  ): {
    @statusCode _: 200;
    @body response: ApiResponse<User>;
  } | {
    @statusCode _: 401;
    @body error: ApiResponse<ErrorDetails>;
  } | {
    @statusCode _: 400;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * 公開ユーザー情報取得
   */
  @get
  @route("/{userId}")
  getPublicUserInfo(@path userId: string): {
    @statusCode _: 200;
    @body response: ApiResponse<PublicUserInfo>;
  } | {
    @statusCode _: 404;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * ユーザーブロック
   */
  @post
  @route("/block")
  blockUser(
    @header authorization: string,
    @body request: {
      userId: string;
    },
  ): {
    @statusCode _: 201;
    @body response: ApiResponse<BlockedUser>;
  } | {
    @statusCode _: 401;
    @body error: ApiResponse<ErrorDetails>;
  } | {
    @statusCode _: 400;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * ブロックユーザー一覧取得
   */
  @get
  @route("/blocked")
  getBlockedUsers(@header authorization: string, ...PaginationQuery): {
    @statusCode _: 200;
    @body response: PaginatedResponse<BlockedUser>;
  } | {
    @statusCode _: 401;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * ユーザーブロック解除
   */
  @delete
  @route("/block/{userId}")
  unblockUser(@header authorization: string, @path userId: string): {
    @statusCode _: 200;
    @body response: ApiResponse<{}>;
  } | {
    @statusCode _: 401;
    @body error: ApiResponse<ErrorDetails>;
  } | {
    @statusCode _: 404;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * NGワード追加
   */
  @post
  @route("/ng-words")
  addNgWord(
    @header authorization: string,
    @body request: NgWordCreateRequest,
  ): {
    @statusCode _: 201;
    @body response: ApiResponse<NgWord>;
  } | {
    @statusCode _: 401;
    @body error: ApiResponse<ErrorDetails>;
  } | {
    @statusCode _: 400;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * NGワード一覧取得
   */
  @get
  @route("/ng-words")
  getNgWords(@header authorization: string, ...PaginationQuery): {
    @statusCode _: 200;
    @body response: PaginatedResponse<NgWord>;
  } | {
    @statusCode _: 401;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * NGワード削除
   */
  @delete
  @route("/ng-words/{ngWordId}")
  deleteNgWord(@header authorization: string, @path ngWordId: string): {
    @statusCode _: 200;
    @body response: ApiResponse<{}>;
  } | {
    @statusCode _: 401;
    @body error: ApiResponse<ErrorDetails>;
  } | {
    @statusCode _: 404;
    @body error: ApiResponse<ErrorDetails>;
  };
}
