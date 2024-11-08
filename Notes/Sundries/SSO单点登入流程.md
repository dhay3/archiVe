# SSO单点登入流程

- 用户登入后, 通过JWT将通过一定规则生成一个token, 不建议将重要信息放入token中, 一般存放唯一标识符

  这里为什么不在登入后就将用户信息存放入cookie中是因为, cookie存放信息不安全

- 将token存入Cookie中

- 创建拦截器, 拦截所有请求, 在请求前判断一下Cookie中是否有token, 如果有token就在**headers**中添加

  一个token属性

- 登入调用后端api, 通过JWT获取唯一表示符, 通过唯一标识符查询数据库, 将结果返回给前端, 前端将用户信息存储到Cookie中

  ```java
      public static String getMemberIdByJwtToken(HttpServletRequest request) {
          String jwtToken = request.getHeader("token");
          if (StringUtils.isEmpty(jwtToken)) return "";
          Jws<Claims> claimsJws = Jwts.parser().setSigningKey(APP_SECRET).parseClaimsJws(jwtToken);
          Claims claims = claimsJws.getBody();
          return (String) claims.get("id");
      }
  ```

  ```java
      @GetMapping("/token")
      public ResponseBo getUserInfoByToken(HttpServletRequest request) {
          //调用jwt工具类获取header中的token
          String memberId = JwtUtils.getMemberIdByJwtToken(request);
          //根据用户id获取用户信息
          Member member = memberService.getById(memberId);
          System.out.println(member);
          return ResponseBo.ok().data("user", member);
      }
  ```

- 从Cookie中获取用户信息

