# golang Token

## Header

    Header 包含了一些元數據，至少會表明 token 類型以及 簽名方法。比如

    {
        "typ" : "JWT",
        "alg" : "HS256"
    }
    type: 必需。 token 類型，JWT 表示是 JSON Web Token.

alg: 必需。 token 所使用的簽名算法，可用的值在 這裡 有規定。

## Claims (Payload)

    Claims 部分包含了一些跟這個 token 有關的重要信息。 JWT 標準規定了一些字段，下面節選一些字段:

    iss: The issuer of the token，token 是給誰的
    sub: The subject of the token，token 主題
    exp: Expiration Time。 token 過期時間，Unix 時間戳格式
    iat: Issued At。 token 創建時間， Unix 時間戳格式
    jti: JWT ID。針對當前 token 的唯一標識
    除了規定的字段外，可以包含其他任何 JSON 兼容的字段。

    Payload 示例:

    {
        "iss": "mozillazg.com",
        "exp": 1435055117,
        "user_id": 1,
        "foo": "bar"
    }

## Signature

    JWT 標準遵照 JSON Web Signature (JWS) 標準來生成簽名。簽名主要用於驗證 token 是否有效，是否被篡改。簽名時可以 這些算法進行簽名，比如 HMAC SHA-256:

    content = base64url_encode(Header) + '.' + base64url_encode(Claims)
    signature = hmacsha256.hash(content)
    說到這裡有一點需要特別注意的是，默認情況下，JWT 中信息都是明文的，即 Claims 的內容並沒有 被加密，可以通過 base64url_decode(text) 的方式解碼得到 Claims 。所以，不要在 Claims 裡包含敏感信息，如果一定要包含敏感信息的話，記得先將 Claims 的內容進行加密（比如，使用 JSON Web Encryption (JWE) 標准進行加密） 然後在進行 base64url_encode 操作。

參考: https://mozillazg.com/2015/06/hello-jwt.html
