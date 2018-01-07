<!-- content srart -->
<div class="am-g am-g-fixed blog-fixed blog-content">
    <div class="am-u-sm-12">
        <article class="am-article blog-article-p">
            <div class="am-article-hd">
                <h1 class="am-article-title blog-text-center">{{.post.Title}}</h1>
                <p class="am-article-meta blog-text-center">
                    <span><a href="#" class="blog-color">article &nbsp;</a></span>-
                    <span><a href="#">@嘟嘟噜 &nbsp;</a></span>-
                    <span><a href="#">{{date .post.PushTime "M/d/Y H:i:s"}}</a></span>
                </p>
            </div>
            <hr>
            <div class="am-article-bd">
                <img src="{{replace_image_src .image.RealPath}}?imageView2/1/w/1920/h/356" alt=""
                     class="blog-entry-img blog-article-margin">
            {{str2html .content}}
            </div>
        </article>

        <div class="am-g blog-article-widget blog-article-margin">
            <div class="am-u-lg-4 am-u-md-5 am-u-sm-7 am-u-sm-centered blog-text-center">
                <a href="/reward?postId={{.post.Id}}" class="am-btn am-btn-default">打赏支持博主</a>
            </div>
        </div>
        <hr>

        <div class="am-g blog-article-widget blog-article-margin">
            <div class="am-u-lg-4 am-u-md-5 am-u-sm-7 am-u-sm-centered blog-text-center">
                <span class="am-icon-tags"> &nbsp;</span><a href="#">标签</a> , <a href="#">TAG</a> , <a href="#">啦啦</a>
                <hr>
                <a href=""><span class="am-icon-qq am-icon-fw am-primary blog-icon"></span></a>
                <a href=""><span class="am-icon-wechat am-icon-fw blog-icon"></span></a>
                <a href=""><span class="am-icon-weibo am-icon-fw blog-icon"></span></a>
            </div>
        </div>

        <hr>
    {{/*<div class="am-g blog-author blog-article-margin">*/}}
    {{/*<div class="am-u-sm-3 am-u-md-3 am-u-lg-2">*/}}
    {{/*<img src="/static/assets/i/f15.jpg" alt="" class="blog-author-img am-circle">*/}}
    {{/*</div>*/}}
    {{/*<div class="am-u-sm-9 am-u-md-9 am-u-lg-10">*/}}
    {{/*<h3><span>作者 &nbsp;: &nbsp;</span><span class="blog-color">amazeui</span></h3>*/}}
    {{/*<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>*/}}
    {{/*</div>*/}}
    {{/*</div>*/}}
    {{/*<hr>*/}}
    {{/*<ul class="am-pagination blog-article-margin">*/}}
    {{/*<li class="am-pagination-prev"><a href="#" class="">&laquo; 一切的回顾</a></li>*/}}
    {{/*<li class="am-pagination-next"><a href="">不远的未来 &raquo;</a></li>*/}}
    {{/*</ul>*/}}

    {{/*<hr>*/}}

    {{/*<form class="am-form am-g">*/}}
    {{/*<h3 class="blog-comment">评论</h3>*/}}
    {{/*<fieldset>*/}}
    {{/*<div class="am-form-group am-u-sm-4 blog-clear-left">*/}}
    {{/*<input type="text" class="" placeholder="名字">*/}}
    {{/*</div>*/}}
    {{/*<div class="am-form-group am-u-sm-4">*/}}
    {{/*<input type="email" class="" placeholder="邮箱">*/}}
    {{/*</div>*/}}

    {{/*<div class="am-form-group am-u-sm-4 blog-clear-right">*/}}
    {{/*<input type="password" class="" placeholder="网站">*/}}
    {{/*</div>*/}}

    {{/*<div class="am-form-group">*/}}
    {{/*<textarea class="" rows="5" placeholder="一字千金"></textarea>*/}}
    {{/*</div>*/}}

    {{/*<p><button type="submit" class="am-btn am-btn-default">发表评论</button></p>*/}}
    {{/*</fieldset>*/}}
    {{/*</form>*/}}

    {{/*<hr>*/}}
    </div>
</div>
<!-- content end -->