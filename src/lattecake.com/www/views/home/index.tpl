<!--start-main-->
<div class="header-bottom">
    <div class="container">
        <div class="logo wow fadeInDown" data-wow-duration=".8s" data-wow-delay=".2s">
            <h1><a href="index.html">DUDULU BLOG</a></h1>
            <p><label class="of"></label>LET'S MAKE A BEST SELF<label class="on"></label></p>
        </div>
    </div>
</div>
<!-- banner -->

<div class="banner">
    <div class="container">
        <h2>若没有过往 我们该怎样成长</h2>
        <p>人生最要紧的就是要保持微笑。生命如此美妙，有太多的事，都值得微笑以对。</p>
        <p>在我们认识的人中，有了误会，有了得失时，就会想起初见时的美丽。或者,那天在某个特定的地方,故地重游,突然发现多年未见的你,一下子就回到了初次相见时的情景,初相遇,那是怎样一种让人难以忘怀的美丽啊!</p>
        <!--<a href="singlepage.html">READ MORE</a>-->
    </div>
</div>
<div class="services w3l wow fadeInDown" data-wow-duration=".8s" data-wow-delay=".2s">
    <div class="container">
        <div class="grid_3 grid_5">
            <div class="bs-example bs-example-tabs" role="tabpanel" data-example-id="togglable-tabs">
                <ul id="myTab" class="nav nav-tabs" role="tablist">
                    <li role="presentation" class="active"><a href="#expeditions" id="expeditions-tab" role="tab"
                                                              data-toggle="tab" aria-controls="expeditions"
                                                              aria-expanded="true">LEARN</a></li>
                    <li role="presentation" class=""><a href="#safari" role="tab" id="safari-tab" data-toggle="tab"
                                                        aria-controls="safari">TRAVEL</a></li>
                    <li role="presentation" class=""><a href="#trekking" role="tab" id="trekking-tab" data-toggle="tab"
                                                        aria-controls="trekking">MUSIC</a></li>
                </ul>
                <div id="myTabContent" class="tab-content">
                    <div role="tabpanel" class="tab-pane fade" id="expeditions" aria-labelledby="expeditions-tab">
                        {{range $key, $post := .Learns}}
                        <div class="col-md-4 col-sm-5 tab-image">
                            <a href="/post/{{$post.Id}}">
                                {{range $k, $image := $post.Images}}
                                <!--{{$key}}-->
                                {{if eq (0) ($k)}}
                                <img src="{{replace_image_src $image.RealPath}}?imageView2/1/w/350/h/220"
                                     class="img-responsive" alt="{{$image.ImageName}}">
                                {{end}}
                                {{end}}
                            </a>
                        </div>
                        {{end}}

                        <div class="clearfix"></div>
                    </div>

                    <div role="tabpanel" class="tab-pane fade" id="safari" aria-labelledby="safari-tab">
                        {{range $key, $post := .TravelPosts}}
                        <div class="col-md-4 col-sm-5 tab-image">
                            <a href="/post/{{$post.Id}}">
                                {{range $k, $image := $post.Images}}
                                <!--{{$key}}-->
                                {{if eq (0) ($k)}}
                                <img src="{{replace_image_src $image.RealPath}}?imageView2/1/w/350/h/220"
                                     class="img-responsive" alt="{{$image.ImageName}}">
                                {{end}}
                                {{end}}
                            </a>
                        </div>
                        {{end}}
                        <div class="clearfix"></div>
                    </div>
                    <div role="tabpanel" class="tab-pane fade active in" id="trekking" aria-labelledby="trekking-tab">
                        {{range $key, $post := .Musics}}
                        <div class="col-md-4 col-sm-5 tab-image">
                            <a href="/post/{{$post.Id}}">
                                {{range $k, $image := $post.Images}}
                                <!--{{$key}}-->
                                {{if eq (0) ($k)}}
                                <img src="{{replace_image_src $image.RealPath}}?imageView2/1/w/350/h/220"
                                     class="img-responsive" alt="{{$image.ImageName}}">
                                {{end}}
                                {{end}}
                            </a>
                        </div>
                        {{end}}
                        <div class="clearfix"></div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- technology-left -->
<div class="technology">
    <div class="container">
        <div class="col-md-9 technology-left">
            <div class="tech-no">
                <!-- technology-top -->

                <div class="tc-ch wow fadeInDown" data-wow-duration=".8s" data-wow-delay=".2s">

                    <div class="tch-img">
                        <a href="/post/{{.StarPost.Post.Id}}">
                            {{range $key, $image := .StarPost.Post.Images}}
                            {{if eq (0) ($key)}}
                            <img src="{{replace_image_src $image.RealPath}}?imageView/1/w/750/h/350"
                                 class="img-responsive" alt="{{$image.ImageName}}">
                            {{end}}
                            {{end}}
                        </a>
                    </div>

                    <h3><a href="/post/{{.StarPost.Post.Id}}">{{.StarPost.Post.Title}}</a></h3>
                    <h6>BY <a href="/about"> 嘟嘟噜 </a>{{date .StarPost.Post.CreatedAt "M/d/Y"}}.</h6>
                    <p>{{.StarPost.Post.Description}}</p>
                    <div class="bht1">
                        <a href="/post/{{.StarPost.Post.Id}}">Continue Reading</a>
                    </div>
                    <div class="clearfix"></div>
                </div>
                <div class="clearfix"></div>
                <!-- technology-top -->
                <!-- technology-top -->
                <div class="w3ls">
                    {{range $k, $recommend := .MusicRecommend}}
                    <div class="col-md-6 w3ls-left wow fadeInDown" data-wow-duration=".8s" data-wow-delay=".2s">
                        <div class="tc-ch">
                            <div class="tch-img">
                                <a href="/post/{{$recommend.Post.Id}}">
                                    {{range $key, $image := $recommend.Post.Images}}
                                    {{if eq (0) ($key)}}
                                    <img src="{{replace_image_src $image.RealPath}}?imageView/1/w/350/h/250"
                                         class="img-responsive" alt="{{$image.ImageName}}">
                                    {{end}}
                                    {{end}}
                                </a>
                            </div>

                            <h3><a href="/post/{{$recommend.Post.Id}}">{{$recommend.Post.Title}}</a></h3>
                            <h6>BY <a href="/about">嘟嘟噜 </a>{{date .StarPost.Post.CreatedAt "M/d/Y"}}.</h6>
                            <p>{{substr $recommend.Post.Description 0 40}}</p>
                            <div class="bht1">
                                <a href="/post/{{$recommend.Post.Id}}">Read More</a>
                            </div>
                            <div class="clearfix"></div>
                        </div>
                    </div>
                    {{end}}
                    <div class="clearfix"></div>
                </div>

                <!-- technology-top -->
                {{range $k, $recommend := .TravelRecommend}}
                <div class="wthree">
                    <div class="col-md-6 wthree-left wow fadeInDown" data-wow-duration=".8s" data-wow-delay=".2s">
                        <div class="tch-img">
                            <a href="/post/{{$recommend.Post.Id}}">
                                {{range $key, $image := $recommend.Post.Images}}
                                {{if eq (0) ($key)}}
                                <img src="{{replace_image_src $image.RealPath}}?imageView/1/w/350/h/220"
                                     class="img-responsive" alt="{{$image.ImageName}}">
                                {{end}}
                                {{end}}
                            </a>
                        </div>
                    </div>
                    <div class="col-md-6 wthree-right wow fadeInDown" data-wow-duration=".8s" data-wow-delay=".2s">
                        <h3><a href="/post/{{$recommend.Post.Id}}">{{$recommend.Post.Title}}</a></h3>
                        <h6>BY <a href="/about">嘟嘟噜 </a>{{date .StarPost.Post.CreatedAt "M/d/Y"}}.</h6>
                        <p>{{substr $recommend.Post.Description 0 70}}</p>
                        <div class="bht1">
                            <a href="/post/{{$recommend.Post.Id}}">Read More</a>
                        </div>
                        <div class="clearfix"></div>

                    </div>
                    <div class="clearfix"></div>
                </div>
                {{end}}
            </div>
        </div>

        <!-- technology-right -->
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
        <div class="clearfix"></div>
        <!-- technology-right -->
    </div>
</div>