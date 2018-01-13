<!-- content srart -->
<div class="am-g am-g-fixed blog-fixed">
    <div class="am-u-md-12 am-u-sm-12">

    {{range $k, $post := .posts}}
        <article class="am-g blog-entry-article">
            <div class="am-u-lg-6 am-u-md-12 am-u-sm-12 blog-entry-img">
                <img src="{{replace_image_src $post.Image}}?imageView/1/w/1280/h/720" alt="" class="am-u-sm-12">
            </div>
            <div class="am-u-lg-6 am-u-md-12 am-u-sm-12 blog-entry-text">
                <span><a href="/post/{{$post.Id}}" class="blog-color">article &nbsp;</a></span>
                <span> @嘟嘟噜 &nbsp;</span>
                <span>{{date $post.PushTime "M/d/Y"}}</span>
                <h1><a href="/post/{{$post.Id}}">{{$post.Title}}</a></h1>
                <p>{{substr $post.Description 0 50}}</p>
                <p><a href="/post/{{$post.Id}}" class="blog-continue">continue reading</a></p>
            </div>
        </article>

    {{end}}
        <ul class="am-pagination">
            <li class="am-pagination-prev"><a href="/posts/{{.prev}}">&laquo; Prev</a></li>
            <li class="am-pagination-next"><a href="/posts/{{.next}}">Next &raquo;</a></li>
        </ul>
    </div>

</div>
<!-- content end -->