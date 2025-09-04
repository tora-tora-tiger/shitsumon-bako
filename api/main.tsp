import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "./models/common.tsp";
import "./models/user.tsp";
import "./models/question.tsp";
import "./routes/users.tsp";
import "./routes/questions.tsp";
import "./routes/files.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;

/**
 * 質問箱サービス API
 */
@service
@server("http://localhost:8080", "Development server")
@server("https://api.shitsumon-bako.com", "Production server")
namespace ShitsumonBako;

/**
 * 共通エラーハンドリング
 */
@error
model Error {
  code: int32;
  message: string;
}
