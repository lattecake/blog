<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="description" content="LatteCake个人博客,拿铁味的蛋糕,是一名码农,擅长各种技术.">
    <meta name="keywords" content="博客,技术,个人博客,个人技术博客,拿铁蛋糕,拿铁味的摩卡,lattecake">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    <title>{{.title}} LatteCake</title>
    <meta name="renderer" content="webkit">
    <meta http-equiv="Cache-Control" content="no-siteapp"/>
    <link rel="icon" type="image/png" href="/static/favicon.ico">
    <meta name="mobile-web-app-capable" content="yes">
    <link rel="icon" sizes="192x192" href="/static/favicon.ico">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-status-bar-style" content="black">
    <meta name="apple-mobile-web-app-title" content="Amaze UI"/>
    <link rel="apple-touch-icon-precomposed" href="/static/favicon.ico">
    <meta name="msapplication-TileImage" content="/static/favicon.ico">
    <meta name="msapplication-TileColor" content="#0e90d2">
    <link rel="stylesheet" href="https://lattecake.oss-cn-beijing.aliyuncs.com/static/blog/css/amazeui.min.css">
    <link rel="stylesheet" href="https://lattecake.oss-cn-beijing.aliyuncs.com/static/blog/css/app.css">
</head>
<body id="blog-article-sidebar">

<!-- header start -->
<header class="am-g am-g-fixed blog-fixed blog-text-center blog-header">
    <div class="am-u-sm-8 am-u-sm-centered">
    {{/*<img width="200" src="http://s.amazeui.org/media/i/brand/amazeui-b.png" alt="LatteCake Logo"/>*/}}
        <h2 class="am-hide-sm-only">那时候刚好下着雨，柏油路面湿冷冷的，还闪烁着青、黄、红颜色的灯火。</h2>
    </div>
</header>
<!-- header end -->
<hr>

{{template "./default_navigation.tpl" .}}

{{.LayoutContent}}



<footer class="blog-footer">
    <div class="am-g am-g-fixed blog-fixed am-u-sm-centered blog-footer-padding">
        <div class="am-u-sm-12 am-u-md-4- am-u-lg-4">
            <h3>Get in touch with</h3>
            <p class="am-text-sm">坚持背负一身破铜烂铁把自己累的够呛;那沉甸甸的重担里装满了无可救药的偏执和毫无希望的等待！</p>
        </div>
        <div class="am-u-sm-12 am-u-md-4- am-u-lg-4">
            <h3>社交账号</h3>
            <p>
            {{/*<a href=""><span class="am-icon-qq am-icon-fw am-primary blog-icon blog-icon"></span></a>*/}}
                <a href="https://github.com/lattecake" target="_blank"><span
                        class="am-icon-github am-icon-fw blog-icon blog-icon"></span></a>
                <a href="https://weibo.com/solacowa"><span class="am-icon-weibo am-icon-fw blog-icon blog-icon"></span></a>
            {{/*<a href=""><span class="am-icon-reddit am-icon-fw blog-icon blog-icon"></span></a>*/}}
            {{/*<a href=""><span class="am-icon-weixin am-icon-fw blog-icon blog-icon"></span></a>*/}}
            </p>
            <h3>Credits</h3>
            <p>别让自己的心塞太满 要留出点空间让阳光照进来！</p>
        </div>
        <div class="am-u-sm-12 am-u-md-4- am-u-lg-4">
            <h1>曾经沧海桑田，已是物是人非。</h1>
            <h3>Heroes</h3>
            <p>
                很久以后我们才会明白，所有人来到你的生命里自有它的意义。人最难的就是学会怎么平静的面对离别，而这个世界的吊诡之处就在于：当你平静面对离别时，那个人已经在你心里永远不会走了...
            </p>
        </div>
    </div>
    <div class="blog-text-center">© LatteCake 2018 All right reserved. By dudulu.me</div>
</footer>


<!--[if (gte IE 9)|!(IE)]><!-->
<script src="https://lattecake.oss-cn-beijing.aliyuncs.com/static/blog/js/jquery.min.js"></script>
<!--<![endif]-->
<!--[if lte IE 8 ]>
<script src="http://libs.baidu.com/jquery/1.11.3/jquery.min.js"></script>
<script src="http://cdn.staticfile.org/modernizr/2.8.3/modernizr.js"></script>
<script src="https://lattecake.oss-cn-beijing.aliyuncs.com/static/blog/js/amazeui.ie8polyfill.min.js"></script>
<![endif]-->
<script src="https://lattecake.oss-cn-beijing.aliyuncs.com/static/blog/js/amazeui.min.js"></script>
<!-- <script src="/static/assets/js/app.js"></script> -->
</body>
</html>