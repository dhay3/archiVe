# springboot Xss(跨站脚本攻击)

### #依赖

```xml
        <dependency>
            <groupId>org.jsoup</groupId>
            <artifactId>jsoup</artifactId>
            <version>1.13.1</version>
        </dependency>
        <!-- BooleanUtils在该依赖下 -->
        <dependency>
            <groupId>org.apache.commons</groupId>
            <artifactId>commons-lang3</artifactId>
            <version>3.10</version>
        </dependency>
```

### #工具类

```java

/**
 * xss过滤工具
 */
public class JsoupUtils {
    //设置白名单
    private static final Whitelist WHITELIST = Whitelist.basicWithImages();
    //配置过滤参数不对代码格式化
    private static final Document
            .OutputSettings OUTPUT_SETTINGS = new Document
            //默认开启,关闭输入的代码格式化
            .OutputSettings().prettyPrint(false);

    static {
        //为标签添加属性,使用伪标签(:all表示所有标签),这里指白名单中的标签
        //允许富文本编辑器设置行内样式
        WHITELIST.addAttributes(":all", "style");
    }

    /**
     * content是用户输入的内容,没有baseUri,所以设置空
     * 过滤,如果不需要baseUri 就使用空字符串
     * 从不信任的html片段中截取信任的片段
     */
    public static String clean(String content) {
        return Jsoup.clean(content, "", WHITELIST, OUTPUT_SETTINGS);
    }
    /*
     这里能发现事件被过滤了  
    public static void main(String[] args) {
       String text = "<a href=\"http://www.baidu.com/a\" onclick=\"alert(1);\">sss</a><script>alert(0);</script>sss";
       System.out.println(clean(text));
    }
     */
}

```

### #request包装类

```java

/**
 * 核心
 * 过滤http请求中参数包含的恶意字符
 * 需要重写getParameter,getParameterValues,getHeader
 */
public class XssHttpServletRequestWrapper extends HttpServletRequestWrapper {
    //原始的请求
    public HttpServletRequest orgRequest;
    //是否包含富文本
    private boolean isIncludeRichText;


    public XssHttpServletRequestWrapper(HttpServletRequest request, boolean isIncludeRichText) {
        super(request);
        orgRequest = request;
        this.isIncludeRichText = isIncludeRichText;
    }

    public boolean isIncludeRichText() {
        return isIncludeRichText;
    }

    public void setIncludeRichText(boolean includeRichText) {
        isIncludeRichText = includeRichText;
    }

    /**
     * 过滤请求头
     */
    @Override
    public String getHeader(String name) {
        JsoupUtils.clean(name);
        String header = super.getHeader(name);
        if (!StringUtils.isEmpty(header)) {
            return JsoupUtils.clean(name);
        }
        return header;
    }

    /**
     * 过滤请求的参数和值
     * 覆盖getParameter方法，将参数名和参数值都做xss过滤。
     * 如果需要获得原始的值，则通过super.getParameterValues(name)来获取
     * getParameterNames,getParameterValues和getParameterMap也可能需要覆盖
     */
    @Override
    public String getParameter(String name) {
        boolean condition = Objects.equals("content", name) || name.endsWith("WithHtml");
        //如果请求的参数为content或是以WithHtml结尾的,且不包含富文本
        if (condition && !isIncludeRichText) {
            //不过滤参数
            return super.getParameter(name);
        }
        //过滤参数
        JsoupUtils.clean(name);
        String value = super.getParameter(name);
        //如果值不为null和空字符串""( " "不算空字符串因为就是判断长度)过滤值
        if (!StringUtils.isEmpty(value)) {
            JsoupUtils.clean(value);
        }
        return value;
    }

    /**
     * 过滤单个参数多个值
     * 如复选框
     */
    @Override
    public String[] getParameterValues(String name) {
        String[] values = super.getParameterValues(name);
        for (int i = 0; i < values.length; i++) {
            //过滤值后重新赋值
            values[i] = JsoupUtils.clean(values[i]);
        }
        return values;
    }

    public HttpServletRequest getOrgRequest() {
        return orgRequest;
    }

    public void setOrgRequest(HttpServletRequest orgRequest) {
        this.orgRequest = orgRequest;
    }

    /**
     * 获取原始的request请求
     */
    public static HttpServletRequest getOrgRequest(HttpServletRequest request) {
        if (request instanceof XssHttpServletRequestWrapper) {
            return ((XssHttpServletRequestWrapper) request).getOrgRequest();
        }
        return request;
    }
}

```

### #filter

```java

/**
 * XssFilter过滤Xss请求的入口
 * 拦截防止xss
 */
@Slf4j
public class XssFilter implements Filter {
    //LoggerFactory log = LoggerFactory.getLogger(XssFilter.class)
    //是否包含富文本内容
    public static boolean IS_INCLUDE_RICH_TEXT = false;

    public List<String> excludes = new ArrayList<>();

    @Override
    public void init(FilterConfig filterConfig) throws ServletException {
        log.debug("----------------- xss filter init ----------------");
        //获取filter中的初始参数
        String isRichText = filterConfig.getInitParameter("isIncludeRichText");
        if (!StringUtils.isEmpty(isRichText)) {
            //将字符串转为布尔
            IS_INCLUDE_RICH_TEXT = BooleanUtils.toBoolean(isRichText);
        }
        String temp = filterConfig.getInitParameter("excludes");
        if (!StringUtils.isEmpty(temp)) {
            String[] url = temp.split(",");
            //spring工具类
            Assert.notNull(url, "exclude不能为null");
            excludes.addAll(Arrays.asList(url));
        }
    }

    @Override
    public void doFilter(ServletRequest request, ServletResponse response, FilterChain chain) throws IOException, ServletException {
        log.debug("----------------------xss filter is open----------------------");
        HttpServletRequest req = (HttpServletRequest) request;
        HttpServletResponse resp = (HttpServletResponse) response;
        if (handleExcludeURL(req, resp)) {
            //包含exclude的url片段放行
            chain.doFilter(request, response);
            return;
        }
        //不包含exclude的url片段
        //将request包装到XssHttpServletRequestWrapper
        XssHttpServletRequestWrapper xssRequest = new XssHttpServletRequestWrapper(
                (HttpServletRequest) request, IS_INCLUDE_RICH_TEXT);
        //放行,交给spring处理
        chain.doFilter(xssRequest, response);
    }

    private boolean handleExcludeURL(HttpServletRequest request, HttpServletResponse response) {
        //spring工具类,判断集合是否为空集合或集合为null
        if (CollectionUtils.isEmpty(excludes)) {
            return false;
        }
        //获取请求的servletPath,不带协议+ip+端口+项目名
        String servletPath = request.getServletPath();
        for (String pattern : excludes) {
            //这里的^表示匹配开头
            Pattern p = Pattern.compile("^" + pattern);
            //String实现CharSequence, 用pattern去匹配servlet
            Matcher matcher = p.matcher(servletPath);
            //判断请求的servletPath中是否有匹配pattern的,只要有一个就返回true
            if (matcher.find()) {
                return true;
            }
        }
        return false;
    }
}

```

### #配置类

```java

@Configuration
public class JsoupConf {
    /**
     * 注册jsoup Filter
     */
    @Bean
    public FilterRegistrationBean<XssFilter> xssfilterRegistrationBean() {
        FilterRegistrationBean<XssFilter> filterRegistrationBean =
                new FilterRegistrationBean<>(new XssFilter());
		//filterRegistrationBean.setFilter(new XssFilter());
        //设置在filter中执行的顺序,为第一执行
        filterRegistrationBean.setOrder(1);
        //指明filter是否开启,默认开启
        filterRegistrationBean.setEnabled(true);
        //拦截所有请求
        filterRegistrationBean.addUrlPatterns("/*");
        HashMap<String, String> initParmas = new HashMap<>();
        //路径自行替换为static下的,首页
        initParmas.put("excludes", "/favicon.ico,/img/*,/js/*,/css/*,/index");
        initParmas.put("isIncludeRichText", "true");
        //设置filter的init-param
        filterRegistrationBean.setInitParameters(initParmas);
        return filterRegistrationBean;
    }
}

```

