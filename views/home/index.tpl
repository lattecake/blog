<!-- banner start -->
<div class="am-g am-g-fixed blog-fixed am-u-sm-centered blog-article-margin">
    <div data-am-widget="slider" class="am-slider am-slider-b1" data-am-slider='{&quot;controlNav&quot;:false}'>
        <ul class="am-slides">
        {{range $k, $post := .starPosts}}
            <li>
                <img src="{{replace_image_src $post.Image}}?imageView/1/w/1280/h/320">
                <div class="blog-slider-desc am-slider-desc ">
                    <div class="blog-text-center blog-slider-con">
                        {{/*<span><a href="/post/{{$post.Id}}" class="blog-color">Article &nbsp;</a></span>*/}}
                        <h1 class="blog-h-margin"><a href="/post/{{$post.Id}}">{{$post.Title}}</a></h1>
                        <p>{{substr $post.Description 0 24}}
                        </p>
                        <span class="blog-bor">{{date $post.PushTime "M/d/Y"}}</span>
                        <br><br><br><br><br><br><br>
                    </div>
                </div>
            </li>
        {{end}}

        </ul>
    </div>
</div>
<!-- banner end -->

<!-- content srart -->
<div class="am-g am-g-fixed blog-fixed">
    <div class="am-u-md-8 am-u-sm-12">

    {{range $k, $post := .posts}}

        <article class="am-g blog-entry-article">
            <div class="am-u-lg-6 am-u-md-12 am-u-sm-12 blog-entry-img">
                <a href="/post/{{$post.Id}}">
                    <img src="{{replace_image_src $post.Image}}?imageView/1/w/1280/h/720" alt="" class="am-u-sm-12">
                </a>
            </div>
            <div class="am-u-lg-6 am-u-md-12 am-u-sm-12 blog-entry-text">
                <span><a href="" class="blog-color">article &nbsp;</a></span>
                <span> @嘟嘟噜 &nbsp;</span>
                <span>{{date $post.PushTime "M/d/Y"}}</span>
                <h1><a href="/post/{{$post.Id}}">{{$post.Title}}</a></h1>
                <p>{{substr $post.Description 0 30}}</p>
                <p><a href="/post/{{$post.Id}}" class="blog-continue">continue reading</a></p>
            </div>
        </article>
    {{end}}
        <ul class="am-pagination">
        {{/*<li class="am-pagination-prev"><a href="">&laquo; Prev</a></li>*/}}
            <li class="am-pagination-next"><a href="/posts/2?s=home">More &raquo;</a></li>
        </ul>
    </div>

    <div class="am-u-md-4 am-u-sm-12 blog-sidebar">
        <div class="blog-sidebar-widget blog-bor">
            <h2 class="blog-text-center blog-title"><span>薛定谔的猿</span></h2>
            <img src="https://lattecake.oss-cn-beijing.aliyuncs.com/static%2Fimages%2Fweixin%2Fqrcode_for_gh_354bc8e8b814_1280.jpg"
                 alt="about me"
                 class="blog-entry-img">
            <p>不跑马拉松的摄影师不是好程序员</p>
            <p>
                How does the world look through your eyes?
            </p>
        </div>
        <div class="blog-sidebar-widget blog-bor">
            <h2 class="blog-text-center blog-title"><span>Contact ME</span></h2>
            <p>
                <a href="https://i.qq.com/?rd=1"><span class="am-icon-qq am-icon-fw am-primary blog-icon"></span></a>
                <a href="https://github.com/lattecake"><span class="am-icon-github am-icon-fw blog-icon"></span></a>
                <a href="https://weibo.com/solacowa/"><span class="am-icon-weibo am-icon-fw blog-icon"></span></a>
                <a href="http://wx.lattecake.com/"><span class="am-icon-weixin am-icon-fw blog-icon"></span></a>
            </p>
        </div>
        <div class="blog-clear-margin blog-sidebar-widget blog-bor am-g ">
            <h2 class="blog-title"><span>TAG cloud</span></h2>
            <div class="am-u-sm-12 blog-clear-padding">
                <a href="" class="blog-tag">码农</a>
                <a href="" class="blog-tag">Gopher</a>
                <a href="" class="blog-tag">开源</a>
            </div>
        </div>
        <div class="blog-sidebar-widget blog-bor">
            <h2 class="blog-title"><span>么么哒</span></h2>
            <ul class="am-list">
                <li><a href="#">每个人都有一个死角， 自己走不出来，别人也闯不进去。</a></li>
                <li><a href="#">多赚点钱，赚的钱多了，很多事自然就变了。</a></li>
                <li><a href="#">每个人都有一道伤口， 或深或浅，盖上布，以为不存在。</a></li>
            </ul>
        </div>
    </div>
</div>
<!-- content end -->
