
<!-- nav start -->
<nav class="am-g am-g-fixed blog-fixed blog-nav">
    <button class="am-topbar-btn am-topbar-toggle am-btn am-btn-sm am-btn-success am-show-sm-only blog-button" data-am-collapse="{target: '#blog-collapse'}" ><span class="am-sr-only">导航切换</span> <span class="am-icon-bars"></span></button>

    <div class="am-collapse am-topbar-collapse" id="blog-collapse">
        <ul class="am-nav am-nav-pills am-topbar-nav">
            <li {{if eq "HomeController_Index" .Active}}class="am-active"{{end}}><a href="/">Home</a></li>
            <li {{if eq "HomeController_About" .Active}}class="am-active"{{end}}><a href="/about">About</a></li>
            <li {{if eq "PostController_GetAll" .Active}}class="am-active"{{end}}><a href="/posts">Posts</a></li>
            <li {{if eq "HomeController_Reward" .Active}}class="am-active"{{end}}><a href="/reward">Reward</a></li>
            {{/*<li><a href="/archives">Archives</a></li>*/}}
        </ul>
        {{/*<form class="am-topbar-form am-topbar-right am-form-inline" role="search">*/}}
            {{/*<div class="am-form-group">*/}}
                {{/*<input type="text" class="am-form-field am-input-sm" placeholder="搜索">*/}}
            {{/*</div>*/}}
        {{/*</form>*/}}
    </div>
</nav>
<!-- nav end -->
<hr>