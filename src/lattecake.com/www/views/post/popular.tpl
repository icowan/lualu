<!-- technology-right -->
<div class="col-md-3 technology-right">


    <div class="blo-top1">

        <div class="tech-btm">
            <div class="search-1">
                <form action="#" method="post">
                    <input type="search" name="Search" value="Search" onfocus="this.value = '';"
                           onblur="if (this.value == '') {this.value = 'Search';}" required="">
                    <input type="submit" value=" ">
                </form>
            </div>
            <h4>Popular Posts </h4>


            {{range $key, $val := .Popular}}
            <div class="blog-grids">
                <div class="blog-grid-left">
                    <a href="singlepage.html"><img src="/static/images/t2.jpg" class="img-responsive" alt=""></a>
                </div>
                <div class="blog-grid-right">

                    <h5><a href="singlepage.html">{{$val.Title}}</a></h5>
                </div>
                <div class="clearfix"></div>
            </div>
            {{end}}


            <div class="insta">
                <h4>Instagram</h4>
                <ul>

                    <li><a href="singlepage.html"><img src="/static/images/t1.jpg" class="img-responsive" alt=""></a>
                    </li>
                    <li><a href="singlepage.html"><img src="/static/images/m1.jpg" class="img-responsive" alt=""></a>
                    </li>
                    <li><a href="singlepage.html"><img src="/static/images/f1.jpg" class="img-responsive" alt=""></a>
                    </li>
                    <li><a href="singlepage.html"><img src="/static/images/m2.jpg" class="img-responsive" alt=""></a>
                    </li>
                    <li><a href="singlepage.html"><img src="/static/images/f2.jpg" class="img-responsive" alt=""></a>
                    </li>
                    <li><a href="singlepage.html"><img src="/static/images/t2.jpg" class="img-responsive" alt=""></a>
                    </li>
                    <li><a href="singlepage.html"><img src="/static/images/f3.jpg" class="img-responsive" alt=""></a>
                    </li>
                    <li><a href="singlepage.html"><img src="/static/images/t3.jpg" class="img-responsive" alt=""></a>
                    </li>
                    <li><a href="singlepage.html"><img src="/static/images/m3.jpg" class="img-responsive" alt=""></a>
                    </li>
                </ul>
            </div>

            <p>Lorem ipsum ex vix illud nonummy, novum tation et his. At vix scripta patrioque scribentur, at
                pro</p>
        </div>


    </div>


</div>
<div class="clearfix"></div>
<!-- technology-right -->