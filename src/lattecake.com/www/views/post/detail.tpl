<div class="banner-1">

</div>

<!-- technology-left -->
<div class="technology">
    <div class="container">
        <div class="col-md-9 technology-left">
            <div class="agileinfo">

                <h2 class="w3">{{.Post.Title}}</h2>
                <div class="single" style="line-height: 2.2em;">
                    {{range $key, $image := .Post.Images}}
                    {{if eq (0) ($key)}}
                    <img src="{{replace_image_src $image.RealPath}}?imageView2/1/w/736/h/356" class="pull-left"
                         alt="{{$image.RealPath}}"/>
                    {{end}}
                    {{end}}
                    <div class="b-bottom">
                        <!--<h5 class="top">{{.Post.Title}}</h5>-->
                        {{str2html .Content}}

                        <p>
                            本文地址: <a href="http://lattecake.com/post/{{.Post.Id}}" target="_blank">https://lattecake.com/post/{{.Post.Id}}</a>
                        </p>
                        <p>{{date .Post.PushTime "M/d/Y H:i:s"}} <a class="span_link" href="#"><span
                                class="glyphicon glyphicon-comment"></span>{{.Post.Reviews}}
                        </a><a class="span_link" href="#"><span class="glyphicon glyphicon-eye-open"></span>{{.Post.ReadNum}}
                        </a></p>

                    </div>
                </div>


                <!--<div class="response">-->
                <!--<h4>Responses</h4>-->
                <!--<div class="media response-info">-->
                <!--<div class="media-left response-text-left">-->
                <!--<a href="#">-->
                <!--<img src="/static/images/sin1.jpg" class="img-responsive" alt="">-->
                <!--</a>-->
                <!--</div>-->
                <!--<div class="media-body response-text-right">-->
                <!--<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit,There are many variations of-->
                <!--passages of Lorem Ipsum available,-->
                <!--sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>-->
                <!--<ul>-->
                <!--<li>Jun 21, 2016</li>-->
                <!--<li><a href="#">Reply</a></li>-->
                <!--</ul>-->
                <!--<div class="media response-info">-->
                <!--<div class="media-left response-text-left">-->
                <!--<a href="#">-->
                <!--<img src="/static/images/sin2.jpg" class="img-responsive" alt="">-->
                <!--</a>-->
                <!--</div>-->
                <!--<div class="media-body response-text-right">-->
                <!--<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit,There are many-->
                <!--variations of passages of Lorem Ipsum available,-->
                <!--sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>-->
                <!--<ul>-->
                <!--<li>July 17, 2016</li>-->
                <!--<li><a href="#">Reply</a></li>-->
                <!--</ul>-->
                <!--</div>-->
                <!--<div class="clearfix"></div>-->
                <!--</div>-->
                <!--</div>-->
                <!--<div class="clearfix"></div>-->
                <!--</div>-->
                <!--</div>-->
                <!--<div class="coment-form">-->
                <!--<h4>Leave your comment</h4>-->
                <!--<form action="#" method="post">-->
                <!--<input type="text" value="Name " name="name" onfocus="this.value = '';"-->
                <!--onblur="if (this.value == '') {this.value = 'Name';}" required="">-->
                <!--<input type="email" value="Email" name="email" onfocus="this.value = '';"-->
                <!--onblur="if (this.value == '') {this.value = 'Email';}" required="">-->
                <!--<input type="text" value="Website" name="websie" onfocus="this.value = '';"-->
                <!--onblur="if (this.value == '') {this.value = 'Website';}" required="">-->
                <!--<textarea onfocus="this.value = '';"-->
                <!--onblur="if (this.value == '') {this.value = 'Your Comment...';}" required="">Your Comment...</textarea>-->
                <!--<input type="submit" value="Submit Comment">-->
                <!--</form>-->
                <!--</div>-->
                <div class="clearfix"></div>
            </div>
        </div>
        <div class="col-md-3 technology-right">


            <div class="blo-top1">

                <div class="tech-btm">
                    <!--<div class="search-1">-->
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
                                <img src="{{replace_image_src $image.RealPath}}?imageView2/1/w/65/h/65"
                                     class="img-responsive" alt="replace_image_url">
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
    </div>
    <div class="clearfix"></div>
</div>