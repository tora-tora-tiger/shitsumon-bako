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
model ProfileBase {
  id: string;
  type: ProfileType;
  name: string;
  value: string;
  order: int32;
  ...Timestamps;
}

union ProfileType {
  "url",
  "text",
  "email",
}

/**
 * プロフィール項目： Twitter
 */
model ProfileTwitter extends ProfileBase {
  type: "text";
  name: "Twitter";
}

model ProfileSite extends ProfileBase {
  type: "url";
  name: "Webサイト";
}

model ProfileEmail extends ProfileBase {
  type: "email";
  name: "メールアドレス";
}

/**
 * プロフィール項目
 */
union ProfileField {
  ProfileTwitter,
  ProfileSite,
  ProfileEmail,
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
  id: Identificator;
  password: string;
}

union Identificator {
  displayId: string;
  email: string;
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
  displayId?: string;
  email?: string;
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
