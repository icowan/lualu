<div class="banner-1">

</div>

<!-- technology-left -->
<div class="technology">
    <div class="container">
        <div class="col-md-9 technology-left">
            <div class="w3agile-1">
                <div class="welcome">
                    <div class="welcome-top heading">
                        <h2 class="w3">About</h2>
                        <div class="welcome-bottom">
                            <img src="https://lattecake.oss-cn-beijing.aliyuncs.com/static/images/about_me.jpg" class="img-responsive" alt="">
                            <p>不跑马拉松的摄影师不是好程序员（＾ω＾）, 男, 一名91年天蝎座(光棍节那天 ∑( ° △ °|||)︴)的码农！12年毕业于遥远的南方某学校, 11年来北京并入行,
                                从业已5年有余。</p>
                            <p>喜欢点一杯摩卡咖啡, 静静地坐在咖啡馆学习学习，或写点什么。喜欢在每个好天气的周末疯狂的跑步，或许是想在忘掉什么。偶尔有些挑食,立志要长肉的瘦子(会不会被人嫌弃)。</p>
                            <p>座右铭：与其用泪水悔恨昨天，不如用汗水拼搏今天。当眼泪流尽的时候，留下的应该是坚强。</p>
                            <p>爱好：摄影、动漫、旅游、网球、马拉松</p>
                        </div>
                    </div>
                </div>
                <div class="team">
                    <h3 class="team-heading">Other</h3>
                    <div class="team-grids">
                        <div class="col-md-6 team-grid">
                            <div class="team-grid1">
                                <img src="https://lattecake.oss-cn-beijing.aliyuncs.com/static%2Fimages%2Fweixin%2Fqrcode_for_gh_354bc8e8b814_1280.jpg" alt=" " class="img-responsive">
                                <div class="p-mask">
                                    <p>一起努力一起学习.</p>
                                </div>
                            </div>
                            <h5>聪聪实验室<span>欢迎关注</span></h5>
                        </div>
                        <div class="col-md-6 team-grid">
                            <div class="team-grid1">
                                <img src="https://lattecake.oss-cn-beijing.aliyuncs.com/static%2Fimages%2Freward%2Fweixin-RMB-xxx.JPG" alt=" " class="img-responsive">
                                <div class="p-mask">
                                    <p>支持微信、支付宝.</p>

                                </div>
                            </div>
                            <h5>欢迎打赏<span>多多益善</span></h5>
                        </div>


                        <div class="clearfix"></div>
                    </div>
                </div>
            </div>
        </div>

        <div class="col-md-3 technology-right">


            <div class="blo-top1">

                <div class="tech-btm">
                    <!--<div class="search-1 wow fadeInDown" data-wow-duration=".8s" data-wow-delay=".2s">-->
                    <!--<form action="#" method="post">-->
                    <!--<input type="search" name="Search" value="Search" onfocus="this.value = '';"-->
                    <!--onblur="if (this.value == '') {this.value = 'Search';}" required="">-->
                    <!--<input type="submit" value=" ">-->
                    <!--</form>-->
                    <!--</div>-->
                    <h4>Popular Posts </h4>

                    {{range $key, $val := .PopularPosts}}
                    <div class="blog-grids">
                        <div class="blog-grid-left">
                            <a href="/post/{{$val.Id}}">
                                {{range $key, $image := $val.Images}}
                                {{if eq (0) ($key)}}
                                <img src="{{replace_image_src $image.RealPath}}?imageView2/1/w/65/h/65" class="img-responsive" alt="replace_image_url">
                                {{end}}
                                {{end}}
                            </a>
                        </div>
                        <div class="blog-grid-right">

                            <h5><a href="/post/{{$val.Id}}">{{$val.Title}}</a></h5>
                        </div>
                        <div class="clearfix"></div>
                    </div>
                    {{end}}

                    <div class="insta">
                        <h4>Links</h4>
                        <ul>

                            <li><a href="http://xcgx.me/" target="_blank"><img
                                    src="https://lattecake.oss-cn-beijing.aliyuncs.com/static/images/avatar/xcgx-me.jpeg"
                                    class="img-responsive"
                                    alt=""></a>
                            </li>
                            <li><a href="http://dudulu.me/" target="_blank"><img
                                    src="https://ofbudvg4c.qnssl.com/images/avatar11821294_950965884967105_1200550191_n.jpg?imageView2/1/w/50/h/50"
                                    class="img-responsive"
                                    alt=""></a>
                            </li>
                            <li><a href="http://ylsc633.com/" target="_blank"><img
                                    src="https://lattecake.oss-cn-beijing.aliyuncs.com/static%2Fimages%2Favatar%2Fylsc633.jpg"
                                    class="img-responsive"
                                    alt=""></a>
                            </li>
                            <li><a href="http://tellu.top/" target="_blank"><img
                                    src="https://lattecake.oss-cn-beijing.aliyuncs.com/static/images/avatar/%40V%60NQ%60L_%25%60_JD%24%5DM%60YENCCS.jpg"
                                    class="img-responsive"
                                    alt=""></a>
                            </li>
                            <li><a href="http://szhshp.org/" target="_blank"><img
                                    src="https://ofbudvg4c.qnssl.com/static/avatars/portrait.jpg"
                                    class="img-responsive"
                                    alt=""></a>
                            </li>

                        </ul>
                    </div>

                    <!--<p>Lorem ipsum ex vix illud nonummy, novum tation et his. At vix scripta patrioque scribentur, at-->
                    <!--pro</p>-->
                </div>


            </div>

    </div>
</div>