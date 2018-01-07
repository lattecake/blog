<!-- banner start -->
<div class="am-g am-g-fixed blog-fixed am-u-sm-centered blog-article-margin">
    <div data-am-widget="slider" class="am-slider am-slider-b1" data-am-slider='{&quot;controlNav&quot;:false}'>
        <ul class="am-slides">
            <li>
                <img src="/static/assets/i/b1.jpg">
                <div class="blog-slider-desc am-slider-desc ">
                    <div class="blog-text-center blog-slider-con">
                        <span><a href="" class="blog-color">Article &nbsp;</a></span>
                        <h1 class="blog-h-margin"><a href="">总在思考一句积极的话</a></h1>
                        <p>那时候刚好下着雨，柏油路面湿冷冷的，还闪烁着青、黄、红颜色的灯火。
                        </p>
                        <span class="blog-bor">2015/10/9</span>
                        <br><br><br><br><br><br><br>
                    </div>
                </div>
            </li>
            <li>
                <img src="/static/assets/i/b2.jpg">
                <div class="am-slider-desc blog-slider-desc">
                    <div class="blog-text-center blog-slider-con">
                        <span><a href="" class="blog-color">Article &nbsp;</a></span>
                        <h1 class="blog-h-margin"><a href="">总在思考一句积极的话</a></h1>
                        <p>那时候刚好下着雨，柏油路面湿冷冷的，还闪烁着青、黄、红颜色的灯火。
                        </p>
                        <span>2015/10/9</span>
                    </div>
                </div>
            </li>
            <li>
                <img src="/static/assets/i/b3.jpg">
                <div class="am-slider-desc blog-slider-desc">
                    <div class="blog-text-center blog-slider-con">
                        <span><a href="" class="blog-color">Article &nbsp;</a></span>
                        <h1 class="blog-h-margin"><a href="">总在思考一句积极的话</a></h1>
                        <p>那时候刚好下着雨，柏油路面湿冷冷的，还闪烁着青、黄、红颜色的灯火。
                        </p>
                        <span>2015/10/9</span>
                    </div>
                </div>
            </li>
            <li>
                <img src="/static/assets/i/b2.jpg">
                <div class="am-slider-desc blog-slider-desc">
                    <div class="blog-text-center blog-slider-con">
                        <span><a href="" class="blog-color">Article &nbsp;</a></span>
                        <h1 class="blog-h-margin"><a href="">总在思考一句积极的话</a></h1>
                        <p>那时候刚好下着雨，柏油路面湿冷冷的，还闪烁着青、黄、红颜色的灯火。
                        </p>
                        <span>2015/10/9</span>
                    </div>
                </div>
            </li>
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
                    {{/*<img src="{{map_get .postImages $post.Id}}" alt="" class="am-u-sm-12">*/}}
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
            <li class="am-pagination-prev"><a href="">&laquo; Prev</a></li>
            <li class="am-pagination-next"><a href="">Next &raquo;</a></li>
        </ul>
    </div>

    <div class="am-u-md-4 am-u-sm-12 blog-sidebar">
        <div class="blog-sidebar-widget blog-bor">
            <h2 class="blog-text-center blog-title"><span>About ME</span></h2>
            <img src="https://lattecake.oss-cn-beijing.aliyuncs.com/static/images/about_me.jpg" alt="about me"
                 class="blog-entry-img">
            <p>不跑马拉松的摄影师不是好程序员</p>
            <p>
                男, 一名91年天蝎座(光棍节那天 ∑( ° △ °|||)︴)的码农！毕业于遥远的南方某学校, 11年来北京并从业已5年有余。
            </p>
            <p>喜欢点一杯咖啡, 然后静静地坐在某个角落写点什么。喜欢在每个好天气的周末疯狂的跑步。偶尔有些挑食,立志要长肉的瘦子(会不会被人嫌弃)。</p>
            <p>与其用泪水悔恨昨天，不如用汗水拼搏今天。当眼泪流尽的时候，留下的应该是坚强。</p>
            <p>爱好：摄影、动漫、旅游、网球、马拉松</p>
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
                <li><a href="#">我把最深沉的秘密放在那里。</a></li>
                <li><a href="#">你不懂我，我不怪你。</a></li>
                <li><a href="#">每个人都有一道伤口， 或深或浅，盖上布，以为不存在。</a></li>
            </ul>
        </div>
    </div>
</div>
<!-- content end -->
