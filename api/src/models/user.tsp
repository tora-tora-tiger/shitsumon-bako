import "@typespec/http";
import "@typespec/rest";
import "./common.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;

namespace ShitsumonBako;

/**
 * ユーザープロフィール情報
 */
model User {
  // 内部ID
  id: string;

  // ユーザーID（表示用、一意）
  displayId: string;

  email: string;
  displayName: string;
  iconImage?: ImageFile;
  profileFields?: ProfileField[];
  isActive: boolean;
  ...Timestamps;
}

/**
 * 公開用ユーザー情報
 */
model PublicUserInfo {
  displayid: string;
  displayName: string;
  iconImage?: ImageFile;
  profileFields?: ProfileField[];
}

/**
 * プロフィール項目ベース
 */
model ProfileFieldBase {
  id: string;
  type: ProfileFieldType;
  name: string;
  value: string;
  order: int32;
}

union ProfileFieldType {
  "url",
  "text",
  "email",
}

/**
 * プロフィール項目： Twitter
 */
model ProfileFieltTwitter extends ProfileFieldBase {
  type: "text";
  name: "Twitter";
}

model ProfileFieldSite extends ProfileFieldBase {
  type: "url";
  name: "Webサイト";
}

model ProfileFieldEmail extends ProfileFieldBase {
  type: "email";
  name: "メールアドレス";
}

/**
 * プロフィール項目
 */
union ProfileField {
  ProfileFieltTwitter,
  ProfileFieldSite,
  ProfileFieldEmail,
}

/**
 * ユーザー登録リクエスト
 */
model UserRegistrationRequest {
  email: string;
  password: string;
  displayId: string;
  displayName: string;
}

/**
 * ログインリクエスト
 */
model LoginRequest {
  email: string;
  password: string;
}

/**
 * ログインレスポンス
 */
model LoginResponse {
  user: User;
  accessToken: string;
  refreshToken: string;
}

/**
 * プロフィール更新リクエスト
 */
model ProfileUpdateRequest {
  displayName?: string;
  iconImageId?: string;
  profileFields?: ProfileFieldUpdate[];
}

/**
 * プロフィールフィールド更新
 */
model ProfileFieldUpdate {
  id?: string;
  name: string;
  value: string;
  order: int32;
  isVisible: boolean;
}

/**
 * ブロックユーザー情報
 */
model BlockedUser {
  id: string;
  blockedUser: PublicUserInfo;
  ...Timestamps;
}

/**
 * NGワード
 */
model NgWord {
  id: string;
  word: string;
  isActive: boolean;
  ...Timestamps;
}

/**
 * NGワード追加リクエスト
 */
model NgWordCreateRequest {
  word: string;
}
